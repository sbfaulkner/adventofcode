require 'minitest/autorun'
require_relative '../present'

class PresentTest < Minitest::Test
  def test_2x3x4_paper_required
    present = Present.new('2x3x4')
    assert_equal(58, present.paper_required)
  end

  def test_2x3x4_ribbon_required
    present = Present.new('2x3x4')
    assert_equal(34, present.ribbon_required)
  end

  def test_1x1x10_paper_required
    present = Present.new('1x1x10')
    assert_equal(43, present.paper_required)
  end

  def test_1x1x10_ribbon_required
    present = Present.new('1x1x10')
    assert_equal(14, present.ribbon_required)
  end
end
