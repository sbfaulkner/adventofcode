# frozen_string_literal: true

require 'bundler/gem_tasks'
require 'rake/testtask'

Rake::TestTask.new(:test) do |t|
  t.libs << 'test'
  t.libs << 'lib'
  t.test_files = FileList['test/**/test_*.rb']
end

require 'rubocop/rake_task'

RuboCop::RakeTask.new

task default: %i[test rubocop]

require 'adventofcode'

task :day01 do
  puts 'Part 1:', Adventofcode::Day01.sum
  puts 'Part 2:', Adventofcode::Day01.sum(spelled: true)
end

task :day02 do
  puts 'Part 1:', Adventofcode::Day02.sum(bag: { red: 12, green: 13, blue: 14 })
  # puts 'Part 2:', Adventofcode::Day02.sum(spelled: true)
end
