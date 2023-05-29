#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#define MAX_LINE 1064
#define MAX_WORD 64
void parse(char * line,char **);
void  execute_f(char **);

int main(int argc, char *argv[]){
	if(argc!=2){
		fprintf(stderr, "\nNo command file was passed to program, try again!\n" );
	}
    FILE * fin=fopen(argv[1],"r");
    if(fin==NULL){
        fprintf(stderr,"\nError Opening the file\n");
        exit(EXIT_FAILURE);

    }
    char line[MAX_LINE];
    char samplecmd[MAX_WORD];
    strcpy(line,""); //cleaning the buffer 
    strcpy(samplecmd,""); //cleaning the sample 
    char *command[MAX_WORD]; //holding one command designed for  execvp

    int i; //indices

    while(fgets(line,MAX_LINE,fin)!=NULL){
       	for (i = 0; i < strlen(line)-2; i++){
    		if((line[i]=='e')
    			&&(line[i+1]=='n')
    			&&(line[i+2]=='d')){
    			if(i>0){
    				if (line[i-1]==' ')line[i-1]='\0';
    			}else
    		    	line[i]='\0';
    		}
    	}
    	if(strlen(line)==0)continue;
    	printf("\n Using system: %s\n",line);
    	system(line);
        printf("\n\t-----||-----\n");
        sleep(3);
        printf("Using execlp: %s\n",line);
        parse(line,command);
		execute_f(command);
       	strcpy(line,""); //cleaning the buffer 
    	 

    }


	return EXIT_SUCCESS;
}

void  parse(char *line, char **argv)
{
     while (*line != '\0') {       /* if not the end of line ....... */ 
          while (*line == ' ' || *line == '\t' || *line == '\n')
               *line++ = '\0';     /* replace white spaces with 0    */
          *argv++ = line;          /* save the argument position     */
          while (*line != '\0' && *line != ' ' && 
                 *line != '\t' && *line != '\n') 
               line++;             /* skip the argument until ...    */
     }
     *argv = '\0';                 /* mark the end of argument list  */
}
void  execute_f(char **argv){
     pid_t  pid;
     int    status;

     if ((pid = fork()) < 0) {     /* fork a child process           */
          printf("*** ERROR: forking child process failed\n");
          exit(1);
     }
     else if (pid == 0) {          /* for the child process:         */
          if (execvp(*argv, argv) < 0) {     /* execute the command  */
               printf("*** ERROR: exec failed\n");
               exit(1);
          }
     }
     else {                                  /* for the parent:      */
          while (wait(&status) != pid)       /* wait for completion  */
               ;
     }
}