#!/bin/python3

import os
import sys

#
# Complete the timeConversion function below.
#
def timeConversion(s):
    #
    # Write your code here.
    #
    s = s.split(':')
    hr = ''
    if s[2][2:] == 'AM':
        if s[0] == '12':
            hr = '00'
        else:
            hr = s[0]
    elif s[2][2:] == 'PM':
        if s[0] == '12':
            hr = '12'
        else:
            hr = str(12+int(s[0]))
    return hr+':'+s[1]+':'+''.join(s[2][:2])

if __name__ == '__main__':
    f = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)

    f.write(result + '\n')

    f.close()
