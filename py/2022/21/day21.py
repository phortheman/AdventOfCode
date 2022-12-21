"""
Advent of Code 2022 Day 21

Date: 12/21/2022
Author: phortheman

"""

from pathlib import Path
from copy import deepcopy

def reverseOperator( operator: str ):
    match operator:
        case "+":
            return "-"
        case "-":
            return "+"
        case "*":
            return "/"
        case "/":
            return "*"

def getMonkeyNumber( inputDict: dict[str,str], key: str ) -> int:
    try:
        return int(inputDict[key])
    except ValueError:
        first, operator, second = inputDict[key].split()

        first = getMonkeyNumber(inputDict, first)
        second = getMonkeyNumber(inputDict, second)

        if operator == "/": operator = "//"
        
        inputDict[key] = eval( f"{first} {operator} {second}" )
        return inputDict[key]

def readInput( inputDict: dict ):
    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        for line in file.readlines():
            inputDict[line[:4]] = line[6:].strip()

def main():
    monkeys = dict()
    readInput( monkeys )

    print( f"The number root will yell out is: {getMonkeyNumber( deepcopy(monkeys), 'root' )}" )


if __name__ == "__main__":
    main()