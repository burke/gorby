#include "gorby.h"

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

static VALUE _gorby_eval(VALUE arg) {
  return rb_eval_string((char*)arg);
}

VALUE gorby_eval(char *str) {
  int exception = 0;
  VALUE ret = rb_protect(_gorby_eval, (VALUE)str, &exception);
  if (exception) {
    return Qfalse;
  } else {
    return ret;
  }
}
