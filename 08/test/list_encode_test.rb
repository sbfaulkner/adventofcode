require 'minitest/autorun'
require_relative '../list'

class ListTest < Minitest::Test
  def setup
    @list = List.new(:encode)
  end

  def test_size_of_empty_string
    count = @list.read <<-'INPUT'
      ""
    INPUT
    assert_equal(2, count)
    assert_equal(6, @list.first.size)
  end

  def test_size_of_string
    count = @list.read <<-'INPUT'
      "abc"
    INPUT
    assert_equal(5, count)
    assert_equal(9, @list.first.size)
  end

  def test_size_of_string_with_escaped_quote
    count = @list.read <<-'INPUT'
      "aaa\"aaa"
    INPUT
    assert_equal(10, count)
    assert_equal(16, @list.first.size)
  end

  def test_size_of_string_with_escaped_backslash
    count = @list.read <<-'INPUT'
      "\\"
    INPUT
    assert_equal(4, count)
    assert_equal(10, @list.first.size)
  end

  def test_size_of_string_with_escaped_hex_character
    count = @list.read <<-'INPUT'
      "\x27"
    INPUT
    assert_equal(6, count)
    assert_equal(11, @list.first.size)
  end
end
