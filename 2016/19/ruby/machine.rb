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
    @replacements = replacements
  end

  def calibrate
    generate.size
  end

  def fabricate
    fabricated = {}
    molecules  = [Molecule.new]

    (1..Float::INFINITY).find do |count|
      current   = molecules
      molecules = []

      STDERR.puts "#{count}: #{current.size} molecule(s)"
      STDERR.puts current.inspect

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
      replacement.each_molecule(molecule) do |m|
        result[m.to_s] ||= m unless block_given? && !yield(m)
      end
    end
    result.values
  end
end
