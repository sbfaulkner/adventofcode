# --- Day 17: No Such Thing as Too Much ---
#
# The elves bought too much eggnog again - 150 liters this time. To fit it all into your refrigerator, you'll need to
# move it into smaller containers. You take an inventory of the capacities of the available containers.
#
# For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If you need to store 25 liters, there
# are four ways to do it:
#
# - 15 and 10
# - 20 and 5 (the first 5)
# - 20 and 5 (the second 5)
# - 15, 5, and 5
#
# Filling all containers entirely, how many different combinations of containers can exactly fit all 150 liters of
# eggnog?
#
# --- Part Two ---
#
# While playing with all the containers in the kitchen, another load of eggnog arrives! The shipping and receiving
# department is requesting as many containers as you can spare.
#
# Find the minimum number of containers that can exactly fit all 150 liters of eggnog. How many different ways can you
# fill that number of containers and still hold exactly 150 litres?
#
# In the example above, the minimum number of containers was two. There were three ways to use that many containers, and
# so the answer there would be 3.
#

INPUT_PATH = File.expand_path('input', __dir__)

require 'benchmark'

answer = nil

time = Benchmark.realtime do
  containers = File.open(INPUT_PATH).each_line.map(&:to_i)
  answer = (1..containers.size).map do |size|
    containers.combination(size).count { |set| set.reduce(&:+) == 150 }
  end.reduce(&:+)
end

STDERR.printf "Part 1: answer=%d (%.3fms elapsed)\n", answer, time * 1000

containers = File.open(INPUT_PATH).each_line.map(&:to_i)
answer = (1..containers.size).map do |size|
  count = containers.combination(size).count { |set| set.reduce(&:+) == 150 }
  break count if count > 0
end

STDERR.printf "Part 2: answer=%d (%.3fms elapsed)\n", answer, time * 1000
