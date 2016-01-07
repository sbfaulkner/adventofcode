class Grid
  def initialize
    @lights = Array.new(1000) { Array.new(1000, false) }
  end

  def configure(instructions)
    instructions.each_line do |instruction|
      process instruction
    end
  end

  def count
    @lights.flatten.count(true)
  end

  private

  def process(instruction)
    if match = instruction.match(/(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)/)
      command = match.captures.first
      x1, y1, x2, y2 = match.captures[1,4].map(&:to_i)

      for x in x1..x2
        for y in y1..y2
          case command
          when 'turn on'
            @lights[x][y] = true
          when 'turn off'
            @lights[x][y] = false
          when 'toggle'
            @lights[x][y] = !@lights[x][y]
          end
        end
      end
    end
  end
end
