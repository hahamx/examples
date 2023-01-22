import asyncio

 
async def client_connected(reader, writer):
    # Communicate with the client with
    # reader/writer streams.  For example:
    await reader.readline()

async def main5(host, port):
    """
    开始接受连接，直到协程被取消。 serve_forever 任务的取消将导致服务器被关闭。
    如果服务器已经在接受连接了，这个方法可以被调用。每个 Server 对象，
    仅能有一个 serve_forever 任务。
    """
    srv = await asyncio.start_server(
        client_connected, host, port)
    await srv.serve_forever()

def hello_world(loop):
    """
    A callback to print 'Hello World' and stop the event loop
    """
    print('He,WOrld')
    loop.stop()

def init_loop():
    """
    loop run forever 安排回调
    """

    loop = asyncio.get_event_loop()

    # Schedule a call to hello_world 计划一个回调
    loop.call_soon(hello_world, loop)

    try:
        loop.run_forever()
    finally:
        loop.close()

# 使用 call later 展示当前日期
import datetime
def display_date(end_time, loop):
    """
    回调
    递归调用
    """
    print(datetime.datetime.now())
    if (loop.time() + 1.0 ) < end_time:
        # 每 2 秒调用一次
        loop.call_later(2, display_date, end_time, loop)
    else:
        loop.stop()

def start_display():
    loop = asyncio.get_event_loop()
    # 为display_date 计划一个 第一次调用
    end_time = loop.time() + 5.0
    loop.call_soon(display_date, end_time, loop)

    # 使用 loop.stop() 阻塞中断所有的 调用
    try:
        loop.run_forever()
    finally:
        loop.close()


# 监控一个文件描述符的 读事件
from socket import socketpair

# 创建一对 读写 对象
rsock, wsock = socketpair()
loop = asyncio.get_event_loop()

def reader():
    """
    读事件
    """
    data = rsock.recv(100)
    print("Received:", data.decode())
    # 完成,未注册 文件描述符
    loop.remove_reader(rsock)
    # 停止事件循环
    loop.stop()

def start_reader_event():
    """

    """
    # 注册 文件描述符到 读取事件
    loop.add_reader(rsock, reader)
    # simulate 模拟 从网络接收数据
    loop.call_soon(wsock.send, "aaadddd".encode())

    try:
        # 运行事件循环
        loop.run_forever()
    finally:
        #  完成并关闭 sockets 和 事件循环
        rsock.close()
        wsock.close()
        loop.close()


## 为SIGINT SIGTERM 设置信号处理器
## 使用 loop.add_signal_handler()方法为信号 SIGINT SIGTERM注册处理程序
import functools
import os
import signal

def ask_exit(signame, loop):
    print("got signal %s: exit" % signame)
    loop.stop()

async def main():
    """
        > kill -1 36242
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36242: send SIGINT or SIGTERM to exit.
        Hangup

        > kill -3 36271
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36271: send SIGINT or SIGTERM to exit.
        got signal SIGINT: exit
        Traceback (most recent call last):
          File "/data/code/asynctest/asynceventloop.py", line 130, in <module>
            asyncio.run(main())
          File "/usr/local/python39/lib/python3.9/asyncio/runners.py", line 44, in run
            return loop.run_until_complete(main)
          File "/usr/local/python39/lib/python3.9/asyncio/base_events.py", line 640, in run_until_complete
            raise RuntimeError('Event loop stopped before Future completed.')
        RuntimeError: Event loop stopped before Future completed.

        > kill -3 36280
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36280: send SIGINT or SIGTERM to exit.
        Quit (core dumped)

        > kill -4 36298
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36298: send SIGINT or SIGTERM to exit.
        Illegal instruction (core dumped)

        >  kill -5 36325
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36325: send SIGINT or SIGTERM to exit.
        Trace/breakpoint trap (core dumped)

        > kill -6 36333
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36333: send SIGINT or SIGTERM to exit.
        Aborted (core dumped)

        > kill -7 36341
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36341: send SIGINT or SIGTERM to exit.
        Bus error (core dumped)

        >kill -8 36353
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36353: send SIGINT or SIGTERM to exit.
        Floating point exception (core dumped)

        > kill -9 36361
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36361: send SIGINT or SIGTERM to exit.
        Killed

        > kill -s SIGINT 36368
        # python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36368: send SIGINT or SIGTERM to exit.
        got signal SIGINT: exit
        Traceback (most recent call last):
          File "/data/code/asynctest/asynceventloop.py", line 130, in <module>
            asyncio.run(main())
          File "/usr/local/python39/lib/python3.9/asyncio/runners.py", line 44, in run
            return loop.run_until_complete(main)
          File "/usr/local/python39/lib/python3.9/asyncio/base_events.py", line 640, in run_until_complete
            raise RuntimeError('Event loop stopped before Future completed.')
        RuntimeError: Event loop stopped before Future completed.

       >kill -s SIGTERM 36525
       python asynceventloop.py
        Event loop running for 1 hour, press Ctrl+C to interrupt.
        pid 36525: send SIGINT or SIGTERM to exit.
        got signal SIGTERM: exit
        Traceback (most recent call last):
          File "/data/code/asynctest/asynceventloop.py", line 130, in <module>
            asyncio.run(main())
          File "/usr/local/python39/lib/python3.9/asyncio/runners.py", line 44, in run
            return loop.run_until_complete(main)
          File "/usr/local/python39/lib/python3.9/asyncio/base_events.py", line 640, in run_until_complete
            raise RuntimeError('Event loop stopped before Future completed.')
        RuntimeError: Event loop stopped before Future completed.
    """
    loop = asyncio.get_running_loop()
    print("Event loop running for 1 hour, press Ctrl+C to interrupt.")
    print(f"pid {os.getpid()}: send SIGINT or SIGTERM to exit.")
    for signame in {'SIGINT', 'SIGTERM'}:
        # 使用 loop.add_signal_handler() 方法为信号 SIGINT 和 SIGTERM 注册处理程序
        loop.add_signal_handler(
            getattr(signal, signame),
            functools.partial(ask_exit, signame, loop))

    await asyncio.sleep(36)

async def test():
    print("never scheduled")

async def main11():
    test()

if __name__ == '__main__':
    # init_loop()

    # start_display()
    # start_reader_event()

    # 处理输入信号
    print()
    asyncio.run(main11(), debug=True)


