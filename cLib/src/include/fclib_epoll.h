/*
 * fclib_epoll.h
 *
 *  Created on: Dec 4, 2016
 *      Author: mayi
 */

#ifndef SRC_INCLUDE_FCLIB_EPOLL_H_
#define SRC_INCLUDE_FCLIB_EPOLL_H_

#include <fclib_errorno.h>
#include <fclib_ds.h>

#define EPOLLWAITMAXEVENTS 32

//#ifdef __linux__

struct fclib_epoll_meta{
	int epollfd;
	unsigned int fds;
	fc_list_ptr_h fd_list_head;

	unsigned int wait_max_event;
	int wait_time_out;

	struct epoll_event * events;
};



//#endif

#endif /* SRC_INCLUDE_FCLIB_EPOLL_H_ */
