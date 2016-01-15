require 'minitest/autorun'
require_relative '../reindeer'

class ReindeerTest < Minitest::Test
  def setup
    @reindeer = Reindeer.load <<-INPUT
      Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
      Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
    INPUT
    @comet, @dancer = @reindeer
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

  def test_travel_accumulates
    1000.times do
      @comet.travel(1)
      @dancer.travel(1)
    end
    assert_equal(1120, @comet.position)
    assert_equal(1056, @dancer.position)
  end

  def test_award_1_point
    assert_equal(1, @comet.award)
  end

  def test_award_2_points
    @comet.award
    assert_equal(2, @comet.award)
  end

  def test_award_5_points
    5.times { @comet.award }
    assert_equal(5, @comet.points)
  end

  def test_race_1s
    Reindeer.race(@reindeer, duration: 1)
    assert_equal(0, @comet.points)
    assert_equal(1, @dancer.points)
  end

  def test_race_140s
    Reindeer.race(@reindeer, duration: 140)
    assert_equal(1, @comet.points)
    assert_equal(139, @dancer.points)
  end

  def test_race_1000s
    winner = Reindeer.race(@reindeer, duration: 1000)
    assert_equal(312, @comet.points)
    assert_equal(689, @dancer.points)
    assert_equal([@dancer], winner)
  end
end
