"""
Advent of Code 2022 Day 8

Date: 12/8/2022
Author: phortheman

"""
from pathlib import Path

GRID = []

def isVisible( row: int, col:int ) -> bool:
    tree = GRID[row][col]
    top = True
    bottom = True
    left = True
    right = True

    for leftIterator in range( col ):
        if GRID[row][leftIterator] >= tree:
            left = False
            break
    
    for rightIterator in range( col + 1, len(GRID[row]) ):
        if GRID[row][rightIterator] >= tree:
            right = False
            break

    for topIterator in range( row ):
        if GRID[topIterator][col] >= tree:
            top = False
            break
    
    for bottomIterator in range( row + 1, len(GRID) ):
        if GRID[bottomIterator][col] >= tree:
            bottom = False
            break

    return top or bottom or left or right

def getScenicScore( row: int, col: int ) -> int:
    tree = GRID[row][col]

    left = 0
    right = 0
    top = 0
    bottom = 0

    for leftIterator in reversed( range( col ) ):
        left += 1
        if GRID[row][leftIterator] >= tree:
            break

    for rightIterator in range( col + 1, len(GRID[row]) ):
        right += 1
        if GRID[row][rightIterator] >= tree:
            break

    for topIterator in reversed( range( row ) ):
        top += 1
        if GRID[topIterator][col] >= tree:
            break

    for bottomIterator in range( row + 1, len(GRID) ):
        bottom += 1
        if GRID[bottomIterator][col] >= tree:
            break

    return left * right * top * bottom

with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:

    for row in file.readlines():
        GRID.append( list( map(int, row.strip()) ) )

    maxRow = len(GRID) - 1
    maxCol = len(GRID[0]) - 1

    visibleCount = (len(GRID) * 2) + (len(GRID[0]) * 2 - 4 )

    scenicScore = 0

    for i in range(len(GRID)):
        if i == 0 or i == maxRow:
            continue
        for j in range( len(GRID[i]) ):
            if j == 0 or j == maxCol:
                continue

            visibleCount += isVisible( i, j )
            treeScore = getScenicScore( i, j )
            if treeScore > scenicScore:
                scenicScore = treeScore

    print( f"The number of tree visible is: {visibleCount}" )
    print( f"The best scenic score is: {scenicScore}")
