n = ARGV[0].to_i
puts "\n\nRuby only:"
puts "Sum for n=1..#{n} of 1/n ="

def harmonic(j)
  1.0 / j
end

sum = 0.0
i = 1
while i <= n do
  sum += harmonic(i)
  i += 1
end

print sum
