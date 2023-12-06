# frozen_string_literal: true

require "strscan"

module Adventofcode
  module Day04
    extend self

    INPUT = File.join(__dir__, "day04", "input.txt")

    def sum(input = File.open(INPUT))
      input.each_line.sum do |line|
        # scratch_card: winning numbers | numbers
        _card, numbers = line.chomp.split(/: +/)

        # winning numbers and numbers
        winning, numbers = numbers.split(/ +\| +/)

        # convert to arrays of integers
        winning = winning.split(/ +/).map(&:to_i)
        numbers = numbers.split(/ +/).map(&:to_i)

        # intersection of winning numbers and numbers
        winners = (winning & numbers)

        # 2 to the power of the number of winners
        winners.empty? ? 0 : 2**(winners.length - 1)
      end
    end
  end
end
