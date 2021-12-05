#include <stdio.h>
#include <pthread.h>

#define NTNX_HASH_GET_API_VERSION   0
#define NTNX_HASH_COMPUTE           1

#define HASH_DEVICE_PATH "/dev/ntnx_hash"
#define DIGEST_BUF_SIZE             33

#define ERR_NO_MEM 1

#define ERR_DEV_GET_VER 1001
#define ERR_UNSPORTED_VER 1002
#define ERR_COMP_ERR 1003

#define ERR_MUTEX_ERR 1004

typedef struct ntnx_hash {
  // fd of the handler of the char device
  int dev_fd;
  pthread_mutex_t mutex;
} ntnx_hash_t;

struct ntnx_hash_compute {
  void *buf; // pointer to the area for hashing
  size_t len; // length of area for checksumming
  void *hash; // pointer to the area for the computed hash
};

#define unlikely(x)     __builtin_expect(!!(x), 0)

ntnx_hash_t *ntnx_hash_setup(void);
char *ntnx_hash_compute(ntnx_hash_t *ctx, void *buf, size_t len);
int ntnx_hash_destroy(ntnx_hash_t *ctx);
