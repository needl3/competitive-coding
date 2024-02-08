#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'gradingStudents' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts INTEGER_ARRAY grades as parameter.
#

def gradingStudents(grades):
    # Write your code here
    rounded = []
    for i in grades:
        multiple = i//10*10+5
        if i < 38 or i%5 == 0:
            rounded.append(i)
        elif i > multiple:
            multiple += 5
            if multiple-i < 3:
                rounded.append(multiple)
            else:
                rounded.append(i)
        elif i < multiple:
            if multiple-i < 3:
                rounded.append(multiple)
            else:
                rounded.append(i)
    return rounded
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    grades_count = int(input().strip())

    grades = []

    for _ in range(grades_count):
        grades_item = int(input().strip())
        grades.append(grades_item)

    result = gradingStudents(grades)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
