require 'minitest/autorun'
require_relative '../reindeer'

class ReindeerTest < Minitest::Test
  def setup
    @comet, @dancer = Reindeer.load <<-INPUT
      Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
      Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
    INPUT
  end

  def test_comet
    assert_equal('Comet', @comet.name)
    assert_equal(14, @comet.velocity)
    assert_equal(10, @comet.stamina)
    assert_equal(127, @comet.rest)
  end

  def test_dancer
    assert_equal('Dancer', @dancer.name)
    assert_equal(16, @dancer.velocity)
    assert_equal(11, @dancer.stamina)
    assert_equal(162, @dancer.rest)
  end

  def test_travel_1s
    time = 1
    assert_equal(14, @comet.travel(time))
    assert_equal(16, @dancer.travel(time))
  end

  def test_travel_10s
    time = 10
    assert_equal(140, @comet.travel(time))
    assert_equal(160, @dancer.travel(time))
  end

  def test_travel_11s
    time = 11
    assert_equal(140, @comet.travel(time))
    assert_equal(176, @dancer.travel(time))
  end

  def test_travel_1000s
    time = 1000
    assert_equal(1120, @comet.travel(time))
    assert_equal(1056, @dancer.travel(time))
  end
end
