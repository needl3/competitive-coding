#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'diagonalDifference' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY arr as parameter.
#

def diagonalDifference(arr):
    global n
    # Write your code here
    l_sum = []
    r_sum = []
    for i in range(n):
        for j in range(n):
            if i == j:
                l_sum.append(arr[i][j])
            if i+j == n-1:
                r_sum.append(arr[i][j])
    return abs(sum(l_sum)-sum(r_sum))


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    arr = []

    for _ in range(n):
        arr.append(list(map(int, input().rstrip().split())))

    result = diagonalDifference(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
