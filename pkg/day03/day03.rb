require "strscan"

class Day03
  REGEX_INSTRUCTION = /(?<instruction>do|don't|mul)\((?<args>\d+,\d+)?\)/
  def initialize(path)
    fullpath = File.expand_path(path, __dir__)

    @instructions = File.readlines(fullpath).each_with_object([]) do |line, instructions|
      scanner = StringScanner.new(line)
      while scanner.scan_until(REGEX_INSTRUCTION) do
        instructions << scanner.named_captures
      end
    end
  end

  def part1
    @instructions.select do |instruction|
      instruction["instruction"] == 'mul'
    end.map do |instruction|
      instruction["args"].split(',').map(&:to_i).reduce(:*)
    end.reduce(:+)
  end

  def part2
    enabled = true

    @instructions.map do |instruction|
      case instruction["instruction"]
      when "do"
        enabled = true
        nil
      when "don't"
        enabled = false
        nil
      when "mul"
        instruction["args"].split(',').map(&:to_i).reduce(:*) if enabled
      end
    end.compact.reduce(:+)
  end
end
