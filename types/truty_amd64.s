#include "textflag.h"
TEXT ·int2bool(SB),NOSPLIT,$0
    MOVQ len+16(FP),AX
    MOVQ out+24(FP),DX
    MOVQ data+0(FP),BX
    VZEROALL
    start:
    VMOVUPD (BX),Y0
    VPCMPEQQ Y0,Y1,Y0
    VPMOVMSKB Y0,CX
    NOTQ CX
    MOVD CX,(DX)
    ADDQ $32,BX
    ADDQ $4,DX
    SUBQ $4,AX
    CMPQ AX,$0
    JGT  start
    RET

TEXT ·float2bool(SB),NOSPLIT,$0
    MOVQ len+16(FP),AX
    MOVQ out+24(FP),DX
    MOVQ data+0(FP),BX
    VZEROALL
    start:
    VMOVUPD (BX),Y0
    VPCMPEQQ Y0,Y1,Y0
    VPMOVMSKB Y0,CX
    NOTQ CX
    MOVD CX,(DX)
    ADDQ $32,BX
    ADDQ $4,DX
    SUBQ $4,AX
    CMPQ AX,$0
    JGT  start
    RET

TEXT ·boolAnd(SB),NOSPLIT,$0
    MOVQ len+16(FP),AX
    MOVQ data+0(FP),BX
    MOVQ data2+24(FP),CX
    start:
    VMOVUPD (BX),Y0
    VMOVUPD (CX),Y1
    VPAND Y0,Y1,Y2
    VMOVUPD Y2,(CX)
    ADDQ $32,BX
    ADDQ $32,CX
    SUBQ $32,AX
    CMPQ AX,$0
    JGT  start
    RET
