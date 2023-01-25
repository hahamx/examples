import sys
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