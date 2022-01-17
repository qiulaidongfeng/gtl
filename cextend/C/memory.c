#include <stdlib.h>
#include <string.h>

void* Malloc(size_t size)
{
    return malloc(size) ;
}

void Free(void *ptr)
{
    return free(ptr) ;
}

void Memcpy(void *dest, void *stc, size_t n)
{
    memcpy(dest,stc,n) ;
    return ;
}
