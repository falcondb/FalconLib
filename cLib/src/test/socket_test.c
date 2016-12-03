/*
 * socket_test.c
 *
 *  Created on: Nov 30, 2016
 *      Author: mayi
 */

#include <fclib_common.h>
#include <fclib_socket.h>
#include <fclib_printf.h>

#define PORTNUM 3333

int main(int argc, char ** argv) {
	char c;

	while ((c = getopt(argc, argv, "sc")) != -1) {
		switch (c) {
		case 's':
			server_run();
			break;
		case 'c':
			client_run();
			break;
		default:
			exit(1);
		}
	}
}

int server_run() {
	int rval;
	struct sock_meta_IPV4 meta;
	int clilen;
	struct client_sock_meta client_meta;
	struct socket_package * pkg;
	struct socket_connection sc = {.socketfd = -1};

	memset(&meta, 0, STRLEN(sock_meta_IPV4));
	meta.base.domain = AF_INET;
	meta.base.port = PORTNUM;

	memset(&client_meta, 0, sizeof(client_meta));

	rval = create_socket((struct sock_meta*) &meta);
	rval = linsten_socket((struct sock_meta*) &meta, 0);

	if (accept_socket(&meta, &client_meta, &sc) == -1) {
		rval = errno;
		goto error;
	}

	NEW_SOCKET_PACKAGE(pkg);

	if((rval = recvPackage(&sc, &meta, &client_meta, pkg, 0) == -1)){
		rval = errno;
		goto error;
	}
	((char *)pkg->body)[pkg->size] = '\0';
	printf("Body: %s\n", (char *)pkg->body);

error:
	FREE_SOCKET_PACKAGE(pkg);
	return rval;
}

int client_run() {
	int rval;
	struct sock_meta_IPV4 meta;
	char msg[] = "Client message";
	struct socket_package * pkg;

	memset(&meta, 0, STRLEN(sock_meta_IPV4));
	meta.base.domain = AF_INET;
	meta.base.port = PORTNUM;
	rval = connect_socket((struct sock_meta*) &meta, "127.0.0.1");



	NEW_SOCKET_PACKAGE(pkg);

	strncpy(pkg->body, msg, strlen(msg));
	pkg->size = strlen(msg);

	rval = sendPackage(&meta, pkg, 0);

	FREE_SOCKET_PACKAGE(pkg);
	return rval;
}
