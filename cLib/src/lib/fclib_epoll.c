/*
 * fclib_epoll.c
 *
 *  Created on: Dec 4, 2016
 *      Author: mayi
 */

//#ifdef __linux__
#include <sys/epoll.h>
#include <fclib_epoll.h>
#include <fclib_common.h>

struct fclib_epoll_meta * create_epoll( int max_event, int time_out ){
	int rval = 0;
	struct fclib_epoll_meta * meta;
	if(unlikely(meta = malloc(sizeof(struct fclib_epoll_meta)) == NULL))
		goto error;

	if(unlikely(rval = epoll_create(0)) == -1)
		goto error;

	meta->epollfd = rval;
	LIST_INIT(&meta->fd_list_head);
	meta->fds = 0;
	meta->wait_max_event = max_event > 0? max_event : EPOLLWAITMAXEVENTS;
	meta->wait_time_out = time_out > 0? time_out : -1;

	if(unlikely((meta->events = calloc(meta->wait_max_event, sizeof(struct epoll_event))) == NULL)){
		rval = free_epoll(meta);
		meta = NULL;
		goto error;
	}
error:
	return meta;
}

int free_epoll(struct fclib_epoll_meta * meta ){
	int rval = 0;
	if(!meta)
		return 0;

	if(unlikely(meta->fds > 0 || LIST_FIRST(&meta->fd_list_head))){
		rval = ERROR_FREE_EPOLL_NOT_EMPTY;
		goto error;
	}

	if(unlikely((rval = close(meta->epollfd) == -1)))
			goto error;

	free(meta->events);
	free(meta);

error:
	return rval;
}
int add_epoll_event(struct fclib_epoll_meta * epoll, int listen_fd, int flags){
	int rval = 0;
	struct epoll_event * ev;
	assert(epoll);
	assert(listen_fd >= 0);
	struct fc_list_ptr_t * list_node;

	/* remember to reclaim the space for the epoll_event */
	ev = malloc(sizeof(struct epoll_event));
	assert(ev);

	ev->events = flags < 0 ? EPOLLIN | EPOLLOUT | EPOLLPRI | EPOLLERR | EPOLLHUP :
			flags == 0 ? EPOLLIN | EPOLLOUT : flags;
	ev->data.fd = listen_fd;

	if(unlikely(rval = epoll_ctl(epoll->epollfd, EPOLL_CTL_ADD, listen_fd, ev)) == -1 )
		goto error;

	FC_LIST_PTR_CRE_NODE(list_node, (void *)ev);
	assert(list_node);

	LIST_INSERT_HEAD(&epoll->fd_list_head, list_node, entries);
	epoll->fds ++;

error:
	return rval;
}

int remove_epoll_event(struct fclib_epoll_meta * epoll, int listen_fd){
	int rval = 0;
	struct epoll_event * ev = NULL;
	struct fc_list_ptr_t * node = NULL;
	assert(epoll);
	assert(listen_fd >= 0);

	if(unlikely(rval = epoll_ctl(epoll->epollfd, EPOLL_CTL_DEL, listen_fd, ev)) == -1 )
		goto error;

	LIST_FOREACH(node, &epoll->fd_list_head, entries)
	{
		if(((struct epoll_event *)node->val)->data.fd == listen_fd){
			free((struct epoll_event *)(node->val));
			LIST_REMOVE(node, entries);
			epoll->fds --;
			break;
		}
	}
error:
	return rval;
}

int wait_epoll(struct fclib_epoll_meta * epoll){
	int rval = 0;

	if(unlikely((rval = epoll_wait(epoll->epollfd, epoll->events ,epoll->wait_max_event, epoll->wait_time_out)) == -1))
		goto error;
error:
	return rval;
}
epoll_wait(int epfd, struct epoll_event *events,
                      int maxevents, int timeout);
//#endif
