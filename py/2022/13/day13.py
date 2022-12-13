"""
Advent of Code 2022 Day 13

Date: 12/13/2022
Author: phortheman

"""
from pathlib import Path
import ast

def evaluatePair( pair: tuple ) -> int:

    leftList = pair[0]
    rightList = pair[1]

    iterations = min( len(leftList), len(rightList) )

    for i in range( iterations ):
        leftValue = leftList[i]
        rightValue = rightList[i]

        # Ensure each value is an integer
        if type( leftValue ) == int and type( rightValue ) == int:

            # Right order
            if leftValue < rightValue:
                return 1
            
            # Not right order
            elif rightValue < leftValue:
                return -1

            # Still need to evaluate
            else:
                continue

        # Both values are lists
        elif type(leftValue) == list and type(rightValue) == list:
            result = evaluatePair( (leftValue,rightValue) )
            if result != 0:
                return result

        # Left is a list
        elif type(leftValue) == list:
            result = evaluatePair( (leftValue, [rightValue]) )
            if result != 0:
                return result

        # Right is a list
        elif type(rightValue) == list:
            result = evaluatePair( ([leftValue], rightValue) )
            if result != 0:
                return result
    
    # Needs to continue
    if iterations == len(leftList) and iterations == len(rightList):
        return 0

    # Right order
    elif iterations == len(leftList):
        return 1

    # Not right order
    elif iterations == len(rightList):
        return -1

    return 0

def bubbleSort( value: list ):
    n = len( value )
    swapped = False

    for i in range( n - 1 ):
        for j in range( 0, n - i - 1 ):
            if evaluatePair( (value[j], value[j+1] ) ) == -1:
                swapped = True
                value[j], value[j+1] = value[j+1], value[j]

        if not swapped:
            return

def main():

    message = []
    correctIndexSum = 0
    leftSide = None
    rightSide = None
    currentIndex = 1

    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        for line in file.readlines():

            # New pair
            if line == '\n':
                leftSide = None
                rightSide = None
                currentIndex += 1
                continue

            if leftSide == None:
                leftSide = ast.literal_eval(line)
            elif rightSide == None:
                rightSide = ast.literal_eval(line)

                if evaluatePair( (leftSide, rightSide) ) == 1:
                    message.append( leftSide )
                    message.append( rightSide )
                    correctIndexSum += currentIndex
                else:
                    message.append( rightSide )
                    message.append( leftSide )


    print( f"The sum of the indices of the pairs is: {correctIndexSum}")

    # Add divider packets
    message.append( [[2]] )
    message.append( [[6]] )

    bubbleSort( message )

    divider1Index, divider2Index = 0, 0
    for i in range( len(message) ):
        if message[i] == [[2]]:
            divider1Index = i + 1
        if message[i] == [[6]]:
            divider2Index = i + 1
        
        if divider1Index != 0 and divider2Index != 0:
            break

    print( f"The decoder key is: {divider1Index * divider2Index}")

if __name__ == "__main__":
    main()
