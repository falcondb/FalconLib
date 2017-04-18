[ $# < 1 ] && exit -1

reset
## add -pg if gprof for profiling
gcc $@ -o test -ggdb3 -I ../include/ -DDEBUG -fprofile-arcs -ftest-coverage

time ./test

gcov $@
