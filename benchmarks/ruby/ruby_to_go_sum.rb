require 'ffi'

module GoHarmonic
  extend FFI::Library
  ffi_lib "go_only/libharmonic.#{FFI::Platform::LIBSUFFIX}"
  attach_function :HarmonicSum, [:int], :double
end

n = ARGV[0].to_i
puts "\n\nRuby with 1 time Go HarmonicSum call:"
puts "Sum for n=1..#{n} of 1/n ="
print GoHarmonic.HarmonicSum(n)
