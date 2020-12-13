@interface Cell : NSObject

@property (readonly) int gridSerialNumber;

@property (readonly) int x;
@property (readonly) int y;

@property (readonly) int rackID;

@property (readonly) int powerLevel;

- (instancetype) initWithGridSerialNumber:(int)gridSerialNumber x:(int)x y:(int)y;

@end
