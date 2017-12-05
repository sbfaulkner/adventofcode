require 'minitest/autorun'
require_relative '../password'

class PasswordTest < Minitest::Test
  def test_increment_xx
    password = Password.new('xx')
    password.increment
    assert_equal('xy', password.value)
  end

  def test_increment_xxz
    password = Password.new('xxz')
    password.increment
    assert_equal('xya', password.value)
  end

  def test_increment_xxzzz
    password = Password.new('xxzzz')
    password.increment
    assert_equal('xyaaa', password.value)
  end

  def test_abc
    password = Password.new('abc')
    assert password.straight?
  end

  def test_abd
    password = Password.new('abd')
    refute password.straight?
  end

  def test_i
    password = Password.new('i')
    assert password.confusing?
  end

  def test_o
    password = Password.new('o')
    assert password.confusing?
  end

  def test_l
    password = Password.new('l')
    assert password.confusing?
  end

  def test_x
    password = Password.new('x')
    refute password.confusing?
  end

  def test_aa
    password = Password.new('aa')
    refute password.pairs?
  end

  def test_aabb
    password = Password.new('aabb')
    assert password.pairs?
  end

  def test_hijklmmn
    password = Password.new('hijklmmn')
    refute password.valid?
  end

  def test_abbceffg
    password = Password.new('abbceffg')
    refute password.valid?
  end

  def test_abbcegjk
    password = Password.new('abbcegjk')
    refute password.valid?
  end

  def test_abcdffaa
    password = Password.new('abcdffaa')
    assert password.valid?
  end

  def test_ghjaabcc
    password = Password.new('ghjaabcc')
    assert password.valid?
  end

  def test_abcdefgh_next
    password = Password.new('abcdefgh').next
    assert_equal('abcdffaa', password.value)
  end

  def test_ghijklmn_next
    password = Password.new('ghijklmn').next
    assert_equal('ghjaabcc', password.value)
  end
end
