N=16000

echo
echo =======================================
echo --- Ruby to Go localhost HTTP cost ---
echo =======================================

time c_only/main $N
time go_only/main $N

time ruby ruby/ruby_only.rb $N
time ruby ruby/ruby_to_go_adds.rb $N
time go_to_c_adds/main $N

go run go_server/main.go &
sleep 1
time go_to_go_http/main $N
curl -s localhost:8000/stop

go run go_server/main.go &
sleep 1
time ruby ruby/ruby_to_go_http.rb $N
curl -s localhost:8000/stop
