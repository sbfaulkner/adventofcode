class Player
  def initialize(digits)
    @digits = digits.scan(/(\d)(\1*)/).map(&:join)
  end

  def say
    @digits.map { |segment| "#{segment.length}#{segment[0]}" }.join
  end
end
