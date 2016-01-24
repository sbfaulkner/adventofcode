class Replacement
  def self.load(source)
    source.each_line.with_object([]) do |replacement, result|
      break result unless match = replacement.match(/([a-z]+) => ([a-z]+)/i)
      result << new(match[1], match[2])
    end
  end

  def initialize(pattern, replacement)
    @pattern     = Regexp.new(pattern)
    @replacement = replacement
  end

  attr_reader :pattern, :replacement
end
