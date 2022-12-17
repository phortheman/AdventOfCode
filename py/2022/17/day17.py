"""
Advent of Code 2022 Day 17

Date: 12/17/2022
Author: phortheman

"""
from pathlib import Path
from copy import deepcopy

SHAPES = {
    0: [
        ["@", "@", "@", "@"]
    ],
    1: [
        [None, "@", None],
        ["@",  "@", "@"],
        [None, "@", None]
    ],
    2: [
        [None, None, "@"],
        [None, None, "@"],
        [ "@",  "@", "@"]
    ],
    3: [
        ["@"],
        ["@"],
        ["@"],
        ["@"]
    ],
    4: [ 
        ["@", "@"],
        ["@", "@"]
    ]
}

class Board():
    def __init__(self, rowLength) -> None:
        self.grid = []
        self.highestPoint = 0
        self.colLength = rowLength

    # Add rows to the "top"
    def addRow(self, num = 1):
        for _ in range(num):
            self.grid.append( [None] * self.colLength )

    # Returns true if the shape can fit in the provided position
    # Row, Col is the top left corner of the shape
    def scanShape( self, shape, row, col ):

        # Creating copy of this object for visualization
        testBoard = deepcopy( self )

        # Out of bounds checks
        if not 0 <= col < testBoard.colLength: # Out of bounds horizontally
            return False
        elif not 0 <= row < len( testBoard.grid ): # Out of bound vertically
            return False
        elif not 0 <= row - ( len(SHAPES[shape]) - 1): # Out of bounds vertically accounting for shape length
            return False
        elif col + len(SHAPES[shape][0]) > testBoard.colLength: # Out of bounds horizontally accounting for shape length
            return False
        
        # See if any of the cells will intersect with a rock
        for i in range( len(SHAPES[shape]) ): # Row
            for j in range( len(SHAPES[shape][i]) ): # Col
                if SHAPES[shape][i][j] == None:
                    continue
                
                if testBoard.grid[row-i][col+j] == "#":
                    return False
                
                testBoard.grid[row-i][col+j] = "@"
        
        # testBoard.printBoard()
        return True


    # Flips the None values with the Rock (#) values
    def settleShape( self, shape, row, col ):
        if self.highestPoint <= row:
            self.highestPoint = row + 1

        for i in range( len(SHAPES[shape]) ): # Row
            for j in range( len(SHAPES[shape][i]) ): # Col
                if SHAPES[shape][i][j] == None:
                    continue
                
                if self.grid[row-i][col+j] == None:
                    self.grid[row-i][col+j] = "#"
                else:
                    print( "Something went wrong!")
                    return False
        
        # self.printBoard()

    # Finds the lowest point possible and then trims off everything below it
    # TODO: Solution to make part 2 possible
    def trimGrid( self ):
        lowestCol = 0
        for i in range( len(self.grid ) ):
            pass
        pass


    # Visualizer
    def printBoard(self):
        for i in reversed(range( len(self.grid) ) ):
            print( "|", end="" )
            for j in range( len( self.grid[i] ) ):
                if self.grid[i][j] == None:
                    print( ".", end="")
                else:
                    print( self.grid[i][j], end="" )
            print( "|" )
        print( "+-------+")
        print()


def readInput( output: list ):
    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        output.extend( [*file.readline().strip()] )

def main():

    instructions = list()
    board = Board(7)

    readInput( instructions )

    #instructions = [*">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"]

    rockCount = 0
    instructionTracker = 0
    currentCol = 2
    currentRow = 3
    board.addRow( 4 )

    # Main loop for creating rocks
    while rockCount < 2022: # rockCount < 2022 or 11 for short test

        # Out of bounds guard
        if instructionTracker == len(instructions):
            instructionTracker = 0

        # Determine the direction the shape is moving
        if instructions[instructionTracker] == ">":
            directionCol = 1
        else: 
            directionCol = -1

        instructionTracker += 1

        # Follow instruction
        if board.scanShape( rockCount % 5, currentRow, currentCol + directionCol ):
            currentCol += directionCol

        # Shape moves down
        if board.scanShape( rockCount % 5, currentRow - 1, currentCol ):
            currentRow -= 1

        # Shape can't move down so it settles
        else:
            board.settleShape( rockCount % 5, currentRow, currentCol )
            rockCount += 1
            currentCol = 2
            currentRow = board.highestPoint + 2 + len(SHAPES[rockCount%5])
            while len( board.grid ) <= currentRow:
                board.addRow()

    print( f"The highest point is: {board.highestPoint}")


if __name__ == "__main__":
    main()