require 'minitest/autorun'
require_relative '../molecule'

class MoleculeTest < Minitest::Test
  def test_hash_equal
    hash = Molecule.new('A').hash
    assert_equal hash, Molecule.new('A').hash
  end

  def test_hash_not_equal
    hash = Molecule.new('A').hash
    refute_equal hash, Molecule.new('B').hash
  end

  def test_equality
    molecule = Molecule.new('A')
    assert Molecule.new('A').eql?(molecule)
  end

  def test_inequality
    molecule = Molecule.new('A')
    refute Molecule.new('B').eql?(molecule)
  end

  def test_to_s
    assert 'HOHOHO', Molecule.new('HOHOHO').to_s
  end

  def test_each
    molecules = %w(HOHH HHOH HHHO).each

    Molecule.new('HHH').each('H', 'HO') do |m|
      assert_equal molecules.next, m.to_s
    end
  end
end
