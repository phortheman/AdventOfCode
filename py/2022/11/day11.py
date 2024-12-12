"""
Advent of Code 2022 Day 11

Date: 12/11/2022
Author: phortheman

"""
from pathlib import Path
from math import lcm
import copy

class Monkey:

    def __init__(self) -> None:
        self.items: list = []
        self.operation: str = ""
        self.modulusValue: int = 0
        self.trueID: str = ""
        self.falseID: str = ""
        self.inspectCount: int = 0

    def __str__(self) -> str:
        return ("Monkey\n"
        f"Items: {self.items}\n"
        f"Operation: {self.operation}\n"
        f"Modulus: {self.modulusValue}\n"
        f"True ID: {self.trueID}\n"
        f"False ID: {self.falseID}\n"
        f"Inspect: {self.inspectCount}")

    def __repr__(self) -> str:
        return ("Monkey\n"
        f"Items: {self.items}\n"
        f"Operation: {self.operation}\n"
        f"Modulus: {self.modulusValue}\n"
        f"True ID: {self.trueID}\n"
        f"False ID: {self.falseID}\n"
        f"Inspect: {self.inspectCount}")

def getMonkeyBusiness( monkeyDict: dict ) -> int:
    topTwo = []
    for monkey in monkeyDict.values():
        topTwo.append(monkey.inspectCount)
    topTwo.sort(reverse=True)
    return topTwo[0] * topTwo[1]

def roundRulesPart1( monkeyDict: dict, rounds: int ):
    for monkey in monkeyDict.values():
        while( len(monkey.items) > 0 ):
            old = monkey.items.pop(0)
            new = eval( monkey.operation )

            new //= 3
            monkey.inspectCount += 1

            recievingMonkey = monkey.falseID

            if new % monkey.modulusValue == 0:
                recievingMonkey = monkey.trueID

            monkeyDict[recievingMonkey].items.append( new )

def roundRulesPart2( monkeyDict: dict, rounds: int, leastCommonMultiple: int ):
    for monkey in monkeyDict.values():
        while( len(monkey.items) > 0 ):
            old = monkey.items.pop(0)
            new = eval( monkey.operation )

            new %= leastCommonMultiple # Pain
            monkey.inspectCount += 1

            recievingMonkey = monkey.falseID

            if new % monkey.modulusValue == 0:
                recievingMonkey = monkey.trueID

            monkeyDict[recievingMonkey].items.append( new )

def main():
    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        
        startingMonkeys = {}
        currentMonkey = ""
        
        for line in file.readlines():
            cleanedUpLine = line.replace(",", "").replace(":", "")
            values = cleanedUpLine.split()

            if( len( values ) == 0 ):
                continue
            
            match values[0]:
                case "Monkey":
                    currentMonkey = values[1]
                    startingMonkeys[currentMonkey] = Monkey()

                case "Starting":
                    startingMonkeys[currentMonkey].items = list( map( int, values[2:] ) )

                case "Operation":
                    startingMonkeys[currentMonkey].operation = ' '.join(values[3:])

                case "Test":
                    startingMonkeys[currentMonkey].modulusValue = int( values[3] )

                case "If":
                    if( values[1] == "true" ):
                        startingMonkeys[currentMonkey].trueID = values[-1]
                    elif( values[1] == "false" ):
                        startingMonkeys[currentMonkey].falseID = values[-1]

                case _:
                    pass

    leastCommonMultiple = lcm(*[monkey.modulusValue for monkey in startingMonkeys.values()] )
    part1Monkeys = copy.deepcopy( startingMonkeys )
    part2Monkeys = copy.deepcopy( startingMonkeys )

    for _ in range( 20 ):
        roundRulesPart1( part1Monkeys, 20 )
    print( f"The level of Monkey Business for part 1 is: {getMonkeyBusiness( part1Monkeys )}" )

    for _ in range( 10_000 ):
        roundRulesPart2( part2Monkeys, 10_000, leastCommonMultiple )
    print( f"The level of Monkey Business for part 2 is: {getMonkeyBusiness( part2Monkeys )}" )

if __name__ == "__main__":
    main()