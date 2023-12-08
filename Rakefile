# frozen_string_literal: true

require "bundler/gem_tasks"
require "rake/testtask"

Rake::TestTask.new(:test) do |t|
  t.libs << "test"
  t.libs << "lib"
  t.test_files = FileList["test/**/test_*.rb"]
end

require "rubocop/rake_task"

RuboCop::RakeTask.new

task default: [:test, :rubocop]

require "adventofcode"

def run(heading, &block)
  print("#{heading}: ")
  start_at = Time.now
  print(yield)
ensure
  puts " [#{Time.now - start_at}s]"
end

desc "Day 1: Trebuchet?!"
task :day01 do
  run("Part 1") { Adventofcode::Day01.sum }
  run("Part 2") { Adventofcode::Day01.sum(spelled: true) }
end

desc "Day 2: Cube Conundrum"
task :day02 do
  run("Part 1") { Adventofcode::Day02.sum }
  run("Part 2") { Adventofcode::Day02.sum_of_power }
end

desc "Day 3: Gear Ratios"
task :day03 do
  run("Part 1") { Adventofcode::Day03::Schematic.new.sum }
  run("Part 2") { Adventofcode::Day03::Schematic.new.sum_gears }
end

desc "Day 4: Scratchcards"
task :day04 do
  run("Part 1") { Adventofcode::Day04.sum }
  run("Part 2") { Adventofcode::Day04.count }
end

desc "Day 5: If You Give A Seed A Fertilizer"
task :day05 do
  run("Part 1") { Adventofcode::Day05::Almanac.new.lowest_location }
  run("Part 2") { Adventofcode::Day05::Almanac.new(ranges: true).lowest_location }
end

desc "Day 6: Wait For It"
task :day06 do
  run("Part 1") { Adventofcode::Day06::Race.load.map(&:winners).reduce(&:*) }
  run("Part 2") { Adventofcode::Day06::Race.load(split: false).map(&:winners).reduce(&:*) }
end

desc "Day 7: Camel Cards"
task :day07 do
  run("Part 1") { Adventofcode::Day07::Camel.total_winnings(Adventofcode::Day07::Camel.load) }
  run("Part 2") { Adventofcode::Day07::Camel.total_winnings(Adventofcode::Day07::Camel.load(jokers: true)) }
end
