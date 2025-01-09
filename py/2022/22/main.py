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
    RIGHT = [1, 0]
    DOWN = [0, 1]
    LEFT = [-1, 0]
    UP = [0, -1]


class Point:
    def __init__(self, x, y, direction: Direction) -> None:
        self.x = x
        self.y = y
        self.facing: Direction = direction

    def __str__(self) -> str:
        return f"({self.x}, {self.y}): {self.facing.name}"

    def turn(self, clockwise: bool):
        match self.facing:
            case Direction.RIGHT:
                if clockwise:
                    self.facing = Direction.DOWN
                else:
                    self.facing = Direction.UP
            case Direction.DOWN:
                if clockwise:
                    self.facing = Direction.LEFT
                else:
                    self.facing = Direction.RIGHT
            case Direction.LEFT:
                if clockwise:
                    self.facing = Direction.UP
                else:
                    self.facing = Direction.DOWN
            case Direction.UP:
                if clockwise:
                    self.facing = Direction.RIGHT
                else:
                    self.facing = Direction.LEFT

    def simMove(self) -> tuple[int, int]:
        return self.x + self.facing.value[0], self.y + self.facing.value[1]

    def move(self, x, y):
        self.x = x
        self.y = y

    def calcPassword(self) -> int:
        return (
            (1000 * (self.y + 1))
            + (4 * (self.x + 1))
            + list(Direction).index(self.facing)
        )


class MapRow:
    # Lower bound is the lowest col value while upper bound is the highest
    # Inclusive
    def __init__(self) -> None:
        self.lowerBound: int = None
        self.upperBound: int = None
        self.values = []

    def __repr__(self) -> str:
        return self.values

    def __str__(self) -> str:
        returnedString = ""
        for _ in range(self.lowerBound):
            returnedString += " "
        for i in self.values:
            returnedString += i
        return returnedString

    def __getitem__(self, key):
        return self.values[key - self.lowerBound]

    def __setitem__(self, key, value):
        self.values[key - self.lowerBound] = value

    def inRange(self, x: int) -> bool:
        return self.lowerBound <= x <= self.upperBound

    def isWall(self, x: int) -> bool:
        return self.values[x - self.lowerBound] == "#"


def readInputPart1(map: dict, instructions: list):
    bParseMap = True
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        currentRowNumber = 0
        for line in f.readlines():
            if line == "\n":
                bParseMap = False
                continue
            if bParseMap:
                row = MapRow()
                bParsingData = False
                for i in range(len(line)):
                    if line[i] == " ":
                        if bParsingData:
                            row.upperBound = i - 1

                        bParsingData = False
                    elif line[i] == "\n":
                        if bParsingData:
                            row.upperBound = i - 1

                        map[currentRowNumber] = row
                        currentRowNumber += 1
                    else:
                        if row.lowerBound is None:
                            row.lowerBound = i
                        bParsingData = True
                        row.values.append(line[i])
            else:
                holder = ""
                for element in line:
                    if element.isnumeric():
                        holder += element
                    else:
                        instructions.append(int(holder))
                        holder = ""
                        instructions.append(element)

                if holder != "":
                    instructions.append(int(holder))


def findNextAvailablePosition(
    map: dict[int, MapRow], y: int, x: int, direction: Direction
) -> tuple[int, int]:
    if direction == Direction.UP:
        while y + 1 in map.keys() and map[y + 1].inRange(x):
            y += 1
    elif direction == Direction.DOWN:
        while y - 1 in map.keys() and map[y - 1].inRange(x):
            y -= 1
    elif direction == Direction.LEFT:
        x = map[y].upperBound
    elif direction == Direction.RIGHT:
        x = map[y].lowerBound

    return x, y


def part1():
    map = dict()
    instructions = []
    readInputPart1(map, instructions)

    point = Point(map[0].lowerBound, 0, Direction.RIGHT)

    for step in instructions:
        if isinstance(step, int):
            for _ in range(step):
                nextX, nextY = point.simMove()
                if nextY not in map.keys() or not map[nextY].inRange(nextX):
                    nextX, nextY = findNextAvailablePosition(
                        map, nextY, nextX, point.facing
                    )
                if not map[nextY].isWall(nextX):
                    point.move(nextX, nextY)
                else:
                    break
        else:
            point.turn(step == "R")

    print(point.calcPassword())


if __name__ == "__main__":
    part1()
