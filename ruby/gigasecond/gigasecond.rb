# frozen_string_literal: true

module BookKeeping
  VERSION = 6
end

# Returns a Time that is a gigasecond from Time t
class Gigasecond
  def self.from(t)
    t + 1_000_000_000
  end
end
