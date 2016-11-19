[ $# != 1 ] && exit -1

reset
## add -pg if gprof for profiling
gcc $1 -o test -ggdb3 -I ../include/ -DDEBUG

./test
