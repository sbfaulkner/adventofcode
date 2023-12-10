# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day09
    class TestDay09 < Minitest::Test
      INPUT = <<~HEREDOC
        0 3 6 9 12 15
        1 3 6 10 15 21
        10 13 16 21 30 45
      HEREDOC

      def test_next
        assert_equal(114, Adventofcode::Day09::History.load(INPUT).map(&:next).sum)
      end

      def test_previous
        assert_equal(2, Adventofcode::Day09::History.load(INPUT).map(&:previous).sum)
      end
    end
  end
end
