#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <math.h>

int checkScore(int arr[150][150],int size,int y,int x){
    int flags[4] = {-1,-1,-1,-1},score = 1;
    for(int i=0;i<size;i++){
        //check right
        if(i>x){
            if(arr[y][i] >= arr[y][x]){
                if(flags[2] == -1)flags[2] = i-x;
            } else if(i==size-1 && flags[2] == -1)flags[2] = i-x;
        }
        //check bottom
        if(i>y){
            if(arr[i][x] >= arr[y][x]){
                if(flags[3] == -1)flags[3] = i-y;
            }else if(i==size-1 && flags[3] == -1)flags[3] = i-y;
        }
    }
    for(int i=size-1;i>=0;i--){
        //check left
        if(i<x){
            if(arr[y][i] >= arr[y][x]){
                if(flags[0] == -1)flags[0] = x-i;
            }else if(i==0 && flags[0] == -1)flags[0] = x;
        }
        //check top
        if(i<y){
            if(arr[i][x] >= arr[y][x]){
                if(flags[1] == -1)flags[1] = y-i;
            }else if(i==0 && flags[1] == -1)flags[1] = y;
        }
    }

    for (int i = 0; i < 4; i++) {
        score*=abs(flags[i]);
    }
    
    return score;
}

int main(){
    int grid[150][150];
    char line[150];
    int gridSize,cnt=0,max=0;

    FILE *input;
    input = fopen("2022/inputs/day8inp.txt","r");

    while(fgets(line,150,input) != NULL){
        for(int i=0;i<strlen(line);i++)grid[cnt][i] = line[i] - '0';
        cnt++;
    }
    gridSize = cnt;

    for(int i=1;i<gridSize-1;i++){
        for(int j=1;j<gridSize-1;j++){
            int s = checkScore(grid,gridSize,i,j);
            max = s>max?s:max;
        }
    }

    printf("Max scenic score : %d\n",max);

    return 0;
}

