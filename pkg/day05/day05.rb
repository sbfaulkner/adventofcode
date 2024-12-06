class Day05
  def initialize(path)
    fullpath = File.expand_path(path, __dir__)
    @rows = File.readlines(fullpath).map(&:chomp)
  end

  def part1
    0
  end

  def part2
    0
  end

  private

  # Add private helper methods here
end
