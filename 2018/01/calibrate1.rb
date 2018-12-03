#!/usr/bin/env ruby
# frozen_string_literal: true

INPUT = File.expand_path('input', __dir__)

frequency = File.readlines(INPUT).map(&:to_i).reduce(0, &:+)

STDOUT.puts "Frequency: #{frequency}"
