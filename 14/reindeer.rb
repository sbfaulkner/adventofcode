class Reindeer
  REINDEER_REGEX = %r{([A-Z][a-z]+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds}

  def self.load(source)
    source.each_line.map do |reindeer|
      name, velocity, stamina, rest = reindeer.match(REINDEER_REGEX).captures
      Reindeer.new(name: name, velocity: velocity.to_i, stamina: stamina.to_i, rest: rest.to_i)
    end
  end

  def initialize(name:, velocity:, stamina:, rest:)
    @name     = name
    @velocity = velocity
    @stamina  = stamina
    @rest     = rest
  end

  attr_reader :name, :rest, :stamina, :velocity

  def travel(time)
    periods = time / (@stamina + @rest)
    remainder = time % (@stamina + @rest)
    periods += [remainder.to_f / @stamina, 1].min
    periods * @stamina * velocity
  end
end
