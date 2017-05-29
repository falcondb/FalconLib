[ $# != 1 ] && exit -1

reset
gcc $1 -o test -ggdb3 -I ../include/ -DDEBUG
./test
