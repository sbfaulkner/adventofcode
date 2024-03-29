# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day05
    class TestDay05 < Minitest::Test
      INPUT = <<~HEREDOC
        seeds: 79 14 55 13

        seed-to-soil map:
        50 98 2
        52 50 48

        soil-to-fertilizer map:
        0 15 37
        37 52 2
        39 0 15

        fertilizer-to-water map:
        49 53 8
        0 11 42
        42 0 7
        57 7 4

        water-to-light map:
        88 18 7
        18 25 70

        light-to-temperature map:
        45 77 23
        81 45 19
        68 64 13

        temperature-to-humidity map:
        0 69 1
        1 0 69

        humidity-to-location map:
        60 56 37
        56 93 4
      HEREDOC

      def test_lowest_location
        assert_equal(35, Adventofcode::Day05::Almanac.new(INPUT).lowest_location)
      end

      def test_lowest_location_with_ranges
        assert_equal(46, Adventofcode::Day05::Almanac.new(INPUT, ranges: true).lowest_location)
      end
    end
  end
end
