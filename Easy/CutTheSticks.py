#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the cutTheSticks function below.
def cutTheSticks(arr):
    length = len(arr)
    arr_length = []
    while True:
        print(length)
        arr_length.append(length)
        shortest = min(arr)
        count_shortest = 0
        #To determine the last set of numbers
        diff = 0
        prev = None
        for i in arr:
            if prev == None:
                prev = i
            elif prev != i:
                diff += 1
            if i == shortest:
                count_shortest += 1
            else:
                i -= shortest
        for i in range(count_shortest):
            arr.remove(shortest)
            length -= 1
        if diff == 0:
            break
    return arr_length


            


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    result = cutTheSticks(arr)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
