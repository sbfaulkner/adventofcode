require 'minitest/autorun'
require_relative '../circuit'

class CircuitTest < Minitest::Test
  # 123 -> x
  # 456 -> y
  # x AND y -> d
  # x OR y -> e
  # x LSHIFT 2 -> f
  # y RSHIFT 2 -> g
  # NOT x -> h
  # NOT y -> i

  # d: 72
  # e: 507
  # f: 492
  # g: 114
  # h: 65412
  # i: 65079
  # x: 123
  # y: 456

  def setup
    @circuit = Circuit.new
  end

  def test_signal
    @circuit.connect <<-INPUT
      123 -> x
    INPUT
    assert_equal(123, @circuit.signal('x'))
  end

  def test_indirect_signal
    @circuit.connect <<-INPUT
      123 -> x
      x -> y
    INPUT
    assert_equal(123, @circuit.signal('y'))
  end

  def test_and
    @circuit.connect <<-INPUT
      123 -> x
      456 -> y
      x AND y -> d
    INPUT
    assert_equal(72, @circuit.signal('d'))
  end

  def test_or
    @circuit.connect <<-INPUT
      123 -> x
      456 -> y
      x OR y -> e
    INPUT
    assert_equal(507, @circuit.signal('e'))
  end

  def test_lshift
    @circuit.connect <<-INPUT
      123 -> x
      x LSHIFT 2 -> f
    INPUT
    assert_equal(492, @circuit.signal('f'))
  end

  def test_rshift
    @circuit.connect <<-INPUT
      456 -> y
      y RSHIFT 2 -> g
    INPUT
    assert_equal(114, @circuit.signal('g'))
  end

  def test_rshift_zero
    @circuit.connect <<-INPUT
      32768 RSHIFT 1 -> r
    INPUT
    assert_equal(16_384, @circuit.signal('r'))
  end

  def test_16bit_shifting
    @circuit.connect <<-INPUT
      65535 -> a
      a LSHIFT 1 -> b
      b RSHIFT 1 -> c
    INPUT
    assert_equal(32_767, @circuit.signal('c'))
  end

  def test_not
    @circuit.connect <<-INPUT
      123 -> x
      NOT x -> h
    INPUT
    assert_equal(65_412, @circuit.signal('h'))
  end
end
