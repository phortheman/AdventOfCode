"""
Advent of Code 2022 Day 4

Date: 12/4/2022
Author: phortheman

"""
from pathlib import Path

def checkIfWithinRange( firstInput: tuple, secondInput: tuple ) -> int:
    firstRange = range( firstInput[0], firstInput[1]+1 )
    secondRange = range( secondInput[0], secondInput[1]+1 )
    if secondInput[0] in firstRange and secondInput[1] in firstRange:
        return 1
    elif firstInput[0] in secondRange and firstInput[1] in secondRange:
        return 1
    return 0

def checkIfOverLap( firstInput: tuple, secondInput: tuple ) -> int:
    firstRange = range( firstInput[0], firstInput[1]+1 )
    secondRange = range( secondInput[0], secondInput[1]+1 )
    if secondInput[0] in firstRange or secondInput[1] in firstRange:
        return 1
    elif firstInput[0] in secondRange or firstInput[1] in secondRange:
        return 1
    return 0

with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
    countWithinRange = 0
    countOverlap = 0

    for pair in file.readlines():
        first, second = pair.split(",")

        firstPair = tuple(map(int, first.split("-") ) )
        secondPair = tuple(map(int, second.split("-") ) )

        countWithinRange += checkIfWithinRange( firstPair, secondPair )
        countOverlap += checkIfOverLap( firstPair, secondPair )

    print(f"Part 1: {countWithinRange}")
    print(f"Part 2: {countOverlap}")