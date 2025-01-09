import argparse
import os
import sys
from math import lcm
import copy


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


class Monkey:
    def __init__(self) -> None:
        self.items: list = []
        self.operation: str = ""
        self.modulusValue: int = 0
        self.trueID: str = ""
        self.falseID: str = ""
        self.inspectCount: int = 0

    def __str__(self) -> str:
        return (
            "Monkey\n"
            f"Items: {self.items}\n"
            f"Operation: {self.operation}\n"
            f"Modulus: {self.modulusValue}\n"
            f"True ID: {self.trueID}\n"
            f"False ID: {self.falseID}\n"
            f"Inspect: {self.inspectCount}"
        )

    def __repr__(self) -> str:
        return (
            "Monkey\n"
            f"Items: {self.items}\n"
            f"Operation: {self.operation}\n"
            f"Modulus: {self.modulusValue}\n"
            f"True ID: {self.trueID}\n"
            f"False ID: {self.falseID}\n"
            f"Inspect: {self.inspectCount}"
        )


def getMonkeyBusiness(monkeyDict: dict) -> int:
    topTwo = []
    for monkey in monkeyDict.values():
        topTwo.append(monkey.inspectCount)
    topTwo.sort(reverse=True)
    return topTwo[0] * topTwo[1]


def roundRulesPart1(monkeyDict: dict, rounds: int):
    for monkey in monkeyDict.values():
        while len(monkey.items) > 0:
            monkey.items.pop(0)
            new = eval(monkey.operation)

            new //= 3
            monkey.inspectCount += 1

            recievingMonkey = monkey.falseID

            if new % monkey.modulusValue == 0:
                recievingMonkey = monkey.trueID

            monkeyDict[recievingMonkey].items.append(new)


def roundRulesPart2(monkeyDict: dict, rounds: int, leastCommonMultiple: int):
    for monkey in monkeyDict.values():
        while len(monkey.items) > 0:
            monkey.items.pop(0)
            new = eval(monkey.operation)

            new %= leastCommonMultiple  # Pain
            monkey.inspectCount += 1

            recievingMonkey = monkey.falseID

            if new % monkey.modulusValue == 0:
                recievingMonkey = monkey.trueID

            monkeyDict[recievingMonkey].items.append(new)


def main():
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        startingMonkeys = {}
        currentMonkey = ""

        for line in f.readlines():
            cleanedUpLine = line.replace(",", "").replace(":", "")
            values = cleanedUpLine.split()

            if len(values) == 0:
                continue

            match values[0]:
                case "Monkey":
                    currentMonkey = values[1]
                    startingMonkeys[currentMonkey] = Monkey()

                case "Starting":
                    startingMonkeys[currentMonkey].items = list(map(int, values[2:]))

                case "Operation":
                    startingMonkeys[currentMonkey].operation = " ".join(values[3:])

                case "Test":
                    startingMonkeys[currentMonkey].modulusValue = int(values[3])

                case "If":
                    if values[1] == "true":
                        startingMonkeys[currentMonkey].trueID = values[-1]
                    elif values[1] == "false":
                        startingMonkeys[currentMonkey].falseID = values[-1]

                case _:
                    pass

    leastCommonMultiple = lcm(
        *[monkey.modulusValue for monkey in startingMonkeys.values()]
    )
    part1Monkeys = copy.deepcopy(startingMonkeys)
    part2Monkeys = copy.deepcopy(startingMonkeys)

    for _ in range(20):
        roundRulesPart1(part1Monkeys, 20)
    print(
        f"The level of Monkey Business for part 1 is: {getMonkeyBusiness(part1Monkeys)}"
    )

    for _ in range(10_000):
        roundRulesPart2(part2Monkeys, 10_000, leastCommonMultiple)
    print(
        f"The level of Monkey Business for part 2 is: {getMonkeyBusiness(part2Monkeys)}"
    )


if __name__ == "__main__":
    main()
