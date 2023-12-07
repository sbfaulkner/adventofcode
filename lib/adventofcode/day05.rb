# frozen_string_literal: true

module Adventofcode
  module Day05
    extend self

    INPUT = File.join(__dir__, "day05", "input.txt")

    class Almanac
      class Map
        Entry = Struct.new(:dest, :src, :len)

        def initialize(lines)
          @map = []

          loop do
            line = lines.next

            break if line.empty?

            @map << Entry.new(*line.split(/ +/).map(&:to_i))
          end

          @map.sort_by!(&:src)
        end

        def [](src)
          entry = bsearch(src)

          return src unless entry

          entry.dest + src - entry.src
        end

        private

        def bsearch(src)
          @map.bsearch do |entry|
            if entry.src > src
              -1
            elsif (entry.src + entry.len) > src
              0
            else
              1
            end
          end
        end
      end

      def initialize(input = File.open(INPUT))
        lines = input.each_line(chomp: true)

        loop do
          line = lines.next

          key, values = line.split(/: */)

          case key
          when "seeds"
            @seeds = values.split(/ +/).map(&:to_i)
          when "seed-to-soil map"
            @seed_to_soil = Map.new(lines)
          when "soil-to-fertilizer map"
            @soil_to_fertilizer = Map.new(lines)
          when "fertilizer-to-water map"
            @fertilizer_to_water = Map.new(lines)
          when "water-to-light map"
            @water_to_light = Map.new(lines)
          when "light-to-temperature map"
            @light_to_temperature = Map.new(lines)
          when "temperature-to-humidity map"
            @temperature_to_humidity = Map.new(lines)
          when "humidity-to-location map"
            @humidity_to_location = Map.new(lines)
          end
        end
      end

      def lowest_location
        @seeds.map { |seed| @seed_to_soil[seed] }
          .map { |soil| @soil_to_fertilizer[soil] }
          .map { |fertilizer| @fertilizer_to_water[fertilizer] }
          .map { |water| @water_to_light[water] }
          .map { |light| @light_to_temperature[light] }
          .map { |temperature| @temperature_to_humidity[temperature] }
          .map { |humidity| @humidity_to_location[humidity] }
          .min
      end
    end
  end
end
