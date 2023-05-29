#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <string.h>
#include <sys/wait.h>
#include <ctype.h>
#define MAX 64
int fd[2]; //pipe variable
void signalHandler(int signal){
	if (signal==SIGUSR1){
		fprintf(stderr, "User signal caught\n");
	}
	return;
}
void reader();
void writer();

int main(int argc, char const *argv[]){
	pipe(fd);
	int status,exited_pid;


	if(signal(SIGUSR1,signalHandler)==SIG_ERR){
		fprintf(stderr, "signal handling error!!!!\n");
		exit(1);
	}


	pid_t pid=fork();
	if(pid==0){
		reader();
	}else{
		if(fork()==0){
			writer();
		}
	}
	exited_pid=wait(&status);
	if(WIFEXITED(status))
	printf("\n P(%d) exited with status %d \n",exited_pid,status);

	exited_pid=wait(&status);
	if(WIFEXITED(status))
	printf("\n P(%d) exited with status %d \n",exited_pid,status);


    printf("Program is exited successfully\n");

	return 0;
}
void reader(){
	int str_length=0;
	char str[MAX];
	int c_n;
	close(fd[1]);
	int i;
	while(1){
		read(fd[0],&str_length,sizeof(int));
		if(str_length!=0)
			c_n=read(fd[0],str,sizeof(char)*str_length);

		if(c_n!=str_length){
			fprintf(stderr, "they was a problem reading\n" );
			exit(1);
		}
		str[str_length]='\0';
		if(strcmp(str,"end")==0)
			break;
		i=0;
		while(str[i]!='\0'){
			str[i]=toupper(str[i]);
			i++;
		}

		printf("\nBy P(%d)READER: %s\n",getpid(),str);

		fflush(stdout);

	}
	printf("\nP(%d) closing.....\n",getpid());
	exit(0);
	return;
}
void writer(){
	int str_length=0;
	char str[MAX];
	int c_n;
	close(fd[0]);
	while(1){
		fprintf(stdout, "\nBy P(%d)WRITER:",getpid());
		fflush(stdin);
		fscanf(stdin,"%s",str);
		str_length=strlen(str);
		write(fd[1],&str_length,sizeof(int));
		c_n=write(fd[1],str,sizeof(char)*str_length);
		if(c_n!=str_length){
			fprintf(stderr, "they was a problem writing\n" );
			exit(1);
		}
		

		sleep(1);
		if(strcmp(str,"end")==0)
			break;


	}
	printf("\nP(%d) closing.....\n",getpid());
	exit(0);
	return;

}