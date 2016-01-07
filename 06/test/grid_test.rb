require 'minitest/autorun'
require_relative '../grid'

class GridTest < Minitest::Test
  def setup
    @grid = Grid.new
  end

  def test_turn_on_all
    @grid.configure('turn on 0,0 through 999,999')
    assert_equal(1_000_000, @grid.count)
  end

  def test_toggle_1000
    @grid.configure('turn on 0,0 through 999,999')
    @grid.configure('toggle 0,0 through 999,0')
    assert_equal(999_000, @grid.count)
  end

  def test_turn_off_4
    @grid.configure('turn on 0,0 through 999,999')
    @grid.configure('turn off 499,499 through 500,500')
    assert_equal(999_996, @grid.count)
  end
end
