# frozen_string_literal: true

module Adventofcode
  module Day06
    class Race
      class << self
        def load(input = File.open(INPUT), split: true)
          lines = input.each_line(chomp: true)
          time = lines.next.split(/: */).last
          distance = lines.next.split(/: */).last
          if split
            times = time.split(/ +/).map(&:to_i)
            distances = distance.split(/ +/).map(&:to_i)
            times.zip(distances).map { |time, distance| new(time: time, distance: distance) }
          else
            [new(time: time.delete(" ").to_i, distance: distance.delete(" ").to_i)]
          end
        end
      end

      def initialize(time:, distance:)
        @time = time
        @distance = distance
      end

      def winners
        # Let T be the total time
        # Let H be the time holding the button down (ie. time spent building potential velocity)
        # Let R be the time the button is released (ie. time spent moving at constant velocity)
        # Let V be the velocity
        # Let D be the distance
        #
        # Basic time/distance/velocity equations:
        # T = H + R
        # D = R * V
        # V = H
        #
        # Substituting:
        # D = (T - H) * H
        # D = T*H - H^2
        #
        # Looks like a quadratic equation...
        # ax^2 + bx + c = 0
        #
        # Rearranging:
        # H^2 - T*H + D = 0
        #
        # x = H
        # a = 1
        # b = -T
        # c = D
        #
        # Application of quadratic formula:
        # x = (-b +/- sqrt(b^2 - 4ac)) / 2a
        #
        # Solving for H:
        # H = (T +/- sqrt(T^2 - 4*D)) / 2

        hmax = ((@time + Math.sqrt(@time**2 - 4 * (@distance + 1))) / 2).floor
        hmin = ((@time - Math.sqrt(@time**2 - 4 * (@distance + 1))) / 2).ceil

        hmax - hmin + 1
      end
    end
  end
end
