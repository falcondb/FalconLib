/*
 * fclib_ds.h
 *
 *  Created on: Nov 15, 2016
 *      Author: mayi
 */

#ifndef SRC_INCLUDE_FCLIB_DS_H_
#define SRC_INCLUDE_FCLIB_DS_H_


#include <sys/queue.h>

struct fc_list_i_t {
	int val;
	LIST_ENTRY(fc_list_i_t) entries;
};

typedef LIST_HEAD(fc_l_i_h, fc_list_i_t) fc_list_i_h;

#define FC_LIST_I_CRE_NODE(node, v) do {\
		node = malloc(sizeof(struct fc_list_i_t)); \
		if (node) node->val = v; \
		}while (0)


#define FC_LIST_PRINT(head, type, field, fmt) do { \
	struct type * node; 	\
	LIST_FOREACH(node, head, field)  \
	fprintf(stdout, "val: %"#fmt"\t", node->val); \
	fprintf(stdout, "\n");\
	} while(0)

#define FC_LIST_FREE(head, type, field) do { \
	struct type * node;		\
	LIST_FOREACH(node, head, field) {	\
		LIST_REMOVE(node, field);	\
		free(node);	\
	}	\
	} while (0)

#endif /* SRC_INCLUDE_FCLIB_DS_H_ */
