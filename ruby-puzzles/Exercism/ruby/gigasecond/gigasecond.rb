require 'date'
# calculates the moment when someone has lived for one gigasecond
class Gigasecond
  def self.from(birthday)
    gigasecond = (10**9)
    end_date = birthday + gigasecond
    end_date
  end
end

module BookKeeping
  VERSION = 5
end
