require 'minitest/autorun'
require_relative '../mine'

class MineTest < Minitest::Test
  def test_abcdef
    mine = Mine.new('abcdef')
    assert_equal(609_043, mine.first)
  end

  def test_pqrstuv
    mine = Mine.new('pqrstuv')
    assert_equal(1_048_970, mine.first)
  end
end
