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
        count_from("AAA", to: /\AZZZ\z/)
      end

      def count_all
        @nodes.keys.grep(/A\z/).map { |node| count_from(node, to: /Z\z/) }.reduce { |a, e| a.lcm(e) }
      end

      private

      def count_from(node, to:)
        steps = 0

        until node.match?(to)
          instruction = @instructions[steps % @instructions.length]

          case instruction
          when "L"
            node = @nodes[node].left
          when "R"
            node = @nodes[node].right
          end

          steps += 1

          break if node.match?(to)
        end

        steps
      end
    end
  end
end
