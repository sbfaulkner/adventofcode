require 'minitest/autorun'
require_relative 'elevator'

class TestElevator < Minitest::Test
  def assert_floor(floor, instructions)
    @elevator.follow(instructions)
    assert_equal(floor, @elevator.floor)
  end

  def setup
    @elevator = Elevator.new
  end

  def test_up
    assert_equal(1, Elevator.direction('('))
  end

  def test_down
    assert_equal(-1, Elevator.direction(')'))
  end

  def test_unknown
    assert_equal(nil, Elevator.direction('?'))
  end

  def test_floor_0
    assert_floor(0, '(())')
  end

  def test_also_floor_0
    assert_floor(0, '()()')
  end

  def test_floor_3
    assert_floor(3, '(((')
  end

  def test_also_floor_3
    assert_floor(3, '(()(()(')
  end

  def test_yet_another_floor_3
    assert_floor(3, '))(((((')
  end

  def test_basement
    assert_floor(-1, '())')
  end

  def test_also_basement
    assert_floor(-1, '))(')
  end

  def test_subfloor_3
    assert_floor(-3, ')))')
  end

  def test_also_subfloor_3
    assert_floor(-3, ')())())')
  end

  def test_find_basement_in_1_instruction
    assert_equal(1, @elevator.follow(')'))
  end

  def test_find_basement_in_5_instructions
    assert_equal(5, @elevator.follow('()())'))
  end

  def test_does_not_count_invalid_instructions
    assert_equal(5, @elevator.follow('  ()())'))
  end
end
