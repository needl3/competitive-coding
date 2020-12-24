#!/bin/python3
import math
import os
import random
import re
import sys
# Complete the repeatedString function below.
def repeatedString(s, n):
    a_in_given = 0
    for i in s:
        if i == 'a':
            a_in_given += 1        
    multiple = n//len(s)
    final_a = multiple*a_in_given
    for i in range(n-len(s)*multiple):
        if s[i] == 'a':
            final_a += 1
    return final_a
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')
    s = input()
    n = int(input())
    result = repeatedString(s, n)
    fptr.write(str(result) + '\n')
    fptr.close()
