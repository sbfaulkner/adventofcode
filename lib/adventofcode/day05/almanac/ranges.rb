# frozen_string_literal: true

module Adventofcode
  module Day05
    class Almanac
      class Ranges
        Entry = Struct.new(:first, :last, :value, :left, :right)

        def initialize
          @root = nil
        end

        def each(&block)
          return enum_for(:each) unless block_given?

          each_entry(@root, &block)
        end

        def insert(first, last, value: nil)
          @root = insert_entry(@root, first, last, value)
        end

        def min
          entry = @root

          return unless entry

          entry = entry.left while entry.left

          entry.first
        end

        private

        def each_entry(entry, &block)
          return unless entry

          each_entry(entry.left, &block)
          yield(entry)
          each_entry(entry.right, &block)
        end

        def find_entry(entry, first, last)
          return unless entry

          if last < entry.first
            find_entry(entry.left, first, last)
          elsif first > entry.last
            find_entry(entry.right, first, last)
          elsif first < entry.first && last > entry.last
            [
              find_entry(entry, first, entry.first - 1),
              entry,
              find_entry(entry, entry.last + 1, last),
            ]
          elsif first < entry.first
            [
              find_entry(entry, first, entry.first - 1),
              entry,
            ]
          elsif last > entry.last
            [
              entry,
              find_entry(entry, entry.last + 1, last),
            ]
          else
            [Entry.new(first, last, entry.value)]
          end
        end

        def insert_entry(entry, first, last, value)
          return Entry.new(first, last, value) unless entry

          if last < entry.first
            entry.left = insert_entry(entry.left, first, last, value)
          elsif first > entry.last
            entry.right = insert_entry(entry.right, first, last, value)
          else
            raise NotImplementedError, "insert_entry into"
          end

          entry
        end
      end
    end
  end
end
