# Zachary LeFevre
# For https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem
import sys
from functools import reduce


def minimumAbsoluteDifference(n, arr):
    sortedArr = sorted(arr)
    smallestDif = sys.maxsize
    dif = 0
    for i in range(len(sortedArr)):
        if(i == len(sortedArr)-1):
            return smallestDif
        else:
            dif = sortedArr[i+1]-sortedArr[i]
        if(dif < smallestDif):
            smallestDif = dif


if __name__ == "__main__":
    n = int(input().strip())
    arr = list(map(int, input().strip().split(' ')))
    result = minimumAbsoluteDifference(n, arr)
    print(result)
