class Password
  STRAIGHTS = 'abcdefghijklmnopqrstuvwxyz'.chars.each_cons(3).map(&:join)

  def initialize(value)
    @value = value
  end

  attr_accessor :value

  def confusing?
    @value =~ /[iol]/
  end

  def increment
    @value.succ!
  end

  def invalid?
    !valid?
  end

  def next
    Password.new(@value).tap do |password|
      loop do
        password.increment
        break if password.valid?
      end
    end
  end

  def pairs?
    @value.scan(/([a-z])(\1)/).uniq.count >= 2
  end

  def straight?
    STRAIGHTS.any? { |straight| @value.include?(straight) }
  end

  def valid?
    straight? && !confusing? && pairs?
  end
end
