# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day06
    class TestDay06 < Minitest::Test
      INPUT = <<~HEREDOC
        Time:      7  15   30
        Distance:  9  40  200
      HEREDOC

      def test_winners
        assert_equal(288, Adventofcode::Day06::Race.load(INPUT).map(&:winners).reduce(&:*))
      end

      def test_winners_ignoring_whitespace
        assert_equal(71503, Adventofcode::Day06::Race.load(INPUT, split: false).map(&:winners).reduce(&:*))
      end
    end
  end
end
