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
        @expanded.last.push(0)
        (@expanded.length - 2).downto(0) do |i|
          @expanded[i].push(@expanded[i].last + @expanded[i + 1].last)
        end
        @expanded.first.last
      end

      def previous
        @expanded.last.unshift(0)
        (@expanded.length - 2).downto(0) do |i|
          @expanded[i].unshift(@expanded[i].first - @expanded[i + 1].first)
        end
        @expanded.first.first
      end
    end
  end
end
