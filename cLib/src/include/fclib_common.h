/*
 * common.h
 *
 *  Created on: Nov 17, 2016
 *      Author: mayi
 */

#ifndef SRC_INCLUDE_FCLIB_COMMON_H_
#define SRC_INCLUDE_FCLIB_COMMON_H_

#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <assert.h>
#include <string.h>
#include <errno.h>
#include <sys/types.h>

#ifdef __GNUC__
#  define CC_ATT(_x) __attribute__ ((_x))
#else
#  define CC_ATT(_x)
#endif

#define likely(x)       __builtin_expect(!!(x),1)
#define unlikely(x)     __builtin_expect(!!(x),0)

#define PTRLEN (sizeof(void *))

#define STRLEN(st) (sizeof(struct st))

#endif /* SRC_INCLUDE_FCLIB_COMMON_H_ */
