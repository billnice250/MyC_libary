#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>
#include <string.h>
#include <ctype.h>
#define MAX_LENGTH 40+1
#define TEMP_FILE "tmp.txt"

static void signalHandler(int sig){
	return;
}
static void receiver();
static void sender(pid_t);

int main(int argc, char *argv[]){
	pid_t producer,consumer; 
	int status,cpid;
	signal(SIGUSR1,signalHandler);
	fprintf(stdout, "\nProgram start (%d)\n",getpid());
	  producer = fork();
	  if (producer== 0) {
	    receiver();

	  } else {
	    consumer= fork();
	    if (consumer== 0) {
	     sender(producer);
	 	}
	  }
	  fprintf (stdout, "Sender   PID %d\n", producer);
	  fprintf (stdout, "Receiver PID %d\n", consumer);

	  cpid = wait(&status);
	  if (WIFEXITED(status)) {
	    printf("Terminated PID %d  messages %d\n", cpid, WEXITSTATUS (status));
	  }

	  cpid = wait(&status);
	  if (WIFEXITED(status)) {
	    printf("Terminated PID %d  messages %d\n", cpid, WEXITSTATUS (status));
	  }


  
  return (0);
}
static void sender(pid_t consumer){
	char word[MAX_LENGTH];
	FILE * fout;
	int sentMsg=0;

	    while(1){
	    	// fflush(stdout);
	    	printf("\n SENDER (%d) enter the desired string : ",getpid());
	    	scanf("%s",word);
	    	fout=fopen(TEMP_FILE,"w");
				if (fout==NULL){
					fprintf(stderr, "File %s wasn't by sender opened \n",TEMP_FILE);
					exit(1);
				}
	    	fprintf(fout, "%d %s",getpid(),word);
	    	// fprintf(stdout, "\nSent to the file  %d %s\n",getpid(),word);

	    	fclose(fout);
	    	sentMsg++;

	    	kill(consumer,SIGUSR1);

	        if (strcmp(word,"end")==0){
	        	break;
	        }
	        
	        pause();
    	}
    	exit(sentMsg);
}

static void receiver(){
	char word[MAX_LENGTH];
	pid_t producer;
	int receiverMsg=0;
	FILE * fin;
	int n;
	    while(1){
	    	pause();
	    	fin=fopen(TEMP_FILE,"r");
			    if (fin==NULL){
					fprintf(stderr, "File %s wasn't by receiver opened\n",TEMP_FILE);
					exit(1);
				}
	    	fscanf(fin,"%d %s",&producer,word);
	    	// fprintf(stdout, "\nRead from file  %d %s\n",producer,word);
	    	fclose(fin);

	    	if (strcmp(word,"end")==0){
	    		break;
	    	}
            
			n=0;
	    	while( word[n]!='\0'){
	    		printf(".");
	    		word[n]=toupper(word[n]);
	    		n++;
	    	}
	    	fprintf(stdout, "\n RECEIVER(%d): %s\n",getpid(),word);
	    	// fflush(stdout);
	    	// sleep(1);
	    	receiverMsg++;  	
	    	kill(producer,SIGUSR1);  
            
	    	
    	}
    	exit(receiverMsg);

}