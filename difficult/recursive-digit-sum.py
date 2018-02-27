# Zachary LeFevre
# For https://www.hackerrank.com/challenges/recursive-digit-sum/problem

import sys
from functools import reduce


def digitSum(n, k):
    strArr = list(n)
    if len(strArr) == 1:
        return strArr[0]
    else:
        sum = reduce((lambda x, y: int(x)+int(y)), strArr) * k
        return digitSum(str(sum), 1)


if __name__ == "__main__":
    n, k = input().strip().split(' ')
    n, k = [str(n), int(k)]
    result = digitSum(n, k)
    print(result)
