# frozen_string_literal: true

module Adventofcode
  module Day07
    class Camel
      class << self
        def load(input = File.open(INPUT), jokers: false)
          input.each_line(chomp: true).map do |line|
            hand, bid = line.split(/ +/)
            new(hand, bid: bid.to_i, jokers: jokers)
          end
        end

        def total_winnings(hands)
          hands.sort.map.with_index { |h, i| h.bid * (i + 1) }.reduce(&:+)
        end
      end

      CARDS = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"]
      CARDS_WITH_JOKERS = ["J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"]

      attr_reader(:bid, :hand)
      # let other hands access to compare
      protected attr_reader(:cards, :sets)

      def initialize(hand, bid:, jokers:)
        @hand = hand
        cards = jokers ? CARDS_WITH_JOKERS : CARDS
        @cards = hand.chars.map { |c| cards.index(c) }
        tally = hand.chars.tally
        bonus = tally.delete("J").to_i if jokers
        @sets = tally.values.sort.reverse
        @sets[0] = @sets[0].to_i + bonus if jokers
        @bid = bid
      end

      def <=>(other)
        [sets, cards] <=> [other.sets, other.cards]
      end
    end
  end
end
