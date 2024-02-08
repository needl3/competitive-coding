#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the sockMerchant function below.
def sockMerchant(n, ar):
    done = []
    counted = []
    for i in ar:
        if i in done:
            pass
        else:
            count = 0
            for j in ar:
                if j == i:
                    count += 1
            done.append(i)
            counted.append(count)
    pairs=0
    for i in counted: 
        pairs += i//2
    return pairs
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    ar = list(map(int, input().rstrip().split()))

    result = sockMerchant(n, ar)

    fptr.write(str(result) + '\n')

    fptr.close()
