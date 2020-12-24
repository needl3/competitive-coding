#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the countingValleys function below.
def countingValleys(n, s):
    prev = None
    count = 0
    array_count = []
    #First change the character to numbers
    for i in s:
        if i == 'U':
            array_count.append(1)
        else:
            array_count.append(-1)
    added = 0
    added_array = [0]
    total_valleys = 0
    for i in array_count:
        added += i
        added_array.append(added)
    print('Added array',added_array)
    prev = 0
    flag = None
    zero_to_subtract = 0
    for i in range(len(added_array)):
        print(added_array[i])
        if added_array[i] > 0:
            flag=False
            print('Flag set False')
        elif added_array[i] < 0:
            flag = True
            print("Flag set True")
        elif added_array[i] == 0:
            if not flag:
                print('Zero Found but flag is False. So, ignoring')
                pass
            else:
                print('Zero Found and flag is True. So, counting')
                total_valleys += 1
    return total_valleys
        
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    s = input()

    result = countingValleys(n, s)

    fptr.write(str(result) + '\n')

    fptr.close()
