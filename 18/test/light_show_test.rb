require 'minitest/autorun'
require_relative '../light_show'

class LightShowTest < Minitest::Test
  LIGHT_SHOW_1 = <<-INPUT
    ###...
    ###...
    ......
    ..###.
    ..###.
    ..###.
  INPUT

  LIGHT_SHOW_2A = <<-INPUT
    .#.#.#
    ...##.
    #....#
    ..#...
    #.#..#
    ####..
  INPUT

  LIGHT_SHOW_2B = <<-INPUT
    ..##..
    ..##.#
    ...##.
    ......
    #.....
    #.##..
  INPUT

  LIGHT_SHOW_2C = <<-INPUT
    ..###.
    ......
    ..###.
    ......
    .#....
    .#....
  INPUT

  LIGHT_SHOW_2D = <<-INPUT
    ...#..
    ......
    ...#..
    ..##..
    ......
    ......
  INPUT

  LIGHT_SHOW_2E = <<-INPUT
    ......
    ......
    ..##..
    ..##..
    ......
    ......
  INPUT

  def test_load
    load_show_1

    lights = [
      [true, true, true, false, false, false],
      [true, true, true, false, false, false],
      [false, false, false, false, false, false],
      [false, false, true, true, true, false],
      [false, false, true, true, true, false],
      [false, false, true, true, true, false]
    ]

    assert_equal lights, @show.lights
  end

  def test_edge_neighbours
    load_show_1

    assert_equal 5, @show.neighbours(0, 1)
  end

  def test_inner_neighbours
    load_show_1

    assert_equal 8, @show.neighbours(4, 3)
  end

  def test_to_s
    load_show_2

    assert_equal @show.lights, LightShow.load(@show.to_s).lights
  end

  def test_1_step
    load_show_2
    @show.step

    assert_equal LightShow.load(LIGHT_SHOW_2B).lights, @show.lights
  end

  def test_2_steps
    load_show_2
    2.times { @show.step }

    assert_equal LightShow.load(LIGHT_SHOW_2C).lights, @show.lights
  end

  def test_3_steps
    load_show_2
    3.times { @show.step }

    assert_equal LightShow.load(LIGHT_SHOW_2D).lights, @show.lights
  end

  def test_4_steps
    load_show_2
    4.times { @show.step }

    assert_equal LightShow.load(LIGHT_SHOW_2E).lights, @show.lights
  end

  def test_count_after_4_steps
    load_show_2
    4.times { @show.step }

    assert_equal 4, @show.count
  end

  private

  def clean(input)
    input.each_line.map(&:strip).join("\n")
  end

  def load_show_1
    @show = LightShow.load(LIGHT_SHOW_1)
  end

  def load_show_2
    @show = LightShow.load(LIGHT_SHOW_2A)
  end
end
