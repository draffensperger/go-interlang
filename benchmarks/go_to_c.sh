N=800000000

echo
echo =======================================
echo --- Go to C interlang call cost ---
echo =======================================

time c_only/main $N
time go_only/main $N
time go_to_c_sum/main $N
time go_to_c_adds/main $N
