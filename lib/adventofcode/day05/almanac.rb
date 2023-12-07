# frozen_string_literal: true

require_relative "almanac/map"
require_relative "almanac/ranges"

module Adventofcode
  module Day05
    class Almanac
      def initialize(input = File.open(INPUT), ranges: false)
        lines = input.each_line(chomp: true)

        loop do
          line = lines.next

          key, values = line.split(/: */)

          case key
          when "seeds"
            @seeds = seed_ranges(*values.split(/ +/).map(&:to_i), ranges: ranges)
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
        soil = @seed_to_soil.transform(@seeds)
        fertilizer = @soil_to_fertilizer.transform(soil)
        water = @fertilizer_to_water.transform(fertilizer)
        light = @water_to_light.transform(water)
        temperature = @light_to_temperature.transform(light)
        humidity = @temperature_to_humidity.transform(temperature)
        location = @humidity_to_location.transform(humidity)
        location.min
      end

      private

      def seed_ranges(*values, ranges:)
        seeds = Ranges.new

        if ranges
          values.each_slice(2) { |first, len| seeds.insert(first, first + len - 1) }
        else
          values.each { |seed| seeds.insert(seed, seed) }
        end

        seeds
      end
    end
  end
end
