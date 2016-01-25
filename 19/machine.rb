require_relative 'replacement'
require_relative 'molecule'

class Machine
  def self.load(source)
    replacements = Replacement.load(source)
    molecule = Molecule.load(source)

    new(molecule, replacements)
  end

  def initialize(molecule, replacements)
    @molecule     = molecule
    @replacements = replacements
  end

  def calibrate
    generate.tap { |r| STDERR.puts "calibrating => #{r.inspect}" }.size
  end

  def generate
    result = []
    @replacements.each do |replacement|
      replacement.each_molecule(@molecule) { |m| result |= [m] }
    end
    result
  end
end
