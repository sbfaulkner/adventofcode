require 'minitest/autorun'
require_relative '../variable_light'

class VariableLightTest < Minitest::Test
  def setup
    @light = VariableLight.new
  end

  def test_initially_off
    assert_equal 0, @light.brightness
  end

  def test_turn_on
    @light.turn_on
    assert_equal 1, @light.brightness
  end

  def test_turn_on_again
    @light.turn_on
    @light.turn_on
    assert_equal 2, @light.brightness
  end

  def test_toggle_on
    @light.toggle
    assert_equal 2, @light.brightness
  end

  def test_turn_off
    @light.turn_on
    @light.turn_off
    assert_equal 0, @light.brightness
  end

  def test_turn_off_again
    @light.turn_on
    @light.turn_off
    @light.turn_off
    assert_equal 0, @light.brightness
  end
end
