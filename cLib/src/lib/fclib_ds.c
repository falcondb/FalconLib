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

int main(int argc, char ** argv) {
	struct fc_list_i_t * node, * new_node;
	int c = 0;
	fc_list_i_h head;

	LIST_INIT(&head);

	for (c = 0; c < 5; ++c) {
		FC_LIST_I_CRE_NODE(node, c);
		if (!node) goto err;
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
err:
	return EXIT_FAILURE;
}
