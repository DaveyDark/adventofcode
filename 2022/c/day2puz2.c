#include <stdio.h>
#include <stdlib.h>

int main(){
    FILE *input;
    char opp,result;
    int score=0;
    input = fopen("2022/inputs/day2inp.txt","r");

    while(fscanf(input,"%c %c\n",&opp,&result) > 0){
        if(result=='Y'){
            score+=3;
            score+=opp=='A'?1:opp=='B'?2:3;
        } else if(result=='X'){
            score+=0;
            score+=opp=='A'?3:opp=='B'?1:2;
        } else {
            score+=6;
            score+=opp=='A'?2:opp=='B'?3:1;
        }
    }

    printf("Final score = %d\n",score);
    
    fclose(input);
    return 0;
}

