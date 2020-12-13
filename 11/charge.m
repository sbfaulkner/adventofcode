#import <Foundation/Foundation.h>
#import "cell.h"

#define GRIDSIZE    300
#define GRIDSERIAL  1308

int main (int argc, const char * argv[])
{
    NSAutoreleasePool *pool = [[NSAutoreleasePool alloc] init];

    NSMutableArray *grid = [[NSMutableArray alloc] initWithCapacity: GRIDSIZE];

    for (int y = 0; y < GRIDSIZE; y++) {
        NSMutableArray *row = [[NSMutableArray alloc] initWithCapacity: GRIDSIZE];

        for (int x = 0; x < GRIDSIZE; x++) {
            [row addObject: [[Cell alloc] initWithGridSerialNumber: GRIDSERIAL x: x y: y]];
        }

        [grid addObject: row];
    }

    int xMax, yMax, sizeMax;
    int totalMax = INT_MIN;

    for (int y = 0; y < GRIDSIZE; y++) {
        for (int x = 0; x < GRIDSIZE; x++) {
            int total = 0;

            for (int s = 0; s < MIN(GRIDSIZE - y, GRIDSIZE - x); s++) {
                for (int wy = 0; wy <= s; wy++) {
                    Cell *c = grid[y+wy][x+s];
                    total += c.powerLevel;
                }

                for (int wx = 0; wx < s; wx++) {
                    Cell *c = grid[y+s][x+wx];
                    total += c.powerLevel;
                }

                if (total > totalMax) {
                    xMax = x;
                    yMax = y;
                    sizeMax = s + 1;
                    totalMax = total;
                }
            }
        }
    }

    NSLog(
        @"For grid serial number %d, the largest total square (with a total power of %d) is %dx%d and has a top-left corner of %d,%d, so its identifier is %d,%d,%d.",
        GRIDSERIAL, totalMax, sizeMax, sizeMax, xMax, yMax, xMax, yMax, sizeMax
    );

    [pool drain];

    return 0;
}
