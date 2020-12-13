using System;
using System.Collections.Generic;

namespace Day10
{
    class Grid
    {
        private Dictionary<int, Dictionary<int, bool>> grid;

        private Nullable<int> maximumX;
        private Nullable<int> maximumY;
        private Nullable<int> minimumX;
        private Nullable<int> minimumY;

        public Grid()
        {
            grid = new Dictionary<int, Dictionary<int, bool>>();
        }

        public int Height()
        {
            return maximumY.Value - minimumY.Value + 1;
        }

        public void Plot((int x, int y) p)
        {
            if (maximumX == null || p.x > maximumX) maximumX = p.x;
            if (maximumY == null || p.y > maximumY) maximumY = p.y;
            if (minimumX == null || p.x < minimumX) minimumX = p.x;
            if (minimumY == null || p.y < minimumY) minimumY = p.y;

            if (!grid.ContainsKey(p.y)) grid.Add(p.y, new Dictionary<int, bool>());

            grid[p.y].TryAdd(p.x, true);
        }

        public void Print()
        {
            for (int y = minimumY.Value; y <= maximumY.Value; y++)
            {
                Dictionary<int, bool> row;
                bool rowPresent = grid.TryGetValue(y, out row);

                for (int x = minimumX.Value; x <= maximumX.Value; x++)
                {
                    bool light = rowPresent && row.ContainsKey(x);

                    Console.Write(light ? "#" : ".");
                }

                Console.Write("\n");
            }
        }
    }
}