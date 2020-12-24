#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the flatlandSpaceStations function below.
def flatlandSpaceStations(n, c):
    maxDist = 0
    c.sort()

    if len(c) == 1:
        return c[0] if c[0] > n-c[0]-1 else n-c[0]-1
    for index, value in enumerate(c):
        if index == 0:
            rightVal = (value+c[index+1])//2
            leftVal = 0
        else:
            if value == c[-1]:
                rightVal = n-1    
            else:
                rightVal = (value+c[index+1])//2
            
            if (value+c[index-1])%2 != 0:
                leftVal = (value+c[index-1])//2+1
            else:
                leftVal = (value+c[index-1])//2

        val = value-leftVal if value-leftVal > rightVal-value else rightVal-value
        maxDist = maxDist if maxDist > val else val
    return maxDist

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nm = input().split()

    n = int(nm[0])

    m = int(nm[1])

    c = list(map(int, input().rstrip().split()))

    result = flatlandSpaceStations(n, c)

    fptr.write(str(result) + '\n')

    fptr.close()
