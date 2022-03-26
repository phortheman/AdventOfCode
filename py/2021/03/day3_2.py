# Advent of Code 2021 Day 3 Part Two
# Author: phortheman
# Date: 03/25/2022

# That first solution was cute but now I have to actually do it


def bitCriteria( inputList, matchBit, pos, max_pos, noCommonBitOverride, onlyCommon ):
    newList = []
    for i in inputList:
        if( i[pos] == matchBit ):
            newList.append( i )

    if ( len( newList ) == 1 ): # All done!
        return newList[0]
    elif ( pos != max_pos ): # Recurse again!
        pos += 1

        newTransposedList = [list(i) for i in zip(*newList)]
        if ( newTransposedList[pos].count('1') > newTransposedList[pos].count('0' ) ):
            if( onlyCommon ): matchBit = '1'
            else: matchBit = '0'
        elif ( newTransposedList[pos].count('1') < newTransposedList[pos].count('0' ) ):
            if( onlyCommon ): matchBit = '0'
            else: matchBit = '1'
        else:
            matchBit = noCommonBitOverride
        
        return bitCriteria( newList, matchBit, pos, max_pos, noCommonBitOverride, onlyCommon )
    else:
        print( "Something went wrong! No value was found" )


with open( "03/input_3.txt", 'r' ) as f:
    listInput = f.read().splitlines()

strGamma = ""
strEpsilon = ""


transposedListInput = [list(i) for i in zip(*listInput)]

for i in transposedListInput:
    if ( i.count('1') > i.count('0' ) ):
        strGamma += '1'
        strEpsilon += '0'
    else:
        strGamma += '0'
        strEpsilon += '1'

strOxyGenRating = bitCriteria( listInput, strGamma[0], 0, len(strGamma), '1', onlyCommon=True )
strCO2ScrubberRating = bitCriteria( listInput, strEpsilon[0], 0, len(strEpsilon), '0', onlyCommon=False )


iGamma = int( strGamma, 2 )
iEpsilon = int( strEpsilon, 2 )
iOxyGenRating = int( strOxyGenRating, 2 )
iCO2ScurbberRating = int( strCO2ScrubberRating, 2 )

print( "Power Level: " + str( iGamma * iEpsilon ) )
print ( "Life Support Rating: " + str( iOxyGenRating * iCO2ScurbberRating ) )