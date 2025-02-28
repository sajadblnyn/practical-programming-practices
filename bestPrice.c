#include <stdio.h>
#include <stdlib.h>
int maxProfit(int *prices, int pricesSize) {
    int profit=0;
    int max_profit=0;
    int min=prices[0];

    for (int i=0;i<pricesSize;i++){

        if(prices[i]<min){
            min=prices[i];
        }else{
            profit=prices[i]-min;

            if(profit>max_profit){
                max_profit=profit;
            }
        }
    }
    return max_profit;
}
void main(){
    int *p;

    p = malloc(sizeof(int)*5);

    p[0]=7;
    p[1]=1;
    p[2]=5;
    p[3]=3;
    p[4]=6;
    
    int m=maxProfit(p,5);

    printf("%d\n",m);
    free(p);
    p=NULL;
}

