# frozen_string_literal: true

module Adventofcode
  module Day06
    class Race
      class << self
        def load(input = File.open(INPUT), split: true)
          lines = input.each_line(chomp: true)
          time = lines.next.split(/: */).last
          distance = lines.next.split(/: */).last
          if split
            times = time.split(/ +/).map(&:to_i)
            distances = distance.split(/ +/).map(&:to_i)
            times.zip(distances).map { |time, distance| new(time: time, distance: distance) }
          else
            [new(time: time.delete(" ").to_i, distance: distance.delete(" ").to_i)]
          end
        end
      end

      def initialize(time:, distance:)
        @time = time
        @distance = distance
      end

      def winners
        count = 0
        hold = Math.sqrt(@distance).floor

        while hold * (@time - hold) > @distance
          count += 1
          hold -= 1
        end

        hold = Math.sqrt(@distance).floor + 1

        while hold * (@time - hold) > @distance
          count += 1
          hold += 1
        end

        count
      end
    end
  end
end
