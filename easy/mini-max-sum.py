# Zachary LeFevre
# For https://www.hackerrank.com/challenges/mini-max-sum/problem
import sys
from functools import reduce


def miniMaxSum(arr):
    sumArr = reduce((lambda x, y: x + y), arr)
    print(sumArr-max(arr), sumArr-min(arr))


if __name__ == "__main__":
    arr = list(map(int, input().strip().split(' ')))
    miniMaxSum(arr)
