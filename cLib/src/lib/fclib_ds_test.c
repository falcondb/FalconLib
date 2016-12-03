/*
 * fclib_ds.c
 *
 *  Created on: Nov 15, 2016
 *      Author: mayi
 *      Test only
 */
#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>
#include <fclib_ds.h>

#include "../include/fclib_common.h"

#define DATASIZE 32
#define MAXKEYLEN 10


static int CC_ATT(unused) test_list() {
	struct fc_list_i_t * node, *new_node;
	int c = 0;
	fc_list_i_h head;

	LIST_INIT(&head);

	for (c = 0; c < 5; ++c) {
		FC_LIST_I_CRE_NODE(node, c);
		if (!node)
			goto err;
		LIST_INSERT_HEAD(&head, node, entries);
	}

	// get the second node, assume the second exists, code for fan only
	node = LIST_NEXT(LIST_NEXT(LIST_FIRST(&head), entries), entries);

	// add 100 as the third
	FC_LIST_I_CRE_NODE(new_node, 100);
	LIST_INSERT_BEFORE(node, new_node, entries);
	FC_LIST_PRINT(&head, fc_list_i_t, entries, d);

	FC_LIST_I_CRE_NODE(new_node, 101);
	LIST_INSERT_HEAD(&head, new_node, entries);
	FC_LIST_PRINT(&head, fc_list_i_t, entries, d);

	FC_LIST_FREE(&head, fc_list_i_t, entries);

	fprintf(stdout, "Byte!\n");

	return EXIT_SUCCESS;
	err: return EXIT_FAILURE;
}

static int CC_ATT(unused) test_hashtable() {
	ENTRY e, *ep = NULL;
	int i, rval = 0;
	char* arrChar, *tmpChar;
	struct hsearch_data *ht;
	int ivs[DATASIZE];
	void *ptr;

	arrChar = calloc(DATASIZE, sizeof(char) * MAXKEYLEN);
	assert(arrChar);

	for (i = 0; i < DATASIZE; ivs[i] = i, ++i)
		;

	tmpChar = arrChar;
	for (i = 0; i < DATASIZE;
			sprintf(tmpChar, "%d", i++), tmpChar += sizeof(char) * MAXKEYLEN)
		;

	ht = calloc(1, sizeof(struct hsearch_data));
	assert(ht);

	if (!(rval = hcreate_r((int) (DATASIZE * 1.25), ht))) {
		perror("Error: ");
		goto cleanup;
	}

	for (i = 0; i < DATASIZE; i++) {
		e.key = arrChar + i * MAXKEYLEN;
		e.data = &ivs[i];
		if (!hsearch_r(e, ENTER, &ep, ht))
			goto cleanup;
		assert(ep);
	}

	FC_HT_PRINT(ht, int, d);

	/* search for random int in the hashtable */
	for (i = 0; i < 10; i++) {
		e.key = arrChar + (random() % DATASIZE) * MAXKEYLEN;
		e.data = NULL;
		hsearch_r(e, FIND, &ep, ht);
#ifdef DEBUG
		if(ep)
		printf("%s->%d\n", ep->key, *((int*)ep->data));
		else
		fprintf(stderr, "%s is not found\n", e.key);
#endif
	}

	cleanup: hdestroy_r(ht);
	free(ht);
	free(arrChar);
	return rval == 0 ? EXIT_SUCCESS : EXIT_FAILURE;
}

int main(int argc, char ** argv) {

	set_debug_traps();

	test_hashtable();
	//test_list();

	return EXIT_SUCCESS;
	err: return EXIT_FAILURE;
}
