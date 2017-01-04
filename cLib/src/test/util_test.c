#include <fclib_util.h>
#include  <signal.h>

int main(int argc, char ** argv) {

	int * ip = 0;
	set_sigal_print_backtrace(SIGSEGV, -1, -1);

	*ip = 1;

	return 0;
}
