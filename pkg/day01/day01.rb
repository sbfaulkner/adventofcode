class Day01
  def initialize(path)
    @lists = [[], []]

    fullpath = File.expand_path(path, __dir__)

    File.readlines(fullpath).map do |line|
      a, b = line.split.map(&:to_i)
      @lists[0] << a
      @lists[1] << b
    end

    @lists[0].sort!
    @lists[1].sort!
  end

  def part1
    @lists[0].zip(@lists[1]).map { |args| args.reduce(&:-).abs }.reduce(&:+)
  end

  def part2
    tally = @lists[1].tally
    @lists[0].map { |id| id * tally.fetch(id, 0) }.reduce(&:+)
  end
end
