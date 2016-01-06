require 'minitest/autorun'
require_relative 'route'

class RouteWithBotTest < Minitest::Test
  def setup
    @route = Route.new(2)
  end

  def test_up_and_down_with_bot
    assert_equal(3, @route.deliver('^v'))
  end

  def test_circle_with_bot
    assert_equal(3, @route.deliver('^>v<'))
  end

  def test_back_and_forth_with_bot
    assert_equal(11, @route.deliver('^v^v^v^v^v'))
  end
end
