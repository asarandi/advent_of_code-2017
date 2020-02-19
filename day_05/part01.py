#!/usr/bin/env python3

# advent of code 2017: day 05, part 01

with open("input.txt") as fp:
    data = [int(x) for x in fp.read().splitlines()]
    fp.close()

i = steps = 0
while i < len(data):
    v = data[i]
    data[i] += 1
    i += v
    steps += 1

print(steps)
