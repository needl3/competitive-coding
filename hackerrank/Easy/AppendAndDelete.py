#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the appendAndDelete function below.
def appendAndDelete(s, t, k):
    index_of_difference = None
    upper = len(s) if len(t) > len(s) else len(t)
    for i in range(upper):
        if s[i] != t[i]:
            index_of_difference = i
            break
    if len(t) > len(s):    
        #Requires append
        if index_of_difference != None:
            return 'Yes' if k>=(len(t)-len(s))+((len(s)-1)-index_of_difference) else 'No'
        else:
            if k>=len(t)-len(s):
                if len(t)-len(s) == 1 or len(t)-len(s) == 3:
                    return 'No'
                else:
                    return 'Yes'
            else:
                return 'No'
    elif len(t) == len(s):
        if index_of_difference == None:
            return 'Yes'
        else:
            return 'Yes' if k >= (len(s)-index_of_difference)*2 else 'No'
    elif len(t)<len(s):
        #Requires deletion
        if index_of_difference == None:    
            return 'Yes' if k >= len(s)-len(t) else 'No'
        else:
            return 'Yes' if k >= (len(s)+len(t)-2*index_of_difference) else 'No'

        

            



if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    t = input()

    k = int(input())

    result = appendAndDelete(s, t, k)

    fptr.write(result + '\n')

    fptr.close()
