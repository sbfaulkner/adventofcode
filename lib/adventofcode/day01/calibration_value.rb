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

    DIGIT_REGEX = /\d/
    NUMBERS_REGEX = Regexp.union(NUMBERS)

    def self.sum(input = File.open(INPUT), spelled: false)
      input.each_line.map do |line|
        Adventofcode::Day01.calibration_value(line.chomp, spelled: spelled)
      end.sum
    end

    def self.calibration_value(line, spelled: false)
      regex = spelled ? Regexp.union(DIGIT_REGEX, NUMBERS_REGEX) : DIGIT_REGEX

      first = line[regex]
      last_index = line.rindex(regex)

      value(first, spelled: spelled) * 10 + value(line[last_index..], spelled: spelled)
    end

    def self.value(digit, spelled:)
      digit = digit.sub(NUMBERS_REGEX) { |word| NUMBERS.index(word) + 1 } if spelled
      digit.to_i
    end
  end
end
