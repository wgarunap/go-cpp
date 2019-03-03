#include <stdio.h>
#include "c.h"

int testFunc(int a, int b){
    printf("Hello C...\n");
    fflush(stdout);
    return a+b;
};

