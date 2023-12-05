module Adventofcode
  module Day01
    INPUT = File.join(__dir__, 'input.txt')

    NUMBERS = %w[
      one
      two
      three
      four
      five
      six
      seven
      eight
      nine
    ]

    NUMBERS_REGEX = Regexp.new(NUMBERS.join('|'))

    def self.sum(input = File.open(INPUT), expand_spelled_numbers: false)
      input.each_line.map do |line|
        Adventofcode::Day01.calibration_value(line, expand_spelled_numbers: expand_spelled_numbers)
      end.sum
    end

    def self.calibration_value(line, expand_spelled_numbers: false)
      line = expand_spelled_numbers(line) if expand_spelled_numbers
      first = line.index(/\d/)
      last = line.rindex(/\d/)
      line[first].to_i * 10 + line[last].to_i
    end

    def self.expand_spelled_numbers(line)
      line.gsub(NUMBERS_REGEX) { |word| NUMBERS.index(word) + 1 }
    end
  end
end
