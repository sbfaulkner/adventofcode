class Ingredient
  REGEX = /([A-Z][a-z]*): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)/

  def self.load(source)
    source.each_line.map do |ingredient|
      name, capacity, durability, flavor, texture, calories = ingredient.match(REGEX).captures
      Ingredient.new(name, capacity: capacity.to_i,
                           durability: durability.to_i,
                           flavor: flavor.to_i,
                           texture: texture.to_i,
                           calories: calories.to_i)
    end
  end

  def initialize(name, calories: 0, capacity: 0, durability: 0, flavor: 0, texture: 0)
    @name       = name
    @capacity   = capacity
    @durability = durability
    @flavor     = flavor
    @texture    = texture
    @calories   = calories
  end

  attr_reader :capacity, :calories, :durability, :flavor, :texture
end
