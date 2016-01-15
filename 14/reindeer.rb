class Reindeer
  REINDEER_REGEX = %r{([A-Z][a-z]+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds}

  def self.load(source)
    source.each_line.map do |reindeer|
      name, velocity, stamina, rest = reindeer.match(REINDEER_REGEX).captures
      Reindeer.new(name: name, velocity: velocity.to_i, stamina: stamina.to_i, rest: rest.to_i)
    end
  end

  def self.race(reindeer, duration: 2503)
    duration.times do
      position = reindeer.map { |r| r.travel(1) }.max
      reindeer.each { |r| r.award if r.position == position }
    end
    points = reindeer.map(&:points).max
    reindeer.select { |r| r.points == points }
  end

  def initialize(name:, velocity:, stamina:, rest:)
    @name     = name
    @velocity = velocity
    @stamina  = stamina
    @rest     = rest
    @points   = 0
    @position = 0
    @time     = 0
  end

  attr_reader :name, :points, :position, :rest, :stamina, :velocity

  def award
    @points += 1
  end

  def travel(time)
    @time += time

    periods = @time / (@stamina + @rest)
    remainder = @time % (@stamina + @rest)
    periods += [remainder.to_f / @stamina, 1].min

    @position = periods * @stamina * velocity
  end
end
