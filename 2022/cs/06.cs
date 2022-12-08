namespace Day6Namespace;

public class Day6 : IPuzzle
{
    public int Day => 6;
    public bool IsTest => false;

    public void Run(string input)
    {
        int part1 = FindStartOfPacket(input, 4);
        Console.WriteLine($"Part1: {part1}");

        int part2 = FindStartOfPacket(input, 14);
        Console.WriteLine($"Part2: {part2}");
    }

    private static int FindStartOfPacket(string input, int packetSize)
    {
        int length = input.Length - packetSize + 1;

        for (int i = 0; i < length; i++)
        {
            ReadOnlySpan<char> span = input.AsSpan(i, packetSize);
            if (IsStartOfPacket(span))
                return i + packetSize;
        }

        return -1;
    }

    private static readonly HashSet<char> tester = new();

    private static bool IsStartOfPacket(ReadOnlySpan<char> value)
    {
        tester.Clear();

        foreach (char c in value)
        {
            if (tester.Add(c) == false)
                return false;
        }

        return true;
    }
}
