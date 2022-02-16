
/*Given a string containing substrings separated by a delimiter (examples: ".","..",etc) the program splints them and store them in a linked list ,dynamically allocated.*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// A linked list node
typedef struct node_s
{
	char* data;
	struct node_s* next;
}node_t;

/* Given a reference (pointer to pointer) to the head of a list and
   an int, inserts a new node on the front of the list. */
void push(node_t** head_ref, char* new_data)
{
	/* 1. allocate node */
	node_t* new_node = (node_t*)malloc(sizeof(node_t));
	/* 2. put in the data  */
	new_node->data = new_data;
	/* 3. Make next of new node as head */
	new_node->next = (*head_ref);
	/* 4. move the head to point to the new node */
	(*head_ref) = new_node;
}


node_t* splitStr(char* str, char delim[]) {
	int init_size = strlen(str);
	node_t* listHead = NULL;
	char* ptr = strtok(str, delim);
	while (ptr != NULL)
	{
		char* str = (char*)malloc(sizeof(char) * strlen(ptr) + 1);
		strcpy(str, ptr);
		push(&listHead, str);
		ptr = strtok(NULL, delim);
	}
	return listHead;
}

void main() {
	char str[] = "a.bb.ccc.dddd.cc.fff";
	char delim[] = ".";
    printf("\nOriginal string: %s",str);
	printf("\nDelimiter string: => '%s'",delim);
	node_t* listHead = splitStr(str, delim);


	printf("\nRead from list:");
	do
	{
		printf("\n\t-%s", listHead->data);
		listHead = listHead->next;
	} while (listHead != NULL);
}

/*
Output:


Original string: a.bb.ccc.dddd.cc.fff
Delimiter string: => '.'
Read from list:
	-fff
	-cc
	-dddd
	-ccc
	-bb
	-a



*/