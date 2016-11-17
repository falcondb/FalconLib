#include <fclib_printf.h>
#include <stdint.h>
#include <inttypes.h>

int main(){

	uint64_t x = 1;
	uint32_t y = -1;

	printf("x: %"PRId64", y: %"
			PRId32"\n", x, y);


//	FC_DEBUG("FC_DEBUG");
//	FC_INFO("info");
//	FC_ERROR("FC_ERROR: %d, %d, %d", 1, 2, 3);
//	return 0;

}
