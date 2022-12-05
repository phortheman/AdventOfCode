"""
Advent of Code 2022 Day 5

Date: 12/5/2022
Author: phortheman

"""
from pathlib import Path
from collections import defaultdict
import copy

CARGO = """[N]             [R]             [C]
[T] [J]         [S] [J]         [N]
[B] [Z]     [H] [M] [Z]         [D]
[S] [P]     [G] [L] [H] [Z]     [T]
[Q] [D]     [F] [D] [V] [L] [S] [M]
[H] [F] [V] [J] [C] [W] [P] [W] [L]
[G] [S] [H] [Z] [Z] [T] [F] [V] [H]
[R] [H] [Z] [M] [T] [M] [T] [Q] [W]
 1   2   3   4   5   6   7   8   9 """

def cleanUpData( item: str ) -> str:
    modifiedData = item.replace( "[", "" )
    modifiedData = modifiedData.replace( "]", "" )
    return modifiedData

def prepCargo( cargoInput: str ) -> dict:
    outputDict = defaultdict(list)
    stackPosition = 1
    stackCount = round( cargoInput.index("\n") / 4 )

    for i in range(0, len(cargoInput), 4):
        crate = cargoInput[i:i+3].strip()
        if stackPosition > stackCount:
            stackPosition = 1

        if crate == "":
            stackPosition += 1
        elif crate.isnumeric():
            break
        else: 
            outputDict[stackPosition].insert(0, cleanUpData(crate))
            stackPosition += 1
    return dict(outputDict)

# The result is the last element of each stack
def calculateResult( cargoStack: dict ) -> str:
    output = ""
    for i in range( 1, len(cargoStacks) + 1 ):
        output += cargoStack[i][-1]

    return output

# Part 1
def popCrates( numberOfPops: int, stackToPush: int, stackToPop: int, cargoStack: dict ):
    for i in range(numberOfPops):
            cargoStack[stackToPush].append( cargoStack[stackToPop].pop() )

# Part 2
def moveCrates( numberCrates: int, stackMoveTo: int, stackMoveFrom: int, cargoStack: dict ):
    cargoStack[stackMoveTo].extend( cargoStack[stackMoveFrom][-numberCrates:] )
    del cargoStack[stackMoveFrom][-numberCrates:]

with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
    cargoStacks = prepCargo(CARGO)
    part1CargoStack = copy.deepcopy(cargoStacks)
    part2CargoStack = copy.deepcopy(cargoStacks)
    # move {numOfCrates} from {popStack} to {pushStack}
    for instruction in file.readlines():
        arguments = instruction.split()

        popCrates( numberOfPops=int(arguments[1]),
                    stackToPop=int(arguments[3]),
                    stackToPush=int(arguments[5]),
                    cargoStack=part1CargoStack
        )

        moveCrates( numberCrates=int(arguments[1]),
                    stackMoveFrom=int(arguments[3]),
                    stackMoveTo=int(arguments[5]),
                    cargoStack=part2CargoStack )

    print( f"Part 1: {calculateResult( part1CargoStack )}" )
    print( f"Part 2: {calculateResult( part2CargoStack )}" )
