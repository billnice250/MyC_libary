/*Given a string  the program the longest chain of a repeated character and returns it's start point and length.*/

#include <stdio.h>
#include <string.h>
void searchStr(char  *str,int* start,int*length ) {
    int amt = 0;
    char current;
    int  max = 0;
    int tmpStart=0;

    for (int i = 0; i < strlen(str); i++) {
        if (current == str[i]){
            amt++;
        }else {
            amt = 1;
            current = str[i];
        }
        if (max < amt){
            max = amt;
            tmpStart=i-max+1;
        }
    }
    *length = max;
    *start = tmpStart;
    return ;
}

void main() {
    char test[] = "abbbbbcewwewqwwwwwwwwwqbbcccccdddeeeee";
    int  start=0;
    int length=0;
    searchStr(test,&start,&length);
    printf("\nInput: '%s'",test);
    printf("\nOutput: the longest chain is \"%c\"\n\t-length=%d\n\t-start=%d",test[start],length,start);

}
/*
Output
Input: 'abbbbbcewwewqwwwwwwwwwqbbcccccdddeeeee'
Output: the longest chain is "w"
	-length=9
	-start=13
*/