/*
 * fclib_util.h
 *
 *  Created on: Oct 16, 2016
 *      Author: mayi
 */

#ifndef SRC_INCLUDE_FCLIB_UTIL_H_
#define SRC_INCLUDE_FCLIB_UTIL_H_

#include <sys/time.h>
#include <signal.h>

#define MIN(m,n) ((m) < (n) ? (m) : (n))
#define MAX(m,n) ((m) > (n) ? (m) : (n))

int readFD(int fd, void * buf, unsigned int size);
int writeFD(int fd, void * buf, unsigned int size);


#define FCLIBDEBUGPAUSEPERIOD 30

#define FCLIBENABLEDEBUGGER if (getenv("FCLIBDEBUGGER")) sleep (FCLIBDEBUGPAUSEPERIOD);

int fclib_set_alarm(struct sigaction *sa, struct itimerval * timer,
		unsigned int sleep_time,  void (*handler)(int) );

int set_sigal_print_backtrace(int signum, int fd, int depth);
void (*print_std_trace) (int signum);

#endif /* SRC_INCLUDE_FCLIB_UTIL_H_ */
