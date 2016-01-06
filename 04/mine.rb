require 'digest'

class Mine
  include Enumerable

  DEFAULT_PREFIX = '00000'

  def initialize(secret, prefix = DEFAULT_PREFIX)
    @secret = secret
    @regex = /\A#{prefix}/
  end

  def each
    (1..Float::INFINITY).each do |value|
      yield value if Digest::MD5.hexdigest("#{@secret}#{value}") =~ @regex
    end
  end
end
