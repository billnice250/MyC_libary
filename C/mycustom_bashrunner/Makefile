CC=gcc
FLAGS= -Wall
DEBUG= -g 

.PHONY: clean liteclean

target:  
	echo "Finalizing the graph app"
	$(CC) $(FLAGS) $(DEBUG)  main.c -o cmdrunner

install: target

	mkdir bin/
	mv cmdrunner ./bin
liteclean:
	rm -rf *.o
	rm -rf *.bin


clean: liteclean
	rm -R ./bin
	rm -R ./tmpDir