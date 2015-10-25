N=80000000

echo
echo =======================================
echo --- Ruby to Go FFI cost ---
echo =======================================

time c_only/main $N
time go_only/main $N

time ruby ruby/ruby_only.rb $N
time ruby ruby/ruby_to_go_sum.rb $N
time ruby ruby/ruby_to_go_adds.rb $N

ruby ruby/ruby_to_go_no_op.rb
