require 'minitest/autorun'
require_relative 'mine'

class MineTest < Minitest::Test
  # If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes
  # (000001dbbfa...), and it is the lowest such number to do so.
  def test_abcdef
    mine = Mine.new('abcdef')
    assert_equal(609_043, mine.first)
  end

  # If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is
  # 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....
  def test_pqrstuv
    mine = Mine.new('pqrstuv')
    assert_equal(1_048_970, mine.first)
  end
end
