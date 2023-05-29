#
#myls makefile
#

CC=gcc
FLAGS= -Wall
#INCLUDEDIR = -I. -I$(includedir)
#MATHLIB= -lm

.PHONY: clean distclean

target: 
	echo "Compiling my Application"
	$(CC) $(FLAGS)  main.c -o myls
install: target
	mkdir bin/
	mv myls ./bin

clean:
	rm -f *.o

distclean: clean
	rm -R ./bin