require 'strscan'

class Molecule
  def self.load(source)
    new(source.gets.strip)
  end

  def initialize(structure)
    @structure = structure
  end

  attr_reader :structure

  def eql?(other)
    other.hash == hash
  end

  def each(pattern, replacement)
    offset = 0
    while match = @structure.match(pattern, offset)
      yield self.class.new("#{match.pre_match}#{replacement}#{match.post_match}")
      offset = match.begin(0) + 1
    end
  end

  def hash
    @structure.hash
  end
end
