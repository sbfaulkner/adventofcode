require 'minitest/autorun'
require_relative '../ingredient'
require_relative '../recipe'

class RecipeTest < Minitest::Test
  def test_new_ingredients
    recipe = Recipe.new

    assert_instance_of(Array, recipe.ingredients)
    assert_empty(recipe.ingredients)
  end

  def test_new_ingredients_added
    sugar = Ingredient.new('sugar')

    recipe = Recipe.new { |r| r.add(1, sugar) }
    amount, ingredient = recipe.ingredients[0]

    assert_equal(1, amount)
    assert_equal(ingredient, sugar)
  end

  def test_combine
    sugar = Ingredient.new('sugar')

    recipes = Recipe.combine([sugar], total: 4)

    assert_instance_of(Array, recipes)
    assert_equal(1, recipes.size)
  end

  # Sugar:     0,1,2,3,4
  # Sprinkles: 4,3,2,1,0
  def test_combine_2_ingredients
    sugar     = Ingredient.new('sugar')
    sprinkles = Ingredient.new('sprinkles')

    recipes = Recipe.combine([sugar, sprinkles], total: 4)

    assert_instance_of(Array, recipes)
    assert_equal(5, recipes.size)
  end

  # Sugar:     0,0,0,0,0,1,1,1,1,2,2,2,3,3,4
  # Sprinkles: 0,1,2,3,4,0,1,2,3,0,1,2,0,1,0
  # Candy:     4,3,2,1,0,3,2,1,0,2,1,0,1,0,0
  def test_combine_3_ingredients
    sugar     = Ingredient.new('sugar')
    sprinkles = Ingredient.new('sprinkles')
    candy     = Ingredient.new('candy')

    recipes = Recipe.combine([sugar, sprinkles, candy], total: 4)

    assert_instance_of(Array, recipes)
    assert_equal(15, recipes.size)
  end

  def test_combined_capacity
    butterscotch, cinnamon = Ingredient.load <<-INPUT
      Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
      Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
    INPUT

    recipe = Recipe.new
    recipe.add(44, butterscotch)
    recipe.add(56, cinnamon)

    assert_equal(68, recipe.capacity)
  end

  def test_combined_durability
    butterscotch, cinnamon = Ingredient.load <<-INPUT
      Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
      Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
    INPUT

    recipe = Recipe.new
    recipe.add(44, butterscotch)
    recipe.add(56, cinnamon)

    assert_equal(80, recipe.durability)
  end

  def test_combined_flavor
    butterscotch, cinnamon = Ingredient.load <<-INPUT
      Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
      Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
    INPUT

    recipe = Recipe.new
    recipe.add(44, butterscotch)
    recipe.add(56, cinnamon)

    assert_equal(152, recipe.flavor)
  end

  def test_combined_texture
    butterscotch, cinnamon = Ingredient.load <<-INPUT
      Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
      Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
    INPUT

    recipe = Recipe.new
    recipe.add(44, butterscotch)
    recipe.add(56, cinnamon)

    assert_equal(76, recipe.texture)
  end

  def test_combined_score
    butterscotch, cinnamon = Ingredient.load <<-INPUT
      Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
      Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
    INPUT

    recipe = Recipe.new
    recipe.add(44, butterscotch)
    recipe.add(56, cinnamon)

    assert_equal(62_842_880, recipe.score)
  end

  def test_combined_score_with_500_calories
    butterscotch, cinnamon = Ingredient.load <<-INPUT
      Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
      Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
    INPUT

    recipe = Recipe.new
    recipe.add(40, butterscotch)
    recipe.add(60, cinnamon)

    assert_equal(500, recipe.calories)
    assert_equal(57_600_000, recipe.score)
  end
end
