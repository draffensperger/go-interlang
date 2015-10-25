require 'ffi'
require 'benchmark'

module GoNoOp
  extend FFI::Library
  ffi_lib "go_to_c_no_op/libgo_no_op.#{FFI::Platform::LIBSUFFIX}"
  attach_function :GoOnlyNoOp, [], :void
end

n = 50000000
puts "\n\nBenchmarking #{n} Ruby to Go no-op FFI calls"
Benchmark.bm do |b|
  b.report { n.times { GoNoOp.GoOnlyNoOp } }
end

def no_op
end

puts "\n\nBenchmarking #{n} Ruby to only no-op function calls"
Benchmark.bm do |b|
  b.report { n.times { no_op } }
end
