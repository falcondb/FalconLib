/*
 * fclib_util.h
 *
 *  Created on: Oct 16, 2016
 *      Author: mayi
 */

#ifndef SRC_INCLUDE_FCLIB_UTIL_H_
#define SRC_INCLUDE_FCLIB_UTIL_H_

#define FCLIBDEBUGPAUSEPERIOD 30

#define likely(x)       __builtin_expect(!!(x),1)
#define unlikely(x)     __builtin_expect(!!(x),0)

#define FCLIBENABLEDEBUGGER if (getenv("FCLIBDEBUGGER")) sleep (FCLIBDEBUGPAUSEPERIOD);

#endif /* SRC_INCLUDE_FCLIB_UTIL_H_ */
