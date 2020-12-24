#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the beautifulTriplets function below.
def beautifulTriplets(d, arr):
    index = 0
    triplet_count = 0
    while index <= n-1:
        count = 1
        temp_triplets = [arr[index]]
        for i in range(index+1,n):
            if temp_triplets[-1] < arr[i] and arr[i]-temp_triplets[-1] == d:
                temp_triplets.append(arr[i])
                count += 1 
            if count == 3:
                break
        if count == 3:
            triplet_count += 1
        index += 1
    return triplet_count



if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nd = input().split()

    n = int(nd[0])

    d = int(nd[1])

    arr = list(map(int, input().rstrip().split()))

    result = beautifulTriplets(d, arr)

    fptr.write(str(result) + '\n')

    fptr.close()
