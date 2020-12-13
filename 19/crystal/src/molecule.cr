class Molecule
  def initialize(@structure = "e")
  end

  def_equals_and_hash @structure

  def replace(pattern, replacement)
    offset = 0
    while match = @structure.match(pattern, offset)
      yield self.class.new("#{match.pre_match}#{replacement}#{match.post_match}")
      offset = match.begin(0) + 1
    end
  end

  def inspect(io)
    @structure.inspect(io)
  end

  def to_s
    @structure
  end
end
