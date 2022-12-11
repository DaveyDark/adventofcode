#include <stdio.h>
#include <string.h>

void search(char sacks[3][50],int *sum){
    for(int i=0; i<strlen(sacks[0]); i++){
        for(int j=0; j<strlen(sacks[1]);j++){
            if(sacks[0][i] == sacks[1][j]){
                for(int k=0; k<strlen(sacks[2]);k++){
                    if(sacks[0][i] == sacks[2][k]){
                        char common=sacks[2][k];
                        if(common>=65 && common <= 90)*sum+=common-38;
                        else *sum += common-96;
                        return;
                    }
                }
            }
        }
    }
}

int main(){
    FILE *input;
    char line[50];
    char sacks[3][50];
    int sum=0,ctr=0;;

    input = fopen("day3inp.txt", "r");

    while(fgets(line,50,input) != NULL){
        strcpy(sacks[ctr],line);
        if(ctr==2){
            ctr=0;
            search(sacks,&sum);
        } else ctr++;
    }

    printf("Sum of badges : %d\n",sum);

    return 0;
}

