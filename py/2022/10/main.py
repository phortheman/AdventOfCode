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


LIT = "#"
DARK = "."

SCREEN = []


def addSignalStrength(cycle: int, register: int, signalDict: dict) -> int:
    # First 20
    if cycle == 20:
        signalDict[cycle] = cycle * register

    # 40 after the initial 20
    elif cycle == 40 * len(signalDict) + 20:
        signalDict[cycle] = cycle * register


def tick(cycle: int, position: int):
    # Force curPixel to be 0-39 based off of the current cycle
    curPixel = cycle % 40

    # If the curPixel is between the 3 pixel wide sprite
    if position - 1 <= curPixel <= position + 1:
        SCREEN.append(LIT)
    else:
        SCREEN.append(DARK)

    # Add the tick
    return cycle + 1


def printScreen():
    for i in range(len(SCREEN)):
        if i % 40 == 0:
            print()
        print(SCREEN[i], end="")


def main():
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        x = 1
        cycle = 0

        signalStrength = {}

        for instruction in f.readlines():
            command, *value = instruction.split()

            match command:
                case "noop":
                    cycle = tick(cycle, x)
                    addSignalStrength(cycle, x, signalStrength)
                case "addx":
                    cycle = tick(cycle, x)
                    addSignalStrength(cycle, x, signalStrength)
                    cycle = tick(cycle, x)
                    addSignalStrength(cycle, x, signalStrength)
                    x += int(value[0])

    # BUG: Only part 2 solution is printed?
    print(f"The Sum of the signal strength is: {sum(signalStrength.values())}")


if __name__ == "__main__":
    main()
