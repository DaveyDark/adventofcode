#include <stdio.h>
#include <string.h>

//NOTE: Solution not working; puzzle unsolved

void stackMove(char *str1,char *str2,int cnt){
    if(cnt>=strlen(str1)){
        int len = strlen(str2);
        strcat(str2,str1);
        str2[len+strlen(str1)]='\0';
        str1[0]='\0';
        return;
    }
    strncat(str2,str1+strlen(str1)-cnt,cnt);
    str2[strlen(str2)+cnt] = '\0';
    str1[strlen(str1)-cnt] = '\0';
}

int main(){
    FILE *input;
    int mode=0,i=0;
    int count,from,to;
    char line[40],stacks[9][50];

    input = fopen("2022/inputs/day5inp.txt","r");

    while(fgets(line,40,input)[0] != '\n'){
        //input to strings
        //TODO
        strcpy(stacks[i], line);
        stacks[i][strlen(stacks[i])-1] = '\0';
        i++;
    }
    for (int k = 0;k<i;k++){
        printf(" [%c] ",stacks[k][strlen(stacks[k])-1]);
    }
    printf ("\n");
    while(fscanf(input, "move %d from %d to %d\n",&count,&from,&to) > 0){
        //process moving
        from--;
        to--;
        // printf("%s\t%s\t%d\n",stacks[from],stacks[to],count);
        if(stacks[from][0] == '\0'){
            continue;
        }
        stackMove(stacks[from],stacks[to], count);
        // printf("%s\t%s\n",stacks[from],stacks[to]);
    }

    printf("\n");
    //print top crates
    printf("\n");
    for (int j=0; j<i ; j++) {
        printf("%c",stacks[j][strlen(stacks[j])-1]);
    }
    printf("\n");

    return 0;
}
