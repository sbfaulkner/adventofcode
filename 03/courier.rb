class Courier
  def initialize
    @x = @y = 0
  end

  def deliver_to(houses)
    houses[@x][@y] += 1
  end

  def move(direction)
    case direction
    when '^'
      @y += 1
    when 'v'
      @y -= 1
    when '>'
      @x += 1
    when '<'
      @x -= 1
    end
  end
end
