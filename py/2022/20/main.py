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


DECRYPTION_KEY = 811_589_153


class Node:
    def __init__(self, value: int) -> None:
        self.value = value
        self.nextNode = None
        self.prevNode = None

    def __repr__(self) -> str:
        return f"{self.value}"

    def __str__(self) -> str:
        return f"{self.value}"


def readInput(encryptedFile: list, decryption=1):
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        for line in f.readlines():
            encryptedFile.append(Node(int(line) * decryption))


# Created a doubly linked circle list and returns the zero node
def doublyLinkNodes(inputList: list) -> Node:
    startNode = inputList[0]
    endNode = inputList[-1]

    zeroNode = None

    startNode.prevNode = endNode
    endNode.nextNode = startNode

    prevNode = startNode

    for i in range(1, len(inputList) - 1):
        currentNode = inputList[i]

        if currentNode.value == 0:
            zeroNode = currentNode

        currentNode.prevNode = prevNode
        prevNode.nextNode = currentNode

        prevNode = currentNode

    endNode.prevNode = prevNode
    prevNode.nextNode = endNode

    return zeroNode


def getGroveCoordinates(zeroNode: Node) -> int:
    currentNode = zeroNode
    result = 0
    for i in range(3001):
        if i % 1000 == 0:
            result += currentNode.value
        currentNode = currentNode.nextNode

    return result


def mixValue(node: Node, interations: int) -> None:
    if node.value == 0:
        return

    # Remove link with current left and right nodes
    node.nextNode.prevNode = node.prevNode
    node.prevNode.nextNode = node.nextNode

    currentNode = node

    if node.value < 0:
        for _ in range(abs(interations)):
            currentNode = currentNode.prevNode
        leftCurrentNode = currentNode.prevNode
        currentNode.prevNode = node
        node.nextNode = currentNode
        node.prevNode = leftCurrentNode
        leftCurrentNode.nextNode = node

    else:
        for _ in range(abs(interations)):
            currentNode = currentNode.nextNode
        rightCurrentNode = currentNode.nextNode
        currentNode.nextNode = node
        node.nextNode = rightCurrentNode
        node.prevNode = currentNode
        rightCurrentNode.prevNode = node


# Doubly linked list with the tail and head connected
def part1():  # 18257
    encryptedFile = []
    readInput(encryptedFile)

    zeroNode = doublyLinkNodes(encryptedFile)

    for node in encryptedFile:
        mixValue(node, node.value)

    print(f"The sum of the grove coordinates are: {getGroveCoordinates(zeroNode)}")


# Horribly inefficient so need to work through a method that isn't a linked list
def part2():
    encryptedFile = []
    readInput(encryptedFile, DECRYPTION_KEY)

    print(encryptedFile)

    zeroNode = doublyLinkNodes(encryptedFile)

    for _ in range(9):
        for node in encryptedFile:
            iterations = node.value % (len(encryptedFile) - 1)
            mixValue(node, iterations)

    print(f"After applying the decryption key: {getGroveCoordinates(zeroNode)}")


def main():
    part1()
    # part2()


if __name__ == "__main__":
    main()
