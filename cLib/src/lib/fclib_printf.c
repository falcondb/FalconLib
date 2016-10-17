
#include <stdio.h>
#include <unistd.h>

#include <fclib_printf.h>

#define NAMELEN 128

int test()
{
	char usrname[NAMELEN] = "Unknown";
#ifdef HAVE_GETLOGIN_R
	if(getlogin_r(usrname, NAMELEN))
		return -1;
#endif

#ifdef FALCONBLINK
	printf("%s", BLINK);
#endif

#ifdef FALCONBOLD
	printf("%s-----\n", PBOLD);
#endif
#if defined (PRINTCOLOR) && defined (PCRED)
	printf("%sYou are:%s\n", PCRED, usrname);
#else
	printf("You are:%s\n", usrname);
#endif /* end of PRINTCOLOR */

#ifdef FALCONPRINT
	printf("%sBye Bye\n",PNRM);
#endif
    return 0;
}
