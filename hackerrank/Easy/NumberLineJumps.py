#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the kangaroo function below.
def kangaroo(x1, v1, x2, v2):
    x1_list = [x1]
    x2_list = [x2] 
    for i in range(1,10000):
        x1_list.append(x1_list[i-1]+v1)
        x2_list.append(x2_list[i-1]+v2)
    for i in range(1,10000):
        if x1_list[i] == x2_list[i]:
            return 'YES'
    return 'NO'
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    x1V1X2V2 = input().split()

    x1 = int(x1V1X2V2[0])

    v1 = int(x1V1X2V2[1])

    x2 = int(x1V1X2V2[2])

    v2 = int(x1V1X2V2[3])

    result = kangaroo(x1, v1, x2, v2)

    fptr.write(result + '\n')

    fptr.close()
