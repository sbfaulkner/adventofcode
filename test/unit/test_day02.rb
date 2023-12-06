# frozen_string_literal: true

require 'test_helper'

module Adventofcode
  module Day02
    class TestDay02 < Minitest::Test
      BAG = { red: 12, green: 13, blue: 14 }

      INPUT = <<~HEREDOC
        Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
        Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
        Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
      HEREDOC

      def test_game_id
        ids = INPUT.each_line.map do |line|
          Adventofcode::Day02::Game.new(line.chomp).id
        end

        assert_equal [1, 2, 3, 4, 5], ids
      end

      def test_game_possible
        possible = INPUT.each_line.map do |line|
          Adventofcode::Day02::Game.new(line.chomp).possible?(BAG)
        end

        assert_equal [true, true, false, false, true], possible
      end

      def test_sum
        assert_equal 8, Adventofcode::Day02.sum(INPUT, bag: BAG)
      end
    end
  end
end
