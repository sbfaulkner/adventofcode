# frozen_string_literal: true

require "strscan"

module Adventofcode
  module Day03
    extend self

    INPUT = File.join(__dir__, "day03", "input.txt")

    SYMBOL_REGEX = /[^.\d]/.freeze
    GEAR = "*"

    class Schematic
      def initialize(input = File.open(INPUT))
        @lines = input.each_line.map(&:chomp)
      end

      def sum
        sum = 0

        scan_part_numbers do |number, _row, _col|
          sum += number.to_i
          next
        end

        sum
      end

      def sum_gears
        potential_gears = Hash.new do |h, k|
          h[k] = Hash.new do |h2, k2|
            h2[k2] = []
          end
        end

        scan_part_numbers do |number, row, col|
          potential_gears[row][col] << number if @lines[row][col] == GEAR
        end

        sum = 0

        potential_gears.each do |_row, cols|
          cols.each do |_col, parts|
            sum += parts.reduce(&:*) if parts.length == 2
          end
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
              @lines[i - 1][min..max].chars.each_with_index do |c, offset|
                block.call(number.to_i, i - 1, min + offset) if c =~ SYMBOL_REGEX
              end
            end

            if index > 0
              # before
              block.call(number.to_i, i, index - 1) if line[index - 1] =~ SYMBOL_REGEX
            end

            if (index + number.length) < line.length
              # after
              block.call(number.to_i, i, index + 1) if line[index + number.length] =~ SYMBOL_REGEX
            end

            next if i >= @lines.length - 1

            # below
            @lines[i + 1][min..max].chars.each_with_index do |c, offset|
              block.call(number.to_i, i + 1, min + offset) if c =~ SYMBOL_REGEX
            end
          end
        end
      end
    end
  end
end
