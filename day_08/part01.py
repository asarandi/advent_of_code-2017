#!/usr/bin/env python3

with open("input.txt") as fp:
    data = fp.read().splitlines()
    fp.close()

def gte(a, b): return True if a >= b else False
def lte(a, b): return True if a <= b else False
def  eq(a, b): return True if a == b else False
def  ne(a, b): return True if a != b else False
def  gt(a, b): return True if a  > b else False
def  lt(a, b): return True if a  < b else False

func = [ gt, gte,  lt,  lte,  eq,   ne]
cond = [">", ">=", "<", "<=", "==", "!="]
registers = {}

res = 0
for line in data:
    r1, op, v1, _, r2, c, v2 = line.split()
    if r1 not in registers: registers[r1] = 0
    if r2 not in registers: registers[r2] = 0
    i = cond.index(c)
    if func[i](registers[r2], int(v2)):
        if op == "inc":
            registers[r1] += int(v1)
        else:            
            registers[r1] -= int(v1)
        res = registers[r1] if registers[r1] > res else res

print("part 1:", registers[max(registers, key=registers.get)])
print("part 2:", res)

