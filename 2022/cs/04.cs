namespace Day4Namespace;

public record struct Range(int Start, int End)
{
    public bool IsFullyContaining(Range other)
    {
        return other.Start >= Start && other.End <= End;
    }

    public bool IsOverlapping(Range other)
    {
        int start = Math.Max(Start, other.Start);
        int end = Math.Min(End, other.End);

        return start <= end;
    }

    public static Range Parse(string value)
    {
        string[] parts = value.Trim().Split('-');

        return new Range(
            int.Parse(parts[0].Trim()),
            int.Parse(parts[1].Trim())
        );
    }
}

class Day4 : IPuzzle
{
    public int Day => 4;
    public bool IsTest => false;

    public void Run(string input)
    {
        List<(Range lhs, Range rhs)> assignmentPairList = CreateAssignmentPairList(input);

        int result1 = 0;
        int result2 = 0;

        foreach ((Range lhs, Range rhs) in assignmentPairList)
        {
            bool isFullyContained = lhs.IsFullyContaining(rhs) || rhs.IsFullyContaining(lhs);

            if (isFullyContained)
                result1++;

            if (isFullyContained || lhs.IsOverlapping(rhs))
                result2++;
        }

        Console.WriteLine($"Part1: {result1}");
        Console.WriteLine($"Part2: {result2}");
    }

    private static List<(Range lhs, Range rhs)> CreateAssignmentPairList(string input)
    {
        var result = new List<(Range lhs, Range rhs)>();

        foreach (string line in input.Split('\n').Select(x => x.Trim()))
        {
            if (line == string.Empty)
                continue;

            string[] pairParts = line.Split(',');

            result.Add((Range.Parse(pairParts[0]), Range.Parse(pairParts[1])));
        }

        return result;
    }
}
