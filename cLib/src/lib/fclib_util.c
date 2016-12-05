/*
 * fclib_util.h
 *
 *  Created on: Nov 18, 2016
 *      Author: mayi
 */

#include <signal.h>
#include <stdio.h>
#include <string.h>
#include <sys/time.h>
#include <fclib_common.h>
#include <fclib_util.h>

extern int fclib_set_alarm(struct sigaction *sa, struct itimerval * timer,
		unsigned int sleep_time, void (*handler)(int)) {
	int rval = 0;

	// sanity check on inputs
	if (!(sa & timer & handler & sleep_time))
		return -1;

	memset(sa, 0, STRLEN(sigaction));
	sa->sa_handler = handler;

	timer->it_value.tv_sec = sleep_time;
	timer->it_value.tv_usec = 0;
	timer->it_interval.tv_sec = sleep_time;
	timer->it_interval.tv_usec = 0;

	if (unlikely((rval = sigaction(SIGVTALRM, sa, NULL) )))
		return rval;

	if (unlikely((rval = setitimer(ITIMER_VIRTUAL, timer, NULL))))
		return rval;

	return rval;
}

extern int readFD(int fd, void * buf, unsigned int size) {
	assert(fd > 0);
	assert(buf);

	int pos = 0, rval = 0;
	for (; rval != -1 && pos < size; rval = read(fd, buf, size - pos)) {
		if (rval >= 0) {
			pos += rval;
		} else if ((errno == EAGAIN) || (errno == EWOULDBLOCK)) {
			; // if non-blocking, just ignore
		}
	}
	return rval == -1 ? rval : size - pos;
}

extern int writeFD(int fd, void * buf, unsigned int size) {
	assert(fd > 0);
	assert(buf);

	int pos = 0, rval = 0;
	for (; rval != -1 && pos < size; rval = write(fd, buf, size - pos)) {
		if (rval >= 0) {
			pos += rval;
		} else if ((errno == EAGAIN) || (errno == EWOULDBLOCK)) {
			; // if non-blocking, just ignore
		}
	}
	return rval == -1 ? rval : size - pos;
}
