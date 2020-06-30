#include "textflag.h"
TEXT ·addIntInt(SB),NOSPLIT,$8
    MOVQ  $0,AX
    MOVQ  data+0(FP),BX
    MOVQ  data2+24(FP),CX
    MOVQ  out+48(FP),DX
    MOVQ  cap+16(FP),BP
    start:
    VMOVUPD (BX),Y0
    ADDQ  $32,BX
    VMOVUPD (CX),Y1
    ADDQ  $32,CX
    VPADDQ Y0,Y1,Y2
    VMOVUPD Y2,(DX)
    ADDQ $32,DX
    ADDQ $4,AX
    CMPQ AX,BP
    JL  start
    RET

TEXT ·subIntInt(SB),NOSPLIT,$8
    MOVQ  $0,AX
    MOVQ  data+0(FP),BX
    MOVQ  data2+24(FP),CX
    MOVQ  out+48(FP),DX
    MOVQ  cap+16(FP),BP
    start:
    VMOVUPD (BX),Y0
    ADDQ  $32,BX
    VMOVUPD (CX),Y1
    ADDQ  $32,CX
    VPSUBQ Y0,Y1,Y2
    VMOVUPD Y2,(DX)
    ADDQ $32,DX
    ADDQ $4,AX
    CMPQ AX,BP
    JL  start
    RET



TEXT ·addIntS(SB),NOSPLIT,$0
    MOVQ  len+40(FP),CX
    MOVQ  data+0(FP),BX
    MOVQ  out+24(FP),DX
    VBROADCASTSD scalar+48(FP),Y1
    MOVQ  $0,AX
    start:
    VMOVUPD (BX),Y0
    ADDQ  $32,BX
    VPADDQ Y0,Y1,Y2
    VMOVUPD Y2,(DX)
    ADDQ  $32,DX
    ADDQ $4,AX
    CMPQ AX,CX
    JLT  start
    RET


TEXT ·addIntFloat(SB),NOSPLIT,$48
    MOVQ  $0,AX
    MOVQ  cap+16(FP),BX
    MOVQ  BX,a-8(SP)
    MOVQ  data+0(FP),BX
    MOVQ  data2+24(FP),CX
    MOVQ  out+48(FP),DX
    start:
    CVTSQ2SD (BX),X0
    MOVD X0,i1-40(SP)
    ADDQ $8,BX
    CVTSQ2SD (BX),X0
    MOVD X0,i2-32(SP)
    ADDQ $8,BX
    CVTSQ2SD (BX),X0
    MOVD X0,i3+-24(SP)
    ADDQ $8,BX
    CVTSQ2SD (BX),X0
    MOVD X0,i4-16(SP)
    ADDQ $8,BX
    VMOVUPD i-40(SP),Y0
    VMOVUPD (CX),Y1
    ADDQ  $32,CX
    VADDPD Y0,Y1,Y2
    VMOVUPD Y2,(DX)
    ADDQ $32,DX
    ADDQ $4,AX
    CMPQ AX,t-8(SP)
    JL  start
    RET

TEXT ·addIntSFloat(SB),NOSPLIT,$0
    MOVQ  $0,AX
    CVTSQ2SD i+48(FP),X0
    VBROADCASTSD X0,Y1
    MOVQ  data+0(FP),BX
    MOVQ  out+24(FP),CX
    MOVQ  cap+16(FP),DX
    start:
    VMOVUPD (BX),Y0
    ADDQ  $32,BX
    VADDPD Y0,Y1,Y2
    VMOVUPD Y2,(CX)
    ADDQ  $32,CX
    ADDQ $4,AX
    CMPQ AX,DX
    JL  start
    RET

TEXT ·addIntFloatS(SB),NOSPLIT,$32
    MOVQ  $0,AX
    MOVQ  data+0(FP),BX
    MOVQ  cap+16(FP),CX
    MOVQ  out+24(FP),DX
    MOVSD f+48(FP),X0
    VBROADCASTSD X0,Y1
    start:
    CVTSQ2SD (BX),X0
    MOVD X0,i1-32(SP)
    ADDQ $8,BX
    CVTSQ2SD (BX),X0
    MOVD X0,i2-24(SP)
    ADDQ $8,BX
    CVTSQ2SD (BX),X0
    MOVD X0,i3-16(SP)
    ADDQ $8,BX
    CVTSQ2SD (BX),X0
    MOVD X0,i4-8(SP)
    ADDQ $8,BX
    VMOVUPD i-32(SP),Y0
    VADDPD Y0,Y1,Y2
    VMOVUPD Y2,(DX)
    ADDQ $32,DX
    ADDQ $4,AX
    CMPQ AX,CX
    JL  start
    RET
