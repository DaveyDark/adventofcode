#include <stdio.h>

int overlaps(int s1,int s2,int e1,int e2){
    //check if range 1 ends before range 2 starts and vice versa
    if(e1<s2 || e2<s1)return 0;
    return 1;
}

int main(){
    FILE *input;
    char line[20];
    int count = 0;
    int range1s,range2s,range1e,range2e;

    input = fopen("2022/inputs/day4inp.txt", "r");

    while(fscanf(input, "%d-%d,%d-%d\n", &range1s,&range1e,&range2s,&range2e) > 0){
        if(overlaps(range1s,range2s,range1e,range2e))count++;
    }

    printf("Overlapping ranges = %d\n",count);
}

