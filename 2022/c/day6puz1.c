#include <stdio.h>
#include <string.h>

#define BUFF_SIZE 14

//common program for puzzle 1 and 2
//change BUFF_SIZE are required

int checkBuffer(char *buff){
    int size=strlen(buff);
    for (int i =0;i<size-1; i++) {
        for(int j =i+1;j<size;j++){
            if(buff[i] == buff[j])return 0;
        }
    }
    return 1;
}

int main(){
    char inp[5000];
    char buff[BUFF_SIZE+1];
    int len,i;
    FILE *input;
    input = fopen("2022/inputs/day6inp.txt", "r");
    fgets(inp, 5000,input);
    len = strlen(inp);
    for (i=BUFF_SIZE; i<len; i++) {
        strncpy(buff, inp+(i-BUFF_SIZE), BUFF_SIZE);
        buff[BUFF_SIZE]='\0';
        if(checkBuffer(buff))break;
    }
    if(i==len-1)printf("Not found\n");
    else printf("Found marker at %d\n",i);
}

