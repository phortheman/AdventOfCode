"""
Advent of Code 2022 Day 25

Date: 12/25/2022
Author: phortheman

"""

from pathlib import Path

def readInput( decimal: int = 0 ):
    with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
        for line in file.readlines():
            decimal += convertSNAFUToDecimal( line.strip() )
    return decimal

def convertDecimalToSNAFU( decimal: int ) -> str:
    output = ""
    while decimal != 0:
        snafuDigit = ((decimal + 2) % 5) - 2
        if snafuDigit == -2: snafuDigit = "="
        elif snafuDigit == -1: snafuDigit = "-"
        else: snafuDigit = str(snafuDigit)
        output = snafuDigit + output
        decimal = ( decimal + 2 ) // 5
    return output

def convertSNAFUToDecimal( snafu: str) -> int:
    decimal = 0
    for d in range(len(snafu)):
        if snafu[d] == "-": snafuDigit = -1
        elif snafu[d] == "=": snafuDigit = -2
        else: snafuDigit = int(snafu[d])
        decimal += ( snafuDigit * ( pow(5, len(snafu) - 1 - d ) ) )
    return decimal

def main():
    sumOfFuel = readInput()
    print( f"The sum of the fuel as a deciaml number is: {sumOfFuel}" )
    print( f"As a SNAFU number: {convertDecimalToSNAFU(sumOfFuel)}")
    pass

if __name__ == "__main__":
    main()