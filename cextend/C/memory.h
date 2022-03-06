#ifndef MEMORY
#define MEMORY
void* Malloc(size_t size) ;
void Free(void* ptr) ;
void Memcpy(uintptr_t dest, uintptr_t src, size_t n) ;
#endif // MEMORY
