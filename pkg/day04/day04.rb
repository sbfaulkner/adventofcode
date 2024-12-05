class Day04
  DIRECTIONS = [
    [-1, -1],
    [-1, 0],
    [-1, 1],
    [0, 1],
    [1, 1],
    [1, 0],
    [1, -1],
    [0, -1],
  ]

  def initialize(path)
    fullpath = File.expand_path(path, __dir__)

    @rows = File.readlines(fullpath).map(&:chomp)
  end

  def part1
    @rows.map.with_index do |row, r|
      row.chars.map.with_index do |ch, c|
        (ch == "X" ? DIRECTIONS.count { |dr, dc| match(r+dr, c+dc, "MAS", dr, dc) } : 0)
      end.sum
    end.sum
  end

  def part2
    0
  end

  private

  def match(r, c, word, dr, dc)
    # matched the word
    return true if word.empty?

    # bounds checking
    return false if r < 0 || c < 0
    return false if r >= @rows.length || c >= @rows[r].length

    # not a match
    return false if @rows[r][c] != word[0]

    match(r+dr, c+dc, word[1..], dr, dc)
  end
end
