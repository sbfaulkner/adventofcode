require 'minitest/autorun'
require_relative '../route'

class RouteTest < Minitest::Test
  def setup
    Route.load <<-INPUT
      London to Dublin = 464
      London to Belfast = 518
      Dublin to Belfast = 141
    INPUT
  end

  def test_load
    distances = {
      'London' => {
        'Dublin' => 464,
        'Belfast' => 518,
      },
      'Dublin' => {
        'Belfast' => 141,
        'London' => 464,
      },
      'Belfast' => {
        'London' => 518,
        'Dublin' => 141,
      }
    }
    assert_equal(distances, Route.distances)
  end

  def test_all
    routes = Route.all
    assert_equal([605, 605, 659, 659, 982, 982], routes.map(&:length).sort)
  end
end
