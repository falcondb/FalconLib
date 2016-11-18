/*
 * fclib_set_alam.c
 *
 *  Created on: Nov 18, 2016
 *      Author: mayi
 */

#include <signal.h>
#include <stdio.h>
#include <string.h>
#include <sys/time.h>
#include <fclib_common.h>

void timer_handler(int signum) {
	static int count = 0;
	printf("timer expired %d times\n", ++count);
}

extern int fclib_set_alarm(struct sigaction *sa, struct itimerval * timer,
		unsigned int sleep_time,  void (*handler)(int) ){
	int rval = 0;

	// sanity check on inputs
	if(!(sa & timer & handler & sleep_time))
		return -1;

	memset(sa, 0, STRLEN(sigaction));
	sa->sa_handler = handler;

	timer->it_value.tv_sec = sleep_time;
	timer->it_value.tv_usec = 0;
	timer->it_interval.tv_sec = sleep_time;
	timer->it_interval.tv_usec = 0;

	if(unlikely((rval = sigaction(SIGVTALRM, sa, NULL) )))
		return rval;

	if(unlikely((rval = setitimer(ITIMER_VIRTUAL, timer, NULL))))
		return rval;

	return rval;
}

int main() {
	struct sigaction sa;
	struct itimerval timer;

	if(fclib_set_alarm(&sa, &timer, 1, &timer_handler)){
		perror("fclib_set_alarm");
		return EXIT_FAILURE;
	}

	while (1)
		;
	return EXIT_SUCCESS;
}
