require "./spec_helper"

describe Molecule do
  it "has a consistent hash value" do
    hash = Molecule.new("A").hash
    Molecule.new("A").hash.should eq hash
  end

  it "has a different hash value" do
    hash = Molecule.new("A").hash
    Molecule.new("B").hash.should_not eq hash
  end

  it "has equality" do
    Molecule.new("A").should eq Molecule.new("A")
  end

  it "has inequality" do
    Molecule.new("B").should_not eq Molecule.new("A")
  end

  it "has uniqueness" do
    arr = [Molecule.new("A")]
    [Molecule.new("A"), Molecule.new("A")].uniq.should eq arr
  end

  it "can be converted to a string" do
    Molecule.new("HOHOHO").to_s.should eq "HOHOHO"
  end

  it "replaces each matched pattern" do
    molecules = %w(HOHH HHOH HHHO).each

    Molecule.new("HHH").replace(/H/, "HO") do |m|
      m.to_s.should eq molecules.next
    end
  end
end
