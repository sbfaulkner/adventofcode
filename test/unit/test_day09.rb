# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day09
    class TestDay09 < Minitest::Test
      def test_count
        input = <<~HEREDOC
          0 3 6 9 12 15
          1 3 6 10 15 21
          10 13 16 21 30 45
        HEREDOC

        assert_equal(114, Adventofcode::Day09::History.load(input).map(&:next).sum)
      end
    end
  end
end
