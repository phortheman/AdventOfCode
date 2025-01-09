
using System.IO;

namespace Solver
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

    public class Solver
    {
        public Solver() { }

        public static List<Instruction> GenerateInput(String filePath)
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

        private static string GetInputFilePath(string[] args)
        {
            for (int i = 0; i < args.Length; i++)
            {
                if (args[i] == "-i" && i + 1 < args.Length)
                {
                    string inputPath = args[i + 1];
                    if (File.Exists(inputPath))
                    {
                        return inputPath;
                    }
                    else
                    {
                        throw new FileNotFoundException($"The specified file does not exist: {inputPath}");
                    }
                }
            }

            string defaultPath = Path.Combine("..", "..", "..", "inputs", "2021", "02", "input.txt");

            if (!File.Exists(defaultPath))
            {
                throw new FileNotFoundException($"The default input file was not found at {defaultPath}");
            }

            return defaultPath;
        }

        static void Main(string[] args)
        {
            string inputFilePath = GetInputFilePath(args);
            List<Instruction> inputData = GenerateInput(inputFilePath);

            Position part1Position = new();
            Position part2Position = new();

            foreach (Instruction instruction in inputData)
            {
                part1Position.RunInstruction(instruction);
                part2Position.RunWithAim(instruction);
            }

            Console.WriteLine($"Part One: {part1Position.horizontal * part1Position.depth}");
            Console.WriteLine($"Part Two: {part2Position.horizontal * part2Position.depth}");
        }

    }

}