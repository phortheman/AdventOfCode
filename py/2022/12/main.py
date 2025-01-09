import argparse
import os
import sys
from collections import deque


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


def breadthFirstSearch(start: tuple, map: list):
    queue = deque([start])
    visited = {start}

    moveCount = 0
    nodesLeftInLayer = 1
    nodesInNextLayer = 0
    reachedEnd = False

    # Cardinal directions
    # 0 = Up
    # 1 = Down
    # 2 = Right
    # 3 = Left
    directionRow = [-1, 1, 0, 0]
    directionCol = [0, 0, 1, -1]

    while queue:
        currentRow, currentCol = queue.popleft()
        currentValue = map[currentRow][currentCol]

        # No skipping :)
        if currentValue == "S":
            currentValue = "a"

        # Check for end
        if currentValue == "E":
            reachedEnd = True
            break

        for i in range(len(directionRow)):
            # Get neighbor coodinates
            neighborRow = currentRow + directionRow[i]
            neighborCol = currentCol + directionCol[i]

            # Check if neighbor is valid
            if not 0 <= neighborRow < len(map):
                continue
            if not 0 <= neighborCol < len(map[0]):
                continue

            # Make the tuple for the position of the neighbor
            neighborPos = (neighborRow, neighborCol)

            # Check if neighbor was visited already
            if neighborPos in visited:
                continue

            # Get the value of the neighbor
            neighborValue = map[neighborRow][neighborCol]

            # No skipping :)
            if neighborValue == "E":
                neighborValue = "z"

            # See if the height is too steep
            if ord(neighborValue) > ord(currentValue) + 1:
                continue

            visited.add(neighborPos)
            queue.append(neighborPos)

            nodesInNextLayer += 1

        # Finished processing this layer
        nodesLeftInLayer -= 1

        # Move on to the next layer
        if nodesLeftInLayer == 0:
            nodesLeftInLayer = nodesInNextLayer
            nodesInNextLayer = 0
            moveCount += 1

    if reachedEnd:
        return moveCount

    # Can't be reached
    return -1


def main():
    input_path = get_puzzle_input()

    heightMap = []
    startPosition = None
    lowestPoints = []

    with open(input_path, "r") as f:
        row = 0
        for line in f.readlines():
            col = []
            for curCol in range(len(line)):
                if line[curCol] == "\n":
                    break

                if line[curCol] == "S":
                    startPosition = (row, curCol)

                elif line[curCol] == "a":
                    lowestPoints.append((row, curCol))

                col.append(line[curCol])

            heightMap.append(col)
            row += 1

    fewestSteps = breadthFirstSearch(startPosition, heightMap)
    fewestStepsAnywhere = fewestSteps

    for lowPoint in lowestPoints:
        result = breadthFirstSearch(lowPoint, heightMap)

        # The search will return -1 if the point can't reach the end
        if result > 0 and result < fewestStepsAnywhere:
            fewestStepsAnywhere = result

    print(f"The fewest steps requried is: {fewestSteps}")
    print(f"The fewest steps from any square is: {fewestStepsAnywhere}")


if __name__ == "__main__":
    main()
