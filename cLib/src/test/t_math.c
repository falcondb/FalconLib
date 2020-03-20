/*
 * t_math.c
 *
 *  Created on: Jan 11, 2017
 *      Author: mayi
 */
#define __USE_GNU
#define _GNU_SOURCE
#include <sys/queue.h>
#include <fclib_ds.h>
#include <stddef.h>
#include <fclib_ds.h>
#include <assert.h>

struct interval{
	int st, end;
};

struct fc_list_inv_t {
	struct interval val;
	LIST_ENTRY(fc_list_inv_t) entries;
};

typedef LIST_HEAD(fc_l_inv_h, fc_list_inv_t) fc_list_inv_h;

static int interval_cmp(void const * v1, void const * v2){
	assert(v1);
	assert(v2);

	return ((struct interval *)v1)->st - ((struct interval *)v2)->st;
}

static mergeIntervals(struct interval * invs, const int cnt, fc_list_inv_h * rst){
	short idx;
	struct fc_list_inv_t * newNode;
	if(!invs || !rst || cnt <= 0)
		return;


	qsort(invs, cnt, sizeof(struct interval), &interval_cmp);

	struct interval cur;
	struct fc_list_inv_t * node;

	for(idx = 1, cur=invs[0]; idx < cnt; ++idx){
		if(invs[idx].st <= cur.end && cur.end < invs[idx].end) {
			cur.end = invs[idx].end;
		}else if(invs[idx].st > cur.end){
			newNode = calloc(1, sizeof(struct fc_list_inv_t));
			newNode->val = cur;
			LIST_INSERT_HEAD(rst, newNode, entries);
			cur = invs[idx];
		}
	}
	newNode = calloc(1, sizeof(struct fc_list_inv_t));
	newNode->val = cur;
	LIST_INSERT_HEAD(rst, newNode, entries);
}

static fc_list_print_free(fc_list_inv_h * head) {
	struct fc_list_inv_t * node;
	LIST_FOREACH(node, head, entries)
		printf("val: %3d,\t %3d\n", node->val.st, node->val.end);

	LIST_FOREACH(node, head, entries){
		LIST_REMOVE(node, entries);
		free(node);
	}
}



static void plusOne(const char d[], const unsigned int dsize,  char rst[]){
	if(!d)
		return;

	char carryon = 1;
	int idx;

	for(idx = 0; idx < dsize; ++idx){
		rst[idx] = (carryon + d[idx])%10;
		carryon = (carryon + d[idx])/10;
	}
	rst[idx] = carryon;
}

void main(){
	const char d[] = {9, 9, 9, 9 , 9};
	char rst[sizeof(d) / sizeof(char) + 1];
	memset(rst, 0, sizeof(d) / sizeof(char) + 1);

	plusOne(d, sizeof(d) / sizeof(char), rst);

	int i;
	for(i=0; i < 6; printf("%1d\t", rst[i++]));
}

//void main(){
//	struct interval data[] = {{1, 2}, {2, 3}, {4,8}, {5,6}, {7,9}, {11, 12}};
//	fc_list_inv_h head;
//
//	LIST_INIT(&head);
//
//	mergeIntervals(data, sizeof(data)/sizeof(struct interval), &head);
//
//	fc_list_print_free(&head);
//
//}
