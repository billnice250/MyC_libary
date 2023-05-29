CC=gcc
FLAGS= -Wall

.PHONY: clean distclean

target:  
	echo "Finalizing the graph app"
	$(CC) $(FLAGS)  main.c -o q2

install: target

	mkdir bin/
	mv q2 ./bin
liteclean:
	rm -rf *.o

clean: liteclean
	rm -R ./bin