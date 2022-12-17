#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int checkVisible(int arr[150][150],int size,int y,int x){
    int flags[4] = {0,0,0,0};
    for(int i=0;i<size;i++){
        //check row
        // if(i==y)goto verticalCheck;
        if(i<x){
            //check left
            if(arr[y][i] >= arr[y][x])flags[0]++;
        } else if(i>x){
            //check right
            if(arr[y][i] >= arr[y][x])flags[2]++;
        }
        if(i<y){
            //check top
            if(arr[i][x] >= arr[y][x])flags[1]++;
        } else if (i>y){
            //check bottom
            if(arr[i][x] >= arr[y][x])flags[3]++;
        }
    }

    for (int i = 0; i < 4; i++) {
        if (flags[i] == 0) {
            return 1;
            break;
        }
    }
    
    return 0;
}

int main(){
    int grid[150][150];
    char line[150];
    int gridSize,cnt=0,visible=0;

    FILE *input;
    input = fopen("2022/inputs/day8inp.txt","r");

    while(fgets(line,150,input) != NULL){
        for(int i=0;i<strlen(line);i++)grid[cnt][i] = line[i] - '0';
        cnt++;
    }
    gridSize = cnt;

    //add edge trees
    visible+= gridSize*2 + (gridSize-2)*2;
    for(int i=1;i<gridSize-1;i++){
        for(int j=1;j<gridSize-1;j++){
            if(checkVisible(grid,gridSize,i,j))visible++;
        }
    }

    printf("Number of visible trees - %d\n",visible);

    return 0;
}

