class Replacement
  def self.load(source)
    result = [] of self

    source.each_line do |replacement|
      break unless match = replacement.match(/([a-z]+) => ([a-z]+)/i)
      result << new(match[1], match[2])
    end

    result
  end

  def initialize(pattern, replacement)
    @pattern     = Regex.new(pattern)
    @replacement = replacement
  end

  def each_molecule(molecule)
    molecule.replace(@pattern, @replacement) do |m|
      yield m
    end
  end
end
