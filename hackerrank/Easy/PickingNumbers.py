#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'pickingNumbers' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY a as parameter.
#

def pickingNumbers(a):
    # Write your code here
    m = 0
    for i in range(len(a)):
        temp = [a[i]]
        print('First element appended: '+str(a[i]))
        count = 2
        for j in range(len(a)):
            if i == j:
                pass
            elif abs(a[i]-a[j]) <=1 and len(temp)==1:
                print('2nd element appended: '+str(a[j]))
                temp.append(a[j])
            elif a[j] in temp:
                print(str(count)+'th element appended: '+str(j))
                temp.append(a[j])
            count += 1
        m = len(temp) if len(temp)>m else m
        del temp[:]
        print('Removed elements from temp',temp)
        print('\n')
    return m
            

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    a = list(map(int, input().rstrip().split()))

    result = pickingNumbers(a)

    fptr.write(str(result) + '\n')

    fptr.close()
