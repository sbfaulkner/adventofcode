# frozen_string_literal: true

module Adventofcode
  module Day02
    extend self

    INPUT = File.join(__dir__, "day02", "input.txt")

    BAG = { red: 12, green: 13, blue: 14 }

    class Game
      attr_reader :id

      class Set
        def initialize(text)
          @set = text.split(", ").each_with_object(Hash.new(0)) do |cubes, set|
            count, color = cubes.split(" ")
            set[color.to_sym] = count.to_i
          end
        end

        def each(&block)
          @set.each(&block)
        end

        def possible?(bag)
          @set.all? { |color, count| bag[color.to_sym] >= count }
        end
      end

      def initialize(text)
        game, sets = text.split(": ")
        @id = game.split(" ").last.to_i
        @sets = sets.split("; ").map { |set| Set.new(set) }
      end

      def possible?(bag = BAG)
        @sets.all? { |set| set.possible?(bag) }
      end

      def minimum_power
        minimum.values.reduce(:*)
      end

      private

      def minimum
        @sets.each_with_object(Hash.new(0)) do |set, minimum|
          set.each do |color, count|
            minimum[color] = count if minimum[color] < count
          end
        end
      end
    end

    def sum(input = File.open(INPUT), bag: BAG)
      input.each_line.map do |line|
        game = Adventofcode::Day02::Game.new(line.chomp)
        game.possible?(bag) ? game.id : 0
      end.sum
    end

    def sum_of_power(input = File.open(INPUT))
      input.each_line.map do |line|
        Adventofcode::Day02::Game.new(line.chomp).minimum_power
      end.sum
    end
  end
end
