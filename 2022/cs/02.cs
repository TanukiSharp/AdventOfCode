namespace Day2Namespace;

class Day02 : IPuzzle
{
    private const int RoundLostScore = 0;
    private const int RoundDrawScore = 3;
    private const int RoundWinScore = 6;

    private record struct HandShapeInfo(int ShapeValue);

    private enum HandShape
    {
        Draw,
        Rock,
        Paper,
        Scissors,
    }

    private enum Strategy
    {
        Lost,
        Draw,
        Win,
    }

    private static HandShape GetHandShape(char letter)
    {
        switch (letter)
        {
            case 'A':
            case 'X':
                return HandShape.Rock;
            case 'B':
            case 'Y':
                return HandShape.Paper;
            case 'C':
            case 'Z':
                return HandShape.Scissors;
        }

        throw new InvalidDataException($"Unknown character '{letter}' in input data.");
    }

    private static Strategy GetStrategy(char letter)
    {
        switch (letter)
        {
            case 'X':
                return Strategy.Lost;
            case 'Y':
                return Strategy.Draw;
            case 'Z':
                return Strategy.Win;
        }

        throw new InvalidDataException($"Invalid strategy character '{letter}' in input data.");
    }

    private static HandShape GetWinner(HandShape opponent, HandShape me)
    {
        // int delta = (int)opponent - (int)me;

        // if (delta == 0)
        //     return HandShape.Draw;

        // if (delta == 1 || delta < -1)
        //     return opponent;

        // return me;

        if (opponent == me)
            return HandShape.Draw;

        if (opponent == HandShape.Rock)
            return me == HandShape.Paper ? me : opponent;

        if (opponent == HandShape.Paper)
            return me == HandShape.Scissors ? me : opponent;

        return me == HandShape.Rock ? me : opponent;
    }

    private static int GetRoundScore(HandShape opponent, HandShape me)
    {
        HandShape winner = GetWinner(opponent, me);

        if (winner == HandShape.Draw)
            return RoundDrawScore;

        if (winner == me)
            return RoundWinScore;

        return RoundLostScore;
    }

    private static HandShape GetWiningHandShape(HandShape opponent)
    {
        if (opponent == HandShape.Rock)
            return HandShape.Paper;
        if (opponent == HandShape.Paper)
            return HandShape.Scissors;
        if (opponent == HandShape.Scissors)
            return HandShape.Rock;

        throw new ArgumentException(nameof(opponent));
    }

    private static HandShape GetLoosingHandShape(HandShape opponent)
    {
        if (opponent == HandShape.Rock)
            return HandShape.Scissors;
        if (opponent == HandShape.Paper)
            return HandShape.Rock;
        if (opponent == HandShape.Scissors)
            return HandShape.Paper;

        throw new ArgumentException(nameof(opponent));
    }

    private static void Part1(List<(char opponentChar, char myChar)> entries)
    {
        int totalScore = 0;

        foreach ((char opponentChar, char myChar) in entries)
        {
            HandShape opponent = GetHandShape(opponentChar);
            HandShape me = GetHandShape(myChar);

            int roundScore = GetRoundScore(opponent, me);
            roundScore += (int)me;

            totalScore += roundScore;
        }

        Console.WriteLine($"My total score in part 1 is {totalScore}.");
    }

    private static HandShape GetPart2HandShape(HandShape opponent, Strategy strategy)
    {
        switch (strategy)
        {
            case Strategy.Lost:
                return GetLoosingHandShape(opponent);
            case Strategy.Draw:
                return opponent;
            case Strategy.Win:
                return GetWiningHandShape(opponent);
        }

        throw new ArgumentException(nameof(strategy));
    }

    private static void Part2(List<(char opponentChar, char myChar)> entries)
    {
        int totalScore = 0;

        foreach ((char opponentChar, char myChar) in entries)
        {
            HandShape opponent = GetHandShape(opponentChar);
            HandShape me = GetPart2HandShape(GetHandShape(opponentChar), GetStrategy(myChar));

            int roundScore = GetRoundScore(opponent, me);
            roundScore += (int)me;

            totalScore += roundScore;
        }

        Console.WriteLine($"My total score in part 2 is {totalScore}.");
    }

    public int Day => 2;
    public bool IsTest => false;

    public void Run(string input)
    {
        int lineNumber = -1;
        List<(char, char)> entries = new();

        foreach (string line in input.Split('\n').Select(x => x.Trim()))
        {
            lineNumber++;

            if (string.IsNullOrWhiteSpace(line))
                continue;

            string[] parts = line.Split(' ');

            if (parts.Length != 2 || parts[0].Length != 1 || parts[1].Length != 1)
                throw new InvalidDataException($"Invalid data at line '{lineNumber}'.");

            entries.Add((parts[0][0], parts[1][0]));
        }

        Part1(entries);
        Part2(entries);
    }
}
