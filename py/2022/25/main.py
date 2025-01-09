import argparse
import os
import sys


def get_puzzle_input():
    """
    Handles input file retrieval with the option to override the default path.
    """
    # Get the day and year directory this script is stored in
    day = os.path.basename(os.path.dirname(os.path.abspath(__file__)))
    year = os.path.basename(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

    # Build the default input path
    root_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), "../../.."))
    default_input_path = os.path.join(root_dir, "inputs", year, day, "input.txt")

    parser = argparse.ArgumentParser(description="Advent of Code Solution")
    parser.add_argument(
        "-i",
        "--input",
        help="Specify a different puzzle input file path",
        default=default_input_path,
    )
    args = parser.parse_args()

    input_path = args.input

    if not os.path.exists(input_path):
        print(f"Error: Input file '{input_path}' does not exist.", file=sys.stderr)
        sys.exit(66)

    return input_path


def readInput(decimal: int = 0):
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        for line in f.readlines():
            decimal += convertSNAFUToDecimal(line.strip())
    return decimal


def convertDecimalToSNAFU(decimal: int) -> str:
    output = ""
    while decimal != 0:
        snafuDigit = ((decimal + 2) % 5) - 2
        if snafuDigit == -2:
            snafuDigit = "="
        elif snafuDigit == -1:
            snafuDigit = "-"
        else:
            snafuDigit = str(snafuDigit)
        output = snafuDigit + output
        decimal = (decimal + 2) // 5
    return output


def convertSNAFUToDecimal(snafu: str) -> int:
    decimal = 0
    for d in range(len(snafu)):
        if snafu[d] == "-":
            snafuDigit = -1
        elif snafu[d] == "=":
            snafuDigit = -2
        else:
            snafuDigit = int(snafu[d])
        decimal += snafuDigit * (pow(5, len(snafu) - 1 - d))
    return decimal


def main():
    sumOfFuel = readInput()
    print(f"The sum of the fuel as a deciaml number is: {sumOfFuel}")
    print(f"As a SNAFU number: {convertDecimalToSNAFU(sumOfFuel)}")
    pass


if __name__ == "__main__":
    main()
