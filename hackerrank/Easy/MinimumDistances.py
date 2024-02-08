#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the minimumDistances function below.
def minimumDistances(a):
    distance = -1
    for i in range(n):
        temp = None
        for j in range(i+1,n):
            if a[i] == a[j]:
                temp = j-i
                break
        if temp != None:
            if distance == -1:
                distance = temp
            else:
                distance = temp if temp < distance else distance
    return distance


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    a = list(map(int, input().rstrip().split()))

    result = minimumDistances(a)

    fptr.write(str(result) + '\n')

    fptr.close()
