make clean
make
# N=1000000000
# N=40000000
# N=20000000
# N=10000000

N=800000000
SLOW=1000

# time GOMAXPROCS=1 go_concurrent/main $N 8
# time GOMAXPROCS=4 go_concurrent/main $N 4
# time GOMAXPROCS=8 go_concurrent/main $N 8
# time GOMAXPROCS=8 go_concurrent/main $N 1

# time GOMAXPROCS=8 go_to_c_concurrent/main $N 8
# time GOMAXPROCS=1 go_to_c_concurrent/main $N 8
# time GOMAXPROCS=8 go_to_c_concurrent/main $N 1

go run go_server/main.go &
sleep 0.1
time ruby ruby/ruby_to_go_http.rb 100
curl -s localhost:8000/stop

# Fast
# time c_only/main $N
# time go_only/main $N
# time go_to_c_sum/main $N
# time ruby ruby/ruby_to_go_sum.rb $N

# Slow
# time ruby ruby/ruby_only.rb $SLOW
# time go_to_c_adds/main $SLOW
# time ruby ruby/ruby_to_go_adds.rb $SLOW



# dyn to go benchmarks
# Ruby to Go ffi adds
# Rugy to Go http adds
# Ruby to Go sum

# concurrency
# Go number server with random delays
# Go concurrent number retriever
