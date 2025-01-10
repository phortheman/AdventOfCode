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

    iInput = []
    iOutput = []
    with open(input_path, "r") as f:
        for line in f:
            iInput.append(int(line))

    for index, value in enumerate(iInput):
        temp = iInput[index:index + 3]
        iOutput.append(sum(temp))

    part1 = 0
    part1Prev = iInput[0]

    part2 = 0
    part2Prev = iOutput[0]

    for i in iInput:
        if part1Prev < i:
            part1 += 1
        part1Prev = i

    for i in iOutput:
        if part2Prev < i:
            part2 += 1
        part2Prev = i

    print("Part 1: ", part1)
    print("Part 2: ", part2)


if __name__ == "__main__":
    main()
