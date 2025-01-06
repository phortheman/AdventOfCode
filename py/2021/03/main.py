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


def bitCriteria(inputList, matchBit, pos, max_pos, noCommonBitOverride, onlyCommon):
    newList = []
    for i in inputList:
        if i[pos] == matchBit:
            newList.append(i)

    if len(newList) == 1:  # All done!
        return newList[0]
    elif pos != max_pos:  # Recurse again!
        pos += 1

        newTransposedList = [list(i) for i in zip(*newList)]
        if newTransposedList[pos].count("1") > newTransposedList[pos].count("0"):
            if onlyCommon:
                matchBit = "1"
            else:
                matchBit = "0"
        elif newTransposedList[pos].count("1") < newTransposedList[pos].count("0"):
            if onlyCommon:
                matchBit = "0"
            else:
                matchBit = "1"
        else:
            matchBit = noCommonBitOverride

        return bitCriteria(
            newList, matchBit, pos, max_pos, noCommonBitOverride, onlyCommon
        )
    else:
        print("Something went wrong! No value was found")


def part2(input_path: str):
    with open(input_path, "r") as f:
        listInput = f.read().splitlines()

    strGamma = ""
    strEpsilon = ""

    transposedListInput = [list(i) for i in zip(*listInput)]

    for i in transposedListInput:
        if i.count("1") > i.count("0"):
            strGamma += "1"
            strEpsilon += "0"
        else:
            strGamma += "0"
            strEpsilon += "1"

    strOxyGenRating = bitCriteria(
        listInput, strGamma[0], 0, len(strGamma), "1", onlyCommon=True
    )
    strCO2ScrubberRating = bitCriteria(
        listInput, strEpsilon[0], 0, len(strEpsilon), "0", onlyCommon=False
    )

    iGamma = int(strGamma, 2)
    iEpsilon = int(strEpsilon, 2)
    iOxyGenRating = int(strOxyGenRating, 2)
    iCO2ScurbberRating = int(strCO2ScrubberRating, 2)

    print("Power Level: " + str(iGamma * iEpsilon))
    print("Life Support Rating: " + str(iOxyGenRating * iCO2ScurbberRating))


def main():
    input_path = get_puzzle_input()

    list_of_Sums = [0] * 12
    iCount = 0

    with open(input_path, "r") as f:
        for line in f:
            bits = list(line)
            list_of_Sums = [
                origVal + int(newVal) for origVal, newVal in zip(list_of_Sums, bits)
            ]
            iCount += 1

    # Part 1
    strGamma = ""
    strEpsilon = ""

    for i in list_of_Sums:
        if i > iCount / 2:
            strGamma += "1"
            strEpsilon += "0"
        else:
            strGamma += "0"
            strEpsilon += "1"

    iGamma = int(strGamma, 2)
    iEpsilon = int(strEpsilon, 2)

    print(iGamma * iEpsilon)

    part2(input_path)


if __name__ == "__main__":
    main()
