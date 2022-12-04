namespace Day1Namespace;

class Day1 : IPuzzle
{
    public int Day { get; } = 1;

    public void Run(string input)
    {
        List<int> calories = new();

        int lineNumber = -1;
        int currentTotalCaloriesCount = 0;

        foreach (string line in input.Split('\n').Select(x => x.Trim()))
        {
            lineNumber++;

            if (line == string.Empty)
            {
                if (currentTotalCaloriesCount > 0)
                {
                    calories.Add(currentTotalCaloriesCount);
                    currentTotalCaloriesCount = 0;
                }

                continue;
            }

            if (int.TryParse(line, out int currentCaloriesCount) == false)
                throw new InvalidDataException($"Invalid numeric value at line {lineNumber}.");

            currentTotalCaloriesCount += currentCaloriesCount;
        }

        if (currentTotalCaloriesCount > 0)
            calories.Add(currentTotalCaloriesCount);

        if (calories.Count == 0)
            throw new InvalidOperationException("No entries recorded.");

        calories.Sort((a, b) => b - a);

        Console.WriteLine($"The Elf with the most calories has {calories[0]} calories.");

        Console.WriteLine($"The total calories of the top three Elves is {calories.Take(3).Sum()} calories.");
    }
}
