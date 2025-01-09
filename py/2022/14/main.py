import argparse
import os
import sys
from collections import defaultdict
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


AIR = "."
ROCK = "#"
SAND = "o"
SPAWN = "+"


# Visualization guide
def displayMap(inputMap: defaultdict):
    minX = min(inputMap.keys())
    maxX = max(inputMap.keys())
    maxY = len(inputMap[minX])
    for j in range(maxY):
        for i in range(minX, maxX + 1):
            print(inputMap[i][j], end="")
        print()


# Fill the whole length of the known y axis max with air
def padYAxis(inputMap: defaultdict, xAxis: int, ySize: int):
    while len(inputMap[xAxis]) <= ySize:
        inputMap[xAxis].append(AIR)


# Draw a X axis line
def drawX(inputMap, startX, endX, yAxis):
    for i in range(min(startX, endX), max(startX, endX) + 1):
        inputMap[i][yAxis] = ROCK


# Draw a Y axis line
def drawY(inputMap, startY, endY, xAxis):
    for i in range(min(startY, endY), max(startY, endY) + 1):
        inputMap[xAxis][i] = ROCK


# Fills the Y Axis with air and then replaces the last element with the floor material
def fillToFloor(inputMap: defaultdict, xAxis: int, maxYAxis: int, floorMaterial=AIR):
    padYAxis(inputMap, xAxis, maxYAxis)
    inputMap[xAxis][maxYAxis] = floorMaterial


def canMoveDown(inputMap, newX, newY, maxY) -> bool:
    if maxY < newY:
        return False
    if inputMap[newX][newY] != AIR:
        return False
    return True


def canMoveLeft(inputMap, newX, newY, maxY, floorMaterial) -> bool:
    if maxY < newY:
        return False

    # Spill over scenario
    if min(inputMap.keys()) > newX:
        fillToFloor(inputMap, xAxis=newX, maxYAxis=maxY, floorMaterial=floorMaterial)

    if inputMap[newX][newY] != AIR:
        return False

    return True


def canMoveRight(inputMap, newX, newY, maxY, floorMaterial) -> bool:
    if maxY < newY:
        return False

    # Spill over scenario
    if max(inputMap.keys()) < newX:
        fillToFloor(inputMap, xAxis=newX, maxYAxis=maxY, floorMaterial=floorMaterial)

    if inputMap[newX][newY] != AIR:
        return False

    return True


def simulateSand(
    inputMap: defaultdict,
    maxY: int,
    floorMaterial=AIR,
    startX: int = 500,
    startY: int = 0,
) -> int:
    x, y = startX, startY
    inputMap[startX][startY] = SPAWN

    unitsOfSand = 0

    while True:
        if canMoveDown(inputMap, x, y + 1, maxY):
            y += 1

        elif canMoveLeft(inputMap, x - 1, y + 1, maxY, floorMaterial):
            x -= 1
            y += 1

        elif canMoveRight(inputMap, x + 1, y + 1, maxY, floorMaterial):
            x += 1
            y += 1

        # The sand cannot move
        elif y == maxY:
            break

        elif inputMap[x][y] == AIR:
            inputMap[x][y] = SAND
            unitsOfSand += 1
            x, y = startX, startY

        elif inputMap[x][y] == SPAWN:
            print("CAN'T SPAWN MORE!")
            unitsOfSand += 1
            break

    return unitsOfSand


def addTrace(inputMap: defaultdict, coord: tuple, prevCoord, maxY):
    x, y = coord

    # Initialize the x axis
    inputMap[x]

    # Get the upper and lower bounds of the x axis
    minX = min(inputMap.keys())
    maxX = max(inputMap.keys())

    # For every x position, ensure the y length is the same for every x axis
    for i in range(minX, maxX + 1):
        padYAxis(inputMap, i, y)
        # while len(inputMap[i]) <= y:
        #     inputMap[i].append(AIR)

    # First trace
    if not isinstance(prevCoord, tuple):
        inputMap[x][y] = ROCK
        return

    prevX, prevY = prevCoord

    if prevX != x:
        drawX(inputMap, startX=x, endX=prevX, yAxis=y)
    elif prevY != y:
        drawY(inputMap, startY=y, endY=prevY, xAxis=x)


def main():
    input_path = get_puzzle_input()

    rawMap = defaultdict(list)
    maxY = 0

    with open(input_path, "r") as f:
        for line in f.readlines():
            traces = line.strip().split(" -> ")
            previousCoord = None
            for coord in traces:
                pos = tuple(map(int, list(coord.split(","))))

                if maxY < pos[1]:
                    maxY = pos[1]

                addTrace(rawMap, pos, previousCoord, maxY)

                previousCoord = pos

    voidMap = deepcopy(rawMap)
    floorMap = deepcopy(rawMap)

    voidYAxisMax = maxY
    floorYAxisMax = maxY + 2

    for xAxis in floorMap.keys():
        fillToFloor(floorMap, xAxis, floorYAxisMax, ROCK)

    voidResult = simulateSand(voidMap, voidYAxisMax)
    floorResult = simulateSand(floorMap, floorYAxisMax, floorMaterial=ROCK)

    print(f"The units of sand that spawned is: {voidResult}")
    print(f"The units of sand that spawned with the floor is: {floorResult}")


if __name__ == "__main__":
    main()
