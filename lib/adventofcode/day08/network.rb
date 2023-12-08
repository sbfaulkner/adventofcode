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

      def each(&block)
        return enum_for(:each) unless block_given?

        node = "AAA"

        while node != "ZZZ"
          @instructions.each_char do |c|
            case c
            when "L"
              dest = @nodes[node].left
            when "R"
              dest = @nodes[node].right
            end

            block.call(node, dest)

            node = dest

            break if node == "ZZZ"
          end
        end
      end
    end
  end
end
