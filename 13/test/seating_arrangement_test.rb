require 'minitest/autorun'
require_relative '../seating_arrangement'

class SeatingArrangementTest < Minitest::Test
  def setup
    SeatingArrangement.load <<-INPUT
      Alice would gain 54 happiness units by sitting next to Bob.
      Alice would lose 79 happiness units by sitting next to Carol.
      Alice would lose 2 happiness units by sitting next to David.
      Bob would gain 83 happiness units by sitting next to Alice.
      Bob would lose 7 happiness units by sitting next to Carol.
      Bob would lose 63 happiness units by sitting next to David.
      Carol would lose 62 happiness units by sitting next to Alice.
      Carol would gain 60 happiness units by sitting next to Bob.
      Carol would gain 55 happiness units by sitting next to David.
      David would gain 46 happiness units by sitting next to Alice.
      David would lose 7 happiness units by sitting next to Bob.
      David would gain 41 happiness units by sitting next to Carol.
    INPUT
  end

  def test_load
    rules = {
      'Alice' => {
        'Bob' => 54,
        'Carol' => -79,
        'David' => -2
      },
      'Bob' => {
        'Alice' => 83,
        'Carol' => -7,
        'David' => -63
      },
      'Carol' => {
        'Alice' => -62,
        'Bob' => 60,
        'David' => 55
      },
      'David' => {
        'Alice' => 46,
        'Bob' => -7,
        'Carol' => 41
      }
    }

    assert_equal(rules, SeatingArrangement.rules)
  end

  def test_all
    arrangements = SeatingArrangement.all
    happiness = [
      -114,
      -114,
      -114,
      -114,
      -114,
      -114,
      -114,
      -114,
      22,
      22,
      22,
      22,
      22,
      22,
      22,
      22,
      330,
      330,
      330,
      330,
      330,
      330,
      330,
      330
    ]
    assert_equal(happiness, arrangements.map(&:happiness).sort)
  end
end
