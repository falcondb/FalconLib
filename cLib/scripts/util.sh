#!/bin/bash

function isInArray () {
	target=$1
	array=( "${@:2}" )
	for item in "${array[@]}"; do
		[ "$target" = "$item" ] && return 1
	done
	return 0
}

# test=("all" "install" "start" "stop")
# isInArray "stop" "${test[@]}"
# echo $?

isInArray "all" $@
ALL_FLAG=$? 

isInArray "install" $@
INSTALL_FLAG=$?

isInArray "start" $@
START_FLAG=$?

isInArray "stop" $@
STOP_FLAG=$?


