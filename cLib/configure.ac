AC_PREREQ([2.69])


AC_INIT([Falcon-CLib], [0.0.1], [abc@gmail.com])
AM_INIT_AUTOMAKE(Falcon-CLib, 0.0.1)
AC_CONFIG_HEADER([src/include/config.h])
AC_OUTPUT(Makefile src/main/Makefile src/lib/Makefile doc/Makefile scripts/Makefile)


Flib_dir=`pwd`
Flib_src="${Flib_dir}/src"
Flib_hdr="${Flib_src}/include"
Flib_lib="${Flib_src}/lib"
Flib_test="${Flib_src}/main"


dnl ####################################
dnl AC_CONFIG_SRCDIR([src/helloWorld.c])

AC_CONFIG_HEADERS([config.h])

dnl # Enable printf color

AC_ARG_ENABLE(falconprint,
[  --enable-falconprint      enables Falcon print features.],
[ case "$enableval" in
  no)
    FalconPrint=no
    ;;
  *)
    FalconPrint=yes
  esac ]
)

if test "x$FalconPrint" = "xyes"; then
	AC_MSG_NOTICE([enables Falcon print features])
  	Flib_cflags="$Flib_cflags -DPRINTCOLOR"
fi



debugmode=yes

AC_ARG_ENABLE(debugmode,
[  --enable-debugmode      enables debugmode.],
[ case "$enableval" in
  no)
    debugmode=no
    ;;
  *)
    debugmode=yes
  esac ]
)

if test "x$debugmode" = "xyes"; then
	AC_MSG_NOTICE([enables debugmode])
  	Flib_cflags="$Flib_cflags -g3"
else
  	Flib_cflags="$Flib_cflags -o3"
fi





# Checks for programs.
AC_PROG_CC
AC_PROG_INSTALL
AC_PROG_LIBTOOL



# Checks for libraries.

# Checks for header files.

Flib_iflags="-I$Flib_hdr"

# Checks for typedefs, structures, and compiler characteristics.

# Checks for library functions.


AC_SUBST(Flib_cflags)
AC_SUBST(Flib_iflags)
AC_SUBST(Flib_lflags)

dnl Invoke libtool
LT_INIT

AC_OUTPUT
