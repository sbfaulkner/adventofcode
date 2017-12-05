require 'minitest/autorun'
require_relative '../machine'

class MachineTest < Minitest::Test
  def test_calibrate
    assert_equal 4, simple_machine.calibrate
  end

  def test_generate_simple
    molecules = %w(HHHH HOHO HOOH OHOH)

    assert_equal molecules, simple_machine.generate.map(&:to_s).sort
  end

  def test_fabricate_hoh
    assert_equal 3, fabrication_machine('HOH').fabricate
  end

  def test_fabricate_hohoho
    assert_equal 6, fabrication_machine('HOHOHO').fabricate
  end

  def test_deconstruct_hoh
    assert_equal 3, fabrication_machine('HOH').deconstruct
  end

  def test_deconstruct_hohoho
    assert_equal 6, fabrication_machine('HOHOHO').deconstruct
  end

  def test_degenerate_hohoho
    molecules = %w(HHOHO HOHHO HOHOH HOHOHe HOHOeO HOHeHO HOeOHO HeHOHO eOHOHO)

    assert_equal molecules, fabrication_machine('HOHOHO').degenerate.map(&:to_s).sort
  end

  private

  def fabrication_machine(target)
    machine(<<-INPUT)
      e => H
      e => O
      H => HO
      H => OH
      O => HH

      #{target}
    INPUT
  end

  def simple_machine
    machine(<<-INPUT)
      H => HO
      H => OH
      O => HH

      HOH
    INPUT
  end

  def machine(source)
    Machine.load(StringIO.new(source))
  end
end
