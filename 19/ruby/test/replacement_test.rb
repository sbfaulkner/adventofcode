require 'minitest/autorun'
require_relative '../replacement'
require_relative '../molecule'

class ReplacementTest < Minitest::Test
  def test_load
    replacements = Replacement.load <<-INPUT
      H => HO
      H => OH

      IGNORETHIS
    INPUT

    assert_instance_of Array, replacements
    assert_equal 2, replacements.size
    assert_instance_of Replacement, replacements.first
  end

  def test_each_molecule
    molecules = %w(HOHH HHOH HHHO).each

    Replacement.new('H', 'HO').each_molecule(Molecule.new('HHH')) do |m|
      assert_equal molecules.next, m.to_s
    end
  end
end
