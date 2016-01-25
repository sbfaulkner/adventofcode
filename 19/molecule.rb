require 'strscan'

class Molecule
  def initialize(structure)
    @structure = structure
  end

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

  def to_s
    @structure
  end
end
