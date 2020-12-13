# --- Day 9: All in a Single Night ---
#
# Every year, Santa manages to deliver all of his presents in a single night.
#
# This year, however, he has some new locations to visit; his elves have provided him the distances between every pair
# of locations. He can start and end at any two (different) locations he wants, but he must visit each location exactly
# once. What is the shortest distance he can travel to achieve this?
#
# For example, given the following distances:
#
# London to Dublin = 464
# London to Belfast = 518
# Dublin to Belfast = 141
#
# The possible routes are therefore:
#
# Dublin -> London -> Belfast = 982
# London -> Dublin -> Belfast = 605
# London -> Belfast -> Dublin = 659
# Dublin -> Belfast -> London = 659
# Belfast -> Dublin -> London = 605
# Belfast -> London -> Dublin = 982
#
# The shortest of these is London -> Dublin -> Belfast = 605, and so the answer is 605 in this example.
#
# What is the distance of the shortest route?
#

INPUT_PATH = File.expand_path('input', __dir__)
INPUT = File.read(INPUT_PATH)

require 'benchmark'
require_relative 'route'

answer = nil

time = Benchmark.realtime do
  Route.load(INPUT)
  route = Route.all.min_by(&:length)
  answer = route.length
end

STDERR.printf "Part 1: answer=%d (%.3fms elapsed)\n", answer, time * 1000

time = Benchmark.realtime do
  Route.load(INPUT)
  route = Route.all.max_by(&:length)
  answer = route.length
end

STDERR.printf "Part 2: answer=%d (%.3fms elapsed)\n", answer, time * 1000
