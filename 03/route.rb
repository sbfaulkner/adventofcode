require_relative 'courier'

class Route
  def initialize(couriers = 1)
    @houses = Hash.new { |h, x| h[x] = Hash.new(0) }
    @couriers = Array.new(couriers) { Courier.new }
  end

  def deliver(directions)
    @couriers.each { |courier| courier.deliver_to(@houses) }

    couriers = @couriers.cycle
    courier = couriers.next

    directions.chars.each do |direction|
      next unless courier.move(direction)
      courier.deliver_to(@houses)
      courier = couriers.next
    end

    @houses.values.map(&:values).flatten.count
  end
end
