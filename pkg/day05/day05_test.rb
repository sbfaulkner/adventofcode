require "minitest/autorun"
require "minitest/focus"
require "minitest/reporters"

Minitest::Reporters.use!

require_relative "day05"

class Day05Test < Minitest::Test
  def test_part1_example1
    input = Day05.new("testdata/example1.txt")
    assert_equal(0, input.part1)
  end

  def test_part1
    input = Day05.new("testdata/input.txt")
    assert_equal(0, input.part1)
  end

  def test_part2_example1
    input = Day05.new("testdata/example1.txt")
    assert_equal(0, input.part2)
  end

  def test_part2
    input = Day05.new("testdata/input.txt")
    assert_equal(0, input.part2)
  end
end
