# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day11
    class TestDay11 < Minitest::Test
      INPUT = <<~HEREDOC
        ...#......
        .......#..
        #.........
        ..........
        ......#...
        .#........
        .........#
        ..........
        .......#..
        #...#.....
      HEREDOC

      def test_image_expansion
        assert_equal(374, Adventofcode::Day11::Image.load(INPUT).expand)
      end

      def test_image_expansion_10_times
        assert_equal(1030, Adventofcode::Day11::Image.load(INPUT).expand(rate: 10))
      end

      def test_image_expansion_100_times
        assert_equal(8410, Adventofcode::Day11::Image.load(INPUT).expand(rate: 100))
      end
    end
  end
end
