/*
 * fclab_printf.h
 *
 *  Created on: Oct 14, 2016
 *      Author: mayi
 */

#ifndef FCLIB_PRINTF_H_
#define FCLIB_PRINTF_H_

#include <stdio.h>

#ifdef FALCONPRINT
	#define PRINTCOLOR
	#define FALCONBOLD
	#define FALCONBLINK
#endif

#ifdef PRINTCOLOR
	#define PCRED  "\x1B[31m"
	#define PCGRN  "\x1B[32m"
	#define PCYEL  "\x1B[33m"
	#define BLINK "\x1B[5m"
#endif

#ifdef FALCONBOLD
	#define PBOLD "\x1B[51m"
#endif


#if defined (PRINTCOLOR) || defined (FALCONBOLD)
	#define PNRM "\x1B[0m"
#endif

#ifdef PRINTCOLOR
	#define FC_DEBUG(fmt, ...) fprintf(stdout, "%s"#fmt"%s\n", PCYEL, ## __VA_ARGS__ , PNRM)
	#define FC_INFO(fmt, ...)  fprintf(stdout, "%s"#fmt"%s\n", PCGRN, ## __VA_ARGS__ , PNRM)
	#define FC_ERROR(fmt, ...) fprintf(stderr, "%s%s"#fmt"%s\n", BLINK, PCRED, ## __VA_ARGS__ , PNRM)
#else
	#define FC_DEBUG(fmt, ...) fprintf(stdout, fmt, ## __VA_ARGS__)
	#define FC_INFO(fmt, ...)  fprintf(stdout, fmt, ## __VA_ARGS__)
	#define FC_ERROR(fmt, ...) fprintf(stderr, fmt ## __VA_ARGS__)
#endif

#endif /* FCLIB_PRINTF_H_ */
