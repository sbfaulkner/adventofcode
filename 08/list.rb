class List
  include Enumerable

  def initialize(action)
    @state = :start
    @list = ''
    @action = action
  end

  def each
    @list.each_line do |line|
      yield line.chomp
    end
  end

  def read(input)
    input.each_char.count do |char|
      case @state
      when :start
        read_token(char)
      when :string
        read_string(char)
      when :string_escape
        read_string_escape(char)
      when :string_hex
        read_string_hex(char)
      when :string_hex2
        read_string_hex2(char)
      end
    end
  end

  private

  def encoding?
    @action == :encode
  end

  def read_string(char)
    case char
    when '"'
      @list << '\""' if encoding?
      @list << "\n"
      @state = :start
    when '\\'
      @list << '\\\\' if encoding?
      @state = :string_escape
    else
      @list << char
    end
  end

  def read_string_escape(char)
    case char
    when 'x'
      @list << char if encoding?
      @state = :string_hex
    else
      @list << '\\' if encoding?
      @list << char
      @state = :string
    end
  end

  def read_string_hex(char)
    @list << char if encoding?
    @hex = char
    @state = :string_hex2
  end

  def read_string_hex2(char)
    @hex << char
    @list << (encoding? ? char : @hex.to_i(16).chr)
    @state = :string
  end

  def read_token(char)
    case char
    when '"'
      @list << '"\"' if encoding?
      @state = :string
    end
  end
end
