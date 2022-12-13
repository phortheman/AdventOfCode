"""
Advent of Code 2022 Day 12

Date: 12/12/2022
Author: phortheman

"""

from pathlib import Path
from collections import deque

def breadthFirstSearch( start: tuple, map: list ):
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
    directionCol = [0, 0, 1, -1 ]

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

        for i in range( len(directionRow) ):
            # Get neighbor coodinates
            neighborRow = currentRow + directionRow[i]
            neighborCol = currentCol + directionCol[i]

            # Check if neighbor is valid
            if not 0 <= neighborRow < len(map):
                continue
            if not 0 <= neighborCol < len(map[0]):
                continue
            
            # Make the tuple for the position of the neighbor
            neighborPos = ( neighborRow, neighborCol )

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
            
            visited.add( neighborPos )
            queue.append( neighborPos )

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

    heightMap = []
    startPosition = None
    lowestPoints = []

    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        
        row = 0
        for line in file.readlines():
            col = []
            for curCol in range(len(line)):
                if line[curCol] == "\n":
                    break

                if line[curCol] == "S":
                    startPosition = (row, curCol)

                elif line[curCol] == "a":
                    lowestPoints.append( (row, curCol) )
                
                col.append(line[curCol])

            heightMap.append( col )
            row += 1

    fewestSteps = breadthFirstSearch( startPosition, heightMap )
    fewestStepsAnywhere = fewestSteps

    for lowPoint in lowestPoints:
            result = breadthFirstSearch( lowPoint, heightMap )

            # The search will return -1 if the point can't reach the end
            if result > 0 and result < fewestStepsAnywhere:
                fewestStepsAnywhere = result

    print( f"The fewest steps requried is: {fewestSteps}")
    print( f"The fewest steps from any square is: {fewestStepsAnywhere}")

if __name__ == "__main__":
    main()
