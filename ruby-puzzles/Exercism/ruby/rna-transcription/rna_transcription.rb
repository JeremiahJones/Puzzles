class Complement
  def self.of_dna(strand)
    strand_array = strand.split('')
    # strand.each_char do |c|
    #   case c
    #     when 'G'
    #       strand[c].gsub!(strand[c],'C')
    #     when 'C'
    #       c.gsub!(c,'G')
    #       strand[c].gsub!(strand[c],'G')
    #     when 'T'
    #       c.gsub!(c,'A')
    #       strand[c].gsub!(strand[c],'A')
    #     when 'A'
    #       c.gsub!(c,'U')
    #       strand[c].gsub!(strand[c],'U')
    #     else
    #       'Not a proper letter'
    #   end
    # end
    strand_array.each do |s|
      s.gsub!('G', 'C')
      s.gsub!('C', 'G')
      s.gsub!('T', 'A')
      s.gsub!('A', 'U')
    end
   end
end

module BookKeeping
  VERSION = 4
end
