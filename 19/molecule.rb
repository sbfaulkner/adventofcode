require 'strscan'

class Molecule
  def self.load(source)
    new(source.gets.strip)
  end

  def initialize(structure)
    @structure = structure
  end

  def eql?(molecule)
    molecule.hash == hash
  end

  def generate(replacements)
    replacements.each_with_object([]) do |replacement, result|
      offset = 0
      while match = @structure.match(replacement.pattern, offset)
        result << self.class.new("#{match.pre_match}#{replacement.replacement}#{match.post_match}")
        offset = match.begin(0) + 1
      end
    end.uniq
  end

  def hash
    @structure.hash
  end
end
