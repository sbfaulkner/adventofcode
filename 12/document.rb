require 'json'

class Document
  def initialize(source, ignore: nil)
    @contents = JSON.load(source)
    @ignore   = ignore
  end

  def numbers
    values.select { |value| value.is_a?(Numeric) }
  end

  def sum
    numbers.reduce(0, &:+)
  end

  private

  def values(from: @contents)
    case from
    when Hash
      values_from_hash(from)
    when Array
      values_from_array(from)
    else
      Array(from)
    end.flatten
  end

  def values_from_array(array)
    array.map { |item| values(from: item) }
  end

  def values_from_hash(hash)
    hash.map do |_key, item|
      return [] if @ignore && @ignore == item
      values(from: item)
    end
  end
end
