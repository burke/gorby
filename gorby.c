#include "ruby.h"

static int _argc;
static char **_argv;
static char *gorby = "gorby";
void boot_vm() {
  _argc = 0;
  _argv = &gorby;
  ruby_sysinit(&_argc, &_argv);
  {
    RUBY_INIT_STACK;
    ruby_init();
  }
}
