namespace Day7Namespace;

public class Day7 : IPuzzle
{
    public int Day => 7;
    public bool IsTest => false;

    public void Run(string input)
    {
        ParseInput(input);

        Part1();
        Part2();
    }

    private void Part1()
    {
        Console.WriteLine($"Part1: {Part1Core(root)}");
    }

    private int Part1Core(MyDirectory dir)
    {
        int result = dir.TotalSize <= 100_000 ? dir.TotalSize : 0;

        foreach (MyDirectory sub in dir.Directories)
            result += Part1Core(sub);

        return result;
    }

    private void Part2()
    {
        const int maxSize = 70_000_000;
        const int neededForUpdate = 30_000_000;

        int availableSize = maxSize - root.TotalSize;
        int needToFree = neededForUpdate - availableSize;

        MyDirectory bestDirectory = root;
        FindDirectory(ref bestDirectory, root, needToFree);

        Console.WriteLine($"Part2: {bestDirectory.TotalSize}");
    }

    private void FindDirectory(ref MyDirectory best, MyDirectory current, int size)
    {
        foreach (MyDirectory dir in current.Directories)
        {
            if (dir.TotalSize >= size && dir.TotalSize < best.TotalSize)
                best = dir;
            FindDirectory(ref best, dir, size);
        }
    }

    private MyDirectory root = new MyDirectory("/", null);
    private MyDirectory? cwd;

    private void ParseInput(string input)
    {
        cwd = root;

        foreach (string line in input.Split('\n').Select(x => x.Trim()))
        {
            if (string.IsNullOrWhiteSpace(line))
                continue;

            if (line.StartsWith("$ "))
                ParseCommand(line.Substring(2));
            else
                ParseFsEntry(line);
        }
    }

    private void ParseCommand(string command)
    {
        string[] parts = command.Split(' ');

        if (parts[0] == "cd")
        {
            string dir = parts[1];

            if (dir == "/")
                cwd = root;
            else if (dir == "..")
                cwd = cwd!.Parent;
            else
                cwd = cwd!.EnsureSubDirectory(dir);
        }

        // Nothing to do for ls command.
    }

    private void ParseFsEntry(string line)
    {
        string[] parts = line.Split(' ');

        if (parts[0] == "dir")
            cwd!.EnsureSubDirectory(parts[1]);
        else
        {
            int size = int.Parse(parts[0]);
            cwd!.Files.Add(new MyFile(parts[1], size, cwd!));
        }
    }
}

public class MyDirectory
{
    public string Name { get; }
    public MyDirectory? Parent { get; }
    public int TotalSize { get; private set; }
    public List<MyDirectory> Directories { get; } = new();
    public List<MyFile> Files { get; } = new();

    public MyDirectory(string name, MyDirectory? parent)
    {
        Name = name;
        Parent = parent;
    }

    public MyDirectory EnsureSubDirectory(string name)
    {
        MyDirectory? existing = Directories.Find(x => x.Name == name);

        if (existing != null)
            return existing;

        MyDirectory newDir = new(name, this);
        Directories.Add(newDir);

        return newDir;
    }

    public void UpdateTotalSize(int deltaSize)
    {
        TotalSize += deltaSize;
        if (Parent != null)
            Parent.UpdateTotalSize(deltaSize);
    }
}

public class MyFile
{
    public MyDirectory Parent { get; }
    public string Name { get; }
    public int Size { get; }

    public MyFile(string name, int size, MyDirectory parent)
    {
        Name = name;
        Size = size;
        Parent = parent;

        Parent.UpdateTotalSize(size);
    }
}
