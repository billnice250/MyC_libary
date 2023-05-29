#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>
#include <fcntl.h>
#define MAX_VALUES 12
int swapp(int ,int, int );
void exchange_sort(int ,int ,int );
int  read_val(int file,int index);
int main(int argc, char const *argv[]){
    int i,temp=0;
    if (argc!=2){
        fprintf(stderr, "Please! Enter the name of data file\n" );
        exit(EXIT_FAILURE);
    }


    int fdata=open(argv[1],O_RDWR);
    if (fdata==-1){
        fprintf(stderr, "Couldn't create or open the file\n" );
        exit(EXIT_FAILURE);
    }
     lseek(fdata,0*sizeof(int),SEEK_SET);
     
    fprintf(stdout, "sorting %d from %d to %d\n",fdata,0,MAX_VALUES);
        for (i = 0; i <MAX_VALUES; ++i){
            int count=read(fdata,&temp,sizeof(int));
            if(count!=-1)
            printf("%d \n",temp);
        }




    exchange_sort(0,MAX_VALUES,fdata);


    lseek(fdata,0*sizeof(int),SEEK_SET);

    fprintf(stdout, "\n---------after sort------\n");
        for (i = 0; i <MAX_VALUES; ++i){
            int count=read(fdata,&temp,sizeof(int));
            if(count!=-1)
            printf("%d ",temp);
        }
    printf("\n");
    close(fdata);

	return 0;
}
int swapp(int i_x,int i_y,int file){
    //reading
    lseek(file,i_x*sizeof(int),SEEK_SET);
    int x;
    int count;
    count =read(file,&x,sizeof(int));
    lseek(file,i_y*sizeof(int),SEEK_SET);
    int y;
    count=read(file,&y,sizeof(int));

    //writting

    lseek(file,i_x*sizeof(int),SEEK_SET);
    count=write(file,&y,sizeof(int));
    lseek(file,i_y*sizeof(int),SEEK_SET);
    count=write(file,&x,sizeof(int));

    return count;
}
void exchange_sort(int from ,int to, int values){
    int i,j;


    for (i = from; i < to-1; ++i){


        for (j = from; j < to-i-1; ++j){
                  if (fork() > 0) {
                    // Father wait
                    wait ((int *) 0);
                  } else {
                    // Child compare and exchange
                    if (read_val(values,j)>read_val(values,j+1)){
                    swapp(j,j+1,values);
                    }
                    exit (0);
                  }

        }
    }




    return;
}
int  read_val(int file,int index){
    int value=0;
    lseek(file,index*sizeof(int),SEEK_SET);
    int count=read(file,&value,sizeof(int));
    if(count<0){
        fprintf(stderr, "Error reading %d value in the file \n",index);
        exit(1);
    }
    return value;

}