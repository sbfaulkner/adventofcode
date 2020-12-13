#!/usr/bin/env ruby
# frozen_string_literal: true

require 'set'

INPUT = File.expand_path('input', __dir__)

changes = File.readlines(INPUT).map(&:to_i)
count = changes.size
frequencies = Set.new
frequency = 0
index = 0

until frequencies.include?(frequency)
  frequencies.add(frequency)

  change = changes[index % count]
  frequency += change

  index += 1
end

STDERR.puts "Frequency: #{frequency}"
