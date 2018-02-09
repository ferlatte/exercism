# frozen_string_literal: true

module BookKeeping
  VERSION = 3
end

# Hamming
class Hamming
  def self.compute(a, b)
    raise ArgumentError, 'a and b must be the same length' if a.length != b.length
    distance = 0
    aa = a.chars
    ba = b.chars
    aa.each_index { |x| distance += 1 if aa[x] != ba[x] }
    distance
  end
end
