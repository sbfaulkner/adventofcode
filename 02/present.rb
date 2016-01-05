class Present
  attr_reader :length, :width, :height

  def initialize(dimensions)
    @length, @width, @height = dimensions.split('x').map(&:to_i).sort.reverse
  end

  def paper_required
    length * width * 2 + length * height * 2 + width * height * 3
  end

  def ribbon_required
    2 * (width + height) + ribbon_required_for_bow
  end

  private

  def ribbon_required_for_bow
    length * width * height
  end
end
