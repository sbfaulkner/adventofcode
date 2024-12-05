class Day04
  def initialize(path)
    fullpath = File.expand_path(path, __dir__)

    @rows = File.readlines(fullpath).map(&:chomp)
  end

  def part1
    @rows.map.with_index do |row, r|
      row.chars.map.with_index do |_, c|
        count_xmas(r, c)
      end.sum
    end.sum
  end

  def part2
    @rows[1..-2].map.with_index do |row, r|
      row.chars[1..-2].map.with_index do |_,c|
        x_mas?(r + 1, c + 1) ? 1 : 0
      end.sum
    end.sum
  end

  private

  XMAS_DIRECTIONS = [
    [-1, -1],
    [-1, 0],
    [-1, 1],
    [0, 1],
    [1, 1],
    [1, 0],
    [1, -1],
    [0, -1],
  ]

  def count_xmas(r, c)
    return 0 unless @rows[r][c] == "X"
    XMAS_DIRECTIONS.count { |dr, dc| match(r+dr, c+dc, "MAS", dr, dc) }
  end

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

  X_MAS_DIRECTIONS = [
    [-1, -1],
    [-1, 1],
    [1, 1],
    [1, -1],
  ]

  def x_mas?(r, c)
    return false unless @rows[r][c] == "A"

    X_MAS_DIRECTIONS.each_with_index do |(dr, dc), i|
      next if @rows[r+dr][c+dc] != "M"
      dr, dc = X_MAS_DIRECTIONS[(i+1)%4]
      next if @rows[r+dr][c+dc] != "M"
      dr, dc = X_MAS_DIRECTIONS[(i+2)%4]
      next if @rows[r+dr][c+dc] != "S"
      dr, dc = X_MAS_DIRECTIONS[(i+3)%4]
      next if @rows[r+dr][c+dc] != "S"

      return true
    end

    false
  end
end
