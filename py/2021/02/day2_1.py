# Advent of Code 2021 Day 2 Part One
# Author: phortheman
# Date: 12/17/2021

iInput = []
with open( "02/input_2.txt", 'r' ) as f:
    for line in f: 
        iInput.append( line )

horizontal = 0
depth = 0

for instruction in iInput:
    temp = instruction.split()
    direction = temp[0]
    value = int(temp[1])
    if( direction == "forward" ):
        horizontal += value
    elif( direction == "down" ):
        depth += value
    elif( direction == "up" ):
        depth -= value

print( horizontal * depth )