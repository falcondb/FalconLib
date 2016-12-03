/*
 * fclib_socker.c
 *
 *  Created on: Nov 30, 2016
 *      Author: mayi
 */

#include<fclib_socket.h>
#include<fclib_common.h>

/* the code here is only for Linux OS */
//#ifdef __linux__
int create_socket(struct sock_meta* socket_meta){
		assert(socket_meta);

		int rval = 0;
		struct sockaddr_in socket_addr;

		if(unlikely(socket_meta->port <= 0)){
			rval = ERROR_CREAT_SOCKET_PORT;
			goto error;
		}

		switch(socket_meta->domain){
			case AF_INET:
				rval = create_socket_IPV4_i((struct sock_meta_IPV4 *)socket_meta);
				break;
			case AF_INET6:
				rval = create_socket_IPV6_i((struct sock_meta_IPV6 *)socket_meta);
				break;
			default:
				rval = ERROR_CREAT_SOCKET_DOMAIN;
				goto error;
				break;
		}
	error:
		return rval;
}

static int create_socket_IPV4_i(struct sock_meta_IPV4 * meta){
	int rval = 0;

	meta->base.type = meta->base.type <= 0 ? SOCK_STREAM: meta->base.type ;
	meta->base.protocol = meta->base.protocol < 0 ? 0 : meta->base.protocol;

	if( unlikely((meta->base.socketfd = socket(AF_INET, meta->base.type, meta->base.protocol)) == -1)){
		rval = ERROR_CREAT_SOCKET_FAILURE;
		goto error;
	}
	meta->socket_addr.sin_family = AF_INET;
	meta->socket_addr.sin_addr.s_addr = htonl(INADDR_ANY);
	meta->socket_addr.sin_port = htons(meta->base.port);

	if( unlikely( (bind (meta->base.socketfd, (struct sockaddr *) &(meta->socket_addr), sizeof(meta->socket_addr))) == -1)){
		rval = ERROR_BIND_SOCKET_FAILURE;
		goto error;
	}

error:
	return rval;
}

static int create_socket_IPV6_i(struct sock_meta_IPV6 * meta){
	int rval = 0;

//	meta->base.type = meta->base.type <= 0 ? SOCK_STREAM: meta->base.type ;
//	meta->base.protocol = meta->base.protocol < 0 ? 0 : meta->base.protocol;
//
//
//	if( unlikely((meta->base.socketfd = socket(AF_INET6, meta->base.type, meta->base.protocol)) == -1)){
//		rval = ERROR_CREAT_SOCKET_FAILURE;
//		goto error;
//	}
//	meta->socket_addr.sin6_family = AF_INET6;
//	meta->socket_addr.sin6_addr = in6addr_any; // htonl(INADDR_ANY);
//	meta->socket_addr.sin6_port = htons(meta->base.port);
//
//	if( unlikely( (bind (meta->base.socketfd, (struct sockaddr *) &(meta->socket_addr), sizeof(meta->socket_addr))) == -1)){
//		rval = ERROR_BIND_SOCKET_FAILURE;
//		goto error;
//	}

error:
	return rval;
}

int linsten_socket(struct sock_meta* socket_meta, int max_pending_queue){
	assert(socket_meta);

	if(socket_meta->type != SOCK_STREAM)
		return ERROR_SOCKET_TYPE;

	// need to check the socket type, if it is a stream socket
	if (max_pending_queue <= 0)
		max_pending_queue = MAX_SOCKET_PENDING_QUEUE;

	return listen(socket_meta->socketfd, max_pending_queue);
}

int accept_socket(struct sock_meta* server_meta, struct client_sock_meta* client_meta, struct socket_connection * connection_socket){
	int rval = 0;
	assert(server_meta);
	assert(connection_socket);
	assert(client_meta);

	if(server_meta->type != SOCK_STREAM)
		return ERROR_SOCKET_TYPE;

	rval = accept(server_meta->socketfd, &client_meta->addr, &client_meta->addrLen);

	if(rval >= 0)
		connection_socket->socketfd = rval;

	return rval;
}

int connect_socket(struct sock_meta* socket_meta, const char* server_ip){
	assert(socket_meta);

	int rval = 0;
	struct sockaddr_in socket_addr;

	if(unlikely(socket_meta->port <= 0)){
		rval = ERROR_CREAT_SOCKET_PORT;
		goto error;
	}

	switch(socket_meta->domain){
		case AF_INET:
			rval = connect_socket_IPV4_i((struct sock_meta_IPV4 *)socket_meta, server_ip);
			break;
		case AF_INET6:
			rval = connect_socket_IPV6_i((struct sock_meta_IPV6 *)socket_meta, server_ip);
			break;
		default:
			rval = ERROR_CREAT_SOCKET_DOMAIN;
			goto error;
			break;
	}
error:
	return rval;
}

static int connect_socket_IPV6_i(struct sock_meta_IPV6 * meta, const char * server_ip){
	assert(meta);
	assert(server_ip);
	int rval = 0;

//	meta->socket_addr.sin6_family = AF_INET6;
//	meta->socket_addr.sin6_addr.s_addr = inet_addr(server_ip);
//	meta->socket_addr.sin6_port = htons(meta->base.port);
//
//	if(connect(meta->base.socketfd, (struct sockaddr *) &meta->socket_addr,
//			sizeof(meta->socket_addr))){
//		rval = ERROR_CONNECT_SOCKET_FAILURE;
//		goto error;
//	}
error:
	return rval;
}

static int connect_socket_IPV4_i(struct sock_meta_IPV4 * meta, const char * ip_addr){
	assert(meta);
	assert(ip_addr);
	int rval = 0;

	meta->base.type = meta->base.type <= 0 ? SOCK_STREAM: meta->base.type ;
	meta->base.protocol = meta->base.protocol < 0 ? 0 : meta->base.protocol;

	if( unlikely((meta->base.socketfd = socket(AF_INET, meta->base.type, meta->base.protocol)) == -1)){
		rval = ERROR_CREAT_SOCKET_FAILURE;
		goto error;
	}

	meta->socket_addr.sin_family = AF_INET;
	meta->socket_addr.sin_addr.s_addr = inet_addr(ip_addr);
	meta->socket_addr.sin_port = htons(meta->base.port);

	if( unlikely( (connect (meta->base.socketfd, (struct sockaddr *) &(meta->socket_addr), sizeof(meta->socket_addr))) == -1)){
			rval = ERROR_CONNECT_SOCKET_FAILURE;
			goto error;
		}

error:
	return rval;
}

void* create_socket_metadata(){
	return NULL;
}

int sendPackage(struct sock_meta* target_socket_meta, struct socket_package * package, int flags){
	int rval = 0;
	uint32_t packet_size = 0;
	assert(target_socket_meta);
	assert(package);
	assert(package->size >=0 );

	packet_size = htonl(package->size);
	//TODO: check the bytes of sendto
	if(package->body){
		if(target_socket_meta->type == SOCK_STREAM){
			rval = fclib_sendto(target_socket_meta->socketfd, &packet_size, sizeof(packet_size), flags, NULL, 0);
			rval = fclib_sendto(target_socket_meta->socketfd, package->body, package->size, flags, NULL, 0);
		}else if(target_socket_meta->type == SOCK_DGRAM){
			if(target_socket_meta->domain == AF_INET){
				fclib_sendto(target_socket_meta->socketfd, &packet_size, sizeof(packet_size), flags,
						(struct sockaddr *) & CONVERT2METAIPV4(target_socket_meta)->socket_addr,
						sizeof(CONVERT2METAIPV4(target_socket_meta)->socket_addr));
				fclib_sendto(target_socket_meta->socketfd, package->body, package->size, flags,
						(struct sockaddr *) & CONVERT2METAIPV4(target_socket_meta)->socket_addr,
						sizeof(CONVERT2METAIPV4(target_socket_meta)->socket_addr));
			}else if(target_socket_meta->domain == AF_INET6){
				fclib_sendto(target_socket_meta->socketfd, &packet_size, sizeof(packet_size), flags,
						(struct sockaddr *) & CONVERT2METAIPV6(target_socket_meta)->socket_addr,
						sizeof(CONVERT2METAIPV6(target_socket_meta)->socket_addr));
				fclib_sendto(target_socket_meta->socketfd, package->body, package->size, flags,
						(struct sockaddr *) & CONVERT2METAIPV6(target_socket_meta)->socket_addr,
						sizeof(CONVERT2METAIPV6(target_socket_meta)->socket_addr));
			}
		}
	}
	return rval;
}

int recvPackage(struct socket_connection* sc, struct sock_meta* source_socket_meta, struct client_sock_meta* socket_meta,
		struct socket_package * package, int flags){
	int rval = 0;
	assert(source_socket_meta);
	assert(socket_meta);
	assert(source_socket_meta->socketfd >= 0);
	assert(package);
	assert(package->body);

	if(source_socket_meta->type == SOCK_STREAM){
		if(unlikely((rval = fclib_recvfrom(sc->socketfd, &package->size, sizeof(package->size), flags, NULL, 0)) == -1))
			goto error;
		package->size = ntohl(package->size);
		assert(package->size >= 0);
		if(unlikely((rval = fclib_recvfrom(sc->socketfd, package->body, package->size, flags, NULL, 0)) == -1))
			goto error;
	}else if(source_socket_meta->type == SOCK_DGRAM){
		if(unlikely((rval = fclib_recvfrom(sc->socketfd, &package->size, sizeof(package->size), flags,
				&socket_meta->addr, &socket_meta->addrLen)) == -1))
			goto error;
		package->size = ntohl(package->size);
		assert(package->size >= 0);
		if(unlikely((rval =fclib_recvfrom(sc->socketfd, package->body, package->size, flags,
				&socket_meta->addr, &socket_meta->addrLen)) == -1))
			goto error;
	}
error:
	return rval;
}

extern int fclib_sendto(int fd, void * buf, unsigned int size, int flags, const struct sockaddr *src_addr, socklen_t addrlen) {
	assert(fd > 0);
	assert(buf);

	int pos = 0, rval = 0;
	for (; rval != -1 && pos < size; rval = sendto(fd, buf, size - pos, flags, src_addr, addrlen)) {
		if (rval >= 0) {
			pos += rval;
		} else if ((errno == EAGAIN) || (errno == EWOULDBLOCK)) {
			; // if non-blocking, just ignore
		}
	}
	return rval == -1 ? rval : size - pos;
}

extern int fclib_recvfrom(int fd, void * buf, unsigned int size, int flags, struct sockaddr *src_addr, socklen_t addrlen) {
	assert(fd > 0);
	assert(buf);

	int pos = 0, rval = 0;
	for (; rval != -1 && pos < size; rval = recvfrom(fd, buf, size - pos, flags, src_addr, &addrlen)) {
		if (rval >= 0) {
			pos += rval;
		} else if ((errno == EAGAIN) || (errno == EWOULDBLOCK)) {
			; // if non-blocking, just ignore
		}
	}
	return rval == -1 ? rval : size - pos;
}

//#endif
