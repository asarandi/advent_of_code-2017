#!/usr/bin/env python3

# advent of code 2017: day 04, part 01

with open("input.txt") as fp:
    data = fp.read().splitlines()
    fp.close

res = 0
for line in data:
    s = line.split()
    res = res + 1 if len(s) == len(set(s)) else res

print(res)
