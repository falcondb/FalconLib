#include <fclib_printf.h>

int main(){

	FC_DEBUG("FC_DEBUG");
	FC_INFO("info");
	FC_ERROR("FC_ERROR: %d, %d, %d", 1, 2, 3);
	return 0;

}