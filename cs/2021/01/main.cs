
using System.IO;

namespace Solver
{

    public class Solver
    {
        public static List<int> GenerateInput(String filePath)
        {
            List<int> output = new();

            if (!File.Exists(filePath))
            {
                throw new FileNotFoundException($"The file at {filePath} was not found.");
            }

            using (StreamReader reader = new(filePath))
            {
                String? line;
                while ((line = reader.ReadLine()) != null)
                {
                    output.Add(int.Parse(line));
                }
            }

            return output;
        }

        public static int GetNumberOfIncreases(List<int> input, int lastIndex = 0)
        {
            if (lastIndex == 0)
            {
                lastIndex = input.Count;
            }
            int numOfIncreases = 0;
            for (int i = 1; i < lastIndex; i++)
            {
                if (input[i - 1] < input[i])
                {
                    numOfIncreases++;
                }
            }
            return numOfIncreases;
        }

        public static List<int> GetWindow(List<int> input)
        {
            List<int> window = new();

            for (int i = 0; i < input.Count - 2; i++)
            {
                window.Add(input[i] + input[i + 1] + input[i + 2]);
            }

            return window;
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

            string defaultPath = Path.Combine("..", "..", "..", "inputs", "2021", "01", "input.txt");

            if (!File.Exists(defaultPath))
            {
                throw new FileNotFoundException($"The default input file was not found at {defaultPath}");
            }

            return defaultPath;
        }

        static void Main(string[] args)
        {
            try
            {
                string inputFilePath = GetInputFilePath(args);
                List<int> inputData = GenerateInput(inputFilePath);

                int partOne = GetNumberOfIncreases(inputData);
                int partTwo = GetNumberOfIncreases(GetWindow(inputData));

                Console.WriteLine($"Part One: {partOne}");
                Console.WriteLine($"Part Two: {partTwo}");
            }
            catch (Exception ex)
            {
                Console.WriteLine($"An error occurred: {ex.Message}");
            }
        }

    }

}