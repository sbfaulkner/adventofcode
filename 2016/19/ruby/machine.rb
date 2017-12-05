require_relative 'replacement'
require_relative 'molecule'

class Machine
  def self.load(source)
    replacements = Replacement.load(source)
    molecule = Molecule.new(source.gets.strip)

    new(molecule, replacements)
  end

  def initialize(molecule, replacements)
    @molecule     = molecule
    @replacements = replacements.sort { |a,b| a.replacement.length <=> b.replacement.length }
  end

  def calibrate
    generate.size
  end

  def deconstruct
    deconstructed = {}
    molecules = [@molecule]

    (1..Float::INFINITY).find do |count|
      current   = molecules
      molecules = []

      STDERR.puts "#{count}: #{current.size} molecule(s)"
      # STDERR.puts current.inspect

      current.any? do |molecule|
        results = degenerate(molecule) do |m|
          deconstructed[m.to_s] = true unless deconstructed[m]
        end

        molecules |= results
        results.any? { |m| m.to_s == 'e' }
      end
    end
  end

  def degenerate(molecule = @molecule)
    result = {}
    @replacements.each do |replacement|
      replacement.devolve(molecule) do |m|
        result[m.to_s] ||= m unless block_given? && !yield(m)
      end
    end
    result.values
  end

  def fabricate
    fabricated = {}
    molecules  = [Molecule.new]

    (1..Float::INFINITY).find do |count|
      current   = molecules
      molecules = []

      STDERR.puts "#{count}: #{current.size} molecule(s)"
      # STDERR.puts current.inspect

      current.any? do |molecule|
        results = generate(molecule) do |m|
          fabricated[m.to_s] = true unless fabricated[m]
        end

        molecules |= results
        results.any? { |m| @molecule.eql?(m) }
      end
    end
  end

  def generate(molecule = @molecule)
    result = {}
    @replacements.each do |replacement|
      replacement.evolve(molecule) do |m|
        result[m.to_s] ||= m unless block_given? && !yield(m)
      end
    end
    result.values
  end
end
