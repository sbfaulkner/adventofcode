require 'minitest/autorun'
require_relative '../aunt'

class AuntTest < Minitest::Test
  def test_aunt_attributes
    attributes = { id: 1, name: 'Sue', cats: 3 }
    aunt = Aunt.new(attributes)
    assert_equal(attributes, aunt.attributes)
  end

  def test_aunt_cats
    attributes = { id: 1, name: 'Sue', cats: 3 }
    aunt = Aunt.new(attributes)
    assert_equal(3, aunt.cats)
  end

  def test_aunt_id
    attributes = { id: 1, name: 'Sue', cats: 3 }
    aunt = Aunt.new(attributes)
    assert_equal(1, aunt.id)
  end

  def test_aunt_name
    attributes = { id: 1, name: 'Sue', cats: 3 }
    aunt = Aunt.new(attributes)
    assert_equal('Sue', aunt.name)
  end

  def test_parse
    aunt = Aunt.parse('Sue 132: cars: 1, vizslas: 3, children: 7')
    assert_instance_of(Aunt, aunt)
    assert_equal(132, aunt.id)
    assert_equal('Sue', aunt.name)
    assert_equal(1, aunt.cars)
    assert_equal(3, aunt.vizslas)
    assert_equal(7, aunt.children)
  end

  def test_load
    aunts = Aunt.load <<-INPUT
      Sue 1: goldfish: 9, cars: 0, samoyeds: 9
      Sue 2: perfumes: 5, trees: 8, goldfish: 8
      Sue 3: goldfish: 9, akitas: 1, trees: 5
    INPUT
    assert_instance_of(Array, aunts)
    assert_instance_of(Aunt, aunts[0])
    assert_equal(1, aunts[0].id)
    assert_equal(5, aunts[1].perfumes)
    assert_equal(5, aunts[2].trees)
  end

  def test_find_first
    Aunt.load <<-INPUT
      Sue 1: goldfish: 9, cars: 0, samoyeds: 9
      Sue 2: perfumes: 5, trees: 8, goldfish: 8
      Sue 3: goldfish: 9, akitas: 1, cars: 5
    INPUT
    assert_equal(1, Aunt.find(goldfish: 9).id)
  end

  def test_find_with_missing_attributes
    Aunt.load <<-INPUT
      Sue 1: goldfish: 9, cars: 0, samoyeds: 9
      Sue 2: perfumes: 5, trees: 8, goldfish: 8
      Sue 3: goldfish: 9, akitas: 1, cars: 5
    INPUT
    assert_equal(1, Aunt.find(goldfish: 9, akitas: 1).id)
  end

  def test_find_with_multiple_attributes
    Aunt.load <<-INPUT
      Sue 1: goldfish: 9, cars: 0, samoyeds: 9
      Sue 2: perfumes: 5, trees: 8, goldfish: 8
      Sue 3: goldfish: 9, akitas: 1, cars: 5
    INPUT
    assert_equal(3, Aunt.find(goldfish: 9, akitas: 1, cars: 5).id)
  end
end
