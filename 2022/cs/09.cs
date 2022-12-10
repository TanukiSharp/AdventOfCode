namespace Day9Namespace;

public class Day9 : IPuzzle
{
    public int Day => 9;
    public bool IsTest => false;

    public void Run(string input)
    {
        List<Movement> movements = ParseInput(input);

        Console.WriteLine($"Part1: {Simulate(movements, 2)}");
        Console.WriteLine($"Part2: {Simulate(movements, 10)}");
    }

    private int Simulate(List<Movement> movements, int ropeSize)
    {
        var head = new Movable(null, ropeSize);
        Movable tail = head.FindTail();

        foreach (Movement movement in movements)
            ApplyMovement(head, movement);

        return tail.Visited!.Count;
    }

    private void ApplyMovement(Movable head, Movement movement)
    {
        for (int i = 0; i < movement.Units; i++)
            head.Move(movement.DeltaX, movement.DeltaY);
    }

    private List<Movement> ParseInput(string input)
    {
        return input
            .Split('\n')
            .Select(x => x.Trim())
            .Where(x => x.Length > 0)
            .Select(ParseLine)
            .ToList();
    }

    private Movement ParseLine(string line)
    {
        string[] parts = line.Split(' ');
        int units = int.Parse(parts[1]);

        switch (parts[0])
        {
            case "L":
                return new Movement(-1, 0, units);
            case "U":
                return new Movement(0, -1, units);
            case "R":
                return new Movement(+1, 0, units);
            case "D":
                return new Movement(0, +1, units);
        }

        throw new InvalidDataException($"Unknown direction '{parts[0]}'.");
    }

    private record struct Movement(int DeltaX, int DeltaY, int Units);

    private void Print(Movable head, bool visitedOnly)
    {
        int minX = 0;
        int minY = 0;
        int maxX = 0;
        int maxY = 0;

        Movable tail = head.FindTail();

        foreach (Coord coord in tail.Visited!)
        {
            minX = Math.Min(Math.Min(minX, coord.X), head.X);
            minY = Math.Min(Math.Min(minY, coord.Y), head.Y);
            maxX = Math.Max(Math.Max(maxX, coord.X), head.X);
            maxY = Math.Max(Math.Max(maxY, coord.Y), head.Y);
        }

        FindRopeBoundingBox(head, ref minX, ref minY, ref maxX, ref maxY);

        int width = Math.Abs(maxX) + Math.Abs(minX) + 1;
        int height = Math.Abs(maxY) + Math.Abs(minY) + 1;

        char[][] entries = new char[height][];

        for (int i = 0; i < height; i++)
            entries[i] = new char[width];

        for (int y = minY, ry = 0; y <= maxY; y++, ry++)
        {
            for (int x = minX, rx = 0; x <= maxX; x++, rx++)
            {
                if (visitedOnly == false && TryFindRopeNodeId(head, x, y, out char? id))
                    entries[ry][rx] = id!.Value;
                else if (x == 0 && y == 0 && visitedOnly == false)
                    entries[ry][rx] = 's';
                else if (tail!.Visited.Contains(new Coord(x, y)))
                    entries[ry][rx] = '#';
                else
                    entries[ry][rx] = '.';
            }
            Console.WriteLine(new string(entries[ry]));
        }

        Console.WriteLine();
        Console.WriteLine();
        Console.WriteLine();
    }

    private void FindRopeBoundingBox(Movable node, ref int minX, ref int minY, ref int maxX, ref int maxY)
    {
        minX = Math.Min(minX, node.X);
        minY = Math.Min(minY, node.Y);
        maxX = Math.Max(maxX, node.X);
        maxY = Math.Max(maxY, node.Y);

        if (node.Next != null)
            FindRopeBoundingBox(node.Next, ref minX, ref minY, ref maxX, ref maxY);
    }

    private bool TryFindRopeNodeId(Movable? node, int x, int y, out char? id)
    {
        if (node == null)
        {
            id = null;
            return false;
        }

        if (node.X == x && node.Y == y)
        {
            id = node.ID;
            return true;
        }

        return TryFindRopeNodeId(node.Next, x, y, out id);
    }
}

public record struct Coord(int X, int Y);

public class Movable
{
    public int X { get; private set; }
    public int Y { get; private set; }

    public Movable? Previous { get; }
    public Movable? Next { get; }

    public HashSet<Coord>? Visited { get; }

    public char ID { get; }

    public Movable(Movable? previous, int ropeSize)
        : this(previous, ropeSize, ropeSize)
    {
    }

    private Movable(Movable? previous, int totalRopeSize, int ropeSize)
    {
        if (previous == null)
            ID = 'H';
        else
            ID = (char)((totalRopeSize - ropeSize) + '0');

        Previous = previous;

        if (ropeSize > 1)
            Next = new Movable(this, totalRopeSize, ropeSize - 1);
        else
        {
            Visited = new();
            Visited.Add(new Coord(0, 0));
        }
    }

    public Movable FindTail()
    {
        if (Next == null)
            return this;
        return Next.FindTail();
    }

    public void Move(int dx, int dy)
    {
        X += dx;
        Y += dy;
        Next?.PreviousMoved();
    }

    public void PreviousMoved()
    {
        if (IsAttachedToPrevious() == false)
        {
            int dx = Previous!.X - X;
            int dy = Previous!.Y - Y;

            Move(Math.Sign(dx), Math.Sign(dy));
        }

        if (Next == null)
            Visited!.Add(new Coord(X, Y));
    }

    private bool IsAttachedToPrevious()
    {
        int dx = Math.Abs(Previous!.X - X);
        int dy = Math.Abs(Previous!.Y - Y);

        return dx <= 1 && dy <= 1;
    }

    public override string ToString()
    {
        return $"{ID} ({X}, {Y})";
    }
}
