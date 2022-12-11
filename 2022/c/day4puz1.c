#include <stdio.h>

int contains(int s1,int s2,int e1,int e2){
    //check if first range is contained
    if(s1>=s2 && e1<=e2)return 1;
    //check if second range is contained
    else if(s1<=s2 && e1>=e2)return 1;
    return 0;
}

int main(){
    FILE *input;
    char line[20];
    int count = 0;
    int range1s,range2s,range1e,range2e;

    input = fopen("2022/inputs/day4inp.txt", "r");

    while(fscanf(input, "%d-%d,%d-%d\n", &range1s,&range1e,&range2s,&range2e) > 0){
        if(contains(range1s,range2s,range1e,range2e))count++;
    }

    printf("Contained ranges = %d\n",count);
}

