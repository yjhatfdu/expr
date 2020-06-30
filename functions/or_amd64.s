#include "textflag.h"
TEXT ·orBool(SB),NOSPLIT,$8
    MOVQ  $0,AX
    MOVQ  data+0(FP),BX
    MOVQ  data2+24(FP),CX
    MOVQ  out+48(FP),DX
    MOVQ  len+16(FP),BP
    start:
    MOVQ  data+0(FP),BX
    VMOVUPD (BX),Y0
    ADDQ  $32,BX
    VMOVUPD (CX),Y1
    VPOR Y0,Y1,Y2
    ADDQ  $32,CX
    VMOVUPD Y2,(DX)
    ADDQ $32,DX
    ADDQ $32,AX
    CMPQ AX,BP
    JL  start
    RET
TEXT ·orBoolS(SB),NOSPLIT,$0
    MOVQ $0,AX
    MOVQ  len+40(FP),CX
    MOVQ  data+0(FP),BX
    MOVQ  out+24(FP),DX
    VPBROADCASTB scalar+48(FP),Y1
    start:
    MOVQ  data+0(FP),BX
    VMOVUPD (BX),Y0
    ADDQ  $32,BX
    VPOR Y0,Y1,Y2
    MOVQ  out+24(FP),DX
    VMOVUPD Y2,(DX)
    ADDQ  $32,DX
    ADDQ $32,AX
    CMPQ AX,CX
    JLT  start
    RET
