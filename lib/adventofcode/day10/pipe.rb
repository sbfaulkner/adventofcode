# frozen_string_literal: true

module Adventofcode
  module Day10
    class Pipe
      class << self
        def load(input = File.open(INPUT))
          lines = input.each_line(chomp: true).to_a
          new(lines)
        end
      end

      STARTING_POINT = "S"
      VERTICAL_PIPE = "|"
      HORIZONTAL_PIPE = "-"
      NORTH_AND_EAST_BEND = "L"
      NORTH_AND_WEST_BEND = "J"
      SOUTH_AND_WEST_BEND = "7"
      SOUTH_AND_EAST_BEND = "F"
      GROUND = "."

      NORTH = [0, -1]
      EAST = [1, 0]
      SOUTH = [0, 1]
      WEST = [-1, 0]

      def initialize(lines)
        @lines = lines
        @height = lines.length
        @width = lines.first.length
        @ymax = @height - 1
        @xmax = @width - 1

        @pipes = Array.new(@height * @width)

        @start = find_starting_point

        pos = @start

        @path = []

        until @path.first == pos
          case get_at(*pos)
          when STARTING_POINT
            # all four directions are possible, but only two will be valid
            raise "Unexpected starting point at #{pos}"
          when VERTICAL_PIPE
            next_pos = get_pos(*pos, direction: NORTH)
            next_pos = get_pos(*pos, direction: SOUTH) if @path.last == next_pos
          when HORIZONTAL_PIPE
            next_pos = get_pos(*pos, direction: EAST)
            next_pos = get_pos(*pos, direction: WEST) if @path.last == next_pos
          when NORTH_AND_EAST_BEND
            next_pos = get_pos(*pos, direction: NORTH)
            next_pos = get_pos(*pos, direction: EAST) if @path.last == next_pos
          when NORTH_AND_WEST_BEND
            next_pos = get_pos(*pos, direction: NORTH)
            next_pos = get_pos(*pos, direction: WEST) if @path.last == next_pos
          when SOUTH_AND_WEST_BEND
            next_pos = get_pos(*pos, direction: SOUTH)
            next_pos = get_pos(*pos, direction: WEST) if @path.last == next_pos
          when SOUTH_AND_EAST_BEND
            next_pos = get_pos(*pos, direction: SOUTH)
            next_pos = get_pos(*pos, direction: EAST) if @path.last == next_pos
          when GROUND
            raise "Unexpected ground at #{pos}"
          end

          @path << pos
          pos = next_pos
        end
      end

      def furthest
        @path.length / 2
      end

      def enclosed
        # ray from point to exterior
        # if it intersects the path an odd number of times, it's enclosed
        # first optimization...
        # reduce points considered to those within the bounding box of the path
        # likely optimization...
        # draw the ray towards whichever side (of the bounding box) is closest
        # possible optimization...
        # if it hits a point known to be enclosed (or not) - it is as well
        # possible optimization...
        # flood fill areas known to be enclosed (or not)

        enclosed = 0

        @lines.each_with_index do |line, y|
          line.each_char.with_index do |_tile, x|
            # not enclosed if it's on the path
            next if @path.include?([x, y])

            line[x] = GROUND # treat junk as ground

            intersections = 0

            previous = nil

            y.times do |i|
              next unless @path.include?([x, y - i - 1])

              tile = get_at(x, y - i - 1)
              next if tile == GROUND || tile == VERTICAL_PIPE

              intersections += 1 if tile == HORIZONTAL_PIPE ||
                (tile == SOUTH_AND_EAST_BEND && previous == NORTH_AND_WEST_BEND) ||
                (tile == SOUTH_AND_WEST_BEND && previous == NORTH_AND_EAST_BEND)

              previous = tile
            end

            enclosed += 1 if intersections.odd?
          end
        end

        enclosed
      end

      private

      NORTH_FACING = [VERTICAL_PIPE, NORTH_AND_EAST_BEND, NORTH_AND_WEST_BEND]
      SOUTH_FACING = [VERTICAL_PIPE, SOUTH_AND_EAST_BEND, SOUTH_AND_WEST_BEND]
      EAST_FACING = [HORIZONTAL_PIPE, NORTH_AND_EAST_BEND, SOUTH_AND_EAST_BEND]
      WEST_FACING = [HORIZONTAL_PIPE, NORTH_AND_WEST_BEND, SOUTH_AND_WEST_BEND]

      def find_starting_point
        @lines.each_with_index do |line, y|
          line.each_char.with_index do |tile, x|
            if tile == STARTING_POINT
              line[x] = resolve_starting_point(x, y)
              return [x, y]
            end
          end
        end
      end

      def resolve_starting_point(x, y)
        north = SOUTH_FACING.include?(get_at(x, y, direction: NORTH))
        south = NORTH_FACING.include?(get_at(x, y, direction: SOUTH))
        east = WEST_FACING.include?(get_at(x, y, direction: EAST))
        west = EAST_FACING.include?(get_at(x, y, direction: WEST))

        case [north, south, east, west]
        when [true, true, false, false]
          VERTICAL_PIPE
        when [false, false, true, true]
          HORIZONTAL_PIPE
        when [true, false, true, false]
          NORTH_AND_EAST_BEND
        when [true, false, false, true]
          NORTH_AND_WEST_BEND
        when [false, true, true, false]
          SOUTH_AND_EAST_BEND
        when [false, true, false, true]
          SOUTH_AND_WEST_BEND
        else
          raise "Unexpected directions for starting point at #{pos}"
        end
      end

      def get_pos(x, y, direction:)
        dx, dy = direction

        x += dx
        y += dy

        [x, y]
      end

      def get_at(x, y, direction: [0, 0])
        x, y = get_pos(x, y, direction: direction)

        return if x < 0 || x > @xmax
        return if y < 0 || y > @ymax

        @lines[y][x]
      end
    end
  end
end
