#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the miniMaxSum function below.
def miniMaxSum(arr):
    sum_arr = []
    for i in range(5):
        temp_sum = 0
        for j in range(5):
            if i == j:
                pass
            else:
                temp_sum += arr[j]
        sum_arr.append(temp_sum)
    print(min(sum_arr),max(sum_arr))
    
if __name__ == '__main__':
    arr = list(map(int, input().rstrip().split()))

    miniMaxSum(arr)
