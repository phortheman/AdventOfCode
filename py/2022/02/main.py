"""
Advent of Code 2022 Day 2

Date: 12/2/2022
Author: phortheman

"""
import argparse
import os
import sys


def get_puzzle_input():
    """
    Handles input file retrieval with the option to override the default path.
    """
    # Get the day and year directory this script is stored in
    day = os.path.basename(os.path.dirname(os.path.abspath(__file__)))
    year = os.path.basename(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

    # Build the default input path
    root_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), "../../.."))
    default_input_path = os.path.join(root_dir, "inputs", year, day, "input.txt")

    parser = argparse.ArgumentParser(description="Advent of Code Solution")
    parser.add_argument(
        "-i", "--input",
        help="Specify a different puzzle input file path",
        default=default_input_path
    )
    args = parser.parse_args()

    input_path = args.input

    if not os.path.exists(input_path):
        print(f"Error: Input file '{input_path}' does not exist.", file=sys.stderr)
        sys.exit(66)

    return input_path


def main():
    input_path = get_puzzle_input()
    # A,X = 1 point
    # B,Y = 2 points
    # C,Z = 3 points
    # Draw = 3 points
    # Win = 6 points

    WIN = 6
    DRAW = 3
    LOSE = 0

    with open(input_path, "r") as f:
        input = f.readlines()

    part1 = 0
    part2 = 0
    for round in input:
        opponent, player = round.split()
        # Part 1
        # A, X = Rock
        if opponent == 'A':
            match player:
                case 'X':
                    part1 += DRAW + 1
                case 'Y':
                    part1 += WIN + 2
                case 'Z':
                    part1 += LOSE + 3

        # B, Y = Paper
        elif opponent == 'B':
            match player:
                case 'X':
                    part1 += LOSE + 1
                case 'Y':
                    part1 += DRAW + 2
                case 'Z':
                    part1 += WIN + 3

        # C, Z = Scissors
        elif opponent == 'C':
            match player:
                case 'X':
                    part1 += WIN + 1
                case 'Y':
                    part1 += LOSE + 2
                case 'Z':
                    part1 += DRAW + 3

        # Part 2
        if opponent == 'A':
            match player:
                case 'X':
                    part2 += LOSE + 3
                case 'Y':
                    part2 += DRAW + 1
                case 'Z':
                    part2 += WIN + 2
        elif opponent == 'B':
            match player:
                case 'X':
                    part2 += LOSE + 1
                case 'Y':
                    part2 += DRAW + 2
                case 'Z':
                    part2 += WIN + 3
        elif opponent == 'C':
            match player:
                case 'X':
                    part2 += LOSE + 2
                case 'Y':
                    part2 += DRAW + 3
                case 'Z':
                    part2 += WIN + 1

    print(part1)
    print(part2)


if __name__ == "__main__":
    main()
