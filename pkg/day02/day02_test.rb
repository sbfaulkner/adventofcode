require "minitest/autorun"
require "minitest/focus"
require "minitest/reporters"

Minitest::Reporters.use!

require_relative "day02"

class Day02Test < Minitest::Test
  def test_part1_example1
    input = Day02.new("testdata/example1.txt")
    assert_equal(2, input.part1)
  end

  def test_part1
    input = Day02.new("testdata/input.txt")
    assert_equal(534, input.part1)
  end

  def test_part2_example1
    input = Day02.new("testdata/example1.txt")
    assert_equal(4, input.part2)
  end

  def test_part2
    input = Day02.new("testdata/input.txt")
    assert_equal(577, input.part2)
  end
end
