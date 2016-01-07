class Light
  attr_accessor :value

  def initialize
    @value = false
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
