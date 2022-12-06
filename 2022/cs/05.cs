using System.Text.RegularExpressions;

namespace Day5Namespace;

record struct Movement(int SourceStackIndex, int TargetStackIndex, int Quantity);

class Day5 : IPuzzle
{
    public int Day { get; } = 5;

    private readonly List<Stack<char>> part1CrateStacks = new();
    private readonly List<Stack<char>> part2CrateStacks = new();
    private readonly List<Movement> movements = new();

    public void Run(string input)
    {
        ParseInput(input);

        Part1();
        Part2();
    }

    private void Part1()
    {
        foreach (Movement movement in movements)
        {
            for (int i = 0; i < movement.Quantity; i++)
            {
                char c = part1CrateStacks[movement.SourceStackIndex].Pop();
                part1CrateStacks[movement.TargetStackIndex].Push(c);
            }
        }

        Console.Write("Part1: ");

        foreach (Stack<char> crateStack in part1CrateStacks)
            Console.Write(crateStack.Peek());

        Console.WriteLine();
    }

    private void Part2()
    {
        Stack<char> tempStack = new();

        foreach (Movement movement in movements)
        {
            for (int i = 0; i < movement.Quantity; i++)
            {
                char c = part2CrateStacks[movement.SourceStackIndex].Pop();
                tempStack.Push(c);
            }

            while (tempStack.Count > 0)
                part2CrateStacks[movement.TargetStackIndex].Push(tempStack.Pop());
        }

        Console.Write("Part2: ");

        foreach (Stack<char> crateStack in part2CrateStacks)
            Console.Write(crateStack.Peek());

        Console.WriteLine();
    }

    private static void ParseCrateLine(string line, List<Stack<char>> crateStacks)
    {
        if (line.StartsWith(" 1 "))
            return;

        int stackIndex = 0;
        int lineIndex = 0;

        while (lineIndex < line.Length)
        {
            if (stackIndex >= crateStacks.Count)
                crateStacks.Add(new Stack<char>());

            if (line[lineIndex] == '[' && line[lineIndex + 2] == ']')
            {
                char c = line[lineIndex + 1];
                crateStacks[stackIndex].Push(c);
            }

            stackIndex++;
            lineIndex += 4;
        }
    }

    private static Regex MovementRegex = new Regex(@"move (\d+) from (\d+) to (\d+)", RegexOptions.Compiled);

    private void ParseMovementLine(string line)
    {
        Match m = MovementRegex.Match(line);

        int quantity = int.Parse(m.Groups[1].Value);
        int from = int.Parse(m.Groups[2].Value);
        int to = int.Parse(m.Groups[3].Value);

        movements.Add(new Movement(from - 1, to - 1, quantity));
    }

    private void ParseInput(string input)
    {
        List<Stack<char>> localCrateStacks = new();

        bool isParsingCrates = true;

        foreach (string line in input.Split('\n').Select(x => x.TrimEnd()))
        {
            if (line == string.Empty)
            {
                isParsingCrates = false;
                continue;
            }

            if (isParsingCrates)
                ParseCrateLine(line, localCrateStacks);
            else
                ParseMovementLine(line);
        }

        // Reconstruct stacks in proper order.
        foreach (Stack<char> localCrateStack in localCrateStacks)
        {
            Stack<char> crateStack = new();

            while (localCrateStack.Count > 0)
                crateStack.Push(localCrateStack.Pop());

            part1CrateStacks.Add(crateStack);
            part2CrateStacks.Add(new Stack<char>(crateStack.Reverse()));
        }
    }
}
