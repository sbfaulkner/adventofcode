# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day10
    class TestDay10 < Minitest::Test
      SQUARE = <<~HEREDOC
        .....
        .S-7.
        .|.|.
        .L-J.
        .....
      HEREDOC

      COMPLEX = <<~HEREDOC
        ..F7.
        .FJ|.
        SJ.L7
        |F--J
        LJ...
      HEREDOC

      def test_square_furthest
        assert_equal(4, Adventofcode::Day10::Pipe.load(SQUARE).furthest)
      end

      def test_complex_furthest
        assert_equal(8, Adventofcode::Day10::Pipe.load(COMPLEX).furthest)
      end
    end
  end
end
