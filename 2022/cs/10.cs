using System.Text;

namespace Day10Namespace;

class Day10 : IPuzzle
{
    public int Day => 10;
    public bool IsTest => false;

    public void Run(string input)
    {
        List<InstructionBase> instructions = ParseInstructions(input);

        Part1(instructions);
        Part2(instructions);
    }

    private void Part1(List<InstructionBase> instructions)
    {
        Dictionary<string, int> registers = new()
        {
            ["X"] = 1,
        };

        var cpu = new CPU(instructions, registers);

        int totalSignalStrength = 0;

        foreach (int cycleToMonitor in CyclesToMonitor())
        {
            while (cpu.CycleStart())
            {
                try
                {
                    if (cpu.CycleNumber == cycleToMonitor)
                    {
                        int signalStrength = cpu.CycleNumber * registers["X"];
                        totalSignalStrength += signalStrength;
                        break;
                    }
                }
                finally
                {
                    cpu.CycleEnd();
                }
            }
        }

        Console.WriteLine($"Part1: {totalSignalStrength}");
    }

    private void Part2(List<InstructionBase> instructions)
    {
        foreach (InstructionBase instruction in instructions)
            instruction.Reset();

        Dictionary<string, int> registers = new()
        {
            ["X"] = 1,
        };

        var cpu = new CPU(instructions, registers);
        var crt = new CRT();

        while (cpu.CycleNumber < 240)
        {
            if (cpu.CycleStart() == false)
                break;

            crt.Render(cpu.CycleNumber, registers["X"]);

            cpu.CycleEnd();
        }

        Console.WriteLine("Part2:");
        crt.Swap();
    }

    private List<InstructionBase> ParseInstructions(string input)
    {
        List<InstructionBase> result = new();

        foreach (string line in input.Split('\n').Select(x => x.Trim()))
        {
            if (line.Length == 0)
                break;

            result.Add(ParseInstruction(line));
        }

        return result;
    }

    private InstructionBase ParseInstruction(string line)
    {
        string[] parts = line.Split(' ');

        switch (parts[0])
        {
            case "noop": return new NoopInstruction();
            case "addx": return new AddInstruction(int.Parse(parts[1]));
        }

        throw new InvalidDataException($"Unknown instruction '{parts[0]}'.");
    }

    private IEnumerable<int> CyclesToMonitor()
    {
        int value = 20;
        yield return value;
        while (value < 220)
        {
            value += 40;
            yield return value;
        }
    }
}

public class CRT
{
    private readonly StringBuilder backBuffer = new();

    private const int width = 40;
    private int lineCounter = 0;

    public void Render(int cycleNumber, int spritePosition)
    {
        // The -1 below is because rendering from cycle 1, not 0,
        // because the CPU needs to be ticked before rendering,
        // causing a cycle to pass.
        int drawPosition = (cycleNumber - 1) % width;

        if (drawPosition >= spritePosition - 1 && drawPosition <= spritePosition + 1)
            backBuffer.Append('#');
        else
            backBuffer.Append('.');

        lineCounter++;

        if (lineCounter == width)
        {
            lineCounter = 0;
            backBuffer.AppendLine();
        }
    }

    public void Swap()
    {
        Console.WriteLine(backBuffer);
    }
}

public class CPU
{
    private readonly List<InstructionBase> instructions;
    private readonly Dictionary<string, int> registers;

    private int pc;

    public int CycleNumber { get; private set; }

    public CPU(List<InstructionBase> instructions, Dictionary<string, int> registers)
    {
        if (instructions.Count == 0)
            throw new ArgumentException("There are no instructions to execute.");

        this.instructions = instructions;
        this.registers = registers;
    }

    public bool CycleStart()
    {
        if (pc >= instructions.Count)
            return false;

        CycleNumber++;

        return true;
    }

    public void CycleEnd()
    {
        if (instructions[pc].End(registers))
            pc++;
    }
}

public abstract class InstructionBase
{
    protected int Counter { get; private set; }
    protected int CycleCount { get; }

    public InstructionBase(int cycleCount)
    {
        CycleCount = cycleCount;
    }

    protected virtual void Done(Dictionary<string, int> registers)
    {
    }

    public bool End(Dictionary<string, int> registers)
    {
        Counter++;

        if (Counter == CycleCount)
        {
            Done(registers);
            return true;
        }

        return false;
    }

    public void Reset()
    {
        Counter = 0;
    }
}

public class NoopInstruction : InstructionBase
{
    public NoopInstruction()
        : base(1)
    {
    }
}

public class AddInstruction : InstructionBase
{
    private readonly int delta;

    public AddInstruction(int delta)
        : base(2)
    {
        this.delta = delta;
    }

    protected override void Done(Dictionary<string, int> registers)
    {
        const string registerName = "X";
        registers.TryGetValue(registerName, out int value);
        registers[registerName] = value + delta;
    }
}
