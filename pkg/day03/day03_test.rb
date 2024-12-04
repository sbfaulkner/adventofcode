require "minitest/autorun"
require "minitest/focus"
require "minitest/reporters"

Minitest::Reporters.use!

require_relative "day03"

class Day03Test < Minitest::Test
  def test_part1_example1
    input = Day03.new("testdata/example1.txt")
    assert_equal(161, input.part1)
  end

  def test_part1
    input = Day03.new("testdata/input.txt")
    assert_equal(189600467, input.part1)
  end

  def test_part2_example2
    input = Day03.new("testdata/example2.txt")
    assert_equal(48, input.part2)
  end

  def test_part2
    input = Day03.new("testdata/input.txt")
    assert_equal(107069718, input.part2)
  end
end
