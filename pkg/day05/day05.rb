class Day05
  def initialize(path)
    fullpath = File.expand_path(path, __dir__)

    @rules = Hash.new { |hash, key| hash[key] = Set.new }
    @updates = []

    lines = File.readlines(fullpath).each

    loop do
      line = lines.next.chomp
      break if line.empty?

      before, after = line.split("|").map(&:to_i)
      @rules[before] << after
    end

    loop do
      line = lines.next.chomp
      break if line.empty?

      @updates << line.split(",").map(&:to_i)
    end
  end

  def part1
    sum = 0

    @updates.each do |update|
      next unless valid_update?(update)
      sum += update[update.size/2]
    end

    sum
  end

  def part2
    sum = 0

    @updates.each do |update|
      next if valid_update?(update)
      update.sort! do |a, b|
        if @rules[a].include?(b)
          -1
        elsif @rules[b].include?(a)
          1
        else
          0
        end
      end
      sum += update[update.size/2]
    end

    sum
  end

  private

  def valid_update?(update)
    preceding = Set.new
    update.each_with_index do |page, index|
      return false if @rules[page].intersect?(preceding)
      preceding << page
    end
    true
  end
end
