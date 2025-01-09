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


AT_MOST = 100_000
TOTAL_SPACE = 70_000_000
UPDATE_SIZE = 30_000_000
NODE_SIZE = {}


class Node:
    def __init__(self, name, parent) -> None:
        self.name = name
        self.parent = parent
        self.data = []
        self.children = []
        self.size = 0

    def __repr__(self) -> str:
        return self.name

    def addChild(self, node):
        self.children.append(node)

    def addData(self, data: int):
        self.data.append(data)

    def getChild(self, name: str):
        for child in self.children:
            if child.name == name:
                return child


def sumOfData(node: Node) -> int:
    size = sum(node.data)
    for child in node.children:
        size += sumOfData(child)
    node.size = size
    NODE_SIZE[node.name] = node.size
    return size


def sumOfAtMost() -> int:
    calcSumAtMost = 0
    for size in NODE_SIZE.values():
        if size < AT_MOST:
            calcSumAtMost += size
    return calcSumAtMost


def sizeOfDirectoryToDelete(targetSize: int) -> int:
    curSize = UPDATE_SIZE
    for size in NODE_SIZE.values():
        if targetSize < size < curSize:
            curSize = size
    return curSize


def main():
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        root = Node(f.readline().split()[2], None)
        workingNode = root
        for line in f.readlines():
            args = line.split()
            if args[1] == "cd":
                match args[2]:
                    case "..":
                        workingNode = workingNode.parent
                    case _:
                        workingNode = workingNode.getChild(args[2])
            elif args[0] == "dir":
                workingNode.addChild(Node(args[1], workingNode))
            elif args[0].isnumeric():
                workingNode.addData(int(args[0]))

    usedSpace = sumOfData(root)
    freeSpace = TOTAL_SPACE - usedSpace
    print(f"The total size: {usedSpace}")
    print(f"The total size at most 100,000: {sumOfAtMost()}")
    print(
        f"The size of the directory to delete: {sizeOfDirectoryToDelete(UPDATE_SIZE - freeSpace)}"
    )


if __name__ == "__main__":
    main()
