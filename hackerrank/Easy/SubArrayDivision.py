#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the birthday function below.
def birthday(s, d, m):
    num = 0
    #Make a list of subsets of set s having length of m
    subsets = []
    if len(s) > 1:
        print('Greater')
        for i in range(len(s)-m+1):
            temp = [s[i]]
            for j in range(i+1,m+i):
                temp.append(s[j])
            print('Temp is: ', temp)
            subsets.append(temp)
    else:
        print('Less')
        num = 1 if s[0] == d and m == len(s) else 0


    #Check for elements summing to d
    for i in subsets:
        print('Summing: ', i, ' = ', sum(i))
        num += 1 if sum(i) == d else 0
        
    #Return the number of elements matching the condition
    return num

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    s = list(map(int, input().rstrip().split()))

    dm = input().rstrip().split()

    d = int(dm[0])

    m = int(dm[1])

    result = birthday(s, d, m)

    fptr.write(str(result) + '\n')

    fptr.close()
