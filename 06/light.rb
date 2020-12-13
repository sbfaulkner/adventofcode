class Light
  def initialize
    @value = false
  end

  def off?
    !@value
  end

  def on?
    @value
  end

  def toggle
    @value = !@value
  end

  def turn_off
    @value = false
  end

  def turn_on
    @value = true
  end
end
