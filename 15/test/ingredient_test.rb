require 'minitest/autorun'
require_relative '../ingredient'

class IngredientTest < Minitest::Test
  def setup
    @ingredients = Ingredient.load <<-INPUT
      Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
      Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
    INPUT
  end

  def test_load
    assert_instance_of(Array, @ingredients)
    assert_equal(2, @ingredients.size)
    assert_instance_of(Ingredient, @ingredients.first)
  end
end
