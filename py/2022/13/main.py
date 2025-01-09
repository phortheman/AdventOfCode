import argparse
import os
import sys
import ast


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


def evaluatePair(pair: tuple) -> int:
    leftList = pair[0]
    rightList = pair[1]

    iterations = min(len(leftList), len(rightList))

    for i in range(iterations):
        leftValue = leftList[i]
        rightValue = rightList[i]

        # Ensure each value is an integer
        if type(leftValue) is int and type(rightValue) is int:
            # Right order
            if leftValue < rightValue:
                return 1

            # Not right order
            elif rightValue < leftValue:
                return -1

            # Still need to evaluate
            else:
                continue

        # Both values are lists
        elif type(leftValue) is list and type(rightValue) is list:
            result = evaluatePair((leftValue, rightValue))
            if result != 0:
                return result

        # Left is a list
        elif type(leftValue) is list:
            result = evaluatePair((leftValue, [rightValue]))
            if result != 0:
                return result

        # Right is a list
        elif type(rightValue) is list:
            result = evaluatePair(([leftValue], rightValue))
            if result != 0:
                return result

    # Needs to continue
    if iterations == len(leftList) and iterations == len(rightList):
        return 0

    # Right order
    elif iterations == len(leftList):
        return 1

    # Not right order
    elif iterations == len(rightList):
        return -1

    return 0


def bubbleSort(value: list):
    n = len(value)
    swapped = False

    for i in range(n - 1):
        for j in range(0, n - i - 1):
            if evaluatePair((value[j], value[j + 1])) == -1:
                swapped = True
                value[j], value[j + 1] = value[j + 1], value[j]

        if not swapped:
            return


def main():
    input_path = get_puzzle_input()

    message = []
    correctIndexSum = 0
    leftSide = None
    rightSide = None
    currentIndex = 1

    with open(input_path, "r") as f:
        for line in f.readlines():
            # New pair
            if line == "\n":
                leftSide = None
                rightSide = None
                currentIndex += 1
                continue

            if leftSide is None:
                leftSide = ast.literal_eval(line)
            elif rightSide is None:
                rightSide = ast.literal_eval(line)

                if evaluatePair((leftSide, rightSide)) == 1:
                    message.append(leftSide)
                    message.append(rightSide)
                    correctIndexSum += currentIndex
                else:
                    message.append(rightSide)
                    message.append(leftSide)

    print(f"The sum of the indices of the pairs is: {correctIndexSum}")

    # Add divider packets
    message.append([[2]])
    message.append([[6]])

    bubbleSort(message)

    divider1Index, divider2Index = 0, 0
    for i in range(len(message)):
        if message[i] == [[2]]:
            divider1Index = i + 1
        if message[i] == [[6]]:
            divider2Index = i + 1

        if divider1Index != 0 and divider2Index != 0:
            break

    print(f"The decoder key is: {divider1Index * divider2Index}")


if __name__ == "__main__":
    main()
