
gorby:
	gcc -DHAVE_STRUCT_TIMESPEC -I/opt/boxen/rbenv/versions/1.9.3-p385-perf/include/ruby-1.9.1 -L'/opt/boxen/rbenv/versions/1.9.3-p385-perf/lib' -Wl,-u,_objc_msgSend -lruby-static -lpthread -ldl -lobjc -o gorby gorby.c
