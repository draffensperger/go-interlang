require 'ffi'

module GoAdder
  extend FFI::Library
  ext = FFI::Platform::LIBSUFFIX
  ffi_lib "./go_adder/libadder.#{ext}"
  attach_function :Add, [:long, :long], :long
end

puts "Ruby says: about to call Go ..."
total = GoAdder.Add(41, 1)
puts "Ruby says: result is #{total}"
