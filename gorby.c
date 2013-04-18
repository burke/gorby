#undef RUBY_EXPORT
#include "ruby.h"
#ifdef HAVE_LOCALE_H
#include <locale.h>
#endif

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

int main(int argc, char **argv)
{
  boot_vm();
  rb_eval_string("puts 'qqqqqwoooooo'");
  //rb_funcall(rb_cObject, rb_intern("puts"), 1, rb_str_new2("qqwoooo"));
  return 0;
}
