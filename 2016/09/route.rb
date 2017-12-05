class Route
  def self.distances
    @distances ||= {}
  end

  def self.load(paths)
    @distances = Hash.new { |h, k| h[k] = {} }.tap do |distances|
      paths.each_line do |path|
        origin, destination, distance = path.match(/([a-z]+) to ([a-z]+) = (\d+)/i).captures
        distances[origin][destination] = distance.to_i
        distances[destination][origin] = distance.to_i
      end
    end
  end

  def self.all
    @distances.keys.permutation.map { |cities| new(cities) }
  end

  def initialize(cities)
    @cities = cities
  end

  def length
    @length ||= @cities.each_cons(2).map { |o, d| self.class.distances.fetch(o).fetch(d) }.reduce(&:+)
  end
end
