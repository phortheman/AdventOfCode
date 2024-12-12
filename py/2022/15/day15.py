"""
Advent of Code 2022 Day 15

Date: 12/15/2022
Author: phortheman

"""
from pathlib import Path
from z3 import *

def manhattanDistance( x0, y0, x1, y1 ):
    return abs(x0 - x1) + abs( y0 - y1 )

def main():
    
    rawData = dict()

    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        for line in file.readlines():
            lineData = line.strip().replace(":", "").replace(",", "").split()
            sensorX, sensorY = lineData[2][2:], lineData[3][2:]
            beaconX, beaconY = lineData[8][2:], lineData[9][2:]

            rawData[ (int(sensorX), int(sensorY))] = manhattanDistance( int(sensorX), int(sensorY), int(beaconX), int(beaconY) )

    part1Target = 2_000_000
    part2Target = 4_000_000

    
    part1Result = (
        max( sensor[0] - abs(part1Target-sensor[1]) + distance for sensor, distance in rawData.items() )
        - min( sensor[0] + abs(part1Target-sensor[1]) - distance for sensor, distance in rawData.items() )
    )

    solver = Solver()
    distressX, distressY = Int('x'), Int('y')

    solver.add( 0 <= distressX, distressX <= part2Target )
    solver.add( 0 <= distressY, distressY <= part2Target )

    for sensor, distance in rawData.items():
        sensorX, sensorY = sensor
        solver.add( Abs(sensorX - distressX) + Abs( sensorY - distressY ) > distance )

    solver.check()
    part2Result = solver.model()

    print( f"The number of position at {part1Target} that can't have a beacon is: {part1Result}" )
    print( f"The tuning frequency is: {part2Result[distressX].as_long() * 4_000_000 + part2Result[distressY].as_long()}" )

if __name__ == "__main__":
    main()