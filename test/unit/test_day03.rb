# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day03
    class TestDay03 < Minitest::Test
      INPUT = <<~HEREDOC
        467..114..
        ...*......
        ..35..633.
        ......#...
        617*......
        .....+.58.
        ..592.....
        ......755.
        ...$.*....
        .664.598..
      HEREDOC

      def test_sum
        schematic = Adventofcode::Day03::Schematic.new(INPUT)

        assert_equal(4361, schematic.sum)
      end

      def test_sum_gears
        schematic = Adventofcode::Day03::Schematic.new(INPUT)

        assert_equal(467835, schematic.sum_gears)
      end
    end
  end
end
