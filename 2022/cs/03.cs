class Day03 : IPuzzle
{
    public int Day { get; } = 3;

    private record struct Rucksack(IEnumerable<char> Compartment1, IEnumerable<char> Compartment2)
    {
        public char FindDuplicateItem()
        {
            return Compartment1
                .Intersect(Compartment2)
                .Single();
        }

        public HashSet<char> GetUniqueItems()
        {
            return Compartment1
                .Concat(Compartment2)
                .ToHashSet();
        }
    }

    public void Run(string input)
    {
        List<Rucksack> rucksacks = CreateRucksackList(input);
        Part1(rucksacks);
        Part2(rucksacks);
    }

    private static void Part1(List<Rucksack> rucksacks)
    {
        int result = 0;

        foreach (Rucksack rucksack in rucksacks)
        {
            char duplicate = rucksack.FindDuplicateItem();
            result += GetItemPriority(duplicate);
        }

        Console.WriteLine($"Part1: {result}");
    }

    private static void Part2(List<Rucksack> rucksacks)
    {
        int priorities = 0;

        for (int i = 0; i < rucksacks.Count; i += 3)
        {
            char badge = FindBadge(rucksacks, i);
            priorities += GetItemPriority(badge);
        }

        Console.WriteLine($"Part2: {priorities}");
    }

    private static char FindBadge(List<Rucksack> rucksacks, int index)
    {
        const int groupSize = 3;
        List<HashSet<char>> allPossibilities = new();

        for (int i = 0; i < groupSize; i++)
            allPossibilities.Add(rucksacks[index + i].GetUniqueItems());

        Dictionary<char, int> common = new();

        foreach (HashSet<char> possibilities in allPossibilities)
        {
            foreach (char item in possibilities)
            {
                common.TryGetValue(item, out int count);
                common[item] = count + 1;
            }
        }

        foreach (KeyValuePair<char, int> kv in common)
        {
            if (kv.Value == 3)
                return kv.Key;
        }

        throw new InvalidDataException("Could not find badge");
    }

    private static int GetItemPriority(char c)
    {
        if (c >= 'a' && c <= 'z')
            return c - 'a' + 1;

        if (c >= 'A' && c <= 'Z')
            return c - 'A' + 27;

        throw new InvalidDataException($"Unknown item '{c}'.");
    }

    private static List<Rucksack> CreateRucksackList(string input)
    {
        var result = new List<Rucksack>();

        foreach (string line in input.Split('\n').Select(x => x.Trim()))
        {
            if (line == string.Empty)
                continue;

            result.Add(
                new Rucksack(
                    line.Substring(0, line.Length / 2),
                    line.Substring(line.Length / 2)
                )
            );
        }

        return result;
    }
}
