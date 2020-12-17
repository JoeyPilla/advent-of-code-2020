import collections
import math
import re
import sys

lines = [1001938,
"41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,431,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,17,x,19,x,x,x,x,x,x,x,x,x,x,x,863,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29"]

def crt(pairs):
    M = 1
    for x, mx in pairs:
        print(x, mx)
        M *= mx
    total = 0
    for x, mx in pairs:
        b = M // mx
        total += x * b * pow(b, mx-2, mx)
        total %= M
    return total


start = int(lines[0])
pairs = []
for i, n in enumerate(lines[1].split(',')):
    if n == 'x':
        continue
    n = int(n)
    pairs.append((n - i, n))
print(crt(pairs))
