# frozen_string_literal: true

require 'test_helper'

module Adventofcode
  module Day01
    class TestAdventofcode < Minitest::Test
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

      def test_it_expands_spelled_numbers
        assert_equal '219', Adventofcode::Day01.expand_spelled_numbers('two1nine')
        assert_equal '8wo3', Adventofcode::Day01.expand_spelled_numbers('eightwothree')
        assert_equal 'abc123xyz', Adventofcode::Day01.expand_spelled_numbers('abcone2threexyz')
        assert_equal 'x2ne34', Adventofcode::Day01.expand_spelled_numbers('xtwone3four')
        assert_equal '49872', Adventofcode::Day01.expand_spelled_numbers('4nineeightseven2')
        assert_equal 'z1ight234', Adventofcode::Day01.expand_spelled_numbers('zoneight234')
        assert_equal '7pqrst6teen', Adventofcode::Day01.expand_spelled_numbers('7pqrstsixteen')
      end

      def test_it_extracts_the_calibration_value_expanding_spelled_numbers
        assert_equal 29, Adventofcode::Day01.calibration_value('two1nine', expand_spelled_numbers: true)
        assert_equal 83, Adventofcode::Day01.calibration_value('eightwothree', expand_spelled_numbers: true)
        assert_equal 13, Adventofcode::Day01.calibration_value('abcone2threexyz', expand_spelled_numbers: true)
        assert_equal 24, Adventofcode::Day01.calibration_value('xtwone3four', expand_spelled_numbers: true)
        assert_equal 42, Adventofcode::Day01.calibration_value('4nineeightseven2', expand_spelled_numbers: true)
        assert_equal 14, Adventofcode::Day01.calibration_value('zoneight234', expand_spelled_numbers: true)
        assert_equal 76, Adventofcode::Day01.calibration_value('7pqrstsixteen', expand_spelled_numbers: true)
      end

      def test_it_sums_the_calibration_values_expanding_spelled_numbers
        assert_equal 281, Adventofcode::Day01.sum(INPUT2, expand_spelled_numbers: true)
      end
    end
  end
end
