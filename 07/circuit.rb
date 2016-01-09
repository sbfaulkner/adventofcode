class Circuit
  def initialize
    @wires = {}
  end

  def connect(connections)
    connections.each_line do |connection|
      parse(connection)
    end
  end

  def signal(wire)
    return wire if wire.is_a?(Integer)
    return wire.to_i if wire =~ /\d+/

    value = @wires[wire]

    return value if value.is_a?(Integer)

    @wires[wire] =
    case value[0]
    when 'AND'
      signal(value[1]) & signal(value[2])
    when 'OR'
      signal(value[1]) | signal(value[2])
    when 'LSHIFT'
      (signal(value[1]) << signal(value[2])) & 0xFFFF
    when 'RSHIFT'
      signal(value[1]) >> signal(value[2])
    when 'NOT'
      (~ signal(value[1])) & 0xFFFF
    when /\d+/
      value[0].to_i
    else
      signal(value[0])
    end
  end

  private

  def parse(connection)
    case connection.strip
    when /\A([a-z]+|\d+) -> ([a-z]+)\z/
      _, input, wire = Regexp.last_match.to_a
      @wires[wire] = [input]
    when /\A([a-z]+|\d+) (AND|OR|LSHIFT|RSHIFT) ([a-z]+|\d+) -> ([a-z]+)\z/
      _, input1, operator, input2, wire = Regexp.last_match.to_a
      @wires[wire] = [operator, input1, input2]
    when /\A(NOT) ([a-z]+|\d+) -> ([a-z]+)\z/
      _, operator, input, wire = Regexp.last_match.to_a
      @wires[wire] = [operator, input]
    end
  end
end
