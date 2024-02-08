#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the squares function below.
def squares(a, b):
    lower = None
    count = 0
    for i in range(a,b+1):
        if round(pow(i,0.5))**2 == i:
            lower = round(pow(i,0.5))
            break
    if lower != None:
        while lower*lower <= b:
            count += 1
            lower += 1

    return count
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        ab = input().split()

        a = int(ab[0])

        b = int(ab[1])

        result = squares(a, b)

        fptr.write(str(result) + '\n')

    fptr.close()
