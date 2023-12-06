# frozen_string_literal: true

require "strscan"

module Adventofcode
  module Day04
    extend self

    INPUT = File.join(__dir__, "day04", "input.txt")

    def sum(input = File.open(INPUT))
      input.each_line.sum do |line|
        _card, winners = parse_card(line)

        # 2 to the power of the number of winners
        winners.empty? ? 0 : 2**(winners.length - 1)
      end
    end

    def count(input = File.open(INPUT))
      tally = Hash.new(0)

      input.each_line do |line|
        card, winners = parse_card(line)

        # account for the original card
        tally[card - 1] += 1

        # increment the tally of subsequent cards based on # of winners
        winners.length.times do |i|
          tally[card + i] += tally[card - 1]
        end
      end

      tally.values.sum
    end

    def parse_card(line)
      # scratch_card: winning numbers | numbers
      card, numbers = line.chomp.split(/: +/)

      # card number
      card = card.split(/ +/).last.to_i

      # winning numbers and numbers
      winning, numbers = numbers.split(/ +\| +/)

      # convert to arrays of integers
      winning = winning.split(/ +/).map(&:to_i)
      numbers = numbers.split(/ +/).map(&:to_i)

      # intersection of winning numbers and numbers
      winners = (winning & numbers)

      [card, winners]
    end
  end
end
