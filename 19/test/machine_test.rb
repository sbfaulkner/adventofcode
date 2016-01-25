require 'minitest/autorun'
require_relative '../machine'

class MachineTest < Minitest::Test
  def test_calibrate
    assert_equal 4, machine.calibrate
  end

  def test_generate
    molecules = %w(HHHH HOHO HOOH OHOH)

    assert_equal molecules, machine.generate.map(&:to_s).sort
  end

  private

  def source
    StringIO.new(<<-INPUT)
      H => HO
      H => OH
      O => HH

      HOH
    INPUT
  end

  def machine
    Machine.load(source)
  end
end
