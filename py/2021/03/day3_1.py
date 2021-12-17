# Advent of Code 2021 Day 3 Part One
# Author: phortheman
# Date: 12/17/2021

# If the sum of the bits is more than half of the number of lines then 1 is the most common bit
# otherwise 0 is the most common bit

list_of_Sums = [0] * 12
iCount = 0
with open( "03/input_3.txt", 'r' ) as f:
    for line in f: 
        bits = list( line )
        list_of_Sums = [ origVal + int( newVal ) for origVal, newVal in zip( list_of_Sums, bits ) ]
        iCount += 1

strGamma = ""
strEpsilon = ""

for i in list_of_Sums:
    if( i > iCount/2 ):
        strGamma += "1"
        strEpsilon+= "0"
    else: 
        strGamma += "0"
        strEpsilon += "1"

iGamma = int( strGamma, 2 )
iEpsilon = int( strEpsilon, 2 )

print( iGamma * iEpsilon )
