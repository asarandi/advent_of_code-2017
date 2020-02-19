#!/usr/bin/env python3

import sys

sys.setrecursionlimit(100000)

with open("input.txt") as fp:
    data = fp.read().strip()
    fp.close()

res = 0
def rec(s, g=False, k=0):
    global res
    if len(s) == 0:
        return res
    c = s[0]    
    if g == False:
        if c == '{': return rec(s[1:], False, k+1)
        if c == '}': return rec(s[1:], False, k-1)
        if c == ',': return rec(s[1:], False, k)
        if c == '<': return rec(s[1:], True, k)
    else:
        if c == '!': return rec(s[2:], True, k)
        if c == '>': return rec(s[1:], False, k)
        res += 1; return rec(s[1:], True, k)
    print(f"error [{s}]")
    return res

print(rec(data))

