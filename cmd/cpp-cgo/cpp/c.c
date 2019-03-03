#include <stdio.h>
#include "c.h"

int64 testFunc(int64 a, int64 b){
    printf("Hello\n");
    fflush(stdout);
    return a+b;
};

