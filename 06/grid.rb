require_relative 'light'

class Grid
  def initialize
    @lights = Array.new(1000) { Array.new(1000, Light.new) }
  end

  def configure(instructions)
    instructions.each_line do |instruction|
      next unless match = instruction.match(/(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)/)

      command = match.captures.first
      x1, y1, x2, y2 = match.captures[1..4].map(&:to_i)

      perform command, x1, y1, x2, y2
    end
  end

  def count
    @lights.flatten.count(&:value)
  end

  private

  def perform(command, x1, y1, x2, y2)
    method = command.tr(' ', '_')

    (x1..x2).each do |x|
      (y1..y2).each do |y|
        @lights[x][y].send method
      end
    end
  end
end
