#import <Foundation/Foundation.h>
#import "cell.h"

@implementation Cell

- (instancetype) initWithGridSerialNumber:(int)gridSerialNumber x:(int)x y:(int)y {
    self = [super init];

    if (self) {
        _gridSerialNumber = gridSerialNumber;

        _x = x;
        _y = y;

        _rackID = _x + 10;

        _powerLevel = (int)((_rackID * _y + _gridSerialNumber) * _rackID % 1000 / 100) - 5;
    }

    return self;
}

@end
