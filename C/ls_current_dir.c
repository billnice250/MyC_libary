#include<stdio.h>
#include<dirent.h>
//
#include<string.h>
#include<unistd.h>
#include<stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>

#define MAX_PATH 256
//global variables
struct stat buff;
struct dirent *de;      //pointer for directory entry
int i=0;
void list_all(char * path_name);

int main(int argn, char * argv[]){
	if(argn!=2){
		fprintf(stderr,"\nInvalid number of arguments\n");
		exit(EXIT_FAILURE);
	}
	list_all(argv[1]);


  return EXIT_SUCCESS;
}
void list_all(char * path_name){
		DIR* drin=opendir(path_name);
			
		if(drin==NULL){
			fprintf(stderr,"\ncould not open the given directory\n");
			exit(EXIT_FAILURE);
		}  
		while((de=readdir(drin))!=NULL){
			

                                if(strcmp(de->d_name,".")==0){
				fprintf(stdout,"%s:\n",path_name);
				}
                                 if(strcmp(de->d_name,".")==0 
				|| strcmp(de->d_name,"..")==0){
				continue;
				}
                                char new_path[MAX_PATH];

                                strcpy(new_path,"");
                                strcat(new_path,path_name);
				strcat(new_path,de->d_name);
				strcat(new_path,"/");
				//fprintf(stdout,"\n %d) path = %s \n",i++,new_path);
                                fprintf(stdout," %s\n",de->d_name);
                                if(opendir(new_path)==NULL)continue;
				list_all(new_path);

		}
	
	       closedir(drin);
     return ;
}
