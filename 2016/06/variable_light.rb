class VariableLight
  def initialize
    @value = 0
  end

  def brightness
    @value
  end

  def toggle
    @value += 2
  end

  def turn_off
    return @value if @value == 0
    @value -= 1
  end

  def turn_on
    @value += 1
  end
end
