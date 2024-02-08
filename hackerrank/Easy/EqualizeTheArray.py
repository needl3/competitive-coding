#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the equalizeArray function below.
def equalizeArray(arr):
    counted = []
    m = 0
    for i in range(n):
        count = 0
        if arr[i] not in counted:
            for j in range(n):
                if arr[j] == arr[i]:
                    count += 1
            counted.append(arr[i])
        m = count if count > m else m
    return n-m


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    result = equalizeArray(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
