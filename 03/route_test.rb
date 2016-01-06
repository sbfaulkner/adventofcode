require 'minitest/autorun'
require_relative 'route'

class RouteTest < Minitest::Test
  def setup
    @route = Route.new
  end

  def test_east
    assert_equal(2, @route.deliver('>'))
  end

  def test_circle
    assert_equal(4, @route.deliver('^>v<'))
  end

  def test_back_and_forth
    assert_equal(2, @route.deliver('^v^v^v^v^v'))
  end
end
