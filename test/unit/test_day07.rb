# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day07
    class TestDay07 < Minitest::Test
      INPUT = <<~HEREDOC
        32T3K 765
        T55J5 684
        KK677 28
        KTJJT 220
        QQQJA 483
      HEREDOC

      def test_winners
        assert_equal(6440, Adventofcode::Day07::Camel.total_winnings(Adventofcode::Day07::Camel.load(INPUT)))
      end

      def test_winners_with_jokers
        assert_equal(
          5905,
          Adventofcode::Day07::Camel.total_winnings(Adventofcode::Day07::Camel.load(INPUT, jokers: true)),
        )
      end
    end
  end
end
