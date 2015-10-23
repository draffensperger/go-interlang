puts "Ruby only: n(n+1)/2 the long way.."

n = ARGV[0].to_i || 100

total = 0
i = 1
while i <= n do
  total += i
  i += 1
end

puts "Sum of 1 .. #{n} = #{total}"
