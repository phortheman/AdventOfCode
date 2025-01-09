import argparse
import os
import sys
from copy import deepcopy


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


def reverseOperator(operator: str):
    match operator:
        case "+":
            return "-"
        case "-":
            return "+"
        case "*":
            return "/"
        case "/":
            return "*"


def getMonkeyNumber(inputDict: dict[str, str], key: str) -> int:
    try:
        return int(inputDict[key])
    except ValueError:
        first, operator, second = inputDict[key].split()

        first = getMonkeyNumber(inputDict, first)
        second = getMonkeyNumber(inputDict, second)

        if operator == "/":
            operator = "//"

        inputDict[key] = eval(f"{first} {operator} {second}")
        return inputDict[key]


def readInput(inputDict: dict):
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        for line in f.readlines():
            inputDict[line[:4]] = line[6:].strip()


def main():
    monkeys = dict()
    readInput(monkeys)

    print(
        f"The number root will yell out is: {getMonkeyNumber(deepcopy(monkeys), 'root')}"
    )


if __name__ == "__main__":
    main()
