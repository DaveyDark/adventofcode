#include <stdio.h>
#include <string.h>

int main(){
    FILE *input;
    char line[50],common;
    int len1,len2,sum=0;

    input = fopen("2022/inputs/day3inp.txt", "r");

    while(fgets(line,50,input) != NULL){
        len1 = strlen(line)/2;
        len2 = strlen(line)-len1;
        char container1[len1+1],container2[len2+1];
        strncpy(container1, line,len1);
        strncpy(container2, line+len1,len2);
        container1[len1] = '\0';
        container2[len2] = '\0';
        int tmp = sum;
        
        for (int i = 0; i < strlen(container1); i++) {
            for (int j = 0; j < strlen(container2); j++) {
                if(container2[j] == container1[i]){
                    common = container1[i];
                    goto postLoop;
                }
            }
        }

        postLoop:
        if(common>=65 && common <= 90)sum+=common-38;
        else sum += common-96;
    }

    printf("Sum of properties : %d\n",sum);

    return 0;
}

