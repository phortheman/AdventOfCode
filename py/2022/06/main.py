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


PACKET_SIZE = 4
MESSAGE_SIZE = 14


def main():
    input_path = get_puzzle_input()

    with open(input_path, "r") as f:
        datastream = f.readline()
        startOfPacketMarker = 0
        startOfMessageMarker = 0
        for i in range(len(datastream)):
            if (
                startOfPacketMarker == 0
                and len(set(datastream[i:i + PACKET_SIZE])) == PACKET_SIZE
            ):
                startOfPacketMarker = i + PACKET_SIZE

            if (
                startOfMessageMarker == 0
                and len(set(datastream[i:i + MESSAGE_SIZE])) == MESSAGE_SIZE
            ):
                startOfMessageMarker = i + MESSAGE_SIZE

            if startOfPacketMarker != 0 and startOfMessageMarker != 0:
                break

    print(f"Start-of-packet marker: {startOfPacketMarker}")
    print(f"Start-of-message marker: {startOfMessageMarker}")


if __name__ == "__main__":
    main()
