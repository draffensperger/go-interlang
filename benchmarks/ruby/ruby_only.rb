n = ARGV[0].to_i
puts "\n\nRuby only:"
puts "Sum for n=1..#{n} of 1/n ="

sum = 0.0
i = 1
while i <= n do
  sum += 1.0 / i
  i += 1
end

print sum
