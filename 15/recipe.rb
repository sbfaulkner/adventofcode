class Recipe
  def self.combine(ingredients, total: 100)
    if ingredients.size > 1
      (0..total).map do |amount|
        combine(ingredients[1..-1], total: total - amount).map do |combination|
          Recipe.new do |recipe|
            recipe.add(amount, ingredients[0])
            combination.ingredients.each do |a, i|
              recipe.add(a, i)
            end
          end
        end
      end.flatten
    else
      [Recipe.new { |r| r.add(total, ingredients[0]) }]
    end
  end

  def initialize
    @ingredients = []
    @capacity = @durability = @flavor = @texture = 0
    @score = 0
    yield self if block_given?
  end

  attr_reader :ingredients

  def add(amount, ingredient)
    ingredients << [amount, ingredient]
  end

  def capacity
    [@ingredients.map { |a, i| a * i.capacity }.reduce(0, &:+), 0].max
  end

  def durability
    [@ingredients.map { |a, i| a * i.durability }.reduce(0, &:+), 0].max
  end

  def flavor
    [@ingredients.map { |a, i| a * i.flavor }.reduce(0, &:+), 0].max
  end

  def score
    capacity * durability * flavor * texture
  end

  def texture
    [@ingredients.map { |a, i| a * i.texture }.reduce(0, &:+), 0].max
  end
end
