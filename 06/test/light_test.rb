require 'minitest/autorun'
require_relative '../light'

class LightTest < Minitest::Test
  def setup
    @light = Light.new
  end

  def test_initially_off
    assert @light.off?
  end

  def test_turn_on
    @light.turn_on
    assert @light.on?
  end

  def test_turn_on_again
    @light.turn_on
    @light.turn_on
    assert @light.on?
  end

  def test_toggle_on
    @light.toggle
    assert @light.on?
  end

  def test_toggle_off
    @light.turn_on
    @light.toggle
    assert @light.off?
  end

  def test_turn_off
    @light.turn_on
    @light.turn_off
    assert @light.off?
  end

  def test_turn_off_again
    @light.turn_on
    @light.turn_off
    @light.turn_off
    assert @light.off?
  end
end
