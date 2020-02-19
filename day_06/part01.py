#!/usr/bin/env python3

# advent of code 2019: day 06, part 01

mem = [5,1,10,0,1,7,13,14,3,12,8,10,7,12,0,6]
s = set()
while True:
    s.add(tuple(mem))
    i = mem.index(max(mem))
    v = mem[i]
    mem[i] = 0
    while v > 0:
        i = (i + 1) % len(mem)
        mem[i] += 1
        v -= 1
    if tuple(mem) in s:
        break
print(len(s))    
