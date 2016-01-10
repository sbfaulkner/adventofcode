require 'minitest/autorun'
require_relative '../player'

class PlayerTest < Minitest::Test
  def test_1
    assert_equal('11', Player.new('1').say)
  end

  def test_11
    assert_equal('21', Player.new('11').say)
  end

  def test_21
    assert_equal('1211', Player.new('21').say)
  end

  def test_1211
    assert_equal('111221', Player.new('1211').say)
  end

  def test_111221
    assert_equal('312211', Player.new('111221').say)
  end
end
