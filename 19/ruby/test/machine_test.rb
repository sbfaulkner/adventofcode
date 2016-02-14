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
