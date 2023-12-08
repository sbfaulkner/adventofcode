# frozen_string_literal: true

module Adventofcode
  module Day06
    class Race
      class << self
        def load(input = File.open(INPUT))
          lines = input.each_line(chomp: true)
          times = lines.next.split(/: */).last.split(/ +/).map(&:to_i)
          distances = lines.next.split(/: */).last.split(/ +/).map(&:to_i)
          times.zip(distances).map { |time, distance| new(time: time, distance: distance) }
        end
      end

      def initialize(time:, distance:)
        @time = time
        @distance = distance
      end

      def winners
        (1...@time).map { |hold| (@time - hold) * hold > @distance }.count(true)
      end
    end
  end
end
