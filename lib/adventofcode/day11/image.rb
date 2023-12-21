# frozen_string_literal: true

module Adventofcode
  module Day11
    class Image
      class << self
        def load(input = File.open(INPUT))
          lines = input.each_line(chomp: true).to_a
          new(lines)
        end
      end

      Galaxy = Struct.new(:x, :y)

      def initialize(lines)
        @height = lines.length
        @width = lines.first.length

        @empty_rows = Array.new(@height, true)
        @empty_columns = Array.new(@width, true)

        @galaxies = []

        lines.each_with_index do |line, y|
          line.each_char.with_index do |char, x|
            next unless char == "#"

            galaxy = Galaxy.new(x, y)

            @empty_rows[y] = false
            @empty_columns[x] = false

            @galaxies << galaxy
          end
        end
      end

      def expand(rate: 2)
        @galaxies.combination(2).map do |(a, b)|
          distance(a, b, rate: rate - 1)
        end.sum
      end

      def distance(a, b, rate:)
        x1 = [a.x, b.x].min
        x2 = [a.x, b.x].max
        y1 = [a.y, b.y].min
        y2 = [a.y, b.y].max

        x2 - x1 + y2 - y1 + @empty_columns[x1..x2].count(&:itself) * rate + @empty_rows[y1..y2].count(&:itself) * rate
      end
    end
  end
end
