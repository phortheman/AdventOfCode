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


# A BingoField is a data struct that have the value of the field and if it is was pulled
class BingoField:
    def __init__(self, value) -> None:
        self.value = value
        self.match = False

    def checkIfMatch(self, checkValue):
        if self.value == checkValue:
            self.match = True

    def __repr__(self) -> str:
        return str(self.value) + ":" + str(self.match)

    def __str__(self) -> str:
        return str(self.value) + ":" + str(self.match)


# A BingoBoard is a data struct that contains a 2D list of BingoFields
# And contains a dictionary of all the values and their coordinates on the 2D list
class BingoBoard:
    def __init__(self, twoDimListOfFields):
        self.valueIndexes = {}  # Key: value of the field, Value: "x,y" coord on the board
        self.board = twoDimListOfFields.copy()

        for xIndex, xItem in enumerate(self.board):
            for yIndex, yItem in enumerate(xItem):
                self.valueIndexes[yItem.value] = f"{xIndex},{yIndex}"

    def __repr__(self):
        return self.board

    def checkBoardForValue(self, drawnValue):
        onBoard = self.valueIndexes.get(drawnValue)
        if onBoard is not None:
            [row, col] = [int(i) for i in onBoard.split(",")]
            field = self.board[row][col]
            field.match = True
            return True
        else:
            return False

    def checkIfBoardWon(self):
        victoryCheck = 0
        for row in self.board:
            for field in row:
                if field.match:
                    victoryCheck += 1
                    if victoryCheck == 5:
                        return True
                else:
                    victoryCheck = 0
                    break

        transposedBoard = [list(i) for i in zip(*self.board)]
        for col in transposedBoard:
            for field in col:
                if field.match:
                    victoryCheck += 1
                    if victoryCheck == 5:
                        return True
                else:
                    victoryCheck = 0
                    break

        return False

    def sumOfUnmatchedFields(self):
        sum = 0
        for row in self.board:
            for field in row:
                if not field.match:
                    sum += int(field.value)
        return sum


def part1(input_path):
    boardsList = []
    with open(input_path, "r") as f:
        # Drawn numbers in reverse order to allow for .pop() call
        drawnNumberList = [x for x in reversed(f.readline().strip().split(","))]

        f.readline()
        tempList = []
        for line in f:
            if line == "\n":  # This indicates a new board
                curBoard = BingoBoard(tempList)
                boardsList.append(curBoard)
                tempList.clear()
            else:
                tempList.append([BingoField(x) for x in line.split()])

    curBoard = BingoBoard(tempList)
    boardsList.append(curBoard)
    tempList.clear()

    # Start Drawing numbers
    while len(drawnNumberList) != 0:
        winningBoard = None
        drawnNumber = drawnNumberList.pop()
        boardsToCheckForWin = []
        for board in boardsList:
            if board.checkBoardForValue(drawnNumber):
                boardsToCheckForWin.append(board)

        for board in boardsToCheckForWin:
            if board.checkIfBoardWon():
                winningBoard = board
                break

        if winningBoard is not None:
            break

    score = winningBoard.sumOfUnmatchedFields()
    score = score * int(drawnNumber)

    print("Part 1: ", score)


def part2(input_path):
    boardsList = []
    with open(input_path, "r") as f:
        # Drawn numbers in reverse order to allow for .pop() call
        drawnNumberList = [x for x in reversed(f.readline().strip().split(","))]

        f.readline()  # The next line is empty so skip over it
        tempList = []
        for line in f:
            if line == "\n":  # This indicates a new board
                curBoard = BingoBoard(tempList)
                boardsList.append(curBoard)
                tempList.clear()
            else:
                tempList.append([BingoField(x) for x in line.split()])

    curBoard = BingoBoard(tempList)
    boardsList.append(curBoard)
    tempList.clear()

    winningBoards = []
    # Start Drawing numbers
    while len(drawnNumberList) != 0:
        drawnNumber = drawnNumberList.pop()
        boardsToCheckForWin = []
        for board in boardsList:
            if board.checkBoardForValue(drawnNumber):
                boardsToCheckForWin.append(board)

        for board in boardsToCheckForWin:
            if board.checkIfBoardWon():
                winningBoards.append(board)
                boardsList.remove(board)  # Don't check won boards anymore

        if len(boardsList) == 0:  # We ran out of boards so all boards won
            break

    score = winningBoards.pop().sumOfUnmatchedFields()
    score = score * int(drawnNumber)

    print("Part 2: ", score)


def main():
    input_path = get_puzzle_input()

    part1(input_path)
    part2(input_path)


if __name__ == "__main__":
    main()
