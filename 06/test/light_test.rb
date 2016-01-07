require 'minitest/autorun'
require_relative '../light'

class LightTest < Minitest::Test
  def setup
    @light = Light.new
  end

  def test_initially_off
    refute @light.value
  end

  def test_turn_on
    @light.turn_on
    assert @light.value
  end

  def test_turn_on_again
    @light.turn_on
    @light.turn_on
    assert @light.value
  end

  def test_toggle_on
    @light.toggle
    assert @light.value
  end

  def test_toggle_off
    @light.turn_on
    @light.toggle
    refute @light.value
  end

  def test_turn_off
    @light.turn_on
    @light.turn_off
    refute @light.value
  end

  def test_turn_off_again
    @light.turn_on
    @light.turn_off
    @light.turn_off
    refute @light.value
  end
end
