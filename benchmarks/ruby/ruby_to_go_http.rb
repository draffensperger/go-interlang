require 'net/http'

n = ARGV[0].to_i
puts "\n\nRuby with n times HTTP calls:"
puts "Sum for n=1..#{n} of 1/n ="

sum = 0.0
i = 1
http = Net::HTTP.new('localhost', 8000)

while i <= n do
  harmonic = http.request(Net::HTTP::Get.new('/' + i.to_s)).body.to_f
  sum += harmonic
  i += 1
end

print sum
