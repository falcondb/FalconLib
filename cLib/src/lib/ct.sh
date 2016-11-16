[ $# != 1 ] && exit -1

gcc $1 -o test -ggdb3 -I ../include/
./test
