N=8000000000

echo
echo =======================================
echo --- Go to C and concurrency ---
echo =======================================

time go_only/main $N
time c_only/main $N

time GOMAXPROCS=1 go_concurrent/main $N 1
time GOMAXPROCS=1 go_to_c_concurrent/main $N 1

time GOMAXPROCS=1 go_concurrent/main $N 8
time GOMAXPROCS=1 go_to_c_concurrent/main $N 8

time GOMAXPROCS=8 go_concurrent/main $N 8
time GOMAXPROCS=8 go_to_c_concurrent/main $N 8

time GOMAXPROCS=4 go_concurrent/main $N 4
time GOMAXPROCS=4 go_to_c_concurrent/main $N 4

time GOMAXPROCS=8 go_concurrent/main $N 1
time GOMAXPROCS=8 go_to_c_concurrent/main $N 1
