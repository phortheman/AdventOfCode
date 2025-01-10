import argparse
import os
import sys
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


# Parses the line until it starts to get instructions. Returns the cargo structure
def parse_cargo(lines: list[str]) -> (int, list[str]):
    cargo = []
    for num, line in enumerate(lines):
        # If the first character of the line is a space we've finished parsing the cargo
        if line.startswith(" "):

            # Reverse the stack so we can use pop and append later
            for stack in cargo:
                stack = stack.reverse()

            # Return line num plus 2 so we start at the first move instruction
            return num + 2, cargo

        stack = 0
        for i in range(1, len(line), 4):
            stack += 1

            # Create the cargo structure. 1 based
            while len(cargo) <= stack:
                cargo.append([])

            # If a letter isn't parsed then there isn't any cargo in this stack
            if not line[i].isalpha():
                continue

            cargo[stack].append(line[i])


def main():
    input_path = get_puzzle_input()

    part1, part2 = "", ""
    with open(input_path, "r") as f:
        lines = f.readlines()

        # Parse out the cargo data and get the instruction start line number
        start, cargoPart1 = parse_cargo(lines)
        cargoPart2 = copy.deepcopy(cargoPart1)

        # Read the instructions and apply them
        for line in lines[start:]:
            instructions = line.split()
            quantity = int(instructions[1])
            fromStack = int(instructions[3])
            toStack = int(instructions[5])

            # Apply the instruction for part 1
            for _ in range(quantity):
                cargoPart1[toStack].append(cargoPart1[fromStack].pop())

            # Apply the instruction for part 2

            # Calculate the range of the list that will be picked up by the crane
            quantity = len(cargoPart2[fromStack]) - quantity
            if quantity < 0:  # If we get a negative number we will set it to zero because that mean all the crates are moving
                quantity = 0

            crane = cargoPart2[fromStack][quantity:]  # Store what is grabbed by the crane
            cargoPart2[fromStack] = cargoPart2[fromStack][:quantity]  # Remove what was grabbed by the crane
            cargoPart2[toStack] += crane  # Put what is on the crane onto the new stack
            crane.clear()  # Remove the crates from the crane

        for crate in cargoPart1:
            if len(crate) > 0:
                part1 += crate[-1]

        for crate in cargoPart2:
            if len(crate) > 0:
                part2 += crate[-1]

    print("Part 1: ", part1)
    print("Part 2: ", part2)


if __name__ == "__main__":
    main()
