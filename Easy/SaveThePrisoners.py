#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the saveThePrisoner function below.
def saveThePrisoner(prisoners, sweets, chair):
    rem_sweets = sweets%prisoners
    chair_reached = chair+rem_sweets-1
    if chair_reached%prisoners == 0:
        return prisoners
    elif chair_reached%prisoners > 0:
        return chair_reached%prisoners
    else:
        return chair_reached


    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input())

    for t_itr in range(t):
        nms = input().split()

        n = int(nms[0])

        m = int(nms[1])

        s = int(nms[2])

        result = saveThePrisoner(n, m, s)

        fptr.write(str(result) + '\n')

    fptr.close()
