#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define LINE 200+1+1
int main(int argn,char *argv[]) {
    if(argn!=3){
        fprintf(stderr,"\nInvalid number of arguments");
        exit(EXIT_FAILURE);
    }
    FILE * fin=fopen(argv[1],"r");
    FILE * fcopy=fopen(argv[2],"w");
    //resetting buffer array;
    char line[LINE];
    strcpy(line,"");


    if(fin==NULL || fcopy==NULL){
        fprintf(stderr,"\nError Opening the file");
        exit(EXIT_FAILURE);

    }
    while (fgets(line,LINE,fin)!=NULL){
        fputs(line,fcopy);
        strcpy(line,"");
    }

    fclose(fin);
    fclose(fcopy);

    return EXIT_SUCCESS;
}