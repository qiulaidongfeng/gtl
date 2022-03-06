#include <stdlib.h>
#include <string.h>
#include <stdint.h>

void* Malloc(size_t size)
{
    return malloc(size) ;
}

void Free(void *ptr)
{
    return free(ptr) ;
}

void Memcpy(uintptr_t dest, uintptr_t stc, size_t n)
{
    memcpy(dest,stc,n) ;
    return ;
}
