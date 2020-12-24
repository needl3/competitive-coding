#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the migratoryBirds function below.
def migratoryBirds(arr):
    spotted = {1:0,2:0,3:0,4:0,5:0}
    for i in arr:
        spotted[i] += 1
    max = 0
    key = None
    for i in range(1,6):
        if spotted[i] > max:
            max = spotted[i]
            key = i
    return key


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr_count = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    result = migratoryBirds(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
