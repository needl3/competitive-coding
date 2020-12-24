#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the breakingRecords function below.
def breakingRecords(scores):
    max = scores[0]
    min = scores[0]
    score = [0,0]
    for i in scores[1:]:
        if i > max:
            max = i
            score[0] += 1
        elif i < min:
            min = i
            score[1] += 1
        else:
            pass
    return score


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    scores = list(map(int, input().rstrip().split()))

    result = breakingRecords(scores)

    fptr.write(' '.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
