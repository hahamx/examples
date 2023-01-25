
import sys
import time
 
from concurrent.futures import Future
from threading import Thread, Lock
from asyncio import wait_for, wrap_future, get_event_loop
import asyncio
from time import sleep
 

def from_coroutine():

    return sys._getframe(2).f_code.co_flags & 0x380


def which_side():
    if from_coroutine():
        print("White")
    else:
        print("Black")


def spam():
    which_side()


async def aspam():
    which_side()


class Queuey():
    def __init__(self, maxsize):
        self.mutex = Lock()
        self.maxsize = maxsize
        self.items = list()
        self.getters = list()
        self.putters = list()
    
    # 获取时加锁  # 唤醒一个 待添加的 fut 元素
    def get_noblock(self):
        with self.mutex:  
            if self.items:
                
                if self.putters:
                    self.putters.pop(0).set_result(True)
                # 返回队列 左部 第一一个元素
                return self.items.pop(0), None
            else:
                fut = Future()
                self.getters.append(fut)
                return fut, None
    
    # 添加时 加锁
    def put_noblock(self, item):
        with self.mutex:  
            if len(self.items) < self.maxsize:
                self.items.append(item)
                # Wake a getter
                if self.getters:
                    self.getters.pop(0).set_result(
                        self.items.pop(0)
                    )
            else:
                fut = Future()
                self.putters.append(fut)
                return fut

    def get_sync(self):
        item, fut = self.get_noblock()
        if fut:
            item = fut.result()
        return item

    async def get_async(self):
        item, fut = self.get_noblock()
        if fut:
            item = await wait_for(wrap_future(fut), None)
        return item

     # 如果是协程,使用 协程的方式 取值  # 不是协程,使用同步的方式 取值
    def get(self):
        if from_coroutine(): 
            print(f"get async item")
            return self.get_async()
        else: 
            print(f"get sync item:")

            return self.get_sync() if len(self.items) > 0 else None

    def put_sync(self, item):

        while True:
            fut = self.put_noblock(item)
            if fut is None:
                return
            fut.result()

    async def put_async(self, item):

        while True:
            fut = self.put_noblock(item)
            if fut is None:
                return
            await wait_for(wrap_future(fut), None)
    
    # 如果是协程,使用协程的方式 放入值  # 否则 使用同步的方式 添加值
    def put(self, item):

        if from_coroutine():  
            print(f"put async item:", item)
            return self.put_async(item)
        else:  
            print(f"put sync item:", item)
            return self.put_sync(item)


def producer(q, n):

    for i in range(n):
        q.put(i)
    q.put(None)


async def aproducer(q, n):

    for i in range(n): 
        await q.put(i) 
    await q.put(None)

 


async def aconsumer(q):

    while True: 
         
        item = await q.get()  
        if item is None:
            break
        print("Async Got:", item)

if __name__ == '__main__':
 
     
    loop = get_event_loop()
    # case 1 使用 入口判断 代码执行 同步还是异步
    q = Queuey(2)
    # producer 同步 存入
    Thread(target=producer, args=(q, 10)).start()

    time.sleep(1)
    # 异步取出
    loop.run_until_complete(aconsumer(q))
    
    # case 2 异步存取
    # from asyncio import Queue
    #
    # q = Queue()
    # loop.create_task(aconsumer(q))  # 异步取
    # loop.run_until_complete(aproducer(q, 10))  # 异步存