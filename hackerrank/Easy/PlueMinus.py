#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the plusMinus function below.
def plusMinus(arr):
    pos = 0.0
    neg = 0.0
    zer = 0.0
    for i in arr:
        if i > 0:
            pos += 1.0
        elif i < 0:
            neg += 1.0
        else:
            zer += 1.0
    print('{:.6f}'.format(pos/len(arr)))
    print('{:.6f}'.format(neg/len(arr)))
    print('{:.6f}'.format(zer/len(arr)))

if __name__ == '__main__':
    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
