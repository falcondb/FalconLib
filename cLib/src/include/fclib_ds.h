/*
 * fclib_ds.h
 *
 *  Created on: Nov 15, 2016
 *      Author: mayi
 */

#ifndef SRC_INCLUDE_FCLIB_DS_H_
#define SRC_INCLUDE_FCLIB_DS_H_

#include <sys/queue.h>
/* Need to check GNU C LIB in configure phase. */
#define __USE_GNU
#define _GNU_SOURCE
#include <search.h>

/**** double linked list gnu clib ****/

struct fc_list_ptr_t {
	void* val;
	LIST_ENTRY(fc_list_ptr_t) entries;
};

/* double linked list with void pointers */
typedef LIST_HEAD(fc_l_ptr_h, fc_list_ptr_t) fc_list_ptr_h;

#define FC_LIST_PTR_CRE_NODE(node, v) do {\
		node = malloc(sizeof(struct fc_list_ptr_t)); \
		if (node) node->val = v; \
		}while (0)

/* double linked list with integers */
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

/**** Hash table gnu clib ****/
/* create a named table */
#define FC_HT_CRE(ptr, els) do {	\
	ptr = calloc(1,sizeof(struct hsearch_data));	\
	hcreate_r((int) (els * 1.25), ptr);	\
} while(0)

/* add an element e to table ht, result res */
#define FC_HT_ADD(e, res, ht)	hsearch_r(e, ENTER, &res, ht)
/* find an element e in table ht, found element res */
#define FC_HT_FIND(e, res, ht)	hsearch_r(e, FIND, &res, ht)
/* free the space of table ht */
#define FC_HT_FREE(ht)		hdestroy_r(ht)
/* get hashing value from pointer */
#define HT_HT_H(ptr) (*(long*)ptr)
/* get hashed element's key from pointer */
#define HT_HT_K(ptr) ( ((ENTRY *) (ptr + sizeof(long)))->key )
/* get hashed element's value from pointer */
#define HT_HT_V(ptr) ( ((ENTRY *) (ptr + sizeof(long)))->data )
/* get next element's address from pointer */
#define HT_HT_NEXT(ptr) (ptr + sizeof(long) + sizeof(struct entry))
/* print all the elements in hash table, v_type: type of element value, fmt: formate of value */
#define FC_HT_PRINT(ht, v_type, v_fmf) do { \
		void *ptr = ht->table;	\
		unsigned int i;			\
		for (i =  0; i < ht->size; ++i){ \
			if(HT_HT_H(ptr) != 0){ 	\
			printf("Index:%u\tHash:%li\tKey: %s\t\tValue: %"#v_fmf"\n", i, HT_HT_H(ptr), HT_HT_K(ptr), *((v_type *)(HT_HT_V(ptr)))); \
			}	\
			ptr = HT_HT_NEXT(ptr);	\
		}	\
		} while (0)

#endif /* SRC_INCLUDE_FCLIB_DS_H_ */
