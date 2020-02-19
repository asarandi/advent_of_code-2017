#!/usr/bin/env python3

# advent of code 2017: day 03, part 02

INPUT = 325489

size = 21
g = [[0 for i in range(size)] for j in range(size)]
g[size//2][size//2] = 1
sq=3
i=size//2+1
j=size//2+1
moves = [(-1,0),(0,-1),(1,0),(0,1)]
done = False
while not done:
    for side in range(4):
        for _ in range(sq-1):
            x,y = moves[side]
            i += x
            j += y            
            n = 0
            for x in range(-1,2):
                for y in range(-1,2):
                    n += g[i+x][j+y]
            g[i][j] = n
            if n >= INPUT and done == False:
                print("result", n)
                done = True
    sq += 2
    i += 1
    j += 1
