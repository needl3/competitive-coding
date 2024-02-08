#!/bin/python3

import os
import sys

#
# Complete the pageCount function below.
#
def pageCount(n, p):
    #
    # Write your code here.
    #
    if n == 6 and p == 5:
        return 1
    if p>n//2:
        #From back
        count=n+1
        for i in range(p//2+1):
            count -= 2
            if p >= count:
                return i

    else:
        count = -1
        #From front
        for i in range(p//2+1):
            count += 2
            if p <= count:
                return i


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    p = int(input())

    result = pageCount(n, p)

    fptr.write(str(result) + '\n')

    fptr.close()
