PROG	= charge
CC		= gcc
OBJS	= charge.o cell.o
COCOA	= -framework Cocoa

$(PROG) : $(OBJS)
	$(CC) $(COCOA) $(OBJS) -o $(PROG)

clean:
	rm -f $(PROG) $(OBJS)
