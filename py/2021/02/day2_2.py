# Advent of Code 2021 Day 2 Part Two
# Author: phortheman
# Date: 12/17/2021

iInput = []
with open( "02/input_2.txt", 'r' ) as f:
    for line in f: 
        iInput.append( line )

horizontal = 0
depth = 0
aim = 0

for instruction in iInput:
    temp = instruction.split()
    direction = temp[0]
    value = int(temp[1])
    if( direction == "forward" ):
        horizontal += value
        depth += aim * value
    elif( direction == "down" ):
        aim += value
    elif( direction == "up" ):
        aim -= value

print( horizontal * depth )