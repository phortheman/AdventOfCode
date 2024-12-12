"""
Advent of Code 2022 Day 18

Date: 12/18/2022
Author: phortheman

"""
from pathlib import Path
from collections import defaultdict, deque

MAX_X = 0
MAX_Y = 0
MAX_Z = 0

MIN_X = 1
MIN_Y = 1
MIN_Z = 1

def readInput(  xAxis: defaultdict,
                yAxis: defaultdict,
                zAxis: defaultdict ):
    global MAX_X, MAX_Y, MAX_Z, MIN_X, MIN_Y, MIN_Z
    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        for line in file.readlines():
            x, y, z = map(int, line.strip().split(","))

            if x >= MAX_X: MAX_X = x + 1
            if y >= MAX_Y: MAX_Y = y + 1
            if z >= MAX_Z: MAX_Z = z + 1
            if x <= MIN_X: MIN_X = x
            if y <= MIN_Y: MIN_Y = y
            if z <= MIN_Z: MIN_Z = z

            xAxis[x].append( (y, z) )
            yAxis[y].append( (x, z) )
            zAxis[z].append( (x, y) )

# Over engineered why to visualize the droplets
def visualize( **kwargs ):
    for z in range( MIN_Z, MAX_Z ):
        print( f"Z Axis: {z}")

        headers = [*kwargs.keys()]

        rowMaxLength = len( max( headers ) )

        if rowMaxLength < MAX_X + 6:
            rowMaxLength = MAX_X + 6

        for title in headers:
            print( f"| {title:-^{rowMaxLength}}", end="")
        print( " |")
        
        for y in range( MIN_Y - 1, MAX_Y + 1 ):
            for data in kwargs.values():
                line = ""
            
                for x in range( MIN_X - 1, MAX_X + 1 ):
                    if z in data.keys() and (x, y) in data[z]:
                        line += "#"
                    else:
                        line += "."
                print(f"| {line:^{rowMaxLength}}", end="")
            print(" |")

# Shared coordinates on an axis position +- 1 means it isn't exposed
def findExposedSides( inputDict: dict ) -> int:
    sides = 0
    for axis in inputDict.keys():
        for coords in inputDict[axis]:
            if axis + 1 not in inputDict.keys():
                sides += 1
            elif coords not in inputDict[axis+1]:
                sides += 1

            if axis - 1 not in inputDict.keys():
                sides += 1
            elif coords not in inputDict[axis-1]:
                sides += 1
    return sides

# Do a 3 dimentional BFS starting at the min coords minus 1 until there are no more spots to visit
# Returns the number of sides on the exterior and a set of the visited points
def fillWithWaterDroplets( lavaDroplets: dict ) -> tuple[int, set]:
    sides = 0

    start = (MIN_X - 1, MIN_Y - 1, MIN_Z - 1)
    queue = deque( [start] )
    visited = { start }

    directionX = [-1, 1, 0, 0, 0, 0]
    directionY = [0, 0, 1, -1, 0, 0 ]
    directionZ = [0, 0, 0, 0, 1, -1 ]

    while queue:
        currentX, currentY, currentZ = queue.popleft()

        for i in range( len(directionX) ):
            nextX = currentX + directionX[i]
            nextY = currentY + directionY[i]
            nextZ = currentZ + directionZ[i]

            # Boundry check
            if not start[0] <= nextX <= MAX_X:
                continue
            if not start[1] <= nextY <= MAX_Y:
                continue
            if not start[2] <= nextZ <= MAX_Z:
                continue

            nextCoord = ( nextX, nextY, nextZ )

            # Check if we've already added this point
            if nextCoord in visited:
                continue

            # Check if this point is lava
            if nextZ in lavaDroplets.keys() and (nextX, nextY) in lavaDroplets[nextZ]:
                sides += 1
                continue
            
            visited.add( nextCoord )
            queue.append( nextCoord )

    return sides, visited

def main():
    lavaDropletsXAxis = defaultdict(list)
    lavaDropletsYAxis = defaultdict(list)
    lavaDropletsZAxis = defaultdict(list)

    waterDropletsZAxis = defaultdict(list)

    readInput( lavaDropletsXAxis, lavaDropletsYAxis, lavaDropletsZAxis  )

    # visualize( lavaDropletsZAxis )

    totalNumberOfSides = (  findExposedSides( dict(lavaDropletsXAxis) ) +
                            findExposedSides( dict(lavaDropletsYAxis) ) +
                            findExposedSides( dict(lavaDropletsZAxis) ) 
    )

    exteriorSides, waterCoords = fillWithWaterDroplets( dict(lavaDropletsZAxis) )

    for x, y, z in waterCoords:
        waterDropletsZAxis[z].append( (x, y) )

    visualize( Lava=lavaDropletsZAxis, Water=waterDropletsZAxis )

    print( f"The total number of sides exposed to air is: {totalNumberOfSides}" )
    print( f"The number of sides on the exterior is: {exteriorSides}")

if __name__ == "__main__":
    main()