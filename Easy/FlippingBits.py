#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the flippingBits function below.
def flippingBits(n):
    binary = []
    while n > 0:
        binary.append(n%2)
        n = n//2
    print(len(binary))
    for i in range(32-len(binary)):
        binary.append(0)
    new_binary = []
    for i in binary:
        if i == 1:
            new_binary.append(0)
        else:
            new_binary.append(1)
    print(len(new_binary))
    deci = 0
    for i in range(32):
        deci += new_binary[i]*pow(2,i)
    return deci


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        n = int(input())

        result = flippingBits(n)

        fptr.write(str(result) + '\n')

    fptr.close()
