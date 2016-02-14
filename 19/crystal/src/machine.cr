require "./replacement"
require "./molecule"

class Machine
  def self.load(source)
    replacements = Replacement.load(source)
    molecule = Molecule.new(source.gets.to_s.strip)

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
    fabricated = {} of String => Bool
    molecules  = [Molecule.new]

    (1..Int64::MAX).find do |count|
      current   = molecules
      molecules = [] of Molecule

      STDERR.puts "#{count}: #{current.size} molecule(s)"
      STDERR.flush
      STDERR.puts "[#{current.map { |m| m.to_s }.join(", ")}]"
      STDERR.flush

      current.any? do |molecule|
        results = generate(molecule) do |m|
          fabricated[m.to_s] = true unless fabricated[m]?
        end

        STDERR.puts "results => #{results.map { |m| m.to_s }.join(", ")}"
        molecules = molecules | results
        STDERR.puts "molecules => #{molecules.map { |m| m.to_s }.join(", ")}"
        results.any? { |m| @molecule.eql?(m) }
      end
    end
  end

  def generate(molecule = @molecule)
    generate(molecule) { true }
  end

  def generate(molecule = @molecule)
    result = {} of String => Molecule

    @replacements.each do |replacement|
      replacement.each_molecule(molecule) do |m|
        (result[m.to_s]? || (result[m.to_s] = m)) if yield(m)
      end
    end

    result.values
  end
end
