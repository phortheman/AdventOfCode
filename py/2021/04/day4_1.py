# Advent of Code 2021 Day 3 Part One
# Author: phortheman
# Date: 03/26/2022

from pathlib import Path
from BingoClasses import BingoField, BingoBoard


boardsList = []

# Got tired of changing the input dir so now I look for 'input.txt' in the running script's directory
with open( Path(__file__).with_name( 'input.txt' ), 'r' ) as f:
    # Drawn numbers in reverse order to allow for .pop() call
    drawnNumberList = [ x for x in reversed( f.readline().strip().split(',') ) ]

    f.readline() # The next line is empty so skip over it
    tempList = []
    for line in f:
        if( line == '\n' ): # This indicates a new board
            curBoard = BingoBoard( tempList )
            boardsList.append( curBoard )
            tempList.clear()
        else: 
            tempList.append( [ BingoField(x) for x in line.split() ])

curBoard = BingoBoard( tempList )
boardsList.append( curBoard )
tempList.clear()

# Start Drawing numbers
while len(drawnNumberList) != 0:
    winningBoard = None
    drawnNumber = drawnNumberList.pop()
    boardsToCheckForWin = []
    for board in boardsList:
        if( board.checkBoardForValue( drawnNumber ) ):
            boardsToCheckForWin.append(board)
    
    for board in boardsToCheckForWin:
        if( board.checkIfBoardWon() ):
            winningBoard = board
            break
    
    if( winningBoard != None ):
        break

score = winningBoard.sumOfUnmatchedFields()
score = score * int( drawnNumber )

print( score )