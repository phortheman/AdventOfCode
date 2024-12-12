# Advent of Code 2021 Day 1 Part One
# Author: phortheman
# Date: 12/17/2021

iInput = []
with open( "01/input_1.txt", 'r' ) as f:
    for line in f: 
        iInput.append( int(line) )

iCountOfIncrease = 0
iPreviousValue = iInput[0]

for i in iInput:
    if( iPreviousValue < i ):
        iCountOfIncrease += 1
    iPreviousValue = i

print( iCountOfIncrease )