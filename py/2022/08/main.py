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


GRID = []


def isVisible(row: int, col: int) -> bool:
    tree = GRID[row][col]
    top = True
    bottom = True
    left = True
    right = True

    for leftIterator in range(col):
        if GRID[row][leftIterator] >= tree:
            left = False
            break

    for rightIterator in range(col + 1, len(GRID[row])):
        if GRID[row][rightIterator] >= tree:
            right = False
            break

    for topIterator in range(row):
        if GRID[topIterator][col] >= tree:
            top = False
            break

    for bottomIterator in range(row + 1, len(GRID)):
        if GRID[bottomIterator][col] >= tree:
            bottom = False
            break

    return top or bottom or left or right


def getScenicScore(row: int, col: int) -> int:
    tree = GRID[row][col]

    left = 0
    right = 0
    top = 0
    bottom = 0

    for leftIterator in reversed(range(col)):
        left += 1
        if GRID[row][leftIterator] >= tree:
            break

    for rightIterator in range(col + 1, len(GRID[row])):
        right += 1
        if GRID[row][rightIterator] >= tree:
            break

    for topIterator in reversed(range(row)):
        top += 1
        if GRID[topIterator][col] >= tree:
            break

    for bottomIterator in range(row + 1, len(GRID)):
        bottom += 1
        if GRID[bottomIterator][col] >= tree:
            break

    return left * right * top * bottom


def main():
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        for row in f.readlines():
            GRID.append(list(map(int, row.strip())))

        maxRow = len(GRID) - 1
        maxCol = len(GRID[0]) - 1

        visibleCount = (len(GRID) * 2) + (len(GRID[0]) * 2 - 4)

        scenicScore = 0

        for i in range(len(GRID)):
            if i == 0 or i == maxRow:
                continue
            for j in range(len(GRID[i])):
                if j == 0 or j == maxCol:
                    continue

                visibleCount += isVisible(i, j)
                treeScore = getScenicScore(i, j)
                if treeScore > scenicScore:
                    scenicScore = treeScore

    print(f"The number of tree visible is: {visibleCount}")
    print(f"The best scenic score is: {scenicScore}")


if __name__ == "__main__":
    main()
