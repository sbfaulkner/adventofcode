class SeatingArrangement
  RULE_REGEX = /([A-Za-z]+) would (gain|lose) (\d+) happiness units by sitting next to ([A-Za-z]+)/

  def self.rules
    @rules ||= {}
  end

  def self.load(source)
    @rules = Hash.new { |h, k| h[k] = Hash.new(0) }.tap do |rules|
      source.each_line.map do |rule|
        person, action, happiness, neighbour = rule.match(RULE_REGEX).captures
        rules[person][neighbour] = (action == 'gain' ? happiness : "-#{happiness}").to_i
      end
    end
  end

  def self.all(including: nil)
    people = @rules.keys
    people += Array(including)
    people.permutation.map { |assignments| new(assignments) }
  end

  def initialize(assignments)
    @assignments = assignments
    @size        = @assignments.size
  end

  def happiness
    @happiness ||= @assignments.each_with_index.map do |assignment, index|
      happiness = self.class.rules[assignment]
      left      = @assignments[(index - 1) % @size]
      right     = @assignments[(index + 1) % @size]
      happiness[left] + happiness[right]
    end.reduce(&:+)
  end
end
