require "./spec_helper"

describe Replacement do
  it "loads input text" do
    replacements = Replacement.load <<-INPUT
      H => HO
      H => OH

      IGNORETHIS
    INPUT

    replacements.should be_a(Array(Replacement))
    replacements.size.should eq 2
  end

  it "replaces molecules" do
    molecules = %w(HOHH HHOH HHHO).each

    Replacement.new("H", "HO").each_molecule(Molecule.new("HHH")) do |m|
      m.to_s.should eq molecules.next
    end
  end
end
