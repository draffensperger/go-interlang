make clean
make
# N=1000000000
# N=40000000
# N=20000000
# N=10000000

N=800000000
SLOW=1000

# Fast
time c_only/main $N
time go_only/main $N
time go_to_c_sum/main $N
time go_concurrent/main $N
time ruby ruby/ruby_to_go_sum.rb $N

# Slow
time ruby ruby/ruby_only.rb $SLOW
time go_to_c_adds/main $SLOW
time ruby ruby/ruby_to_go_adds.rb $SLOW



# dyn to go benchmarks
# Ruby to Go ffi adds
# Rugy to Go http adds
# Ruby to Go sum

# concurrency
# Go number server with random delays
# Go concurrent number retriever
