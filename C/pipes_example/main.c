#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
int fone[2]; //pipe variable 1
int ftwo[2]; //pipe variable 2
int main(int argc, char const *argv[]){
	pipe(fone);
	pipe(ftwo);
	int status,exited_pid;

    int n=10;
    char x='X';
	pid_t pid=fork();
	if(pid==0){
		//closing all pipe files not used by the child process

		close(fone[0]);
		close(ftwo[1]);
		while(n--){
			if(read(ftwo[0],&x,1)<0)fprintf(stderr, "Error\n");
			fprintf(stdout, "\n==>I'm the child (PID=%06d)\n",getpid());
			if(write(fone[1],&x,1)<0)fprintf(stderr, "Error\n");
		}
		
		exit(0);
	
	}else{
		//closing all pipe files not used by the parent process
		close(fone[1]);
		close(ftwo[0]);

		
		while(n--){
			sleep(1); // notice! that this sleep() doesnt affect the synchronization, it used to separate several sessions
			fprintf(stdout, "\n(%02d)I'm the father (PID=%06d)\n",9-n,getpid());
			if(write(ftwo[1],&x,1)<0)fprintf(stderr, "Error\n");
			if(read(fone[0],&x,1)<0)fprintf(stderr, "Error\n");
		}
	
	}
	exited_pid=wait(&status);
	if(WIFEXITED(status))
	printf("\n P(%d) exited with status %d \n",exited_pid,status);


    printf("\nProgram is exited successfully\n");

	return 0;
}
