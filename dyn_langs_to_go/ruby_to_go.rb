require 'ffi'

module GoAdder
  extend FFI::Library
  ffi_lib "./go_adder/libadder.#{FFI::Platform::LIBSUFFIX}"
  attach_function :Add, [:int, :int], :int
end

puts "Ruby says: about to call Go ..."
total = GoAdder.Add(41, 1)
puts "Ruby says: got the result as: #{total}"
