require "minitest/autorun"
require "minitest/focus"
require "minitest/reporters"

Minitest::Reporters.use!

require_relative "day01"

class Day01Test < Minitest::Test
  def test_part1_example1
    input = Day01.new("testdata/example1.txt")
    assert_equal(11, input.part1)
  end

  def test_part1
    input = Day01.new("testdata/input.txt")
    assert_equal(1651298, input.part1)
  end

  def test_part2_example1
    input = Day01.new("testdata/example1.txt")
    assert_equal(31, input.part2)
  end

  def test_part2
    input = Day01.new("testdata/input.txt")
    assert_equal(21306195, input.part2)
  end
end
