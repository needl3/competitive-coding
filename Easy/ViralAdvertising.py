#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the viralAdvertising function below.
def viralAdvertising(n):
    cum_likes = 0
    temp_shares = 5
    for i in range(1,n+1):
        temp_likes = math.floor(temp_shares//2)
        cum_likes += temp_likes
        temp_shares = temp_likes * 3
    return cum_likes

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    result = viralAdvertising(n)

    fptr.write(str(result) + '\n')

    fptr.close()
