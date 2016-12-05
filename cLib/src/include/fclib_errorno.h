/*
 * fclib_errorno.h
 *
 *  Created on: Dec 4, 2016
 *      Author: mayi
 */

#ifndef SRC_INCLUDE_FCLIB_ERRORNO_H_
#define SRC_INCLUDE_FCLIB_ERRORNO_H_


/* socket related error numbers */
#define ERROR_CREAT_SOCKET_FAILURE 1001;
#define ERROR_BIND_SOCKET_FAILURE 1002;
#define ERROR_CONNECT_SOCKET_FAILURE 1003;
#define ERROR_CREAT_SOCKET_PORT 1101;
#define ERROR_CREAT_SOCKET_DOMAIN 1102;
#define ERROR_SOCKET_TYPE 1103;

/* epoll related error number */
#ifdef __linux__
#define ERROR_CREAT_EPOLL_FAILURE 2001;
#define ERROR_FREE_EPOLL_NOT_EMPTY 2002;

#endif

#endif /* SRC_INCLUDE_FCLIB_ERRORNO_H_ */
