"""
Advent of Code 2022 Day 1

Date: 12/1/2022
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
        "-i",
        "--input",
        help="Specify a different puzzle input file path",
        default=default_input_path,
    )
    args = parser.parse_args()

    input_path = args.input

    if not os.path.exists(input_path):
        print(f"Error: Input file '{input_path}' does not exist.", file=sys.stderr)
        sys.exit(66)

    return input_path


def main():
    input_path = get_puzzle_input()

    elves = [0]

    with open(input_path, "r") as f:
        # Start with first elf
        elf = 0
        for calories in f.readlines():
            # If the line is a new line then start counting the calories for the next elf
            if calories == "\n":
                elf += 1
                elves.append(0)
            else:
                elves[elf] += int(calories)

    print(f"The most calories is: {max(elves)}")
    print(
        f"The sum of the top three calories is: {sum(sorted(elves, reverse=True)[:3])}"
    )


if __name__ == "__main__":
    main()
