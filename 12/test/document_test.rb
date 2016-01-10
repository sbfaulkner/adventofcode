require 'minitest/autorun'
require_relative '../document'

class DocumentTest < Minitest::Test
  def test_array
    document = Document.new('[1,2,3]')
    assert_equal(6, document.sum)
  end

  def test_associative_array
    document = Document.new('{"a":2,"b":4}')
    assert_equal(6, document.sum)
  end

  def test_nested_array
    document = Document.new('[[[3]]]')
    assert_equal(3, document.sum)
  end

  def test_nested_associative_array
    document = Document.new('{"a":{"b":4},"c":-1}')
    assert_equal(3, document.sum)
  end

  def test_associative_array_containing_array
    document = Document.new('{"a":[-1,1]}')
    assert_equal(0, document.sum)
  end

  def test_array_containing_associative_array
    document = Document.new('[-1,{"a":1}]')
    assert_equal(0, document.sum)
  end

  def test_empty_array
    document = Document.new('[]')
    assert_equal(0, document.sum)
  end

  def test_empty_associative_array
    document = Document.new('{}')
    assert_equal(0, document.sum)
  end

  def test_nothing_to_ignore
    document = Document.new('[1,2,3]', ignore: 'red')
    assert_equal(6, document.sum)
  end

  def test_contained_object_ignored
    document = Document.new('[1,{"c":"red","b":2},3]', ignore: 'red')
    assert_equal(4, document.sum)
  end

  def test_entire_structure_ignored
    document = Document.new('{"d":"red","e":[1,2,3,4],"f":5}', ignore: 'red')
    assert_equal(0, document.sum)
  end

  def test_value_in_array_not_ignored
    document = Document.new('[1,"red",5]', ignore: 'red')
    assert_equal(6, document.sum)
  end
end
