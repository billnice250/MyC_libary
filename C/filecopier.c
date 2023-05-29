#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <dirent.h>
#include <sys/types.h>
#include <sys/stats.h>

#define BUFFERSIZE 200


void list_dir(char *src_name,char *dest_name);
void file_copier(char * fin,char * fout);
int main(int argn, char * argv[]){
	if (argn!=3){
		fprintf(stderr,"\n Invalid number of arguments\n");
		return EXIT_FAILURE;
	}
    list_dir(argv[1],argv[2]);


	return EXIT_SUCCESS;
}

void list_dir(char *src_name,char *dest_name){
	struct dirent *src;  
	struct dirent *dest;
	struct stat statbuff_src;
	struct stat statbuff_dest;

	DIR * current_dir = opendir(src_name);
	DIR *  dest_dir = mkdir(dest_name);

	if(lstat(src_name,&statbuff_src)<0){
		fprintf(stderr,"\nError running lstats on src\n");
		exit(EXIT_FAILURE);
	}
	if(lstat(src_name,&statbuff_dest)<0){
		fprintf(stderr,"\nError running lstats on src\n");
		exit(EXIT_FAILURE);
	}
	if (current_dir==NULL){
		fprintf(stderr,"\nError Opening the dir entered\n");
		return;

     }
     if (dest_dir==NULL){
		fprintf(stderr,"\nError creating the new dir\n");
		return;
	 }

	 while(src=readdir(current_dir)!=NULL){
		if(strcmp(de->d_name,".")==0 || strcmp(de->d_name,"..")==0){
				continue;
		}
                        

	 }


return;
}
void file_copier(char * forg,char * fout){
	int fin=open(forg,O_RDONLY);
    int  fcopy=open(fout,O_WRONLY|O_SYNC|O_CREAT|O_TRUNC );
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

}