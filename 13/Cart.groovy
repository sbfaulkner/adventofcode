class Cart {
  enum Turn {
    LEFT, STRAIGHT, RIGHT
  }

  enum Direction {
    UP('^'), RIGHT('>'), DOWN('v'), LEFT('<')

    private final String stringValue

    static Direction fromString(String s) {
      switch (s) {
        case '^':
          Direction.UP
          break
        case '>':
          Direction.RIGHT
          break
        case 'v':
          Direction.DOWN
          break
        case '<':
          Direction.LEFT
          break
      }
    }

    Direction(String value) {
      this.stringValue = value
    }

    String toString() {
      this.stringValue
    }
  }

  Boolean disabled
  int x, y
  Direction direction
  private Turn turn

  Cart(int x, y, String direction) {
    this.disabled = false
    this.x = x
    this.y = y
    this.direction = Direction.fromString(direction)
    this.turn = Turn.LEFT
  }

  void move(Track t) {
    if (this.disabled) return

    switch (this.direction) {
      case Direction.UP:
        this.y--
        break
      case Direction.RIGHT:
        this.x++
        break
      case Direction.DOWN:
        this.y++
        break
      case Direction.LEFT:
        this.x--
        break
    }

    if (t.detectCollision(this.x, this.y)) {
      println("*** Collision @ ${this.x},${this.y} ***")
    }

    switch (t.getAt(this.x, this.y)) {
      case '+':
        switch (this.turn) {
          case Turn.LEFT:
            this.direction--
            break
          case Turn.RIGHT:
            this.direction++
            break
        }
        this.turn++
        break
      case '/':
        switch (this.direction) {
          case Direction.UP:
            this.direction++
            break
          case Direction.RIGHT:
            this.direction--
            break
          case Direction.DOWN:
            this.direction++
            break
          case Direction.LEFT:
            this.direction--
            break
        }
        break
      case '\\':
        switch (this.direction) {
          case Direction.UP:
            this.direction--
            break
          case Direction.RIGHT:
            this.direction++
            break
          case Direction.DOWN:
            this.direction--
            break
          case Direction.LEFT:
            this.direction++
            break
        }
        break
    }
  }
}
