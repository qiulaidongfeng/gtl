#include "textflag.h"

//func Copy_MOVSQ(dest, src unsafe.Pointer, n uint)
// use AX CX DI SI R8 R9 R10 R11 R12 
//R8=dest R9=src R10=n R11=8n R12=1n
TEXT ·Copy_MOVSQ(SB),NOSPLIT|DUPOK,$128-24
//将参数移到寄存器
	MOVQ dest+0(FP),R8
	MOVQ src+8(FP),R9
	MOVQ n+16(FP),R10
	CMPQ R10,$0
	JE ret //复制字节数等于0
	MOVQ R9,SI
	MOVQ R8,DI
	MOVQ R10,CX
	CMPQ CX,$8
	JB copy1no8
//获取8个字节复制次数
	MOVQ R10,0(SP)
	MOVQ $8,AX
	MOVQ AX,8(SP)
	CALL ·div(SB)
	MOVQ 16(SP),R11
	MOVQ 24(SP),R12
//循环从src复制到dest,每次复制8个字节
	MOVQ R11,CX
	CLD
	REP 
	MOVSQ
	CMPQ R12,$0
	JNE copy1
	RET
copy1:
	MOVQ R12,CX
	CLD
	REP 
	MOVSB
	RET
copy1no8:
	CLD
	REP 
	MOVSB
	RET
ret:
	RET
	
//func div(a, b uint64) (c uint64,d uint64)
// use and save AX CX DX 
TEXT ·div(SB),NOSPLIT|DUPOK,$0-32
	MOVQ a+0(FP),AX
	MOVQ b+8(FP),CX
	MOVQ $0,DX
	DIVQ CX
	MOVQ AX,c+16(FP)
	MOVQ DX,d+24(FP)
	RET

//func Copy_SSE_Movups(dest, src unsafe.Pointer, n uint)
// use AX CX DI SI R8 R9 R10 R11 R12 
//R8=dest R9=src R10=n R11=16n R12=1n 
TEXT ·Copy_SSE_Movups(SB),NOSPLIT|DUPOK,$40-24
//将参数移到寄存器
	MOVQ dest+0(FP),R8
	MOVQ src+8(FP),R9
	MOVQ n+16(FP),R10
	MOVQ R8,DI
	MOVQ R9,SI
	MOVQ R10,CX
	CMPQ R10,$0
	JE ret
	CMPQ CX,$16
	JB copy1no16 //如果复制长度小于16
//获取16个字节复制次数
	MOVQ R10,0(SP)
	MOVQ $16,AX
	MOVQ AX,8(SP)
	CALL ·div(SB)
	MOVQ 16(SP),R11
	MOVQ 24(SP),R12
//循环从src复制到dest,每次复制16个字节
	MOVQ R11,CX
copy16:
	MOVUPS 0(SI),X1
	MOVUPS X1,0(DI)
	ADDQ $16,SI
	ADDQ $16,DI
	SUBQ $1,CX
	CMPQ CX,$0
	JNE copy16 
//不为0循环
	CMPQ R12,$0
	JNE copy1
	RET
copy1:
	MOVQ R12,CX
	CLD
	REP 
	MOVSB
	RET	
copy1no16:
	CLD
	REP 
	MOVSB
	RET	
ret:
	RET	
	