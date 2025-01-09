import argparse
import os
import sys
from enum import Enum


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


class Direction(Enum):
    N = [0, -1]
    S = [0, 1]
    E = [1, 0]
    W = [-1, 0]
    NE = [1, -1]
    NW = [-1, -1]
    SE = [1, 1]
    SW = [-1, 1]


DIRECTION_CHECK = [
    [Direction.N, Direction.NE, Direction.NW],
    [Direction.S, Direction.SE, Direction.SW],
    [Direction.W, Direction.NW, Direction.SW],
    [Direction.E, Direction.NE, Direction.SE],
]


def readInput(currentPositions: list):
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        currentY = 0
        for line in f.readlines():
            for x in range(len(line)):
                if line[x] == "#":
                    currentPositions.append((x, currentY))
            currentY += 1


def checkForNeighbor(
    currentPositions: list[tuple[int, int]], elf: tuple[int, int], round: int
) -> tuple[int, int]:
    hasNeighbor = False
    moveDirection = None
    for i in range(len(DIRECTION_CHECK)):
        canMove = False
        for dir in DIRECTION_CHECK[(i + round) % len(DIRECTION_CHECK)]:
            checkPos = (elf[0] + dir.value[0], elf[1] + dir.value[1])
            if checkPos in currentPositions:
                hasNeighbor = True
                canMove = False
                break
            else:
                canMove = True

        if canMove and moveDirection is None:
            moveDirection = DIRECTION_CHECK[(i + round) % len(DIRECTION_CHECK)][0]

        if moveDirection is not None and hasNeighbor:
            return (elf[0] + moveDirection.value[0], elf[1] + moveDirection.value[1])

    # Return None so the elf doesn't move
    return None


def simulateRound(currentPositions: list[tuple[int, int]], round: int):
    # Key: Next Position. Value: Current Position
    nextPositions = dict()

    # If the purposed next position was already rejected, skip it
    rejectedPositions = set()

    # For every position, run the check
    for elf in currentPositions:
        purposedPosition = checkForNeighbor(currentPositions, elf, round)
        if purposedPosition is None:
            continue
        elif purposedPosition in rejectedPositions:
            continue
        elif purposedPosition in nextPositions.keys():
            # Remove newPos from the dict and add to the rejected set
            del nextPositions[purposedPosition]
            rejectedPositions.add(purposedPosition)
        else:
            nextPositions[purposedPosition] = elf

    if len(nextPositions) == 0:
        return False

    # Change the current position for every elf that moved
    for nextPosition, currentPosition in nextPositions.items():
        index = currentPositions.index(currentPosition)
        currentPositions[index] = nextPosition

    return True


def createBoundary(
    currentPositions: list[tuple[int, int]],
) -> tuple[tuple[int, int], tuple[int, int]]:
    minX = currentPositions[0][0]
    minY = currentPositions[0][1]
    maxX = currentPositions[0][0]
    maxY = currentPositions[0][1]

    for x, y in currentPositions:
        if x < minX:
            minX = x
        if x > maxX:
            maxX = x

        if y < minY:
            minY = y
        if y > maxY:
            maxY = y

    return (minX, minY), (maxX, maxY)


def visualize(currentPositions: list[tuple[int, int]], bufferX=0, bufferY=0):
    start, end = createBoundary(currentPositions)
    output = ""

    for y in range(start[1] - bufferY, end[1] + 1 + bufferY):
        for x in range(start[0] - bufferX, end[0] + 1 + bufferY):
            if (x, y) in currentPositions:
                output += "#"
            else:
                output += "."
        output += "\n"

    print(output)


def emptyTiles(currentPositions: list[tuple[int, int]]) -> int:
    start, end = createBoundary(currentPositions)
    diffX = abs(start[0] - end[0]) + 1
    diffY = abs(start[1] - end[1]) + 1
    return (diffX * diffY) - len(currentPositions)


def main():
    currentPositions = []
    readInput(currentPositions)

    currentRound = 0
    while simulateRound(currentPositions, currentRound):
        if currentRound == 9:
            print(
                f"The number of empty ground titles after 10 rounds: {emptyTiles(currentPositions)}"
            )
        currentRound += 1

    visualize(currentPositions)
    print(f"The number of rounds until no more elves move is: {currentRound + 1}")


if __name__ == "__main__":
    main()
