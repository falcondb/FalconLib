/*
 * epoll_test.c
 *
 *  Created on: Dec 4, 2016
 *      Author: mayi
 */

#include <fclib_epoll.h>

int main(int argc, char ** argv) {

	int rval = 0;
	int fd = 10;
	struct fclib_epoll_meta epoll;

	memset(&epoll, 0, sizeof(struct fclib_epoll_meta));

	rval = create_epoll(epoll);

	rval = add_epoll_event(&epoll, fd, 0);

	printf("Return: %d\n", rval);
	return rval;
}
