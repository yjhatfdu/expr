#include "textflag.h"
TEXT ·cvtInt2Float(SB),NOSPLIT,$0
   MOVQ $0,AX
   MOVQ data1+0(FP),BX
   MOVQ data2+24(FP),CX
   MOVQ len+8(FP),DX
   start:
   CVTSQ2SD (BX),X1
   MOVSD X1,(CX)
   ADDQ $8,BX
   ADDQ $8,CX
   INCQ AX
   CMPQ AX,DX
   JLT  start
   RET

TEXT ·cvtFloat2Int(SB),NOSPLIT,$8
   MOVQ $0,AX
   MOVQ data1+0(FP),BX
   MOVQ data2+24(FP),CX
   MOVQ len+8(FP),BP
   start:
   MOVSD (BX),X1
   CVTSD2SQ X1,DX
   MOVQ DX,(CX)
   ADDQ $8,BX
   ADDQ $8,CX
   INCQ AX
   CMPQ AX,BP
   JLT  start
   RET

