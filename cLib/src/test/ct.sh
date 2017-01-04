set -x
[ $# < 1 ] && exit -1

reset
## add -pg if gprof for profiling
gcc $@ -o test -ggdb3 -I ../include/ -DDEBUG

./test
