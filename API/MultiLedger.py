from inspect import trace
from multiprocessing import Process
from time import sleep

def myFunc(num):
    print(3*5*num)

if __name__ == '__main__':
    nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]
    process = []
    for i in nums:
        proc = Process(target=myFunc, args=(i,))
        proc.start()
    for p in process:
        p.join()
    