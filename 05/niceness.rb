module Niceness
  V1 = 1
  V2 = 2

  def nice?(version = V2)
    case version
    when V1
      match(/(?:[aeiou].*){3,}/) && match(/(.)\1/) && !match(/ab|cd|pq|xy/)
    when V2
      match(/(..).*\1/) && match(/(.).\1/)
    else
      raise ArgumentError, "Unsupported version (#{version})"
    end
  end

  def naughty?(version = V2)
    !nice?(version)
  end
end

String.include Niceness
