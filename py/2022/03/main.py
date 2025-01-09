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


def getPriority(char: str) -> int:
    # Lower Case
    if ord(char) > 96:
        return ord(char) - 96
    # Upper Case
    else:
        return ord(char) - 38


def getItemType(contents: str) -> str:
    splitIndex = int(len(contents) / 2)
    for item in contents[:splitIndex]:
        if contents[splitIndex:].find(item) != -1:
            return item


def getGroupPriority(group: list) -> int:
    for item in min(group):
        if (
            group[0].find(item) != -1
            and group[1].find(item) != -1
            and group[2].find(item) != -1
        ):
            return getPriority(item)


def main():
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        prioritySum = 0
        groupPrioritySum = 0
        workingGroup = []
        for rucksack in f.readlines():
            prioritySum += getPriority(getItemType(rucksack[:-1]))

            workingGroup.append(rucksack[:-1])

            if len(workingGroup) > 2:
                groupPrioritySum += getGroupPriority(workingGroup)
                workingGroup = []

    print(prioritySum)
    print(groupPrioritySum)


if __name__ == "__main__":
    main()
