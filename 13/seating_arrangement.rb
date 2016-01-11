class SeatingArrangement
  RULE_REGEX = /([A-Za-z]+) would (gain|lose) (\d+) happiness units by sitting next to ([A-Za-z]+)/

  def self.rules
    @rules ||= {}
  end

  def self.load(source)
    @rules = Hash.new { |h, k| h[k] = {} }.tap do |rules|
      source.each_line.map do |rule|
        person, action, happiness, neighbour = rule.match(RULE_REGEX).captures
        rules[person][neighbour] = (action == 'gain' ? happiness : "-#{happiness}").to_i
      end
    end
  end

  def self.all
    @rules.keys.permutation.map { |people| new(people) }
  end

  def initialize(people)
    @people = people
    @size   = @people.size
  end

  def happiness
    @happiness ||= @people.each_with_index.map do |person, index|
      happiness = self.class.rules[person]
      left      = @people[(index - 1) % @size]
      right     = @people[(index + 1) % @size]
      happiness[left] + happiness[right]
    end.reduce(&:+)
  end
end
