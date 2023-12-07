# frozen_string_literal: true

require_relative "ranges"

module Adventofcode
  module Day05
    class Almanac
      class Map < Ranges
        class << self
          def load(lines)
            new.tap { |map| map.load(lines) }
          end
        end

        def load(lines)
          loop do
            line = lines.next

            break if line.empty?

            dest, src, len = line.split(/ +/).map(&:to_i)

            insert(src, src + len - 1, value: dest - src)
          end
        end

        def transform(intervals)
          transformed = Ranges.new

          intervals.each do |interval|
            first = interval.first
            last = interval.last

            if (mapping = find(first, last))
              first += mapping.value
              last += mapping.value
            end

            transformed.insert(first, last)
          end

          transformed
        end
      end
    end
  end
end
