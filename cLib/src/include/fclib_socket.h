/*
 * fclib_socket.h
 *
 *  Created on: Nov 30, 2016
 *      Author: mayi
 */

#ifndef SRC_INCLUDE_FCLIB_SOCKET_H_
#define SRC_INCLUDE_FCLIB_SOCKET_H_


#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>

#define ERROR_CREAT_SOCKET_FAILURE 1;
#define ERROR_BIND_SOCKET_FAILURE 2;
#define ERROR_CONNECT_SOCKET_FAILURE 3;
#define ERROR_CREAT_SOCKET_PORT 101;
#define ERROR_CREAT_SOCKET_DOMAIN 102;
#define ERROR_SOCKET_TYPE 103;

#ifndef MAX_SOCKET_PENDING_QUEUE
#define MAX_SOCKET_PENDING_QUEUE (64);

#define MAX_SOCKET_PACKET_BODY (4112) // 4K
#endif

//#ifdef __linux__
struct sock_meta{
	int socketfd;
	/* this field determines the struct could be casted to sock_meta_IPV4/6 */
	int domain;
	int type;
	int protocol;
	int port;
};

struct sock_meta_IPV4{
	struct sock_meta base;
	struct sockaddr_in socket_addr;
};

struct sock_meta_IPV6{
	struct sock_meta base;
	struct sockaddr_in6 socket_addr;
};

struct client_sock_meta{
	struct sockaddr addr;
	socklen_t addrLen;
};

struct socket_package{
	uint32_t size;
	void * body;
};

struct socket_connection{
	int socketfd;
};

#define CONVERT2METAIPV4(PTR) ((struct sock_meta_IPV4 *)PTR)
#define CONVERT2METAIPV6(PTR) ((struct sock_meta_IPV6 *)PTR)

#define NEW_SOCKET_PACKAGE(PTR) do {\
		PTR = malloc(sizeof(struct socket_package)); \
		if (PTR) PTR->body = malloc(MAX_SOCKET_PACKET_BODY); \
		}while (0)

#define RESET_SOCKET_PACKAGE(PTR) do {\
		PTR->size = 0; \
		memset(PTR->body, 0, MAX_SOCKET_PACKET_BODY); \
		}while (0)

#define FREE_SOCKET_PACKAGE(PTR) do {\
		if(PTR) free(PTR->body); \
		free(PTR); \
		}while (0)

int create_socket(struct sock_meta* socket_meta);

int linsten_socket(struct sock_meta* socket_meta, int max_pending_queue);

static int create_socket_IPV4_i(struct sock_meta_IPV4 * meta);
static int create_socket_IPV6_i(struct sock_meta_IPV6 * meta);

static int connect_socket_IPV4_i(struct sock_meta_IPV4 * meta, const char * ip);
static int connect_socket_IPV6_i(struct sock_meta_IPV6 * meta, const char * ip);

#endif

//#endif /* SRC_INCLUDE_FCLIB_SOCKET_H_ */
