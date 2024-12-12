"""
Advent of Code 2022 Day 1

Date: 12/1/2022
Author: phortheman

"""

from pathlib import Path

elves = [0]

with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
    
    # Start with first elf
    elf = 0
    for calories in file.readlines():
        # If the line is a new line then start counting the calories for the next elf
        if calories == '\n':
            elf += 1
            elves.append(0)
        else:
            elves[elf] += int(calories)

    print( f"The most calories is: {max(elves)}")
    print( f"The sum of the top three calories is: {sum( sorted( elves, reverse=True )[:3] )}" )
