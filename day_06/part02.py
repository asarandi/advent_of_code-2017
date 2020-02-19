#!/usr/bin/env python3

# advent of code 2019: day 06, part 01

s = set()
mem = [5,1,10,0,1,7,13,14,3,12,8,10,7,12,0,6]
flag = False
res = 0
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
        if flag == False:
            flag = True
            dupe = tuple(mem)
        else:
            if tuple(mem) == dupe:
                break
    if flag:
        res += 1
print(res)    
