#include <stdio.h>
#include <stdlib.h>

int main(){
    FILE *input;
    char opp,resp;
    int score=0;
    input = fopen("day2inp.txt","r");

    while(fscanf(input,"%c %c\n",&opp,&resp) > 0){
        score += (resp=='X')?1:((resp=='Y')?2:3);
        if((opp+23) == resp){
            score+=3;
            continue;
        }
        if(opp=='A')score+=(resp=='Y')?6:0;
        else if(opp=='B')score+=(resp=='Z')?6:0;
        else score+=(resp=='X')?6:0; 
    }

    printf("Final score = %d\n",score);
    
    fclose(input);
    return 0;
}

