class Day02
  def initialize(path)
    fullpath = File.expand_path(path, __dir__)

    @reports = File.readlines(fullpath).map do |line|
      line.split.map(&:to_i)
    end
  end

  def part1
    @reports.count { |report| safe?(report) }
  end

  def part2
    @reports.count { |report| safe?(report) || safe_with_dampener?(report) }
  end

  private

  def safe?(report)
    diffs = report.each_cons(2).map { |a, b| a - b }
    return false if diffs.any? { |diff| diff.abs > 3 || diff == 0 }
    return false if diffs.each_cons(2).any? { |a, b| a * b < 0 }
    true
  end

  def safe_with_dampener?(report)
    report.each_with_index.any? do |_, i|
      dampened = report.select.with_index { |_, j| j != i }
      safe?(dampened)
    end
  end
end
