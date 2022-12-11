#include <stdio.h>
#include <stdlib.h>

const int TOP_COUNT=5;

void shift(int arr[TOP_COUNT],int ind,int val){
    int ctr = TOP_COUNT-1;
    while(ctr>=ind){
        if(ctr == ind)arr[ctr] = val;
        else arr[ctr] = arr[ctr-1];
        ctr--;
    }
}

int main(){
    FILE *input;
    char line[20];
    int cal,sum = 0,top[TOP_COUNT];
    input = fopen("2022/inputs/prog01.1.in","r");

    while(fgets(line,20,input) != NULL ){
        if(line[0]=='\n'){
            for (int i=0; i<TOP_COUNT; i++) {
                if(sum>top[i]){
                    shift(top, i, sum);
                    sum=0;
                    break;
                }
            }
            sum=0;
            continue;
        }
        cal = atoi(line);
        sum+=cal;
        /* printf( "%s",line ); */
    }

    for(int i=0;i<TOP_COUNT;i++){
        printf("%d is #%d\n",top[i],i+1);
    }

    fclose(input);
}

