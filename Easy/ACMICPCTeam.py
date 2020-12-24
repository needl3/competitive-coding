#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the acmTeam function below.
def acmTeam(topic):
    max_sub_count = []
    #iterating over student 1
    for i in range(n-1):
        position_0 = [j for j in range(m) if topic[i][j] == '0']
        len_position_1 = m-len(position_0)
        #For iterating over students to find second companion
        for j in range(i+1,n):
            count = 0
            for k in position_0:
                if topic[j][k] == '1':
                    count += 1
            max_sub_count.append(count+len_position_1)
    max_topics = max(max_sub_count)
    total_ways = 0
    for i in max_sub_count:
        total_ways += 1 if i == max_topics else 0
    return [max_topics,total_ways]




if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nm = input().split()

    n = int(nm[0])

    m = int(nm[1])

    topic = []

    for _ in range(n):
        topic_item = input()
        topic.append(topic_item)

    result = acmTeam(topic)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
