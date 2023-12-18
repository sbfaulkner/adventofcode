# frozen_string_literal: true

require "test_helper"

module Adventofcode
  module Day10
    class TestDay10 < Minitest::Test
      def test_square_furthest
        input = <<~HEREDOC
          .....
          .S-7.
          .|.|.
          .L-J.
          .....
        HEREDOC

        assert_equal(4, Adventofcode::Day10::Pipe.load(input).furthest)
      end

      def test_complex_furthest
        input = <<~HEREDOC
          ..F7.
          .FJ|.
          SJ.L7
          |F--J
          LJ...
        HEREDOC

        assert_equal(8, Adventofcode::Day10::Pipe.load(input).furthest)
      end

      def test_simple_enclosed
        input = <<~HEREDOC
          ...........
          .S-------7.
          .|F-----7|.
          .||.....||.
          .||.....||.
          .|L-7.F-J|.
          .|..|.|..|.
          .L--J.L--J.
          ...........
        HEREDOC

        assert_equal(4, Adventofcode::Day10::Pipe.load(input).enclosed)
      end

      def test_unexposed_enclosed
        input = <<~HEREDOC
          ..........
          .S------7.
          .|F----7|.
          .||....||.
          .||....||.
          .|L-7F-J|.
          .|..||..|.
          .L--JL--J.
          ..........
        HEREDOC

        assert_equal(4, Adventofcode::Day10::Pipe.load(input).enclosed)
      end

      def test_larger_enclosed
        input = <<~HEREDOC
          .F----7F7F7F7F-7....
          .|F--7||||||||FJ....
          .||.FJ||||||||L7....
          FJL7L7LJLJ||LJ.L-7..
          L--J.L7...LJS7F-7L7.
          ....F-J..F7FJ|L7L7L7
          ....L7.F7||L7|.L7L7|
          .....|FJLJ|FJ|F7|.LJ
          ....FJL-7.||.||||...
          ....L---J.LJ.LJLJ...
        HEREDOC

        assert_equal(8, Adventofcode::Day10::Pipe.load(input).enclosed)
      end

      def test_junk_enclosed
        input = <<~HEREDOC
          FF7FSF7F7F7F7F7F---7
          L|LJ||||||||||||F--J
          FL-7LJLJ||||||LJL-77
          F--JF--7||LJLJ7F7FJ-
          L---JF-JLJ.||-FJLJJ7
          |F|F-JF---7F7-L7L|7|
          |FFJF7L7F-JF7|JL---7
          7-L-JL7||F7|L7F-7F7|
          L.L7LFJ|||||FJL7||LJ
          L7JLJL-JLJLJL--JLJ.L
        HEREDOC

        assert_equal(10, Adventofcode::Day10::Pipe.load(input).enclosed)
      end
    end
  end
end
