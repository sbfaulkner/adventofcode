require "./spec_helper"

module MachineBuilder
  def self.fabrication_machine(target)
    machine <<-INPUT
      e => H
      e => O
      H => HO
      H => OH
      O => HH

      #{target}
    INPUT
  end

  def self.simple_machine
    machine <<-INPUT
      H => HO
      H => OH
      O => HH

      HOH
    INPUT
  end

  def self.machine(source)
    Machine.load(MemoryIO.new(source))
  end
end

describe Machine do
  context "with a simple machine" do
    it "calibrates" do
      machine = MachineBuilder.simple_machine
      machine.calibrate.should eq 4
    end

    it "generates molecules" do
      molecules = %w(HHHH HOHO HOOH OHOH)
      machine = MachineBuilder.simple_machine
      machine.generate.map(&.to_s).sort.should eq molecules
    end

    it "fabricates HOH in 3 generations" do
      machine = MachineBuilder.fabrication_machine("HOH")
      machine.fabricate.should eq 3
    end

    it "fabricates HOHOHO in 6 generations" do
      machine = MachineBuilder.fabrication_machine("HOHOHO")
      machine.fabricate.should eq 6
    end
  end
end
