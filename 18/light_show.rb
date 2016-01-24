class LightShow
  def self.load(source, always_on: [])
    new source.each_line.map { |row| row.strip.each_char.map { |ch| ch == '#' } }, always_on: always_on
  end

  def initialize(lights, always_on: [])
    @lights    = lights
    @always_on = always_on.each { |r, c| @lights[r][c] = true }
  end

  attr_reader :lights

  def always_on?(r, c)
    @always_on.include?([r, c])
  end

  def on?(r, c)
    return if r < 0 || c < 0 || r >= @lights.size
    @lights[r][c] || always_on?(r, c)
  end

  def count
    @lights.map { |row| row.count(true) }.reduce(&:+)
  end

  def neighbours(r, c)
    rows = (r - 1)..(r + 1)
    columns = (c - 1)..(c + 1)
    rows.map { |rr| columns.map { |cc| on?(rr, cc) unless rr == r && cc == c }.count(true) }.reduce(&:+)
  end

  def step
    @lights = @lights.map.with_index do |row, r|
      row.map.with_index do |on, c|
        next true if always_on?(r, c)
        count = neighbours(r, c)
        count == 3 || on && count == 2
      end
    end
  end

  def to_s
    @lights.map { |row| row.map { |value| value ? '#' : '.' }.join }.join("\n")
  end
end
