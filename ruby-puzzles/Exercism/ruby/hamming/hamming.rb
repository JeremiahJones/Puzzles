module BookKeeping
  VERSION = 3
end
# Determines the "hamming" count between 2 DNA strands
class Hamming
  def self.compute(strand1, strand2)
    hamming_distance = 0
    strand1_array = strand1.split('')
    strand2_array = strand2.split('')
    raise_strand_exceptions(strand1_array, strand2_array)

    strand1_array.each_with_index do |value1, index1|
      strand2_array.each_with_index do |value2, index2|
        if index2 == index1
          hamming_distance += 1 if value2 != value1
        end
      end
    end
    hamming_distance
  end
end

def raise_strand_exceptions(strand1, strand2)
  raise ArgumentError if strand1.length != strand2.length
end
