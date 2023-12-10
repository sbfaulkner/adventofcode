# frozen_string_literal: true

module Adventofcode
  module Day09
    class History
      class << self
        def load(input = File.open(INPUT))
          input.each_line(chomp: true).map do |line|
            new(*line.split(" ").map(&:to_i))
          end
        end
      end

      def initialize(*values)
        @expanded = [values.dup]
        @expanded << @expanded.last.each_cons(2).map { |a, b| b - a } until @expanded.last.all?(&:zero?)
      end

      def next
        @expanded.last << 0
        (@expanded.length - 1).times.reverse_each do |i|
          @expanded[i - 1] << @expanded[i - 1].last + @expanded[i].last
        end
        @expanded.first.last
      end
    end
  end
end
