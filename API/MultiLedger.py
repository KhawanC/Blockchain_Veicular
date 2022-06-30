import multiprocessing
import random, time, threading
from multiprocessing import Process
from webbrowser import get

global queue

listPlacas = []

def getLetter():
    letras = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
        'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
    letters = random.choice(letras)
    queue.enque(letters)

def getNum():
    nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    num = random.choice(nums)
    queue.enque(str(num))

def ProcessPlaca():
    thredProc = []
    for _ in range(3):
        t = threading.Thread(target=getLetter)
        t.start()
        thredProc.append(t)
    for _ in range(1):
        j = threading.Thread(target=getNum)
        j.start()
        thredProc.append(t)
    for _ in range(1):
        t = threading.Thread(target=getLetter)
        t.start()
        thredProc.append(t)
    for _ in range(2):
        j = threading.Thread(target=getNum)
        j.start()
        thredProc.append(t)
    for t in thredProc:
        t.join()
    placa = str(queue)
    placaFormatada = placa.replace("'", "").replace("[", "").replace("]", "").replace(",", "").replace(" ", "")[::-1]
    print(placaFormatada)

class Queue(object):
 
    def __init__(self):
        self.item = []

    def __str__(self):
        return "{}".format(self.item)

    def __repr__(self):
        return "{}".format(self.item)

    def enque(self, item):
        """
        Insert the elements in queue
        :param item: Any
        :return: Bool
        """
        self.item.insert(0, item)
        return True

    def size(self):
        """
        Return the size of queue
        :return: Int
        """
        return len(self.item)

    def dequeue(self):
        """
        Return the elements that came first
        :return: Any
        """
        if self.size() == 0:
            return None
        else:
            return self.item.pop()

    def peek(self):
        """
        Check the Last elements
        :return: Any
        """
        if self.size() == 0:
            return None
        else:
            return self.item[-1]

    def isEmpty(self):
        """
        Check is the queue is empty
        :return: bool
        """
        if self.size() == 0:
            return True
        else:
            return False
 
queue = Queue()


if __name__ == '__main__':

    # process = []
    # for i in range(1000):
    #     p = Process(target=ProcessPlaca)
    #     p.start()
    #     process.append(p)
    # for p in process:
    #     p.join()
    print('h')
    
    
    