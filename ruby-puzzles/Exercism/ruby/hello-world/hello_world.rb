# Say hello to the world before you begin on this adventure
class HelloWorld
  def self.hello(name = (no_arguments_passed = true, ''))
    if no_arguments_passed
      'Hello, World!'
    else
      "Hello, #{name}!"
    end
  end
end
