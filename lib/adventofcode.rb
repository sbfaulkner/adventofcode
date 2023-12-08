# frozen_string_literal: true

require_relative "adventofcode/version"

Dir[File.join(__dir__, "adventofcode", "day*.rb")].sort.each { |file| require file }

module Adventofcode
  class Error < StandardError; end
  # Your code goes here...
end
