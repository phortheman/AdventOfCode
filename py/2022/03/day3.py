"""
Advent of Code 2022 Day 3

Date: 12/3/2022
Author: phortheman

"""
from pathlib import Path


"""
Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52
"""
def getPriority( char: str ) -> int:
    # Lower Case
    if( ord(char) > 96 ):
        return ord(char) - 96
    # Upper Case
    else:
        return ord(char) - 38

def getItemType( contents: str ) -> str:
    splitIndex = int( len(contents)/2 )
    for item in contents[:splitIndex]:
        if contents[splitIndex:].find(item) != -1:
            return item

def getGroupPriority( group: list )-> int:
    for item in min( group ):
        if group[0].find(item) != -1 and group[1].find(item) != -1 and group[2].find(item) != -1:
            return getPriority(item)

with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
    
    prioritySum = 0
    groupPrioritySum = 0
    workingGroup = []
    for rucksack in file.readlines():

        prioritySum += getPriority( getItemType( rucksack[:-1] ) )

        workingGroup.append(rucksack[:-1])

        if len(workingGroup) > 2:
            groupPrioritySum += getGroupPriority( workingGroup )
            workingGroup = []

print( prioritySum )
print( groupPrioritySum )