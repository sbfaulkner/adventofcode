require 'minitest/autorun'
require_relative '../grid'
require_relative '../variable_light'

class VariableGridTest < Minitest::Test
  def setup
    @grid = Grid.new(VariableLight)
  end

  def test_turn_on_1
    @grid.configure('turn on 0,0 through 0,0')
    assert_equal(1, @grid.brightness)
  end

  def test_toggle_all
    @grid.configure('toggle 0,0 through 999,999')
    assert_equal(2_000_000, @grid.brightness)
  end
end
