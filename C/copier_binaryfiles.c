#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#define BUFFERSIZE 200
int main(int argn,char *argv[]) {
    if(argn!=3){
        fprintf(stderr,"\nInvalid number of arguments\n");
        exit(EXIT_FAILURE);
    }
    int fin=open(argv[1],O_RDONLY);
    int  fcopy=open(argv[2],O_WRONLY|O_SYNC|O_CREAT|O_TRUNC );
    //resetting buffer array;
    char buffer[BUFFERSIZE];
    strcpy(buffer,"");// unnecessary but unsures that the buffer is always empty for the new line
    int nR=0;
    int nW=0;
    fprintf(stdout,"\nCopying file (%d)%s  into (%d)%s \n",fin,argv[1],fcopy,argv[2]);

    if(fin==-1 || fcopy==-1){
        fprintf(stderr,"\nError Opening the file\n");
        exit(EXIT_FAILURE);

    }
    
    fprintf(stdout,"\n Start=> /");
  
    while ((nR=read(fin,buffer,BUFFERSIZE))>0){
        fprintf(stdout,".");
        nW=write(fcopy,buffer,nR);

        
         
	if(nR!=nW){
        fprintf(stderr,"\nError Writting in the file\n");
        exit(EXIT_FAILURE);

	}
        strcpy(buffer,"");// unnecessary but unsures that the buffer is always empty for the new line
    }
    if(nR<0)
     fprintf(stderr,"Write Error.\n"); //to check whether the last operation of read was a success or a failure.

    fprintf(stdout,"/<=finished\n");

   close(fin);
   close(fcopy);

    return EXIT_SUCCESS;
}
