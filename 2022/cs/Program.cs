using System.Text.RegularExpressions;

public interface IPuzzle
{
    void Run(string input);
}

class Program
{
    private static readonly Regex TypeNameRegex = new Regex(@"Day(\d{2})", RegexOptions.Compiled);

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

    private static string ReadInput<T>()
    {
        Match m = TypeNameRegex.Match(typeof(T).Name);

        if (m.Success == false)
            throw new InvalidProgramException("You solution class name must start with 'Day' and be follow by 2 digits representing the day of the AoC problem.");

        string dayNumberStr = m.Groups[1].Value;

        string absoluteInputFilename = FindFile(AppContext.BaseDirectory, $"{dayNumberStr}.txt");

        return File.ReadAllText(absoluteInputFilename);
    }

    private static void Run<T>() where T : IPuzzle, new()
    {
        new T().Run(ReadInput<T>());
    }

    private static void Main(string[] args)
    {
        Run<Day01>();
    }
}
