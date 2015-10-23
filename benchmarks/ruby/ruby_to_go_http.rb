require 'net/http'

n = ARGV[0].to_i
puts "\n\nRuby with n times Go http calls:"
puts "Sum for n=1..#{n} of 1/n ="

sum = 0.0
i = 1
while i <= n do
  harmonic = Net::HTTP.get('localhost', "/#{i}", 8000).to_f
  sum += harmonic
  i += 1
end

print sum
