module Adventofcode
  module Day02
    extend self

    INPUT = File.join(__dir__, 'day02', 'input.txt')

    class Game
      attr_reader :id, :sets

      class Set
        def initialize(text)
          @set = text.split(', ').each_with_object(Hash.new(0)) do |cubes, set|
            count, color = cubes.split(' ')
            set[color.to_sym] = count.to_i
          end
        end

        def possible?(bag)
          @set.all? { |color, count| bag[color.to_sym] >= count }
        end
      end

      def initialize(text)
        game, sets = text.split(': ')
        @id = game.split(' ').last.to_i
        @sets = sets.split('; ').map { |set| Set.new(set) }
      end

      def possible?(bag)
        @sets.all? { |set| set.possible?(bag) }
      end
    end

    def sum(input = File.open(INPUT), bag:)
      input.each_line.map do |line|
        game = Adventofcode::Day02::Game.new(line.chomp)
        game.possible?(bag) ? game.id : 0
      end.sum
    end
  end
end
