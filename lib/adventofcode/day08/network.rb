# frozen_string_literal: true

require "forwardable"

module Adventofcode
  module Day08
    class Network
      include Enumerable
      extend Forwardable

      Entry = Struct.new(:left, :right)

      class << self
        def load(input = File.open(INPUT))
          lines = input.each_line(chomp: true)

          instructions = lines.next
          lines.next # skip blank line

          new(instructions) do |network|
            loop do
              node, left, right = lines.next.scan(/(\w+) = \((\w+), (\w+)\)/).flatten

              network[node] = Entry.new(left, right)
            end
          end
        end
      end

      def_delegators(:@nodes, :[]=, :[])

      def initialize(instructions, &block)
        @instructions = instructions
        @nodes = {}
        block&.call(self)
      end

      def count
        steps = 0

        node = "AAA"

        while node != "ZZZ"
          instruction = @instructions[steps % @instructions.length]

          case instruction
          when "L"
            node = @nodes[node].left
          when "R"
            node = @nodes[node].right
          end

          steps += 1

          break if node == "ZZZ"
        end

        steps
      end
    end
  end
end
