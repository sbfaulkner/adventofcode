using System;
using System.Text.RegularExpressions;

namespace Day10
{
    class Point
    {
        private const string InputPattern = @"^position=<(?<position>[^>]+)> velocity=<(?<velocity>[^>]+)>$";

        private static Regex InputRegex = new Regex(InputPattern, RegexOptions.Compiled | RegexOptions.IgnoreCase);

        public static Point Parse(string inputText)
        {
            Match match = InputRegex.Matches(inputText)[0];

            var position = ParseTuple(match.Groups["position"].Value);
            var velocity = ParseTuple(match.Groups["velocity"].Value);

            return new Point(position, velocity);
        }

        public (int x, int y) position;
        public (int x, int y) velocity;

        public Point((int, int) p, (int, int) v)
        {
            position = p;
            velocity = v;
        }

        public (int x, int y) PositionAt(int time)
        {
            return (position.x + velocity.x * time, position.y + velocity.y * time);
        }

        private static (int x, int y) ParseTuple(string text)
        {
            string[] s = text.Split(',');
            return (Int32.Parse(s[0]), Int32.Parse(s[1]));
        }
    }
}