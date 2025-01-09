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


def checkIfWithinRange(firstInput: tuple, secondInput: tuple) -> int:
    firstRange = range(firstInput[0], firstInput[1] + 1)
    secondRange = range(secondInput[0], secondInput[1] + 1)
    if secondInput[0] in firstRange and secondInput[1] in firstRange:
        return 1
    elif firstInput[0] in secondRange and firstInput[1] in secondRange:
        return 1
    return 0


def checkIfOverLap(firstInput: tuple, secondInput: tuple) -> int:
    firstRange = range(firstInput[0], firstInput[1] + 1)
    secondRange = range(secondInput[0], secondInput[1] + 1)
    if secondInput[0] in firstRange or secondInput[1] in firstRange:
        return 1
    elif firstInput[0] in secondRange or firstInput[1] in secondRange:
        return 1
    return 0


def main():
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        countWithinRange = 0
        countOverlap = 0

        for pair in f.readlines():
            first, second = pair.split(",")

            firstPair = tuple(map(int, first.split("-")))
            secondPair = tuple(map(int, second.split("-")))

            countWithinRange += checkIfWithinRange(firstPair, secondPair)
            countOverlap += checkIfOverLap(firstPair, secondPair)

    print(f"Part 1: {countWithinRange}")
    print(f"Part 2: {countOverlap}")


if __name__ == "__main__":
    main()
