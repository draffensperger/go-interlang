make clean
make

# Add a blank line between make and run
echo

# Give the search path for the library
# LD_LIBRARY_PATH is for Linux, DYLD_LIBRARY_PATH for Mac
LD_LIBRARY_PATH=./go_adder DYLD_LIBRARY_PATH=./go_adder ./main
