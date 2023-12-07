# frozen_string_literal: true

module Adventofcode
  module Day05
    class Almanac
      class Map
        Entry = Struct.new(:first, :last, :offset, :left, :right)

        class << self
          def load(lines)
            new.tap { |map| map.load(lines) }
          end
        end

        def initialize
          @map = nil
        end

        def load(lines)
          loop do
            line = lines.next

            break if line.empty?

            dest, src, len = line.split(/ +/).map(&:to_i)

            @map = insert(@map, src, src + len - 1, dest - src)
          end
        end

        def [](src)
          get(@map, src)
        end

        private

        def get(entry, src)
          return src unless entry

          if src < entry.first
            get(entry.left, src)
          elsif src > entry.last
            get(entry.right, src)
          else
            src + entry.offset
          end
        end

        def insert(entry, first, last, offset)
          return Entry.new(first, last, offset) unless entry

          if last < entry.first
            entry.left = insert(entry.left, first, last, offset)
          elsif first > entry.last
            entry.right = insert(entry.right, first, last, offset)
          else
            raise NotImplementedError, "insert into"
          end

          entry
        end
      end
    end
  end
end
