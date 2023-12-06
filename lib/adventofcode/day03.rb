# frozen_string_literal: true

require "strscan"

module Adventofcode
  module Day03
    extend self

    INPUT = File.join(__dir__, "day03", "input.txt")

    NONSYMBOL_REGEX = /[^.\d]/.freeze

    class Schematic
      def initialize(input = File.open(INPUT))
        @lines = input.each_line.map(&:chomp)
      end

      def sum
        sum = 0

        scan_part_numbers do |number|
          sum += number.to_i
          next
        end

        sum
      end

      private

      def scan_part_numbers(&block)
        @lines.each_with_index.each do |line, i|
          scanner = StringScanner.new(line)
          while scanner.scan_until(/(\d+)/)
            number = scanner.captures[0]
            index = scanner.charpos - number.length

            min = [0, index - 1].max
            max = [line.length - 1, index + number.length].min

            if i > 0
              # above
              block.call(number.to_i) if @lines[i - 1][min..max].chars.any? { |c| c =~ NONSYMBOL_REGEX }
            end

            if index > 0
              # before
              block.call(number.to_i, i, index - 1) if line[index - 1] =~ NONSYMBOL_REGEX
            end

            if (index + number.length) < line.length
              # after
              block.call(number.to_i, i, index + 1) if line[index + number.length] =~ NONSYMBOL_REGEX
            end

            if i < @lines.length - 1
              # below
              block.call(number.to_i) if @lines[i + 1][min..max].chars.any? { |c| c =~ NONSYMBOL_REGEX }
            end
          end
        end
      end
    end
  end
end
