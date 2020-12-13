require 'strscan'

class Aunt
  def self.find(attributes, gt: [], lt: [])
    gt = gt.map(&:to_sym)
    lt = lt.map(&:to_sym)
    @aunts.find do |aunt|
      attributes.all? do |name, value|
        name = name.to_sym
        next true unless aunt.attributes.include?(name)
        if gt.include?(name)
          aunt.attributes[name] > value
        elsif lt.include?(name)
          aunt.attributes[name] < value
        else
          aunt.attributes[name] == value
        end
      end
    end
  end

  def self.load(source)
    @aunts = source.each_line.map { |aunt| parse(aunt) }
  end

  # e.g.
  # Sue 1: goldfish: 9, cars: 0, samoyeds: 9
  def self.parse(aunt)
    scanner = StringScanner.new(aunt)
    if scanner.scan(/\s*([A-Z][a-z]+)\s+(\d+):\s+/)
      attributes = { id: scanner[2].to_i, name: scanner[1] }
      while scanner.scan(/([a-z]+):\s+(\d+)(?:, )?/)
        attributes[scanner[1].to_sym] = scanner[2].to_i
      end
      Aunt.new(attributes)
    end
  end

  def initialize(attributes = {})
    @attributes = attributes.map { |name, value| [name.to_sym, value] }.to_h
  end

  attr_reader :attributes

  def method_missing(method_name, *_args)
    if @attributes.include?(method_name.to_sym)
      @attributes[method_name.to_sym]
    else
      super
    end
  end

  def respond_to_missing?(method_name, _include_private = false)
    @attributes.include?(method_name.to_sym) || super
  end
end
