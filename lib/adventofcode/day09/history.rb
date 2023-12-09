# frozen_string_literal: true

module Adventofcode
  module Day09
    class History
      class << self
        def load(input = File.open(INPUT))
          input.each_line(chomp: true).map do |line|
            new(*line.split(" ").map(&:to_i))
          end
        end
      end

      def initialize(*values)
        @values = values
        @expanded = Hash.new do |rows, r|
          rows[r] = Hash.new do |row, c|
            max = rows[0].length - r

            row[c] = if c == max
              if row[c - 1] == 0
                0
              else
                row[c - 1] + rows[r + 1][c - 1]
              end
            else
              rows[r - 1][c + 1] - rows[r - 1][c]
            end
          end
        end

        values.each_with_index { |value, index| @expanded[0][index] = value }
      end

      def next
        @expanded[0][@expanded[0].length]
      end
    end
  end
end
