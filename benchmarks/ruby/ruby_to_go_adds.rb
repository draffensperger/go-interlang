require 'ffi'

module GoHarmonic
  extend FFI::Library
  ffi_lib "go_only/libharmonic.#{FFI::Platform::LIBSUFFIX}"
  attach_function :AddHarmonic, [:double, :int], :double
end

n = ARGV[0].to_i
puts "\n\nRuby with n times Go AddHarmonic calls:"
puts "Sum for n=1..#{n} of 1/n ="

sum = 0.0
i = 1
while i <= n do
  sum = GoHarmonic.AddHarmonic(sum, i)
  i += 1
end

print sum
