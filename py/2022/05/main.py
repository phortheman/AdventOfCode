import argparse
import os
import sys
from collections import defaultdict
import copy


CARGO = """[N]             [R]             [C]
[T] [J]         [S] [J]         [N]
[B] [Z]     [H] [M] [Z]         [D]
[S] [P]     [G] [L] [H] [Z]     [T]
[Q] [D]     [F] [D] [V] [L] [S] [M]
[H] [F] [V] [J] [C] [W] [P] [W] [L]
[G] [S] [H] [Z] [Z] [T] [F] [V] [H]
[R] [H] [Z] [M] [T] [M] [T] [Q] [W]
 1   2   3   4   5   6   7   8   9 """


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


def cleanUpData(item: str) -> str:
    modifiedData = item.replace("[", "")
    modifiedData = modifiedData.replace("]", "")
    return modifiedData


def prepCargo(cargoInput: str) -> dict:
    outputDict = defaultdict(list)
    stackPosition = 1
    stackCount = round(cargoInput.index("\n") / 4)

    for i in range(0, len(cargoInput), 4):
        crate = cargoInput[i:i + 3].strip()
        if stackPosition > stackCount:
            stackPosition = 1

        if crate == "":
            stackPosition += 1
        elif crate.isnumeric():
            break
        else:
            outputDict[stackPosition].insert(0, cleanUpData(crate))
            stackPosition += 1
    return dict(outputDict)


# The result is the last element of each stack
def calculateResult(cargoStack: dict) -> str:
    output = ""
    for i in range(1, len(cargoStack) + 1):
        output += cargoStack[i][-1]

    return output


# Part 1
def popCrates(numberOfPops: int, stackToPush: int, stackToPop: int, cargoStack: dict):
    for i in range(numberOfPops):
        cargoStack[stackToPush].append(cargoStack[stackToPop].pop())


# Part 2
def moveCrates(
    numberCrates: int, stackMoveTo: int, stackMoveFrom: int, cargoStack: dict
):
    cargoStack[stackMoveTo].extend(cargoStack[stackMoveFrom][-numberCrates:])
    del cargoStack[stackMoveFrom][-numberCrates:]


def main():
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        cargoStacks = prepCargo(CARGO)
        part1CargoStack = copy.deepcopy(cargoStacks)
        part2CargoStack = copy.deepcopy(cargoStacks)
        # move {numOfCrates} from {popStack} to {pushStack}
        for instruction in f.readlines():
            arguments = instruction.split()

            popCrates(
                numberOfPops=int(arguments[1]),
                stackToPop=int(arguments[3]),
                stackToPush=int(arguments[5]),
                cargoStack=part1CargoStack,
            )

            moveCrates(
                numberCrates=int(arguments[1]),
                stackMoveFrom=int(arguments[3]),
                stackMoveTo=int(arguments[5]),
                cargoStack=part2CargoStack,
            )

    print(f"Part 1: {calculateResult(part1CargoStack)}")
    print(f"Part 2: {calculateResult(part2CargoStack)}")


if __name__ == "__main__":
    main()
