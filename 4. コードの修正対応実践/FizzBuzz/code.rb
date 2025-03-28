# RubyのFizzBuzz

def fizzbuzz
  (1..100).each do |i|
    if i % 15 == 0
      puts "FizzBuzz"
    elif i % 3 == 0
      puts "Fizz"
    elif i % 5 == 0
      puts "Buzz"
    else
      puts i
    end
  end
end

fizzbuzz