"""
Advent of Code 2022 Day 24

Date: 12/24/2022
Author: phortheman

"""

from pathlib import Path
from collections import defaultdict
from copy import deepcopy

DIRECTION_X = [-1, 1, 0,  0 ]
DIRECTION_Y = [ 0, 0, 1, -1 ]

class Blizzard:
    def __init__( self, pos: int, direction: str ) -> None:
        self.pos = pos

        match direction:
            case ">":
                self.direction = direction
                self.step = 1
            case "v":
                self.direction = direction
                self.step = 1
            case "<":
                self.direction = direction
                self.step = -1
            case "^":
                self.direction = direction
                self.step = -1
            case _:
                raise ValueError("Only accepted values for direction are 'v,^,<,>'")
    
    def __str__(self) -> str:
        return self.direction

    def __repr__(self) -> str:
        return f"Pos: {self.pos}, Direction: {self.direction}"

    def __eq__(self, __o: object) -> bool:
        return __o == self.pos

class Map:
    def __init__(self) -> None:
        # Key: constant axis position. Value: position on that axis that gets incremented/decremented every minute
        self.xAxisBlizzards = defaultdict(list)
        self.yAxisBlizzards = defaultdict(list)

        self.startPos = None
        self.endPos = None
        self.boundaryX = 0
        self.boundaryY = 0

        # Current elves
        self.elves = set()

    # Return True is an elf has reached the end
    def tick(self) -> bool:
        # Move blizzards first then try to move elves
        for xBlizzards in self.xAxisBlizzards.values():
            for xBliz in xBlizzards:
                xBliz.pos += xBliz.step
                if xBliz.pos == 0: xBliz.pos = self.boundaryY
                elif xBliz.pos > self.boundaryY: xBliz.pos = 1

        for yBlizzards in self.yAxisBlizzards.values():
            for yBliz in yBlizzards:
                yBliz.pos += yBliz.step
                if yBliz.pos == 0: yBliz.pos = self.boundaryX
                elif yBliz.pos > self.boundaryX: yBliz.pos = 1

        # Loop through elves
        currentElves = deepcopy(self.elves)

        for elf in currentElves:

            for i in range( len(DIRECTION_X) ):
                newX = elf[0] + DIRECTION_X[i]
                newY = elf[1] + DIRECTION_Y[i]

                if (newX, newY) == self.endPos:
                    self.elves.clear()
                    self.elves.add( ( newX, newY ) )
                    return True

                # If the position is not in bounds the elves can't move there
                if not 0 < newX < self.boundaryX + 1:
                    continue
                elif not 0 < newY < self.boundaryY + 1:
                    continue

                # Elves would move into the blizzard
                if newX in self.yAxisBlizzards[newY] or newY in self.xAxisBlizzards[newX]:
                    continue

                self.elves.add( (newX, newY) )


            # This elf is currenting in a blizzard so it can't wait
            if elf[1] in self.xAxisBlizzards[elf[0]] or elf[0] in self.yAxisBlizzards[elf[1]]:
                self.elves.remove(elf)

        return False

    def visualize( self ):
        output = ""

        # First line logic
        for i in range( self.boundaryX + 2 ):
            char = "#"
            if self.startPos == ( i, 0 ):
                char = "."
            if ( i, 0 ) in self.elves:
                char = "e"
            output += char
        
        output += "\n"
        for y in range( 1, self.boundaryY + 1 ):
            output += "#"
            for x in range( 1, self.boundaryX + 1 ):
                char = "."
                if ( x, y ) in self.elves:
                    char = "e"
                    output += char
                    continue
                for bliz in self.xAxisBlizzards[x]:
                    if bliz.pos == y:
                        char = bliz.direction
                        break
                for bliz in self.yAxisBlizzards[y]:
                    if bliz.pos == x:
                        if char != ".": char = "2"
                        else: char = bliz.direction
                        break
                output += char
            output += "#\n"

        # Last line logic
        for j in range( self.boundaryX + 2 ):
            char = "#"
            if self.endPos == ( j, self.boundaryY + 1 ):
                char = "."
            if ( j, self.boundaryY + 1 ) in self.elves:
                char = "E"
            output += char

        print( output )


# Return the start position
def readInput( inputMap: Map ):
    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        currentY = 0
        inputMap.startPos = ( file.readline().index("."), 0 )
        for line in file.readlines():
            currentY += 1
            for x in range(len(line)):
                if line[x] == ">" or line[x] == "<":
                    inputMap.yAxisBlizzards[currentY].append( Blizzard( x, line[x] ) )
                elif line[x] == "^" or line[x] == "v":
                    inputMap.xAxisBlizzards[x].append( Blizzard( currentY, line[x] ) )

        
    inputMap.boundaryY = currentY - 1
    inputMap.boundaryX = len( line ) - 2
    inputMap.endPos = ( line.index("."), currentY )


def main():

    map = Map()

    readInput( map )

    start = map.startPos
    end = map.endPos

    map.elves.add( map.startPos )

    timeElapsedFromStart1 = 1

    while not map.tick():
        timeElapsedFromStart1 += 1
        
    print( f"The fasted way requires {timeElapsedFromStart1} minutes" )
    print( "On no! We have to go back!" )

    map.endPos = start
    map.startPos = end

    timeElapsedFromEnd = 1
    while not map.tick():
        timeElapsedFromEnd += 1

    print( f"It took us {timeElapsedFromEnd} minutes to get back to the start" )
    print( "Time to go back to the end!" )

    map.endPos = end
    map.startPos = start

    timeElapsedFromStart2 = 1
    while not map.tick():
        timeElapsedFromStart2 += 1

    print( f"Finally back! It took us {timeElapsedFromStart2} minutes" )

    print( f"The total time elapsed is {timeElapsedFromStart1 + timeElapsedFromEnd + timeElapsedFromStart2}")


if __name__ == "__main__":
    main()