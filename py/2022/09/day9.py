"""
Advent of Code 2022 Day 9

Date: 12/9/2022
Author: phortheman

"""
from pathlib import Path

class Knot:
    def __init__( self ):
        self.x = 0
        self.y = 0

    def __repr__(self) -> str:
        return f"({self.x}, {self.y})"

    def move( self, x: int, y: int ):
        self.x += x
        self.y += y

    def chaseKnot( self, prevKnot ):
        xDiff = self.x - prevKnot.x
        yDiff = self.y - prevKnot.y

        if( abs( xDiff ) < 2 and abs(yDiff) < 2 ):
            return # Don't move
        
        if ( xDiff != 0 and yDiff != 0 ): # Diagonal
            if( xDiff < 0 ):  # Move Right
                self.x += 1
            elif( xDiff > 0 ): # Move Left
                self.x -= 1
            
            if( yDiff < 0 ):   # Move Up
                self.y += 1
            elif( yDiff > 0 ): # Move Down
                self.y -= 1
        else:
            if( xDiff < 0 ):   # Move Right
                self.x += 1
            elif( xDiff > 0 ): # Move Left
                self.x -= 1
            elif( yDiff < 0 ): # Move Up
                self.y += 1
            elif( yDiff > 0 ): # Move Down
                self.y -= 1


with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:

    rope = []
    size = 10 # 2 for part 1, 10 for part 2
    for i in range( size ):
        rope.append( Knot() )

    positionsVisited = set()
    positionsVisited.add( (0,0) )

    for instruction in file.readlines():
        direction, count = instruction.split()

        match direction:
            case "R":
                xStep, yStep = 1, 0
            case "L":
                xStep, yStep = -1, 0
            case "D":
                xStep, yStep = 0, -1
            case "U":
                xStep, yStep = 0, 1
        
        for i in range( int(count) ):

            rope[0].move( xStep, yStep )

            for j in range( 1, len(rope) ):
                rope[j].chaseKnot( rope[j-1] )
            
            positionsVisited.add( (rope[-1].x, rope[-1].y) )

    print( f"The number of positions the tail has visited is: {len( positionsVisited )}" )
