namespace Day8Namespace;

using TreeMap = List<List<int>>;

public class Day8 : IPuzzle
{
    public int Day => 8;
    public bool IsTest => false;

    public void Run(string input)
    {
        var map = ParseInput(input);

        Part1(map);
        Part2(map);
    }

    private void Part1(TreeMap map)
    {
        int height = map.Count;
        int width = map[0].Count;

        HashSet<string> visibilityMap = new();

        for (int x = 0; x < width; x++)
        {
            Scan(map, x, 0, height, 0, +1, visibilityMap);
            Scan(map, x, height - 1, height, 0, -1, visibilityMap);
        }

        for (int y = 0; y < height; y++)
        {
            Scan(map, 0, y, width, +1, 0, visibilityMap);
            Scan(map, width - 1, y, width, -1, 0, visibilityMap);
        }

        Console.WriteLine($"Part1: {visibilityMap.Count}");
    }

    private void Part2(TreeMap map)
    {
	    int bestViewingDistance = 0;

	    int height = map.Count;
	    int width = map[0].Count;

	    for (int x = 0; x < width; x++)
        {
            for (int y = 0; y < height; y++)
            {
                int viewingDistance = 1;
                viewingDistance *= ComputeViewingDistance(map, x, y, width, height, -1, 0);
                viewingDistance *= ComputeViewingDistance(map, x, y, width, height, +1, 0);
                viewingDistance *= ComputeViewingDistance(map, x, y, width, height, 0, -1);
                viewingDistance *= ComputeViewingDistance(map, x, y, width, height, 0, +1);

                if (viewingDistance > bestViewingDistance)
                    bestViewingDistance = viewingDistance;
            }
        }

        Console.WriteLine($"Part2: {bestViewingDistance}");
    }

    private int ComputeViewingDistance(TreeMap map, int x, int y, int width, int height, int dx, int dy)
    {
        int viewingDistance = 0;
        int consideringTreeHeight = map[y][x];

        while (true)
        {
            x += dx;
            y += dy;

            if (x < 0 || y < 0 || x >= width || y >= height)
                break;

            int scanningTreeHeight = map[y][x];
            viewingDistance++;

            if (scanningTreeHeight >= consideringTreeHeight)
                break;
        }

        return viewingDistance;
    }

    private void Scan(TreeMap map, int startX, int startY, int count, int dx, int dy, HashSet<string> visibilityMap)
    {
        int largetTreeHeight = -1;

        int x = startX;
        int y = startY;

        while (count > 0)
        {
            int treeHeight = map[y][x];

            if (treeHeight > largetTreeHeight)
            {
                visibilityMap.Add($"{x}-{y}");
                largetTreeHeight = treeHeight;
            }

            if (treeHeight == 9)
                break;

            count--;

            x += dx;
            y += dy;
        }
    }

    private TreeMap ParseInput(string input)
    {
        TreeMap result = new();

        foreach (string line in input.Split('\n').Select(x => x.Trim()))
        {
            if (line.Length == 0)
                break;

            List<int> treeLine = new();

            foreach (char c in line)
                treeLine.Add((int)(c - '0'));

            result.Add(treeLine);
        }

        return result;
    }
}
