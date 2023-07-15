using System.IO;

class Day_01
{
    static void Main(string[] args)
    {
        List<int> input = GetInput();

        int partOne = GetNumberOfIncreases(input);
        int partTwo = GetNumberOfIncreases(GetWindow(input));

        Console.WriteLine($"The number of increases with the input is {partOne}");
        Console.WriteLine($"The number of increases with the window is {partTwo}");
    }

    static List<int> GetInput()
    {
        List<int> input = new List<int>();

        StreamReader reader = new StreamReader("01/input.txt");
        String line = reader.ReadLine();

        while (line != null)
        {
            input.Add(int.Parse(line));
            line = reader.ReadLine();
        }

        reader.Close();

        return input;
    }

    static int GetNumberOfIncreases(List<int> input, int lastIndex = 0)
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

    static List<int> GetWindow(List<int> input)
    {
        List<int> window = new List<int>();

        for (int i = 0; i < input.Count - 2; i++)
        {
            window.Add(input[i] + input[i + 1] + input[i + 2]);
        }

        return window;
    }
}