"".T.M1 STEXT nosplit size=1 args=0x18 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (interface4.go:10)	TEXT	"".T.M1(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (interface4.go:10)	FUNCDATA	$0, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0000 00000 (interface4.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (interface4.go:10)	FUNCDATA	$5, "".T.M1.arginfo1(SB)
	0x0000 00000 (interface4.go:10)	FUNCDATA	$6, "".T.M1.argliveinfo(SB)
	0x0000 00000 (interface4.go:10)	PCDATA	$3, $1
	0x0000 00000 (interface4.go:12)	RET
	0x0000 c3                                               .
"".T.M2 STEXT nosplit size=1 args=0x18 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (interface4.go:14)	TEXT	"".T.M2(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (interface4.go:14)	FUNCDATA	$0, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0000 00000 (interface4.go:14)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (interface4.go:14)	FUNCDATA	$5, "".T.M2.arginfo1(SB)
	0x0000 00000 (interface4.go:14)	FUNCDATA	$6, "".T.M2.argliveinfo(SB)
	0x0000 00000 (interface4.go:14)	PCDATA	$3, $1
	0x0000 00000 (interface4.go:16)	RET
	0x0000 c3                                               .
"".main STEXT size=268 args=0x0 locals=0x78 funcid=0x0 align=0x0
	0x0000 00000 (interface4.go:23)	TEXT	"".main(SB), ABIInternal, $120-0
	0x0000 00000 (interface4.go:23)	CMPQ	SP, 16(R14)
	0x0004 00004 (interface4.go:23)	PCDATA	$0, $-2
	0x0004 00004 (interface4.go:23)	JLS	258
	0x000a 00010 (interface4.go:23)	PCDATA	$0, $-1
	0x000a 00010 (interface4.go:23)	SUBQ	$120, SP
	0x000e 00014 (interface4.go:23)	MOVQ	BP, 112(SP)
	0x0013 00019 (interface4.go:23)	LEAQ	112(SP), BP
	0x0018 00024 (interface4.go:23)	FUNCDATA	$0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x0018 00024 (interface4.go:23)	FUNCDATA	$1, gclocals·c73a1255c0cd89e7a2b21dfb392a7db1(SB)
	0x0018 00024 (interface4.go:23)	FUNCDATA	$2, "".main.stkobj(SB)
	0x0018 00024 (interface4.go:29)	MOVQ	$17, ""..autotmp_13+88(SP)
	0x0021 00033 (interface4.go:29)	LEAQ	go.string."hello, interface"(SB), CX
	0x0028 00040 (interface4.go:29)	MOVQ	CX, ""..autotmp_13+96(SP)
	0x002d 00045 (interface4.go:29)	MOVQ	$16, ""..autotmp_13+104(SP)
	0x0036 00054 (interface4.go:29)	LEAQ	type."".T(SB), AX
	0x003d 00061 (interface4.go:29)	LEAQ	""..autotmp_13+88(SP), BX
	0x0042 00066 (interface4.go:29)	PCDATA	$1, $0
	0x0042 00066 (interface4.go:29)	CALL	runtime.convT(SB)
	0x0047 00071 (interface4.go:29)	MOVQ	AX, ""..autotmp_38+48(SP)
	0x004c 00076 (interface4.go:32)	MOVQ	$17, ""..autotmp_13+88(SP)
	0x0055 00085 (interface4.go:32)	LEAQ	go.string."hello, interface"(SB), CX
	0x005c 00092 (interface4.go:32)	MOVQ	CX, ""..autotmp_13+96(SP)
	0x0061 00097 (interface4.go:32)	MOVQ	$16, ""..autotmp_13+104(SP)
	0x006a 00106 (interface4.go:32)	LEAQ	""..autotmp_13+88(SP), BX
	0x006f 00111 (interface4.go:32)	LEAQ	type."".T(SB), AX
	0x0076 00118 (interface4.go:32)	PCDATA	$1, $1
	0x0076 00118 (interface4.go:32)	CALL	runtime.convT(SB)
	0x007b 00123 (interface4.go:32)	MOVQ	AX, ""..autotmp_39+40(SP)
	0x0080 00128 (interface4.go:33)	MOVUPS	X15, ""..autotmp_19+72(SP)
	0x0086 00134 (interface4.go:33)	LEAQ	type."".T(SB), CX
	0x008d 00141 (interface4.go:33)	MOVQ	CX, ""..autotmp_19+72(SP)
	0x0092 00146 (interface4.go:33)	MOVQ	""..autotmp_38+48(SP), CX
	0x0097 00151 (interface4.go:33)	MOVQ	CX, ""..autotmp_19+80(SP)
	0x009c 00156 (<unknown line number>)	NOP
	0x009c 00156 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x00a3 00163 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_19+72(SP), CX
	0x00a8 00168 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x00ad 00173 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x00b0 00176 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x00b7 00183 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $2
	0x00b7 00183 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00bc 00188 (interface4.go:34)	MOVUPS	X15, ""..autotmp_21+56(SP)
	0x00c2 00194 (interface4.go:34)	MOVQ	go.itab."".T,"".NonEmptyInterface+8(SB), CX
	0x00c9 00201 (interface4.go:34)	MOVQ	CX, ""..autotmp_21+56(SP)
	0x00ce 00206 (interface4.go:34)	MOVQ	""..autotmp_39+40(SP), CX
	0x00d3 00211 (interface4.go:34)	MOVQ	CX, ""..autotmp_21+64(SP)
	0x00d8 00216 (<unknown line number>)	NOP
	0x00d8 00216 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x00df 00223 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x00e6 00230 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_21+56(SP), CX
	0x00eb 00235 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x00f0 00240 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x00f3 00243 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $0
	0x00f3 00243 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00f8 00248 (interface4.go:35)	MOVQ	112(SP), BP
	0x00fd 00253 (interface4.go:35)	ADDQ	$120, SP
	0x0101 00257 (interface4.go:35)	RET
	0x0102 00258 (interface4.go:35)	NOP
	0x0102 00258 (interface4.go:23)	PCDATA	$1, $-1
	0x0102 00258 (interface4.go:23)	PCDATA	$0, $-2
	0x0102 00258 (interface4.go:23)	CALL	runtime.morestack_noctxt(SB)
	0x0107 00263 (interface4.go:23)	PCDATA	$0, $-1
	0x0107 00263 (interface4.go:23)	JMP	0
	0x0000 49 3b 66 10 0f 86 f8 00 00 00 48 83 ec 78 48 89  I;f.......H..xH.
	0x0010 6c 24 70 48 8d 6c 24 70 48 c7 44 24 58 11 00 00  l$pH.l$pH.D$X...
	0x0020 00 48 8d 0d 00 00 00 00 48 89 4c 24 60 48 c7 44  .H......H.L$`H.D
	0x0030 24 68 10 00 00 00 48 8d 05 00 00 00 00 48 8d 5c  $h....H......H.\
	0x0040 24 58 e8 00 00 00 00 48 89 44 24 30 48 c7 44 24  $X.....H.D$0H.D$
	0x0050 58 11 00 00 00 48 8d 0d 00 00 00 00 48 89 4c 24  X....H......H.L$
	0x0060 60 48 c7 44 24 68 10 00 00 00 48 8d 5c 24 58 48  `H.D$h....H.\$XH
	0x0070 8d 05 00 00 00 00 e8 00 00 00 00 48 89 44 24 28  ...........H.D$(
	0x0080 44 0f 11 7c 24 48 48 8d 0d 00 00 00 00 48 89 4c  D..|$HH......H.L
	0x0090 24 48 48 8b 4c 24 30 48 89 4c 24 50 48 8b 1d 00  $HH.L$0H.L$PH...
	0x00a0 00 00 00 48 8d 4c 24 48 bf 01 00 00 00 48 89 fe  ...H.L$H.....H..
	0x00b0 48 8d 05 00 00 00 00 e8 00 00 00 00 44 0f 11 7c  H...........D..|
	0x00c0 24 38 48 8b 0d 00 00 00 00 48 89 4c 24 38 48 8b  $8H......H.L$8H.
	0x00d0 4c 24 28 48 89 4c 24 40 48 8b 1d 00 00 00 00 48  L$(H.L$@H......H
	0x00e0 8d 05 00 00 00 00 48 8d 4c 24 38 bf 01 00 00 00  ......H.L$8.....
	0x00f0 48 89 fe e8 00 00 00 00 48 8b 6c 24 70 48 83 c4  H.......H.l$pH..
	0x0100 78 c3 e8 00 00 00 00 e9 f4 fe ff ff              x...........
	rel 3+0 t=23 type."".T+0
	rel 3+0 t=23 type."".T+0
	rel 3+0 t=23 type.*os.File+0
	rel 3+0 t=23 type.*os.File+0
	rel 36+4 t=14 go.string."hello, interface"+0
	rel 57+4 t=14 type."".T+0
	rel 67+4 t=7 runtime.convT+0
	rel 88+4 t=14 go.string."hello, interface"+0
	rel 114+4 t=14 type."".T+0
	rel 119+4 t=7 runtime.convT+0
	rel 137+4 t=14 type."".T+0
	rel 159+4 t=14 os.Stdout+0
	rel 179+4 t=14 go.itab.*os.File,io.Writer+0
	rel 184+4 t=7 fmt.Fprintln+0
	rel 197+4 t=14 go.itab."".T,"".NonEmptyInterface+8
	rel 219+4 t=14 os.Stdout+0
	rel 226+4 t=14 go.itab.*os.File,io.Writer+0
	rel 244+4 t=7 fmt.Fprintln+0
	rel 259+4 t=7 runtime.morestack_noctxt+0
"".(*T).M1 STEXT dupok nosplit size=58 args=0x8 locals=0x8 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*T).M1(SB), DUPOK|NOSPLIT|WRAPPER|ABIInternal, $8-8
	0x0000 00000 (<autogenerated>:1)	SUBQ	$8, SP
	0x0004 00004 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0008 00008 (<autogenerated>:1)	LEAQ	(SP), BP
	0x000c 00012 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0010 00016 (<autogenerated>:1)	TESTQ	R12, R12
	0x0013 00019 (<autogenerated>:1)	JNE	41
	0x0015 00021 (<autogenerated>:1)	NOP
	0x0015 00021 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0015 00021 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0015 00021 (<autogenerated>:1)	FUNCDATA	$5, "".(*T).M1.arginfo1(SB)
	0x0015 00021 (<autogenerated>:1)	FUNCDATA	$6, "".(*T).M1.argliveinfo(SB)
	0x0015 00021 (<autogenerated>:1)	PCDATA	$3, $1
	0x0015 00021 (<autogenerated>:1)	TESTQ	AX, AX
	0x0018 00024 (<autogenerated>:1)	JEQ	35
	0x001a 00026 (<autogenerated>:1)	MOVQ	(SP), BP
	0x001e 00030 (<autogenerated>:1)	ADDQ	$8, SP
	0x0022 00034 (<autogenerated>:1)	RET
	0x0023 00035 (<autogenerated>:1)	PCDATA	$1, $1
	0x0023 00035 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0028 00040 (<autogenerated>:1)	XCHGL	AX, AX
	0x0029 00041 (<autogenerated>:1)	LEAQ	16(SP), R13
	0x002e 00046 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0032 00050 (<autogenerated>:1)	JNE	21
	0x0034 00052 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x0038 00056 (<autogenerated>:1)	JMP	21
	0x0000 48 83 ec 08 48 89 2c 24 48 8d 2c 24 4d 8b 66 20  H...H.,$H.,$M.f 
	0x0010 4d 85 e4 75 14 48 85 c0 74 09 48 8b 2c 24 48 83  M..u.H..t.H.,$H.
	0x0020 c4 08 c3 e8 00 00 00 00 90 4c 8d 6c 24 10 4d 39  .........L.l$.M9
	0x0030 2c 24 75 e1 49 89 24 24 eb db                    ,$u.I.$$..
	rel 36+4 t=7 runtime.panicwrap+0
"".(*T).M2 STEXT dupok nosplit size=58 args=0x8 locals=0x8 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*T).M2(SB), DUPOK|NOSPLIT|WRAPPER|ABIInternal, $8-8
	0x0000 00000 (<autogenerated>:1)	SUBQ	$8, SP
	0x0004 00004 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0008 00008 (<autogenerated>:1)	LEAQ	(SP), BP
	0x000c 00012 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0010 00016 (<autogenerated>:1)	TESTQ	R12, R12
	0x0013 00019 (<autogenerated>:1)	JNE	41
	0x0015 00021 (<autogenerated>:1)	NOP
	0x0015 00021 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0015 00021 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0015 00021 (<autogenerated>:1)	FUNCDATA	$5, "".(*T).M2.arginfo1(SB)
	0x0015 00021 (<autogenerated>:1)	FUNCDATA	$6, "".(*T).M2.argliveinfo(SB)
	0x0015 00021 (<autogenerated>:1)	PCDATA	$3, $1
	0x0015 00021 (<autogenerated>:1)	TESTQ	AX, AX
	0x0018 00024 (<autogenerated>:1)	JEQ	35
	0x001a 00026 (<autogenerated>:1)	MOVQ	(SP), BP
	0x001e 00030 (<autogenerated>:1)	ADDQ	$8, SP
	0x0022 00034 (<autogenerated>:1)	RET
	0x0023 00035 (<autogenerated>:1)	PCDATA	$1, $1
	0x0023 00035 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0028 00040 (<autogenerated>:1)	XCHGL	AX, AX
	0x0029 00041 (<autogenerated>:1)	LEAQ	16(SP), R13
	0x002e 00046 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0032 00050 (<autogenerated>:1)	JNE	21
	0x0034 00052 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x0038 00056 (<autogenerated>:1)	JMP	21
	0x0000 48 83 ec 08 48 89 2c 24 48 8d 2c 24 4d 8b 66 20  H...H.,$H.,$M.f 
	0x0010 4d 85 e4 75 14 48 85 c0 74 09 48 8b 2c 24 48 83  M..u.H..t.H.,$H.
	0x0020 c4 08 c3 e8 00 00 00 00 90 4c 8d 6c 24 10 4d 39  .........L.l$.M9
	0x0030 2c 24 75 e1 49 89 24 24 eb db                    ,$u.I.$$..
	rel 36+4 t=7 runtime.panicwrap+0
"".NonEmptyInterface.M1 STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".NonEmptyInterface.M1(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	58
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	85
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(FP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, ""..this+32(FP)
	0x0027 00039 (<autogenerated>:1)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x0027 00039 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0027 00039 (<autogenerated>:1)	FUNCDATA	$5, "".NonEmptyInterface.M1.arginfo1(SB)
	0x0027 00039 (<autogenerated>:1)	FUNCDATA	$6, "".NonEmptyInterface.M1.argliveinfo(SB)
	0x0027 00039 (<autogenerated>:1)	PCDATA	$3, $1
	0x0027 00039 (<autogenerated>:1)	MOVQ	24(AX), CX
	0x002b 00043 (<autogenerated>:1)	MOVQ	BX, AX
	0x002e 00046 (<autogenerated>:1)	PCDATA	$1, $1
	0x002e 00046 (<autogenerated>:1)	CALL	CX
	0x0030 00048 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0035 00053 (<autogenerated>:1)	ADDQ	$16, SP
	0x0039 00057 (<autogenerated>:1)	RET
	0x003a 00058 (<autogenerated>:1)	NOP
	0x003a 00058 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003a 00058 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003a 00058 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x003f 00063 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0044 00068 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0049 00073 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x004e 00078 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0053 00083 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0053 00083 (<autogenerated>:1)	JMP	0
	0x0055 00085 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005a 00090 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 34 48 83 ec 10 48 89 6c 24 08 48  I;f.v4H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 38 48 89 44  .l$.M.f M..u8H.D
	0x0020 24 18 48 89 5c 24 20 48 8b 48 18 48 89 d8 ff d1  $.H.\$ H.H.H....
	0x0030 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24 08 48  H.l$.H....H.D$.H
	0x0040 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08 48 8b  .\$......H.D$.H.
	0x0050 5c 24 10 eb ab 4c 8d 6c 24 18 66 0f 1f 44 00 00  \$...L.l$.f..D..
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type."".NonEmptyInterface+96
	rel 46+0 t=10 +0
	rel 69+4 t=7 runtime.morestack_noctxt+0
"".NonEmptyInterface.M2 STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".NonEmptyInterface.M2(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	58
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	85
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(FP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, ""..this+32(FP)
	0x0027 00039 (<autogenerated>:1)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x0027 00039 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0027 00039 (<autogenerated>:1)	FUNCDATA	$5, "".NonEmptyInterface.M2.arginfo1(SB)
	0x0027 00039 (<autogenerated>:1)	FUNCDATA	$6, "".NonEmptyInterface.M2.argliveinfo(SB)
	0x0027 00039 (<autogenerated>:1)	PCDATA	$3, $1
	0x0027 00039 (<autogenerated>:1)	MOVQ	32(AX), CX
	0x002b 00043 (<autogenerated>:1)	MOVQ	BX, AX
	0x002e 00046 (<autogenerated>:1)	PCDATA	$1, $1
	0x002e 00046 (<autogenerated>:1)	CALL	CX
	0x0030 00048 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0035 00053 (<autogenerated>:1)	ADDQ	$16, SP
	0x0039 00057 (<autogenerated>:1)	RET
	0x003a 00058 (<autogenerated>:1)	NOP
	0x003a 00058 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003a 00058 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003a 00058 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x003f 00063 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0044 00068 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0049 00073 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x004e 00078 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0053 00083 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0053 00083 (<autogenerated>:1)	JMP	0
	0x0055 00085 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005a 00090 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 34 48 83 ec 10 48 89 6c 24 08 48  I;f.v4H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 38 48 89 44  .l$.M.f M..u8H.D
	0x0020 24 18 48 89 5c 24 20 48 8b 48 20 48 89 d8 ff d1  $.H.\$ H.H H....
	0x0030 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24 08 48  H.l$.H....H.D$.H
	0x0040 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08 48 8b  .\$......H.D$.H.
	0x0050 5c 24 10 eb ab 4c 8d 6c 24 18 66 0f 1f 44 00 00  \$...L.l$.f..D..
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type."".NonEmptyInterface+104
	rel 46+0 t=10 +0
	rel 69+4 t=7 runtime.morestack_noctxt+0
type..eq."".T STEXT dupok size=95 args=0x10 locals=0x20 funcid=0x0 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq."".T(SB), DUPOK|ABIInternal, $32-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	68
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$32, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 24(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	24(SP), BP
	0x0014 00020 (<autogenerated>:1)	FUNCDATA	$0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x0014 00020 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0014 00020 (<autogenerated>:1)	FUNCDATA	$5, type..eq."".T.arginfo1(SB)
	0x0014 00020 (<autogenerated>:1)	FUNCDATA	$6, type..eq."".T.argliveinfo(SB)
	0x0014 00020 (<autogenerated>:1)	PCDATA	$3, $1
	0x0014 00020 (<autogenerated>:1)	MOVQ	(AX), DX
	0x0017 00023 (<autogenerated>:1)	CMPQ	(BX), DX
	0x001a 00026 (<autogenerated>:1)	JNE	56
	0x001c 00028 (<autogenerated>:1)	MOVQ	16(AX), CX
	0x0020 00032 (<autogenerated>:1)	MOVQ	8(BX), DX
	0x0024 00036 (<autogenerated>:1)	MOVQ	8(AX), AX
	0x0028 00040 (<autogenerated>:1)	CMPQ	16(BX), CX
	0x002c 00044 (<autogenerated>:1)	JNE	56
	0x002e 00046 (<autogenerated>:1)	MOVQ	DX, BX
	0x0031 00049 (<autogenerated>:1)	PCDATA	$1, $1
	0x0031 00049 (<autogenerated>:1)	CALL	runtime.memequal(SB)
	0x0036 00054 (<autogenerated>:1)	JMP	58
	0x0038 00056 (<autogenerated>:1)	XORL	AX, AX
	0x003a 00058 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003a 00058 (<autogenerated>:1)	MOVQ	24(SP), BP
	0x003f 00063 (<autogenerated>:1)	ADDQ	$32, SP
	0x0043 00067 (<autogenerated>:1)	RET
	0x0044 00068 (<autogenerated>:1)	NOP
	0x0044 00068 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0044 00068 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0044 00068 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0049 00073 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x004e 00078 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0053 00083 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0058 00088 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x005d 00093 (<autogenerated>:1)	PCDATA	$0, $-1
	0x005d 00093 (<autogenerated>:1)	JMP	0
	0x0000 49 3b 66 10 76 3e 48 83 ec 20 48 89 6c 24 18 48  I;f.v>H.. H.l$.H
	0x0010 8d 6c 24 18 48 8b 10 48 39 13 75 1c 48 8b 48 10  .l$.H..H9.u.H.H.
	0x0020 48 8b 53 08 48 8b 40 08 48 39 4b 10 75 0a 48 89  H.S.H.@.H9K.u.H.
	0x0030 d3 e8 00 00 00 00 eb 02 31 c0 48 8b 6c 24 18 48  ........1.H.l$.H
	0x0040 83 c4 20 c3 48 89 44 24 08 48 89 5c 24 10 e8 00  .. .H.D$.H.\$...
	0x0050 00 00 00 48 8b 44 24 08 48 8b 5c 24 10 eb a1     ...H.D$.H.\$...
	rel 50+4 t=7 runtime.memequal+0
	rel 79+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 13  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 13 6e 00 01 00 00 00 00 13  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=22 type.[]interface {}+0
	rel 0+0 t=22 type.error+0
	rel 0+0 t=22 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.string."hello, interface" SRODATA dupok size=16
	0x0000 68 65 6c 6c 6f 2c 20 69 6e 74 65 72 66 61 63 65  hello, interface
go.info."".T.M1$abstract SDWARFABSFCN dupok size=18
	0x0000 05 2e 54 2e 4d 31 00 01 01 13 74 00 00 00 00 00  ..T.M1....t.....
	0x0010 00 00                                            ..
	rel 0+0 t=22 type."".T+0
	rel 13+4 t=31 go.info."".T+0
go.info."".T.M2$abstract SDWARFABSFCN dupok size=18
	0x0000 05 2e 54 2e 4d 32 00 01 01 13 74 00 00 00 00 00  ..T.M2....t.....
	0x0010 00 00                                            ..
	rel 0+0 t=22 type."".T+0
	rel 13+4 t=31 go.info."".T+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*func()- SRODATA dupok size=9
	0x0000 00 07 2a 66 75 6e 63 28 29                       ..*func()
type.*func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9b 90 75 1b 08 08 08 36 00 00 00 00 00 00 00 00  ..u....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 48+8 t=1 type.func()+0
type.func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f6 bc 82 f6 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00                                      ....
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 44+4 t=-32763 type.*func()+0
runtime.interequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.interequal+0
type..namedata.*main.NonEmptyInterface. SRODATA dupok size=25
	0x0000 01 17 2a 6d 61 69 6e 2e 4e 6f 6e 45 6d 70 74 79  ..*main.NonEmpty
	0x0010 49 6e 74 65 72 66 61 63 65                       Interface
type.*"".NonEmptyInterface SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f6 af 20 d1 08 08 08 36 00 00 00 00 00 00 00 00  .. ....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.NonEmptyInterface.+0
	rel 48+8 t=1 type."".NonEmptyInterface+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type..namedata.M1. SRODATA dupok size=4
	0x0000 01 02 4d 31                                      ..M1
type..namedata.M2. SRODATA dupok size=4
	0x0000 01 02 4d 32                                      ..M2
type."".NonEmptyInterface SRODATA dupok size=112
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 4a 37 15 f3 07 08 08 14 00 00 00 00 00 00 00 00  J7..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.interequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.NonEmptyInterface.+0
	rel 44+4 t=5 type.*"".NonEmptyInterface+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".NonEmptyInterface+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.M1.+0
	rel 100+4 t=5 type.func()+0
	rel 104+4 t=5 type..namedata.M2.+0
	rel 108+4 t=5 type.func()+0
type..eqfunc."".T SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq."".T+0
type..namedata.*main.T. SRODATA dupok size=9
	0x0000 01 07 2a 6d 61 69 6e 2e 54                       ..*main.T
type..namedata.*func(*main.T)- SRODATA dupok size=16
	0x0000 00 0e 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 54 29  ..*func(*main.T)
type.*func(*"".T) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 73 30 e0 de 08 08 08 36 00 00 00 00 00 00 00 00  s0.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.T)-+0
	rel 48+8 t=1 type.func(*"".T)+0
type.func(*"".T) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 5b c2 60 8a 02 08 08 33 00 00 00 00 00 00 00 00  [.`....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.T)-+0
	rel 44+4 t=-32763 type.*func(*"".T)+0
	rel 56+8 t=1 type.*"".T+0
type.*"".T SRODATA dupok size=104
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8b 54 ca 19 09 08 08 36 00 00 00 00 00 00 00 00  .T.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 02 00 02 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.T.+0
	rel 48+8 t=1 type."".T+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.M1.+0
	rel 76+4 t=26 type.func()+0
	rel 80+4 t=26 "".(*T).M1+0
	rel 84+4 t=26 "".(*T).M1+0
	rel 88+4 t=5 type..namedata.M2.+0
	rel 92+4 t=26 type.func()+0
	rel 96+4 t=26 "".(*T).M2+0
	rel 100+4 t=26 "".(*T).M2+0
type..namedata.*func(main.T)- SRODATA dupok size=15
	0x0000 00 0d 2a 66 75 6e 63 28 6d 61 69 6e 2e 54 29     ..*func(main.T)
type.*func("".T) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8d 60 38 6a 08 08 08 36 00 00 00 00 00 00 00 00  .`8j...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.T)-+0
	rel 48+8 t=1 type.func("".T)+0
type.func("".T) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 63 a7 41 a9 02 08 08 33 00 00 00 00 00 00 00 00  c.A....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.T)-+0
	rel 44+4 t=-32763 type.*func("".T)+0
	rel 56+8 t=1 type."".T+0
type..namedata.n- SRODATA dupok size=3
	0x0000 00 01 6e                                         ..n
type..namedata.s- SRODATA dupok size=3
	0x0000 00 01 73                                         ..s
type."".T SRODATA dupok size=176
	0x0000 18 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 b5 a4 ef 44 07 08 08 19 00 00 00 00 00 00 00 00  ...D............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 02 00 02 00 40 00 00 00 00 00 00 00  ........@.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 type..eqfunc."".T+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.T.+0
	rel 44+4 t=5 type.*"".T+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".T+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.n-+0
	rel 104+8 t=1 type.int+0
	rel 120+8 t=1 type..namedata.s-+0
	rel 128+8 t=1 type.string+0
	rel 144+4 t=5 type..namedata.M1.+0
	rel 148+4 t=26 type.func()+0
	rel 152+4 t=26 "".(*T).M1+0
	rel 156+4 t=26 "".T.M1+0
	rel 160+4 t=5 type..namedata.M2.+0
	rel 164+4 t=26 type.func()+0
	rel 168+4 t=26 "".(*T).M2+0
	rel 172+4 t=26 "".T.M2+0
go.itab."".T,"".NonEmptyInterface SRODATA dupok size=40
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 b5 a4 ef 44 00 00 00 00 00 00 00 00 00 00 00 00  ...D............
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type."".NonEmptyInterface+0
	rel 8+8 t=1 type."".T+0
	rel 24+8 t=-32767 "".(*T).M1+0
	rel 32+8 t=-32767 "".(*T).M2+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".T.M1.arginfo1 SRODATA static dupok size=11
	0x0000 fe 00 08 fe 08 08 10 08 fd fd ff                 ...........
"".T.M1.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
"".T.M2.arginfo1 SRODATA static dupok size=11
	0x0000 fe 00 08 fe 08 08 10 08 fd fd ff                 ...........
"".T.M2.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·c73a1255c0cd89e7a2b21dfb392a7db1 SRODATA dupok size=14
	0x0000 03 00 00 00 09 00 00 00 00 00 02 00 01 00        ..............
"".main.stkobj SRODATA static size=56
	0x0000 03 00 00 00 00 00 00 00 c8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0020 10 00 00 00 00 00 00 00 e8 ff ff ff 18 00 00 00  ................
	0x0030 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
	rel 36+4 t=5 runtime.gcbits.02+0
	rel 52+4 t=5 runtime.gcbits.02+0
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
"".(*T).M1.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".(*T).M1.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
"".(*T).M2.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".(*T).M2.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·09cf9819fc716118c209c2d2155a3632 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 02 00                    ..........
"".NonEmptyInterface.M1.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".NonEmptyInterface.M1.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
"".NonEmptyInterface.M2.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".NonEmptyInterface.M2.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
type..eq."".T.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
type..eq."".T.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
