class LightShow
  def self.load(source, always_on: [])
    new source.each_line.map { |row| row.strip.each_char.map { |ch| ch == '#' } }, always_on: always_on
  end

  def initialize(lights, always_on: [])
    @lights    = lights
    @always_on = always_on.each { |r, c| @lights[r][c] = true }
  end

  attr_reader :lights

  def on?(r, c)
    return true if @always_on.include?([r, c])
    return unless 0 <= r && r < @lights.size
    row = @lights[r]
    return unless 0 <= c && c < row.size
    row[c]
  end

  def count
    @lights.map { |row| row.count(true) }.reduce(&:+)
  end

  def neighbours(r, c)
    [
      on?(r - 1, c - 1), on?(r - 1, c), on?(r - 1, c + 1),
      on?(r, c - 1), on?(r, c + 1),
      on?(r + 1, c - 1), on?(r + 1, c), on?(r + 1, c + 1)
    ].count(true)
  end

  def step
    @lights = @lights.map.with_index do |row, r|
      row.map.with_index do |on, c|
        count = neighbours(r, c)
        on && count == 2 || count == 3 || @always_on.include?([r, c])
      end
    end
  end

  def to_s
    @lights.map { |row| row.map { |value| value ? '#' : '.' }.join }.join("\n")
  end
end
