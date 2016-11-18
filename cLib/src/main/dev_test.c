#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>
#include <assert.h>
#include <string.h>

#include <errno.h>
/* Need to check GNU C LIB in configure phase. */
#define __USE_GNU
#define _GNU_SOURCE
#include <search.h>

#define DATASIZE 32
#define MAXKEYLEN 10


int main() {

	ENTRY e, *ep = NULL;
	int i, rval = 0;
	char* arrChar, *tmpChar;
	struct hsearch_data *ht;
	int ivs[DATASIZE];
	void *ptr;

	arrChar = calloc(DATASIZE, sizeof(char) * MAXKEYLEN);
	assert(arrChar);

	for(i = 0; i < DATASIZE; ivs[i] = i, ++i);

	tmpChar = arrChar;
	for(i = 0; i < DATASIZE; sprintf(tmpChar, "%d", i++), tmpChar += sizeof(char) * MAXKEYLEN);

	ht = calloc(1,sizeof(struct hsearch_data));
	assert(ht);

	if (!(rval = hcreate_r((int) (DATASIZE * 1.25), ht))){
		perror("Error: ");
		goto cleanup;
	}

	for (i = 0; i < DATASIZE; i++) {
		e.key = arrChar + i * MAXKEYLEN;
		e.data = &ivs[i];
		if(!hsearch_r(e, ENTER, &ep, ht))
			goto cleanup;
		assert(ep);
	}

	FC_HT_PRINT(ht);


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

cleanup:
	hdestroy_r(ht);
	free(ht);
	free(arrChar);
	return rval== 0 ? EXIT_SUCCESS : EXIT_FAILURE;
}

