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

desc "Day 1: Trebuchet?!"
task :day01 do
  puts "Part 1:", Adventofcode::Day01.sum
  puts "Part 2:", Adventofcode::Day01.sum(spelled: true)
end

desc "Day 2: Cube Conundrum"
task :day02 do
  puts "Part 1:", Adventofcode::Day02.sum
  puts "Part 2:", Adventofcode::Day02.sum_of_power
end
