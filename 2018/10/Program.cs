using System;
using System.IO;
using System.Reflection;
using System.Collections.Generic;

namespace Day10
{
    class Program
    {
        private const string InputFile = "input";

        private static string AssemblyPath()
        {
            return new Uri(Assembly.GetExecutingAssembly().CodeBase).LocalPath;
        }

        private static string InputPath()
        {
            return Path.GetFullPath(Path.Combine(SourceDirectory(), InputFile));
        }

        private static string SourceDirectory()
        {
            return Path.GetFullPath(Path.Combine(AssemblyPath(), "../../../.."));
        }

        static void Main(string[] args)
        {
            List <Point> points = new List<Point>();

            using (StreamReader sr = new StreamReader(InputPath()))
            {
                string line;

                while ((line = sr.ReadLine()) != null)
                {
                    points.Add(Point.Parse(line));
                }
            }

            for (int time = 0; ; time++)
            {
                Grid grid = new Grid();

                foreach (Point p in points)
                {
                    grid.Plot(p.PositionAt(time));
                }

                if (grid.Height() == 10) {
                    Console.WriteLine("Time: {0}", time);
                    grid.Print();
                    break;
                }
            }
        }
    }
}
