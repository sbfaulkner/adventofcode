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

  def neighbours?(r, c, min:, max:)
    count = 0
    rows = (r - 1)..(r + 1)
    columns = (c - 1)..(c + 1)

    rows.each do |rr|
      columns.each do |cc|
        next if rr == r && cc == c
        count += 1 if on?(rr, cc)
        break if count > max
      end
      break if count > max
    end

    count >= min && count <= max
  end

  def step
    @lights = @lights.map.with_index do |row, r|
      row.map.with_index do |on, c|
        next true if always_on?(r, c)
        if on
          neighbours?(r, c, min: 2, max: 3)
        else
          neighbours?(r, c, min: 3, max: 3)
        end
      end
    end
  end

  def to_s
    @lights.map { |row| row.map { |value| value ? '#' : '.' }.join }.join("\n")
  end
end
