
using System.IO;

namespace Day_01
{

    public class Day_01_Solver
    {
        public Day_01_Solver(){}

        public List<int> GenerateInput(String filePath)
        {
            List<int> output = new List<int>();

            using (StreamReader reader = new StreamReader(filePath))
            {
                String? line;
                while ( (line = reader.ReadLine() ) != null)
                {
                    output.Add(int.Parse(line));
                }
            }

            return output;
        }

        public int GetNumberOfIncreases(List<int> input, int lastIndex = 0)
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

        public List<int> GetWindow(List<int> input)
        {
            List<int> window = new List<int>();

            for (int i = 0; i < input.Count - 2; i++)
            {
                window.Add(input[i] + input[i + 1] + input[i + 2]);
            }

            return window;
        }
        static void Main(string[] args)
        {
            if( args.Length == 0)
            {
                Console.WriteLine("Please provide the puzzle input as a file");
                return;
            }

            List<int> inputData;

            Day_01_Solver solver = new Day_01_Solver();

            inputData = solver.GenerateInput(args[0]);

            int partOne = solver.GetNumberOfIncreases(inputData);
            int partTwo = solver.GetNumberOfIncreases(solver.GetWindow(inputData));

            Console.WriteLine($"The number of increases with the input is {partOne}");
            Console.WriteLine($"The number of increases with the window is {partTwo}");
        }

    }

}
