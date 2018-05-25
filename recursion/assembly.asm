"".tailRecursion STEXT size=378 args=0x18 locals=0x60
	0x0000 00000 (./main.go:15)	TEXT	"".tailRecursion(SB), $96-24
	0x0000 00000 (./main.go:15)	MOVQ	(TLS), CX
	0x0009 00009 (./main.go:15)	CMPQ	SP, 16(CX)
	0x000d 00013 (./main.go:15)	JLS	368
	0x0013 00019 (./main.go:15)	SUBQ	$96, SP
	0x0017 00023 (./main.go:15)	MOVQ	BP, 88(SP)
	0x001c 00028 (./main.go:15)	LEAQ	88(SP), BP
	0x0021 00033 (./main.go:15)	FUNCDATA	$0, gclocals·cadea2e49003779a155f5f8fb1f0fe78(SB)
	0x0021 00033 (./main.go:15)	FUNCDATA	$1, gclocals·dc990f0ce5208716c4d33a75f4225a61(SB)
	0x0021 00033 (./main.go:15)	MOVQ	"".number+104(SP), AX
	0x0026 00038 (./main.go:17)	MOVQ	AX, ""..autotmp_4+48(SP)
	0x002b 00043 (./main.go:17)	MOVQ	$0, ""..autotmp_3+72(SP)
	0x0034 00052 (./main.go:17)	MOVQ	$0, ""..autotmp_3+80(SP)
	0x003d 00061 (./main.go:17)	LEAQ	type.int(SB), CX
	0x0044 00068 (./main.go:17)	MOVQ	CX, (SP)
	0x0048 00072 (./main.go:17)	LEAQ	""..autotmp_4+48(SP), DX
	0x004d 00077 (./main.go:17)	MOVQ	DX, 8(SP)
	0x0052 00082 (./main.go:17)	PCDATA	$0, $1
	0x0052 00082 (./main.go:17)	CALL	runtime.convT2E64(SB)
	0x0057 00087 (./main.go:17)	MOVQ	24(SP), AX
	0x005c 00092 (./main.go:17)	MOVQ	16(SP), CX
	0x0061 00097 (./main.go:17)	MOVQ	CX, ""..autotmp_3+72(SP)
	0x0066 00102 (./main.go:17)	MOVQ	AX, ""..autotmp_3+80(SP)
	0x006b 00107 (./main.go:17)	LEAQ	go.string."number: %d\n"(SB), AX
	0x0072 00114 (./main.go:17)	MOVQ	AX, (SP)
	0x0076 00118 (./main.go:17)	MOVQ	$11, 8(SP)
	0x007f 00127 (./main.go:17)	LEAQ	""..autotmp_3+72(SP), AX
	0x0084 00132 (./main.go:17)	MOVQ	AX, 16(SP)
	0x0089 00137 (./main.go:17)	MOVQ	$1, 24(SP)
	0x0092 00146 (./main.go:17)	MOVQ	$1, 32(SP)
	0x009b 00155 (./main.go:17)	PCDATA	$0, $1
	0x009b 00155 (./main.go:17)	CALL	log.Printf(SB)
	0x00a0 00160 (./main.go:15)	MOVQ	"".number+104(SP), AX
	0x00a5 00165 (./main.go:15)	MOVQ	"".product+112(SP), CX
	0x00aa 00170 (./main.go:16)	ADDQ	AX, CX
	0x00ad 00173 (./main.go:16)	MOVQ	CX, "".product+112(SP)
	0x00b2 00178 (./main.go:18)	MOVQ	CX, ""..autotmp_6+40(SP)
	0x00b7 00183 (./main.go:18)	MOVQ	$0, ""..autotmp_5+56(SP)
	0x00c0 00192 (./main.go:18)	MOVQ	$0, ""..autotmp_5+64(SP)
	0x00c9 00201 (./main.go:18)	LEAQ	type.int(SB), DX
	0x00d0 00208 (./main.go:18)	MOVQ	DX, (SP)
	0x00d4 00212 (./main.go:18)	LEAQ	""..autotmp_6+40(SP), DX
	0x00d9 00217 (./main.go:18)	MOVQ	DX, 8(SP)
	0x00de 00222 (./main.go:18)	PCDATA	$0, $2
	0x00de 00222 (./main.go:18)	CALL	runtime.convT2E64(SB)
	0x00e3 00227 (./main.go:18)	MOVQ	16(SP), AX
	0x00e8 00232 (./main.go:18)	MOVQ	24(SP), CX
	0x00ed 00237 (./main.go:18)	MOVQ	AX, ""..autotmp_5+56(SP)
	0x00f2 00242 (./main.go:18)	MOVQ	CX, ""..autotmp_5+64(SP)
	0x00f7 00247 (./main.go:18)	LEAQ	go.string."product: %d\n"(SB), AX
	0x00fe 00254 (./main.go:18)	MOVQ	AX, (SP)
	0x0102 00258 (./main.go:18)	MOVQ	$12, 8(SP)
	0x010b 00267 (./main.go:18)	LEAQ	""..autotmp_5+56(SP), AX
	0x0110 00272 (./main.go:18)	MOVQ	AX, 16(SP)
	0x0115 00277 (./main.go:18)	MOVQ	$1, 24(SP)
	0x011e 00286 (./main.go:18)	MOVQ	$1, 32(SP)
	0x0127 00295 (./main.go:18)	PCDATA	$0, $2
	0x0127 00295 (./main.go:18)	CALL	log.Printf(SB)
	0x012c 00300 (./main.go:18)	MOVQ	"".number+104(SP), AX
	0x0131 00305 (./main.go:19)	CMPQ	AX, $1
	0x0135 00309 (./main.go:19)	JNE	326
	0x0137 00311 (./main.go:20)	MOVQ	AX, "".~r2+120(SP)
	0x013c 00316 (./main.go:20)	MOVQ	88(SP), BP
	0x0141 00321 (./main.go:20)	ADDQ	$96, SP
	0x0145 00325 (./main.go:20)	RET
	0x0146 00326 (./main.go:22)	DECQ	AX
	0x0149 00329 (./main.go:22)	MOVQ	AX, (SP)
	0x014d 00333 (./main.go:22)	MOVQ	"".product+112(SP), AX
	0x0152 00338 (./main.go:22)	MOVQ	AX, 8(SP)
	0x0157 00343 (./main.go:22)	PCDATA	$0, $0
	0x0157 00343 (./main.go:22)	CALL	"".tailRecursion(SB)
	0x015c 00348 (./main.go:22)	MOVQ	16(SP), AX
	0x0161 00353 (./main.go:22)	MOVQ	AX, "".~r2+120(SP)
	0x0166 00358 (./main.go:22)	MOVQ	88(SP), BP
	0x016b 00363 (./main.go:22)	ADDQ	$96, SP
	0x016f 00367 (./main.go:22)	RET
	0x0170 00368 (./main.go:22)	NOP
	0x0170 00368 (./main.go:15)	PCDATA	$0, $-1
	0x0170 00368 (./main.go:15)	CALL	runtime.morestack_noctxt(SB)
	0x0175 00373 (./main.go:15)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 5d  dH..%....H;a...]
	0x0010 01 00 00 48 83 ec 60 48 89 6c 24 58 48 8d 6c 24  ...H..`H.l$XH.l$
	0x0020 58 48 8b 44 24 68 48 89 44 24 30 48 c7 44 24 48  XH.D$hH.D$0H.D$H
	0x0030 00 00 00 00 48 c7 44 24 50 00 00 00 00 48 8d 0d  ....H.D$P....H..
	0x0040 00 00 00 00 48 89 0c 24 48 8d 54 24 30 48 89 54  ....H..$H.T$0H.T
	0x0050 24 08 e8 00 00 00 00 48 8b 44 24 18 48 8b 4c 24  $......H.D$.H.L$
	0x0060 10 48 89 4c 24 48 48 89 44 24 50 48 8d 05 00 00  .H.L$HH.D$PH....
	0x0070 00 00 48 89 04 24 48 c7 44 24 08 0b 00 00 00 48  ..H..$H.D$.....H
	0x0080 8d 44 24 48 48 89 44 24 10 48 c7 44 24 18 01 00  .D$HH.D$.H.D$...
	0x0090 00 00 48 c7 44 24 20 01 00 00 00 e8 00 00 00 00  ..H.D$ .........
	0x00a0 48 8b 44 24 68 48 8b 4c 24 70 48 01 c1 48 89 4c  H.D$hH.L$pH..H.L
	0x00b0 24 70 48 89 4c 24 28 48 c7 44 24 38 00 00 00 00  $pH.L$(H.D$8....
	0x00c0 48 c7 44 24 40 00 00 00 00 48 8d 15 00 00 00 00  H.D$@....H......
	0x00d0 48 89 14 24 48 8d 54 24 28 48 89 54 24 08 e8 00  H..$H.T$(H.T$...
	0x00e0 00 00 00 48 8b 44 24 10 48 8b 4c 24 18 48 89 44  ...H.D$.H.L$.H.D
	0x00f0 24 38 48 89 4c 24 40 48 8d 05 00 00 00 00 48 89  $8H.L$@H......H.
	0x0100 04 24 48 c7 44 24 08 0c 00 00 00 48 8d 44 24 38  .$H.D$.....H.D$8
	0x0110 48 89 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7  H.D$.H.D$.....H.
	0x0120 44 24 20 01 00 00 00 e8 00 00 00 00 48 8b 44 24  D$ .........H.D$
	0x0130 68 48 83 f8 01 75 0f 48 89 44 24 78 48 8b 6c 24  hH...u.H.D$xH.l$
	0x0140 58 48 83 c4 60 c3 48 ff c8 48 89 04 24 48 8b 44  XH..`.H..H..$H.D
	0x0150 24 70 48 89 44 24 08 e8 00 00 00 00 48 8b 44 24  $pH.D$......H.D$
	0x0160 10 48 89 44 24 78 48 8b 6c 24 58 48 83 c4 60 c3  .H.D$xH.l$XH..`.
	0x0170 e8 00 00 00 00 e9 86 fe ff ff                    ..........
	rel 5+4 t=16 TLS+0
	rel 64+4 t=15 type.int+0
	rel 83+4 t=8 runtime.convT2E64+0
	rel 110+4 t=15 go.string."number: %d\n"+0
	rel 156+4 t=8 log.Printf+0
	rel 204+4 t=15 type.int+0
	rel 223+4 t=8 runtime.convT2E64+0
	rel 250+4 t=15 go.string."product: %d\n"+0
	rel 296+4 t=8 log.Printf+0
	rel 344+4 t=8 "".tailRecursion+0
	rel 369+4 t=8 runtime.morestack_noctxt+0
"".recursiveChannel STEXT size=158 args=0x18 locals=0x38
	0x0000 00000 (./main.go:26)	TEXT	"".recursiveChannel(SB), $56-24
	0x0000 00000 (./main.go:26)	MOVQ	(TLS), CX
	0x0009 00009 (./main.go:26)	CMPQ	SP, 16(CX)
	0x000d 00013 (./main.go:26)	JLS	148
	0x0013 00019 (./main.go:26)	SUBQ	$56, SP
	0x0017 00023 (./main.go:26)	MOVQ	BP, 48(SP)
	0x001c 00028 (./main.go:26)	LEAQ	48(SP), BP
	0x0021 00033 (./main.go:26)	FUNCDATA	$0, gclocals·a35fa7d7e7129fc330c152d6236a3e5c(SB)
	0x0021 00033 (./main.go:26)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (./main.go:26)	MOVQ	"".number+64(SP), AX
	0x0026 00038 (./main.go:26)	MOVQ	"".product+72(SP), CX
	0x002b 00043 (./main.go:27)	ADDQ	AX, CX
	0x002e 00046 (./main.go:28)	CMPQ	AX, $1
	0x0032 00050 (./main.go:28)	JNE	91
	0x0034 00052 (./main.go:29)	MOVQ	CX, ""..autotmp_3+40(SP)
	0x0039 00057 (./main.go:29)	MOVQ	"".result+80(SP), AX
	0x003e 00062 (./main.go:29)	MOVQ	AX, (SP)
	0x0042 00066 (./main.go:29)	LEAQ	""..autotmp_3+40(SP), AX
	0x0047 00071 (./main.go:29)	MOVQ	AX, 8(SP)
	0x004c 00076 (./main.go:29)	PCDATA	$0, $1
	0x004c 00076 (./main.go:29)	CALL	runtime.chansend1(SB)
	0x0051 00081 (./main.go:30)	MOVQ	48(SP), BP
	0x0056 00086 (./main.go:30)	ADDQ	$56, SP
	0x005a 00090 (./main.go:30)	RET
	0x005b 00091 (./main.go:32)	DECQ	AX
	0x005e 00094 (./main.go:32)	MOVQ	AX, 16(SP)
	0x0063 00099 (./main.go:32)	MOVQ	CX, 24(SP)
	0x0068 00104 (./main.go:32)	MOVQ	"".result+80(SP), AX
	0x006d 00109 (./main.go:32)	MOVQ	AX, 32(SP)
	0x0072 00114 (./main.go:32)	MOVL	$24, (SP)
	0x0079 00121 (./main.go:32)	LEAQ	"".recursiveChannel·f(SB), AX
	0x0080 00128 (./main.go:32)	MOVQ	AX, 8(SP)
	0x0085 00133 (./main.go:32)	PCDATA	$0, $1
	0x0085 00133 (./main.go:32)	CALL	runtime.newproc(SB)
	0x008a 00138 (./main.go:33)	MOVQ	48(SP), BP
	0x008f 00143 (./main.go:33)	ADDQ	$56, SP
	0x0093 00147 (./main.go:33)	RET
	0x0094 00148 (./main.go:33)	NOP
	0x0094 00148 (./main.go:26)	PCDATA	$0, $-1
	0x0094 00148 (./main.go:26)	CALL	runtime.morestack_noctxt(SB)
	0x0099 00153 (./main.go:26)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 81  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 38 48 89 6c 24 30 48 8d 6c 24  ...H..8H.l$0H.l$
	0x0020 30 48 8b 44 24 40 48 8b 4c 24 48 48 01 c1 48 83  0H.D$@H.L$HH..H.
	0x0030 f8 01 75 27 48 89 4c 24 28 48 8b 44 24 50 48 89  ..u'H.L$(H.D$PH.
	0x0040 04 24 48 8d 44 24 28 48 89 44 24 08 e8 00 00 00  .$H.D$(H.D$.....
	0x0050 00 48 8b 6c 24 30 48 83 c4 38 c3 48 ff c8 48 89  .H.l$0H..8.H..H.
	0x0060 44 24 10 48 89 4c 24 18 48 8b 44 24 50 48 89 44  D$.H.L$.H.D$PH.D
	0x0070 24 20 c7 04 24 18 00 00 00 48 8d 05 00 00 00 00  $ ..$....H......
	0x0080 48 89 44 24 08 e8 00 00 00 00 48 8b 6c 24 30 48  H.D$......H.l$0H
	0x0090 83 c4 38 c3 e8 00 00 00 00 e9 62 ff ff ff        ..8.......b...
	rel 5+4 t=16 TLS+0
	rel 77+4 t=8 runtime.chansend1+0
	rel 124+4 t=15 "".recursiveChannel·f+0
	rel 134+4 t=8 runtime.newproc+0
	rel 149+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=270 args=0x0 locals=0x70
	0x0000 00000 (./main.go:35)	TEXT	"".main(SB), $112-0
	0x0000 00000 (./main.go:35)	MOVQ	(TLS), CX
	0x0009 00009 (./main.go:35)	CMPQ	SP, 16(CX)
	0x000d 00013 (./main.go:35)	JLS	260
	0x0013 00019 (./main.go:35)	SUBQ	$112, SP
	0x0017 00023 (./main.go:35)	MOVQ	BP, 104(SP)
	0x001c 00028 (./main.go:35)	LEAQ	104(SP), BP
	0x0021 00033 (./main.go:35)	FUNCDATA	$0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x0021 00033 (./main.go:35)	FUNCDATA	$1, gclocals·3a5f6b21d2b25517271e2f12e2835dc4(SB)
	0x0021 00033 (./main.go:35)	LEAQ	type.chan int(SB), AX
	0x0028 00040 (./main.go:42)	MOVQ	AX, (SP)
	0x002c 00044 (./main.go:42)	MOVQ	$0, 8(SP)
	0x0035 00053 (./main.go:42)	PCDATA	$0, $0
	0x0035 00053 (./main.go:42)	CALL	runtime.makechan(SB)
	0x003a 00058 (./main.go:42)	MOVQ	16(SP), AX
	0x003f 00063 (./main.go:42)	MOVQ	AX, "".result+80(SP)
	0x0044 00068 (./main.go:44)	MOVQ	$4, (SP)
	0x004c 00076 (./main.go:44)	MOVQ	$0, 8(SP)
	0x0055 00085 (./main.go:44)	PCDATA	$0, $1
	0x0055 00085 (./main.go:44)	CALL	"".recursiveChannel(SB)
	0x005a 00090 (./main.go:45)	MOVQ	$0, ""..autotmp_2+72(SP)
	0x0063 00099 (./main.go:45)	MOVQ	"".result+80(SP), AX
	0x0068 00104 (./main.go:45)	MOVQ	AX, (SP)
	0x006c 00108 (./main.go:45)	LEAQ	""..autotmp_2+72(SP), AX
	0x0071 00113 (./main.go:45)	MOVQ	AX, 8(SP)
	0x0076 00118 (./main.go:45)	PCDATA	$0, $0
	0x0076 00118 (./main.go:45)	CALL	runtime.chanrecv1(SB)
	0x007b 00123 (./main.go:45)	MOVQ	""..autotmp_2+72(SP), AX
	0x0080 00128 (./main.go:46)	MOVQ	AX, ""..autotmp_4+64(SP)
	0x0085 00133 (./main.go:46)	MOVQ	$0, ""..autotmp_3+88(SP)
	0x008e 00142 (./main.go:46)	MOVQ	$0, ""..autotmp_3+96(SP)
	0x0097 00151 (./main.go:46)	LEAQ	type.int(SB), AX
	0x009e 00158 (./main.go:46)	MOVQ	AX, (SP)
	0x00a2 00162 (./main.go:46)	LEAQ	""..autotmp_4+64(SP), AX
	0x00a7 00167 (./main.go:46)	MOVQ	AX, 8(SP)
	0x00ac 00172 (./main.go:46)	PCDATA	$0, $2
	0x00ac 00172 (./main.go:46)	CALL	runtime.convT2E64(SB)
	0x00b1 00177 (./main.go:46)	MOVQ	24(SP), AX
	0x00b6 00182 (./main.go:46)	MOVQ	16(SP), CX
	0x00bb 00187 (./main.go:46)	MOVQ	CX, ""..autotmp_3+88(SP)
	0x00c0 00192 (./main.go:46)	MOVQ	AX, ""..autotmp_3+96(SP)
	0x00c5 00197 (./main.go:46)	LEAQ	go.string."recursive: %d\n"(SB), AX
	0x00cc 00204 (./main.go:46)	MOVQ	AX, (SP)
	0x00d0 00208 (./main.go:46)	MOVQ	$14, 8(SP)
	0x00d9 00217 (./main.go:46)	LEAQ	""..autotmp_3+88(SP), AX
	0x00de 00222 (./main.go:46)	MOVQ	AX, 16(SP)
	0x00e3 00227 (./main.go:46)	MOVQ	$1, 24(SP)
	0x00ec 00236 (./main.go:46)	MOVQ	$1, 32(SP)
	0x00f5 00245 (./main.go:46)	PCDATA	$0, $2
	0x00f5 00245 (./main.go:46)	CALL	fmt.Printf(SB)
	0x00fa 00250 (./main.go:47)	MOVQ	104(SP), BP
	0x00ff 00255 (./main.go:47)	ADDQ	$112, SP
	0x0103 00259 (./main.go:47)	RET
	0x0104 00260 (./main.go:47)	NOP
	0x0104 00260 (./main.go:35)	PCDATA	$0, $-1
	0x0104 00260 (./main.go:35)	CALL	runtime.morestack_noctxt(SB)
	0x0109 00265 (./main.go:35)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 f1  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 70 48 89 6c 24 68 48 8d 6c 24  ...H..pH.l$hH.l$
	0x0020 68 48 8d 05 00 00 00 00 48 89 04 24 48 c7 44 24  hH......H..$H.D$
	0x0030 08 00 00 00 00 e8 00 00 00 00 48 8b 44 24 10 48  ..........H.D$.H
	0x0040 89 44 24 50 48 c7 04 24 04 00 00 00 48 c7 44 24  .D$PH..$....H.D$
	0x0050 08 00 00 00 00 e8 00 00 00 00 48 c7 44 24 48 00  ..........H.D$H.
	0x0060 00 00 00 48 8b 44 24 50 48 89 04 24 48 8d 44 24  ...H.D$PH..$H.D$
	0x0070 48 48 89 44 24 08 e8 00 00 00 00 48 8b 44 24 48  HH.D$......H.D$H
	0x0080 48 89 44 24 40 48 c7 44 24 58 00 00 00 00 48 c7  H.D$@H.D$X....H.
	0x0090 44 24 60 00 00 00 00 48 8d 05 00 00 00 00 48 89  D$`....H......H.
	0x00a0 04 24 48 8d 44 24 40 48 89 44 24 08 e8 00 00 00  .$H.D$@H.D$.....
	0x00b0 00 48 8b 44 24 18 48 8b 4c 24 10 48 89 4c 24 58  .H.D$.H.L$.H.L$X
	0x00c0 48 89 44 24 60 48 8d 05 00 00 00 00 48 89 04 24  H.D$`H......H..$
	0x00d0 48 c7 44 24 08 0e 00 00 00 48 8d 44 24 58 48 89  H.D$.....H.D$XH.
	0x00e0 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7 44 24  D$.H.D$.....H.D$
	0x00f0 20 01 00 00 00 e8 00 00 00 00 48 8b 6c 24 68 48   .........H.l$hH
	0x0100 83 c4 70 c3 e8 00 00 00 00 e9 f2 fe ff ff        ..p...........
	rel 5+4 t=16 TLS+0
	rel 36+4 t=15 type.chan int+0
	rel 54+4 t=8 runtime.makechan+0
	rel 86+4 t=8 "".recursiveChannel+0
	rel 119+4 t=8 runtime.chanrecv1+0
	rel 154+4 t=15 type.int+0
	rel 173+4 t=8 runtime.convT2E64+0
	rel 200+4 t=15 go.string."recursive: %d\n"+0
	rel 246+4 t=8 fmt.Printf+0
	rel 261+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=96 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	89
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0024 00036 (<autogenerated>:1)	JLS	47
	0x0026 00038 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002a 00042 (<autogenerated>:1)	ADDQ	$8, SP
	0x002e 00046 (<autogenerated>:1)	RET
	0x002f 00047 (<autogenerated>:1)	JNE	56
	0x0031 00049 (<autogenerated>:1)	PCDATA	$0, $0
	0x0031 00049 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0036 00054 (<autogenerated>:1)	UNDEF
	0x0038 00056 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x003f 00063 (<autogenerated>:1)	PCDATA	$0, $0
	0x003f 00063 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x0044 00068 (<autogenerated>:1)	PCDATA	$0, $0
	0x0044 00068 (<autogenerated>:1)	CALL	log.init(SB)
	0x0049 00073 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x0050 00080 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0054 00084 (<autogenerated>:1)	ADDQ	$8, SP
	0x0058 00088 (<autogenerated>:1)	RET
	0x0059 00089 (<autogenerated>:1)	NOP
	0x0059 00089 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0059 00089 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005e 00094 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 4a 48  dH..%....H;a.vJH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 76 09 48 8b 2c 24 48 83 c4 08 c3 75  ..<.v.H.,$H....u
	0x0030 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01 e8  ................
	0x0040 00 00 00 00 e8 00 00 00 00 c6 05 00 00 00 00 02  ................
	0x0050 48 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb a0  H.,$H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 50+4 t=8 runtime.throwinit+0
	rel 58+4 t=15 "".initdone·+-1
	rel 64+4 t=8 fmt.init+0
	rel 69+4 t=8 log.init+0
	rel 75+4 t=15 "".initdone·+-1
	rel 90+4 t=8 runtime.morestack_noctxt+0
go.string."number: %d\n" SRODATA dupok size=11
	0x0000 6e 75 6d 62 65 72 3a 20 25 64 0a                 number: %d.
go.string."product: %d\n" SRODATA dupok size=12
	0x0000 70 72 6f 64 75 63 74 3a 20 25 64 0a              product: %d.
go.info."".tailRecursion SDWARFINFO size=84
	0x0000 02 22 22 2e 74 61 69 6c 52 65 63 75 72 73 69 6f  ."".tailRecursio
	0x0010 6e 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  n...............
	0x0020 00 00 01 9c 01 05 6e 75 6d 62 65 72 00 01 9c 00  ......number....
	0x0030 00 00 00 05 70 72 6f 64 75 63 74 00 04 9c 11 08  ....product.....
	0x0040 22 00 00 00 00 05 7e 72 32 00 04 9c 11 10 22 00  ".....~r2.....".
	0x0050 00 00 00 00                                      ....
	rel 18+8 t=1 "".tailRecursion+0
	rel 26+8 t=1 "".tailRecursion+378
	rel 47+4 t=28 go.info.int+0
	rel 65+4 t=28 go.info.int+0
	rel 79+4 t=28 go.info.int+0
go.range."".tailRecursion SDWARFRANGE size=0
go.info."".recursiveChannel SDWARFINFO size=90
	0x0000 02 22 22 2e 72 65 63 75 72 73 69 76 65 43 68 61  ."".recursiveCha
	0x0010 6e 6e 65 6c 00 00 00 00 00 00 00 00 00 00 00 00  nnel............
	0x0020 00 00 00 00 00 01 9c 01 05 6e 75 6d 62 65 72 00  .........number.
	0x0030 01 9c 00 00 00 00 05 70 72 6f 64 75 63 74 00 04  .......product..
	0x0040 9c 11 08 22 00 00 00 00 05 72 65 73 75 6c 74 00  ...".....result.
	0x0050 04 9c 11 10 22 00 00 00 00 00                    ....".....
	rel 21+8 t=1 "".recursiveChannel+0
	rel 29+8 t=1 "".recursiveChannel+158
	rel 50+4 t=28 go.info.int+0
	rel 68+4 t=28 go.info.int+0
	rel 85+4 t=28 go.info.chan int+0
go.range."".recursiveChannel SDWARFRANGE size=0
go.string."recursive: %d\n" SRODATA dupok size=14
	0x0000 72 65 63 75 72 73 69 76 65 3a 20 25 64 0a        recursive: %d.
go.info."".main SDWARFINFO size=46
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 04 72 65 73  .............res
	0x0020 75 6c 74 00 04 9c 11 58 22 00 00 00 00 00        ult....X".....
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+270
	rel 41+4 t=28 go.info.chan int+0
go.range."".main SDWARFRANGE size=0
go.info."".init SDWARFINFO size=29
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 00           .............
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+96
go.range."".init SDWARFRANGE size=0
"".initdone· SNOPTRBSS size=1
"".recursiveChannel·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".recursiveChannel+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.03 SRODATA dupok size=1
	0x0000 03                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*chan int- SRODATA dupok size=12
	0x0000 00 00 09 2a 63 68 61 6e 20 69 6e 74              ...*chan int
type.*chan int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ed 7b ed 3b 00 08 08 36 00 00 00 00 00 00 00 00  .{.;...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan int-+0
	rel 48+8 t=1 type.chan int+0
type.chan int SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 91 55 cb 71 02 08 08 32 00 00 00 00 00 00 00 00  .U.q...2........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan int-+0
	rel 44+4 t=6 type.*chan int+0
	rel 48+8 t=1 type.int+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.log. SRODATA dupok size=6
	0x0000 00 00 03 6c 6f 67                                ...log
gclocals·cadea2e49003779a155f5f8fb1f0fe78 SRODATA dupok size=11
	0x0000 03 00 00 00 03 00 00 00 00 00 00                 ...........
gclocals·dc990f0ce5208716c4d33a75f4225a61 SRODATA dupok size=11
	0x0000 03 00 00 00 04 00 00 00 00 0c 03                 ...........
gclocals·a35fa7d7e7129fc330c152d6236a3e5c SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 04 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·3a5f6b21d2b25517271e2f12e2835dc4 SRODATA dupok size=11
	0x0000 03 00 00 00 03 00 00 00 00 01 06                 ...........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
