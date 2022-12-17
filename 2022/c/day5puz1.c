#include <stdio.h>
#include <string.h>

//Had to slightly modify the input for this, the input for the stacks is accepted in this format
//FSAG
//ADF
//ASM
//instead of this
//[G]
//[A] [F] [M]
//[S] [D] [S]
//[F] [A] [A]

void move(char *inp,char *out){
    int len1 = strlen(inp);
    int len2 = strlen(out);
    out[len2] = inp[len1-1];
    out[len2+1] = '\0';
    inp[len1-1] = '\0';
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
        for(;count>0;count--){
            move(stacks[from],stacks[to]);
        }
    }

    printf("\n");
    //print top crates
    for (int j=0; j<i ; j++) {
        printf("Top of %d = %c\n",j,stacks[j][strlen(stacks[j])-1]);
    }

    return 0;
}

