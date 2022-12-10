"""
Advent of Code 2022 Day 10

Date: 12/10/2022
Author: phortheman

"""
from pathlib import Path

LIT = "#"
DARK = "."

SCREEN = []

def addSignalStrength( cycle: int, register: int, signalDict: dict ) -> int:
    # First 20
    if cycle == 20:
        signalDict[cycle] = (cycle * register)

    # 40 after the initial 20
    elif cycle == 40 * len(signalDict) + 20:
        signalDict[cycle] = (cycle * register)

def tick( cycle: int, position: int ):

    # Force curPixel to be 0-39 based off of the current cycle
    curPixel = cycle % 40

    # If the curPixel is between the 3 pixel wide sprite
    if position - 1 <= curPixel <= position + 1:
        SCREEN.append( LIT )
    else:
        SCREEN.append( DARK )

    # Add the tick
    return cycle + 1

def printScreen():
    for i in range( len( SCREEN ) ):
        if i % 40 == 0:
            print()
        print(SCREEN[i], end="")


with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
    x = 1
    cycle = 0

    signalStrength = {}

    for instruction in file.readlines():
        command, *value = instruction.split()

        match command:
            case "noop":
                cycle = tick( cycle, x )
                addSignalStrength( cycle, x, signalStrength)
            case "addx":
                cycle = tick( cycle, x )
                addSignalStrength( cycle, x, signalStrength)
                cycle = tick( cycle, x )
                addSignalStrength( cycle, x, signalStrength)
                x += int(value[0])
    
    print( f"The Sum of the signal strength is: {sum(signalStrength.values())}" )
    printScreen()
