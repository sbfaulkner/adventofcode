# frozen_string_literal: true

require "strscan"

module Adventofcode
  module Day03
    extend self

    INPUT = File.join(__dir__, "day03", "input.txt")

    NONSYMBOL_REGEX = /[^.\d]/.freeze

    class Schematic
      def initialize(input = File.open(INPUT))
        @lines = input.each_line.map(&:chomp).tap { |r| warn("Input has #{r.length} lines") }
      end

      def sum
        sum = 0

        @lines.each_with_index.each do |line, i|
          scanner = StringScanner.new(line)
          while scanner.scan_until(/(\d+)/)
            number = scanner.captures[0]
            index = scanner.charpos - number.length

            min = [0, index - 1].max
            max = [line.length - 1, index + number.length].min

            if index > 0 && line[index - 1] =~ NONSYMBOL_REGEX
              # before
              sum += number.to_i
              next
            end

            if (index + number.length) < line.length && line[index + number.length] =~ NONSYMBOL_REGEX
              # after
              sum += number.to_i
              next
            end

            if i > 0 && @lines[i - 1][min..max].chars.any? { |c| c =~ NONSYMBOL_REGEX }
              # above
              sum += number.to_i
              next
            end

            next unless i < @lines.length - 1 && @lines[i + 1][min..max].chars.any? { |c| c =~ NONSYMBOL_REGEX }

            # below
            sum += number.to_i
            next

            # no match
          end
        end

        sum
      end
    end
  end
end
