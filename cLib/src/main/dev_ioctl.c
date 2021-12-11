#include <stdlib.h>
#include <fcntl.h>
#include <errno.h>
#include <sys/ioctl.h>
#include <unistd.h>
#include <string.h>
#include "dev_ioctl.h"

dev_ioctl_t *dev_ioctl_setup(void) {
  dev_ioctl_t *ntnx;
  ntnx = malloc(sizeof(*ntnx));
  if (unlikely(ntnx == NULL)) {
    errno = ERR_NO_MEM;
    goto err_return;
  }

  ntnx->dev_fd = open(HASH_DEVICE_PATH, O_RDWR);
  if (unlikely(ntnx->dev_fd < 0))
    goto err_free;

  //ntnx->mutex = PTHREAD_MUTEX_INITIALIZER;
  if (unlikely(pthread_mutex_init(&ntnx->mutex, NULL) < 0))
    goto err_close;


  return ntnx;

err_close:
  close(ntnx->dev_fd);
err_free:
  free(ntnx);
  // the errno is neither from open() or free()
err_return:
  return NULL;
}

char *dev_ioctl_compute(dev_ioctl_t *ctx, void *buf, size_t len) {
  unsigned int api_ver;
  int ret;
  struct dev_ioctl_compute comp;

  ret = ioctl(ctx->dev_fd, NTNX_HASH_GET_API_VERSION, &api_ver);
  if (unlikely(ret < 0)) {
    errno = ERR_DEV_GET_VER;
    return NULL;
  }

  if (unlikely(api_ver != 1)) {
    errno = ERR_UNSPORTED_VER;
    return NULL;
  }
  // this piece of memory will be returned to the caller with the digest
  comp.hash = malloc(sizeof(char) * DIGEST_BUF_SIZE);
  if (unlikely(comp.hash == NULL)) {
    errno = ERR_DEV_GET_VER;
    return NULL;
  }

  comp.buf = buf;
  comp.len = len;

  ret = pthread_mutex_lock(&ctx->mutex);
  if (unlikely(ret < 0)) {
    errno = ERR_MUTEX_ERR;
    return NULL;
  }

  ret = ioctl(ctx->dev_fd, NTNX_HASH_COMPUTE, &comp);
  pthread_mutex_unlock(&ctx->mutex);
  if (unlikely(ret == -1 ||
      strlen((char*)comp.hash) != DIGEST_BUF_SIZE - 1)) {
    errno = ERR_COMP_ERR;
    return NULL;
  }

  return (char *)comp.hash;
}

int dev_ioctl_destroy(dev_ioctl_t *ctx) {
  int ret;

  pthread_mutex_destroy(&ctx->mutex);
  ret = close(ctx->dev_fd);
  free(ctx);

  return ret;
}
