# frozen_string_literal: true

require_relative "almanac/map"

module Adventofcode
  module Day05
    class Almanac
      def initialize(input = File.open(INPUT))
        lines = input.each_line(chomp: true)

        loop do
          line = lines.next

          key, values = line.split(/: */)

          case key
          when "seeds"
            @seeds = values.split(/ +/).map(&:to_i)
          when "seed-to-soil map"
            @seed_to_soil = Map.load(lines)
          when "soil-to-fertilizer map"
            @soil_to_fertilizer = Map.load(lines)
          when "fertilizer-to-water map"
            @fertilizer_to_water = Map.load(lines)
          when "water-to-light map"
            @water_to_light = Map.load(lines)
          when "light-to-temperature map"
            @light_to_temperature = Map.load(lines)
          when "temperature-to-humidity map"
            @temperature_to_humidity = Map.load(lines)
          when "humidity-to-location map"
            @humidity_to_location = Map.load(lines)
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
