public interface IPuzzle
{
    int Day { get; }
    void Run(string input);
}

class Program
{
    private static string FindFile(string? basePath, string filename)
    {
        while (true)
        {
            if (basePath == null)
                throw new IOException($"Could not find file '{filename}' anywhere.");

            string absoluteFilename = Path.Combine(basePath, filename);

            if (File.Exists(absoluteFilename))
                return absoluteFilename;

            basePath = Path.GetDirectoryName(basePath);
        }
    }

    private static string ReadInput<T>(T puzzle) where T : IPuzzle
    {
        string filename = $"{puzzle.Day.ToString().PadLeft(2, '0')}.txt";
        string absoluteInputFilename = FindFile(AppContext.BaseDirectory, filename);

        return File.ReadAllText(absoluteInputFilename);
    }

    private static void Run<T>(T puzzle) where T : IPuzzle
    {
        puzzle.Run(ReadInput<T>(puzzle));
    }

    private static void Main(string[] args)
    {
        Run(new Day5Namespace.Day5());
    }
}
