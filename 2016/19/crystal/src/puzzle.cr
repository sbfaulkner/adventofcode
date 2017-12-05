require "benchmark"
require "./machine"

# INPUT_PATH = File.expand_path("input", __DIR__)
#
# answer = nil
#
# time = Benchmark.realtime do
#   machine = Machine.load(File.open(INPUT_PATH))
#   answer = machine.calibrate
# end
#
# STDERR.printf "Part 1: answer=%d (%.3fms elapsed)\n", answer, time.total_milliseconds
#
# time = Benchmark.realtime do
#   machine = Machine.load(File.open(INPUT_PATH))
#   answer = machine.fabricate
# end
#
# STDERR.printf "Part 2: answer=%d (%.3fms elapsed)\n", answer, time.total_milliseconds
