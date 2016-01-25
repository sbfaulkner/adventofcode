class Replacement
  def self.load(source)
    source.each_line.with_object([]) do |replacement, result|
      break result unless match = replacement.match(/([a-z]+) => ([a-z]+)/i)
      result << new(match[1], match[2])
    end
  end

  def initialize(pattern, replacement)
    @pattern     = pattern
    @replacement = replacement
  end

  def each_molecule(molecule, &block)
    molecule.each(@pattern, @replacement, &block)
  end
end
