#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the jumpingOnClouds function below.
def jumpingOnClouds(c):
    index = 0
    jump = 0
    while index < len(c):
        try:
            if c[index+1] == 0:
                try:
                    if c[index+2] == 0:
                        index += 2
                        jump += 1
                    elif c[index+2] == 1:
                        index += 1
                        jump += 1
                except IndexError:
                    index +=1
                    jump += 1
            elif c[index+1] == 1:
                index += 2
                jump += 1
        except IndexError:
            index += 1
    return jump
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    c = list(map(int, input().rstrip().split()))

    result = jumpingOnClouds(c)

    fptr.write(str(result) + '\n')

    fptr.close()
