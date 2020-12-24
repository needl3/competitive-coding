#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the kaprekarNumbers function below.
def kaprekarNumbers(p, q):
    isKaprekar = []
    for i in range(p,q+1):
        length = len(str(i))
        temp = i*i
        if length == 1 and i < 4:
            l = 0
            r = temp
        else:
            temp = str(temp)
            r = int(''.join(temp[-length:]))
            l = int(''.join(temp[:-length]))
        if l+r == i:
            isKaprekar.append(str(i))
    isKaprekar = ' '.join(isKaprekar)
    if len(isKaprekar) != 0:
        print(isKaprekar)
    else:
        print('INVALID RANGE')
            
            

        
if __name__ == '__main__':
    p = int(input())

    q = int(input())

    kaprekarNumbers(p, q)
