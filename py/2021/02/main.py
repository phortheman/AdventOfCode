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

    iInput = []
    with open(input_path, "r") as f:
        for line in f:
            iInput.append(line)

    # Part 1
    horizontal = 0
    depth = 0

    for instruction in iInput:
        temp = instruction.split()
        direction = temp[0]
        value = int(temp[1])
        if direction == "forward":
            horizontal += value
        elif direction == "down":
            depth += value
        elif direction == "up":
            depth -= value

    print(horizontal * depth)

    # Part 2
    horizontal = 0
    depth = 0
    aim = 0

    for instruction in iInput:
        temp = instruction.split()
        direction = temp[0]
        value = int(temp[1])
        if direction == "forward":
            horizontal += value
            depth += aim * value
        elif direction == "down":
            aim += value
        elif direction == "up":
            aim -= value

    print(horizontal * depth)


if __name__ == "__main__":
    main()

