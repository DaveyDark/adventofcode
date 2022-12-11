#include <stdio.h>
#include <stdlib.h>

int main(){
    FILE *input;
    char line[20];
    int cal,sum = 0,greatest = 0;
    input = fopen("2022/inputs/prog01.1.in","r");

    while(fgets(line,20,input) != NULL ){
        if(line[0]=='\n'){
            if(sum>greatest)greatest = sum;
            sum=0;
            continue;
        }
        cal = atoi(line);
        sum+=cal;
        /* printf( "%s",line ); */
    }

    printf("%d is greatest\n",greatest);

    fclose(input);
}

