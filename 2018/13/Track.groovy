import java.nio.file.Path
import java.util.regex.*

class Track {
  private static final cartPattern = Pattern.compile(/[\^v\<\>]/)

  private List track;
  List carts;

  Track(Path input) {
    this.track = []
    this.carts = []

    int y = 0

    input.eachLine {
      Matcher m = cartPattern.matcher(it)

      while (m.find()) {
        int x = m.start()
        String d = m.group(0)

        this.carts.add(new Cart(x, y, d))
      }

      this.track.add(it.tr("^v<>", "||-"))

      y++
    }
  }

  Boolean detectCollision(int x, y) {
    List carts = this.carts.findAll { it.x == x && it.y == y }

    if (carts.size() > 1) {
      carts.each { it.disabled = true }
      true
    }
  }

  Character getAt(int x, y) {
    this.track[y][x]
  }

  Boolean isMultipleCarts() {
    this.carts.count { ! it.disabled } > 1
  }

  void print() {
    this.track.eachWithIndex { row, y ->
      int x = 0
      println row.collectReplacements {
        Cart cart = this.carts.find { c -> c.y == y && c.x == x }
        x++
        if (cart != null) {
          cart.direction.toString();
        }
      }
    }
  }

  void tick() {
    this.carts.each {
      it.move(this)
    }

    this.carts.removeAll { it.disabled }
  }
}
