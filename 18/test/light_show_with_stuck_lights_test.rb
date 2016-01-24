require 'minitest/autorun'
require_relative '../light_show'

class LightShowWithStuckLightsTest < Minitest::Test
  LIGHT_SHOW_A = <<-INPUT
    ##.#.#
    ...##.
    #....#
    ..#...
    #.#..#
    ####.#
  INPUT

  LIGHT_SHOW_B = <<-INPUT
    #.##.#
    ####.#
    ...##.
    ......
    #...#.
    #.####
  INPUT

  LIGHT_SHOW_C = <<-INPUT
    #..#.#
    #....#
    .#.##.
    ...##.
    .#..##
    ##.###
  INPUT

  LIGHT_SHOW_D = <<-INPUT
    #...##
    ####.#
    ..##.#
    ......
    ##....
    ####.#
  INPUT

  LIGHT_SHOW_E = <<-INPUT
    #.####
    #....#
    ...#..
    .##...
    #.....
    #.#..#
  INPUT

  LIGHT_SHOW_F = <<-INPUT
    ##.###
    .##..#
    .##...
    .##...
    #.#...
    ##...#
  INPUT

  def test_1_step
    load_show
    @show.step

    assert_equal LightShow.load(LIGHT_SHOW_B).lights, @show.lights
  end

  def test_2_steps
    load_show
    2.times { @show.step }

    assert_equal LightShow.load(LIGHT_SHOW_C).lights, @show.lights
  end

  def test_3_steps
    load_show
    3.times { @show.step }

    assert_equal LightShow.load(LIGHT_SHOW_D).lights, @show.lights
  end

  def test_4_steps
    load_show
    4.times { @show.step }

    assert_equal LightShow.load(LIGHT_SHOW_E).lights, @show.lights
  end

  def test_5_steps
    load_show
    5.times { @show.step }

    assert_equal LightShow.load(LIGHT_SHOW_F).lights, @show.lights
  end

  def test_count_after_5_steps
    load_show
    5.times { @show.step }

    assert_equal 17, @show.count
  end

  private

  def clean(input)
    input.each_line.map(&:strip).join("\n")
  end

  def load_show
    @show = LightShow.load(LIGHT_SHOW_A, always_on: [[0, 0], [0, 5], [5, 0], [5, 5]])
  end
end
