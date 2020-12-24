#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the workbook function below.
def workbook(n, k, arr):
    page_no = 0
    count = 0
    for i in range(n):
        #Find the no of pages we can go on from here
        temp_page = arr[i]//k+1 if arr[i]%k != 0 else arr[i]//k
        no_of_pages = [page_no+1, page_no+temp_page] if page_no+k < arr[i] else [page_no+1, arr[i]]
        current_question_range = [1,k]
        print('For i: ',i)
        print('Total page for questions: ',temp_page)
        print('Page from to: ',no_of_pages)
        for j in range(no_of_pages[0],no_of_pages[1]+1):
            #Find upper and lower limit of questions
            print("\nCurrent qsn range: ", current_question_range)
            if j >= current_question_range[0] and j <= current_question_range[1]:
                count += 1
            current_question_range = [current_question_range[1]+1, current_question_range[1]+k] if current_question_range[1]+k <= arr[i] else [current_question_range[1]+1,arr[i]]
        page_no += temp_page
        print("\nPage no increased to: ",page_no)
    return count

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nk = input().split()

    n = int(nk[0])

    k = int(nk[1])

    arr = list(map(int, input().rstrip().split()))

    result = workbook(n, k, arr)

    fptr.write(str(result) + '\n')

    fptr.close()
