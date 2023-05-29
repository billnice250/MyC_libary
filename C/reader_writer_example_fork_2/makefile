CC=gcc
FLAGS= -Wall -lm

.PHONY: clean distclean

target:  
	echo "Finalizing the graph app"
	$(CC) $(FLAGS)  main.c -o run

install: target

	mkdir bin/
	mv run ./bin
liteclean:
	rm -rf *.o

clean: liteclean
	rm -R ./bin
