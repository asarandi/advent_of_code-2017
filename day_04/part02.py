#!/usr/bin/env python3

# advent of code 2017: day 04, part 02

def is_valid(s):
    for w1 in s:
        for w2 in s:
            if w1 == w2:
                continue
            if sorted(list(w1)) == sorted(list(w2)):
                return False
    return True            

with open("input.txt") as fp:
    data = fp.read().splitlines()
    fp.close

res = 0
for line in data:
    s = line.split()
    if len(s) != len(set(s)):
        continue
    if is_valid(s):
        res += 1

print(res)
