"""
Advent of Code 2022 Day 2

Date: 12/2/2022
Author: phortheman

"""
from pathlib import Path

# A,X = 1 point
# B,Y = 2 points
# C,Z = 3 points
# Draw = 3 points
# Win = 6 points

WIN = 6
DRAW = 3
LOSE = 0

PART_1 = False

with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:

    score = 0
    for round in file.readlines():
        opponent, player = round.split()
        if PART_1:
            # A, X = Rock
            if opponent == 'A':
                match player:
                    case 'X':
                        score += DRAW + 1
                    case 'Y':
                        score += WIN + 2
                    case 'Z':
                        score += LOSE + 3

            # B, Y = Paper
            elif opponent == 'B':
                match player:
                    case 'X':
                        score += LOSE + 1
                    case 'Y':
                        score += DRAW + 2
                    case 'Z':
                        score += WIN + 3

            # C, Z = Scissors
            elif opponent == 'C':
                match player:
                    case 'X':
                        score += WIN + 1
                    case 'Y':
                        score += LOSE + 2
                    case 'Z':
                        score += DRAW + 3
        else: # Part 2 scoring: 12382
            if opponent == 'A':
                match player:
                    case 'X':
                        score += LOSE + 3
                    case 'Y':
                        score += DRAW + 1
                    case 'Z':
                        score += WIN + 2
            elif opponent == 'B':
                match player:
                    case 'X':
                        score += LOSE + 1
                    case 'Y':
                        score += DRAW + 2
                    case 'Z':
                        score += WIN + 3
            elif opponent == 'C':
                match player:
                    case 'X':
                        score += LOSE + 2
                    case 'Y':
                        score += DRAW + 3
                    case 'Z':
                        score += WIN + 1

    print( score )
