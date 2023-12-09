# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day07
    class TestDay07 < Minitest::Test
      def test_count
        input = <<~HEREDOC
          RL

          AAA = (BBB, CCC)
          BBB = (DDD, EEE)
          CCC = (ZZZ, GGG)
          DDD = (DDD, DDD)
          EEE = (EEE, EEE)
          GGG = (GGG, GGG)
          ZZZ = (ZZZ, ZZZ)
        HEREDOC

        assert_equal(2, Adventofcode::Day08::Network.load(input).count)
      end

      def test_count_with_repeated_instructions
        input = <<~HEREDOC
          LLR

          AAA = (BBB, BBB)
          BBB = (AAA, ZZZ)
          ZZZ = (ZZZ, ZZZ)
        HEREDOC

        assert_equal(6, Adventofcode::Day08::Network.load(input).count)
      end

      def test_count_all
        input = <<~HEREDOC
          LR

          11A = (11B, XXX)
          11B = (XXX, 11Z)
          11Z = (11B, XXX)
          22A = (22B, XXX)
          22B = (22C, 22C)
          22C = (22Z, 22Z)
          22Z = (22B, 22B)
          XXX = (XXX, XXX)
        HEREDOC

        assert_equal(6, Adventofcode::Day08::Network.load(input).count_all)
      end
    end
  end
end
