class Elevator
  attr_reader :floor

  def self.direction(instruction)
    case instruction
    when '('
      1
    when ')'
      -1
    end
  end

  def initialize
    @floor = 0
  end

  def follow(instructions, destination = nil)
    count = 0
    instructions.chars.each do |ch|
      if step = self.class.direction(ch)
        count += 1
        @floor += step
      end
      break if @floor == destination
    end
    count
  end
end
