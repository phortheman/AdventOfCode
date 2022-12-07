"""
Advent of Code 2022 Day 7

Date: 12/7/2022
Author: phortheman

"""
from pathlib import Path

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
        self.children.append( node )
    
    def addData(self, data: int):
        self.data.append( data )

    def getChild(self, name: str):
        for child in self.children:
            if child.name == name:
                return child

def sumOfData( node: Node ) -> int:
    size = sum( node.data )
    for child in node.children:
        size += sumOfData( child )
    node.size = size
    NODE_SIZE[node.name] = node.size
    return size

def sumOfAtMost() -> int:
    calcSumAtMost = 0
    for size in NODE_SIZE.values():
        if size < AT_MOST:
            calcSumAtMost += size
    return calcSumAtMost

def sizeOfDirectoryToDelete( targetSize: int ) -> int:
    curSize = UPDATE_SIZE
    for size in NODE_SIZE.values():
        if targetSize < size < curSize:
            curSize = size
    return curSize

with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:

    root = Node(file.readline().split()[2], None)
    workingNode = root
    for line in file.readlines():
        args = line.split()
        if args[1] == "cd":
            match args[2]:
                case "..":
                    workingNode = workingNode.parent
                case _:
                    workingNode = workingNode.getChild(args[2])
        elif args[0] == "dir":
            workingNode.addChild( Node(args[1], workingNode ) )
        elif args[0].isnumeric():
            workingNode.addData( int( args[0] ) )

    usedSpace = sumOfData(root)
    freeSpace = TOTAL_SPACE - usedSpace
    print( f"The total size: {usedSpace}"  )
    print( f"The total size at most 100,000: {sumOfAtMost()}")
    print( f"The size of the directory to delete: {sizeOfDirectoryToDelete( UPDATE_SIZE - freeSpace )}")