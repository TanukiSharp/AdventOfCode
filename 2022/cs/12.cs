namespace Day12Namespace;

public class Day12 : IPuzzle
{
    public int Day => 12;
    public bool IsTest => false;

    public void Run(string input)
    {
        MapInfo mapInfo = ParseInput(input);

        Part1(mapInfo);
    }

    public static readonly Coord[] Deltas = new Coord[] { new(-1, 0), new(0, -1), new(+1, 0), new(0, +1) };

    private void Part1(MapInfo mapInfo)
    {
        Part1BFS(mapInfo);
        Part2BFS(mapInfo);

        //Part1Dijkstra(mapInfo);
        //Part1AStar(mapInfo);
    }

    private void Part1BFS(MapInfo mapInfo)
    {
        var bfs = new BFS(CreateCanMoveToFunction(mapInfo, false));
        bfs.Run(mapInfo.Start, coord => coord == mapInfo.End);

        List<Coord>? solution = BFS.ConstuctPath(bfs.EndNode);

        Print(solution!, mapInfo);

        Console.WriteLine();
        Console.WriteLine($"Part1: {solution!.Count - 1}");
    }

    private void Part2BFS(MapInfo mapInfo)
    {
        var bfs = new BFS(CreateCanMoveToFunction(mapInfo, true));
        bfs.Run(mapInfo.End, coord =>
        {
            int elevation = mapInfo.HeightMap[coord.X, coord.Y];
            return elevation == 0;
        });

        List<Coord>? solution = BFS.ConstuctPath(bfs.EndNode);

        Print(solution, mapInfo);

        Console.WriteLine();
        Console.WriteLine($"Part2: {solution!.Count - 1}");
    }

    private void FindShortestPathToB(MapInfo mapInfo, BFSNode node, int currentLength, ref int pathLength)
    {
        if (node.Children == null)
            return;

        int elevation = mapInfo.HeightMap[node.Coord.X, node.Coord.Y];

        if (elevation == 1)
        {
            pathLength = Math.Min(pathLength, currentLength);
            return;
        }

        if (elevation < 1)
            throw new InvalidOperationException("Not supposed to go that far.");

        foreach (BFSNode child in node.Children)
            FindShortestPathToB(mapInfo, child, currentLength - 1, ref pathLength);
    }

    private void Part1Dijkstra(MapInfo mapInfo)
    {
        Coord start = mapInfo.End;
        Coord end = mapInfo.Start;

        var dijkstra = new Dijkstra(mapInfo.Size.X, mapInfo.Size.Y, CreateCanMoveToFunction(mapInfo, false));

        List<Coord> solution = dijkstra.Run(start, end);

        Print(solution, mapInfo);

        Console.WriteLine();
        Console.WriteLine($"Part1: {solution.Count}");
    }

    private void Part1AStar(MapInfo mapInfo)
    {
        Func<Coord, double> h = coord =>
        {
            int dx = coord.X - mapInfo.End.X;
            int dy = coord.Y - mapInfo.End.Y;
            return Math.Sqrt(dx * dx + dy * dy);
        };

        var aStar = new AStar(h, CreateCanMoveToFunction(mapInfo, false));
        List<Coord>? solution = aStar.Run(mapInfo.Start, mapInfo.End);

        Print(solution!, mapInfo);

        Console.WriteLine();
        Console.WriteLine($"Part1: {solution!.Count - 1}");
    }

    private Func<Coord, Coord, bool> CreateCanMoveToFunction(MapInfo mapInfo, bool isReversed)
    {
        return (from, to) =>
        {
            if (to.X < 0 || to.X >= mapInfo.Size.X || to.Y < 0 || to.Y >= mapInfo.Size.Y)
                return false;

            int previousElevation = mapInfo.HeightMap[from.X, from.Y];
            int currentElevation = mapInfo.HeightMap[to.X, to.Y];

            int diff = currentElevation - previousElevation;

            if (isReversed == false)
                return diff <= 1;

            return diff >= -1;
        };
    }

    private void Print(List<Coord> solution, MapInfo mapInfo)
    {
        for (int y = 0; y < mapInfo.Size.Y; y++)
        {
            for (int x = 0; x < mapInfo.Size.X; x++)
            {
                char c = (char)(mapInfo.HeightMap[x, y] + 'a');
                if (solution.Contains(new(x, y)))
                    Console.ForegroundColor = ConsoleColor.Red;

                if (x == mapInfo.Start.X && y == mapInfo.Start.Y ||
                    x == mapInfo.End.X && y == mapInfo.End.Y)
                    c = (char)(c - 'a' + 'A');

                Console.Write(c);
                Console.ResetColor();
            }
            Console.WriteLine();
        }
    }

    private MapInfo ParseInput(string input)
    {
        List<List<int>> map = new();

        int startX = -1;
        int startY = -1;

        int endX = -1;
        int endY = -1;

        int y = -1;
        foreach (string line in input.Split('\n').Select(x => x.Trim()))
        {
            if (line.Length == 0)
                break;

            y++;

            List<int> lineHeights = new();

            for (int x = 0; x < line.Length; x++)
            {
                char c = line[x];

                if (c == 'S')
                {
                    startX = x;
                    startY = y;
                    c = 'a';
                }
                else if (c == 'E')
                {
                    endX = x;
                    endY = y;
                    c = 'z';
                }

                lineHeights.Add(c - 'a');
            }

            map.Add(lineHeights);
        }

        int width = map[0].Count;
        int height = map.Count;

        int[,] heightMap = new int[width, height];

        for (int yy = 0; yy < height; yy++)
        {
            for (int xx = 0; xx < width; xx++)
                heightMap[xx, yy] = map[yy][xx];
        }

        return new MapInfo(
            new(startX, startY),
            new(endX, endY),
            new(width, height),
            heightMap
        );
    }
}

public record struct Coord(int X, int Y)
{
    public override string ToString()
    {
        return $"{X}, {Y}";
    }
}

public record struct MapInfo(Coord Start, Coord End, Coord Size, int[,] HeightMap);

public class BFSNode
{
    public Coord Coord { get; }

    private BFSNode? parent;
    public BFSNode? Parent
    {
        get => parent;
        set
        {
            if (parent != null)
                throw new InvalidOperationException("Cannot reassign parent.");
            parent = value;
            if (parent != null)
            {
                if (parent.Children == null)
                    parent.Children = new();
                parent.Children.Add(this);
            }
        }
    }

    public HashSet<BFSNode>? Children { get; private set; }

    public BFSNode(Coord coord)
    {
        Coord = coord;
    }

    public override int GetHashCode()
    {
        return Coord.GetHashCode();
    }

    public override bool Equals(object? obj)
    {
        if (obj is BFSNode n)
            return n.Coord == Coord;
        return false;
    }

    public override string ToString()
    {
        return Coord.ToString();
    }
}

public class BFS
{
    private readonly Func<Coord, Coord, bool> canMoveTo;

    private readonly HashSet<BFSNode> explored = new();
    private readonly Queue<BFSNode> queue = new();

    public BFSNode? StartNode { get; private set; }
    public BFSNode? EndNode { get; private set; }

    public BFS(Func<Coord, Coord, bool> canMoveTo)
    {
        this.canMoveTo = canMoveTo;
    }

    public void Run(Coord start, Predicate<Coord> stop)
    {
        BFSNode startNode = new BFSNode(start);

        StartNode = startNode;

        explored.Add(startNode);
        queue.Enqueue(startNode);

        while (queue.Count > 0)
        {
            BFSNode v = queue.Dequeue();

            if (stop(v.Coord))
            {
                EndNode = v;
                return;
            }

            foreach (Coord delta in Day12.Deltas)
            {
                var toCoord = new Coord(v.Coord.X + delta.X, v.Coord.Y + delta.Y);

                if (canMoveTo(v.Coord, toCoord))
                    Next(v, new BFSNode(toCoord));
            }
        }
    }

    public static List<Coord> ConstuctPath(BFSNode? root)
    {
        if (root == null)
            return new();

        List<Coord> result = new();

        BFSNode? current = root;

        while (current != null)
        {
            result.Add(current.Coord);
            current = current.Parent;
        }

        return result;
    }

    private void Next(BFSNode v, BFSNode w)
    {
        if (explored.Add(w) == false)
            return;

        w.Parent = v;
        queue.Enqueue(w);
    }
}

public class AStar
{
    private readonly Func<Coord, double> h;
    private readonly Func<Coord, Coord, bool> canMoveTo;

    private readonly HashSet<Coord> openSet = new();
    private readonly Dictionary<Coord, Coord> cameFrom = new();
    private readonly Dictionary<Coord, double> gScore = new();
    private readonly Dictionary<Coord, double> fScore = new();

    public AStar(Func<Coord, double> h, Func<Coord, Coord, bool> canMoveTo)
    {
        this.h = h;
        this.canMoveTo = canMoveTo;
    }

    public List<Coord>? Run(Coord start, Coord end)
    {
        openSet.Add(start);

        gScore[start] = 0.0;
        fScore[start] = h(start);

        while (openSet.Count > 0)
        {
            Coord current = FindNodeWithBestScore(openSet, fScore);

            if (current == end)
                return ReconstructPath(current);

            openSet.Remove(current);

            foreach (Coord delta in Day12.Deltas)
            {
                Coord next = new(current.X + delta.X, current.Y + delta.Y);

                if (canMoveTo(current, next))
                    Next(current, next);
            }
        }

        return null;
    }

    private void Next(Coord current, Coord neighbor)
    {
        double tentativeScore = gScore[current];

        double neighborScore = GetScore(neighbor, gScore);

        if (tentativeScore >= neighborScore)
            return;

        cameFrom[neighbor] = current;

        gScore[neighbor] = tentativeScore;
        fScore[neighbor] = tentativeScore + h(neighbor);

        openSet.Add(neighbor);
    }

    private static double GetScore(Coord coord, Dictionary<Coord, double> scores)
    {
        if (scores.TryGetValue(coord, out double score))
            return score;
        return double.PositiveInfinity;
    }

    private List<Coord> ReconstructPath(Coord current)
    {
        List<Coord> totalPath = new() { current };

        while (cameFrom.ContainsKey(current))
        {
            current = cameFrom[current];
            totalPath.Insert(0, current);
        }

        return totalPath;
    }

    private Coord FindNodeWithBestScore(HashSet<Coord> openSet, Dictionary<Coord, double> fScore)
    {
        List<(Coord coord, double score)> temp = new();

        foreach (Coord openSetCoord in openSet)
        {
            if (fScore.TryGetValue(openSetCoord, out double score))
                temp.Add((openSetCoord, score));
        }

        var (coord, _) = temp.OrderBy(x => x.score).First();

        return coord;
    }
}

public class Dijkstra
{
    private readonly Func<Coord, Coord, bool> canMoveTo;

    private readonly HashSet<Coord> nodes = new();
    private readonly Dictionary<Coord, double> dist = new();
    private readonly Dictionary<Coord, Coord> prev = new();

    public Dijkstra(int width, int height, Func<Coord, Coord, bool> canMoveTo)
    {
        this.canMoveTo = canMoveTo;

        Initialize(width, height);
    }

    private void Initialize(int width, int height)
    {
        for (int y = 0; y < height; y++)
        {
            for (int x = 0; x < width; x++)
            {
                Coord node = new Coord(x, y);
                dist[node] = double.PositiveInfinity;
                nodes.Add(node);
            }
        }
    }

    public List<Coord> Run(Coord startNode, Coord targetNode)
    {
        dist[startNode] = 0.0;

        while (nodes.Count > 0)
        {
            Coord u = FindNodeWithMinDistance();
            nodes.Remove(u);

            foreach (Coord delta in Day12.Deltas)
            {
                Coord v = new(u.X + delta.X, u.Y + delta.Y);

                if (canMoveTo(u, v) == false || nodes.Contains(v) == false)
                    continue;

                double alt = dist[u];

                if (alt < dist[v])
                {
                    dist[v] = alt;
                    prev[v] = u;
                }
            }
        }

        return FindShortestPath(targetNode);
    }

    private List<Coord> FindShortestPath(Coord targetNode)
    {
        List<Coord> s = new();

        Coord u = targetNode;

        while (prev.TryGetValue(u, out Coord v))
        {
            s.Insert(0, u);
            u = v;
        }

        return s;
    }

    private Coord FindNodeWithMinDistance()
    {
        List<(Coord node, double distance)> temp = new();

        foreach (Coord node in nodes)
        {
            if (dist.TryGetValue(node, out double distance))
                temp.Add((node, distance));
        }

        var (bestNode, _) = temp.OrderBy(x => x.distance).First();

        return bestNode;
    }
}
