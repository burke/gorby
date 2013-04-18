
libgorby.dylib: gorby.o
	gcc -L/Users/burke/src/g/ruby -Wl,-u,_objc_msgSend -lruby-static -lpthread -ldl -lobjc -dynamiclib -Wl,-headerpad_max_install_names,-undefined,dynamic_lookup,-compatibility_version,1.0,-current_version,1.0 -o libgorby.dylib gorby.o

gorby.o: gorby.c
	gcc -DHAVE_STRUCT_TIMESPEC -I/Users/burke/src/g/ruby/include -o gorby.o -c gorby.c

gorby: gorby.o
	gcc -L/Users/burke/src/g/ruby -Wl,-u,_objc_msgSend -lruby-static -lpthread -ldl -lobjc -o gorby gorby.o
