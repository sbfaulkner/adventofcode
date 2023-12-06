# frozen_string_literal: true

require 'test_helper'

module Adventofcode
  module Day01
    class TestDay01 < Minitest::Test
      INPUT1 = <<~HEREDOC
        1abc2
        pqr3stu8vwx
        a1b2c3d4e5f
        treb7uchet
      HEREDOC

      def test_it_extracts_the_calibration_value
        assert_equal 12, Adventofcode::Day01.calibration_value('1abc2')
        assert_equal 38, Adventofcode::Day01.calibration_value('pqr3stu8vwx')
        assert_equal 15, Adventofcode::Day01.calibration_value('a1b2c3d4e5f')
        assert_equal 77, Adventofcode::Day01.calibration_value('treb7uchet')
      end

      def test_it_sums_the_calibration_values
        assert_equal 142, Adventofcode::Day01.sum(INPUT1)
      end

      INPUT2 = <<~HEREDOC
        two1nine
        eightwothree
        abcone2threexyz
        xtwone3four
        4nineeightseven2
        zoneight234
        7pqrstsixteen
      HEREDOC

      def test_it_extracts_the_calibration_value_with_spelled_numbers
        assert_equal 29, Adventofcode::Day01.calibration_value('two1nine', spelled: true)
        assert_equal 83, Adventofcode::Day01.calibration_value('eightwothree', spelled: true)
        assert_equal 13, Adventofcode::Day01.calibration_value('abcone2threexyz', spelled: true)
        assert_equal 24, Adventofcode::Day01.calibration_value('xtwone3four', spelled: true)
        assert_equal 42, Adventofcode::Day01.calibration_value('4nineeightseven2', spelled: true)
        assert_equal 14, Adventofcode::Day01.calibration_value('zoneight234', spelled: true)
        assert_equal 76, Adventofcode::Day01.calibration_value('7pqrstsixteen', spelled: true)
      end

      def test_it_extracts_ambiguous_values_with_spelled_numbers
        assert_equal 12, Adventofcode::Day01.calibration_value('oneeightwo', spelled: true)
      end

      def test_it_sums_the_calibration_values_with_spelled_numbers
        assert_equal 281, Adventofcode::Day01.sum(INPUT2, spelled: true)
      end
    end
  end
end
