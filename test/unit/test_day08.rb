# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day07
    class TestDay07 < Minitest::Test
      INPUT1 = <<~HEREDOC
        RL

        AAA = (BBB, CCC)
        BBB = (DDD, EEE)
        CCC = (ZZZ, GGG)
        DDD = (DDD, DDD)
        EEE = (EEE, EEE)
        GGG = (GGG, GGG)
        ZZZ = (ZZZ, ZZZ)
      HEREDOC

      INPUT2 = <<~HEREDOC
        LLR

        AAA = (BBB, BBB)
        BBB = (AAA, ZZZ)
        ZZZ = (ZZZ, ZZZ)
      HEREDOC

      def test_count
        assert_equal(2, Adventofcode::Day08::Network.load(INPUT1).count)
      end

      def test_count_with_repeated_instructions
        assert_equal(6, Adventofcode::Day08::Network.load(INPUT2).count)
      end
    end
  end
end
