# Advent of Code 2021 Day 1 Part Two
# Author: phortheman
# Date: 12/17/2021

iInput = []
iOutput = []
with open( "01/input_1.txt", 'r' ) as f:
    for line in f: 
        iInput.append( int(line) )

for index, value in enumerate(iInput):
    temp = iInput[index:index+3]
    iOutput.append( sum(temp) )

        

iCountOfIncrease = 0
iPreviousValue = iOutput[0]

for i in iOutput:
    if( iPreviousValue < i ):
        iCountOfIncrease += 1
    iPreviousValue = i

print( iCountOfIncrease )