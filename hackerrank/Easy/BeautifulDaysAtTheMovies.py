#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the beautifulDays function below.
def beautifulDays(i, j, k):
    num = 0
    for i in range(i,j+1):
        l = i
        rev = 0
        while l != 0:
            rem = l%10
            l = l//10
            rev = rev*10+rem
        if abs(i-rev)%k == 0:
            num += 1

    return num
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    ijk = input().split()

    i = int(ijk[0])

    j = int(ijk[1])

    k = int(ijk[2])

    result = beautifulDays(i, j, k)

    fptr.write(str(result) + '\n')

    fptr.close()
