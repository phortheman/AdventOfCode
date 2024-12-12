
using System.IO;

namespace Day_02
{

    public enum Direction
    {
        UP,
        DOWN,
        FORWARD
    }

    public struct Instruction
    {
        public Direction direction;
        public int units;
    }

    public struct Position
    {
        public int horizontal;
        public int depth;
        public int aim;

        public void RunInstruction(Instruction instruction)
        {
            switch (instruction.direction)
            {
                case Direction.UP:
                    {
                        depth -= instruction.units;
                        break;
                    }
                case Direction.DOWN:
                    {
                        depth += instruction.units;
                        break;
                    }
                case Direction.FORWARD:
                    {
                        horizontal += instruction.units;
                        break;
                    }
            }
        }

        public void RunWithAim(Instruction instruction)
        {
            switch (instruction.direction)
            {
                case Direction.UP:
                    {
                        aim -= instruction.units;
                        break;
                    }
                case Direction.DOWN:
                    {
                        aim += instruction.units;
                        break;
                    }
                case Direction.FORWARD:
                    {
                        horizontal += instruction.units;
                        depth += aim * instruction.units;
                        break;
                    }
            }
        }
    }

    public class Day_02_Solver
    {
        public Day_02_Solver() { }

        public List<Instruction> GenerateInput(String filePath)
        {
            List<Instruction> output = new();

            using (StreamReader reader = new(filePath))
            {
                string? line;
                while ((line = reader.ReadLine()) != null)
                {
                    Instruction instruction = new();
                    string[] fields = line.Split(" ");

                    Enum.TryParse(fields[0].ToUpper(), out instruction.direction);
                    int.TryParse(fields[1], out instruction.units);

                    output.Add(instruction);
                }
            }

            return output;
        }

        static void Main(string[] args)
        {
            if (args.Length == 0)
            {
                Console.WriteLine("Please provide the puzzle input as a file");
                return;
            }

            List<Instruction> inputData;

            Day_02_Solver solver = new();

            inputData = solver.GenerateInput(args[0]);

            Position part1Position = new();
            Position part2Position = new();

            foreach (Instruction instruction in inputData)
            {
                part1Position.RunInstruction(instruction);
                part2Position.RunWithAim(instruction);
            }

            Console.WriteLine($"For part 1: {part1Position.horizontal * part1Position.depth}");
            Console.WriteLine($"For part 2: {part2Position.horizontal * part2Position.depth}");
        }

    }

}
