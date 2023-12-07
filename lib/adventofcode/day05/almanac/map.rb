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
            transform_each(@root, interval.first, interval.last) do |first, last|
              transformed.insert(first, last)
            end
          end

          transformed
        end

        private

        def transform_each(entry, first, last, &block)
          unless entry
            block.call(first, last)
            return
          end

          if last < entry.first
            transform_each(entry.left, first, last, &block)
          elsif first > entry.last
            transform_each(entry.right, first, last, &block)
          else
            transform_each(entry, first, entry.first - 1, &block) if first < entry.first
            block.call([entry.first, first].max + entry.value, [entry.last, last].min + entry.value)
            transform_each(entry, entry.last + 1, last, &block) if last > entry.last
          end
        end
      end
    end
  end
end
