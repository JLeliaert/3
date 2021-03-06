package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"sync"
	"unsafe"
)

// CUDA handle for resize kernel
var resize_code cu.Function

// Stores the arguments for resize kernel invocation
type resize_args_t struct {
	arg_dst    unsafe.Pointer
	arg_Dx     int
	arg_Dy     int
	arg_Dz     int
	arg_src    unsafe.Pointer
	arg_Sx     int
	arg_Sy     int
	arg_Sz     int
	arg_layer  int
	arg_scalex int
	arg_scaley int
	argptr     [11]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for resize kernel invocation
var resize_args resize_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	resize_args.argptr[0] = unsafe.Pointer(&resize_args.arg_dst)
	resize_args.argptr[1] = unsafe.Pointer(&resize_args.arg_Dx)
	resize_args.argptr[2] = unsafe.Pointer(&resize_args.arg_Dy)
	resize_args.argptr[3] = unsafe.Pointer(&resize_args.arg_Dz)
	resize_args.argptr[4] = unsafe.Pointer(&resize_args.arg_src)
	resize_args.argptr[5] = unsafe.Pointer(&resize_args.arg_Sx)
	resize_args.argptr[6] = unsafe.Pointer(&resize_args.arg_Sy)
	resize_args.argptr[7] = unsafe.Pointer(&resize_args.arg_Sz)
	resize_args.argptr[8] = unsafe.Pointer(&resize_args.arg_layer)
	resize_args.argptr[9] = unsafe.Pointer(&resize_args.arg_scalex)
	resize_args.argptr[10] = unsafe.Pointer(&resize_args.arg_scaley)
}

// Wrapper for resize CUDA kernel, asynchronous.
func k_resize_async(dst unsafe.Pointer, Dx int, Dy int, Dz int, src unsafe.Pointer, Sx int, Sy int, Sz int, layer int, scalex int, scaley int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	resize_args.Lock()
	defer resize_args.Unlock()

	if resize_code == 0 {
		resize_code = fatbinLoad(resize_map, "resize")
	}

	resize_args.arg_dst = dst
	resize_args.arg_Dx = Dx
	resize_args.arg_Dy = Dy
	resize_args.arg_Dz = Dz
	resize_args.arg_src = src
	resize_args.arg_Sx = Sx
	resize_args.arg_Sy = Sy
	resize_args.arg_Sz = Sz
	resize_args.arg_layer = layer
	resize_args.arg_scalex = scalex
	resize_args.arg_scaley = scaley

	args := resize_args.argptr[:]
	cu.LaunchKernel(resize_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for resize kernel.
var resize_map = map[int]string{0: "",
	20: resize_ptx_20,
	30: resize_ptx_30,
	35: resize_ptx_35}

// resize PTX code for various compute capabilities.
const (
	resize_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<61>;
	.reg .f32 	%f<21>;
	.reg .s64 	%rd<12>;


	ld.param.u64 	%rd4, [resize_param_0];
	ld.param.u32 	%r11, [resize_param_1];
	ld.param.u32 	%r17, [resize_param_2];
	ld.param.u64 	%rd5, [resize_param_4];
	ld.param.u32 	%r12, [resize_param_5];
	ld.param.u32 	%r13, [resize_param_6];
	ld.param.u32 	%r14, [resize_param_8];
	ld.param.u32 	%r15, [resize_param_9];
	ld.param.u32 	%r16, [resize_param_10];
	.loc 1 8 1
	mov.u32 	%r18, %ctaid.x;
	mov.u32 	%r19, %ntid.x;
	mov.u32 	%r20, %tid.x;
	mad.lo.s32 	%r21, %r19, %r18, %r20;
	.loc 1 9 1
	mov.u32 	%r22, %ntid.y;
	mov.u32 	%r23, %ctaid.y;
	mov.u32 	%r24, %tid.y;
	mad.lo.s32 	%r25, %r22, %r23, %r24;
	.loc 1 11 1
	setp.lt.s32	%p1, %r21, %r11;
	setp.lt.s32	%p2, %r25, %r17;
	and.pred  	%p3, %p1, %p2;
	.loc 1 11 1
	@!%p3 bra 	BB0_11;
	bra.uni 	BB0_1;

BB0_1:
	.loc 1 16 1
	setp.gt.s32	%p4, %r16, 0;
	@%p4 bra 	BB0_3;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	bra.uni 	BB0_10;

BB0_3:
	.loc 1 16 1
	mul.lo.s32 	%r35, %r16, %r25;
	mad.lo.s32 	%r36, %r14, %r13, %r35;
	mul.lo.s32 	%r37, %r12, %r36;
	mad.lo.s32 	%r57, %r15, %r21, %r37;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r58, 0;
	cvta.to.global.u64 	%rd6, %rd5;

BB0_4:
	mul.wide.s32 	%rd7, %r57, 4;
	add.s64 	%rd11, %rd6, %rd7;
	.loc 1 19 1
	setp.lt.s32	%p5, %r15, 1;
	@%p5 bra 	BB0_9;

	.loc 1 16 1
	mul.lo.s32 	%r59, %r15, %r21;
	mov.u32 	%r60, 0;

BB0_6:
	.loc 1 17 1
	mad.lo.s32 	%r47, %r25, %r16, %r58;
	setp.lt.s32	%p6, %r47, %r13;
	.loc 1 22 1
	setp.lt.s32	%p7, %r59, %r12;
	and.pred  	%p8, %p6, %p7;
	.loc 1 22 1
	@!%p8 bra 	BB0_8;
	bra.uni 	BB0_7;

BB0_7:
	.loc 1 23 1
	ld.global.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	.loc 1 24 1
	add.f32 	%f19, %f19, 0f3F800000;

BB0_8:
	add.s64 	%rd11, %rd11, 4;
	.loc 1 19 1
	add.s32 	%r59, %r59, 1;
	.loc 1 19 22
	add.s32 	%r60, %r60, 1;
	.loc 1 19 1
	setp.lt.s32	%p9, %r60, %r15;
	@%p9 bra 	BB0_6;

BB0_9:
	.loc 1 16 22
	add.s32 	%r58, %r58, 1;
	.loc 1 16 1
	setp.lt.s32	%p10, %r58, %r16;
	add.s32 	%r57, %r57, %r12;
	@%p10 bra 	BB0_4;

BB0_10:
	.loc 1 28 106
	mad.lo.s32 	%r56, %r25, %r11, %r21;
	cvta.to.global.u64 	%rd8, %rd4;
	mul.wide.s32 	%rd9, %r56, 4;
	add.s64 	%rd10, %rd8, %rd9;
	.loc 2 3608 3
	div.rn.f32 	%f18, %f20, %f19;
	.loc 1 28 106
	st.global.f32 	[%rd10], %f18;

BB0_11:
	.loc 1 30 2
	ret;
}


`
	resize_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<61>;
	.reg .f32 	%f<21>;
	.reg .s64 	%rd<12>;


	ld.param.u64 	%rd4, [resize_param_0];
	ld.param.u32 	%r11, [resize_param_1];
	ld.param.u32 	%r17, [resize_param_2];
	ld.param.u64 	%rd5, [resize_param_4];
	ld.param.u32 	%r12, [resize_param_5];
	ld.param.u32 	%r13, [resize_param_6];
	ld.param.u32 	%r14, [resize_param_8];
	ld.param.u32 	%r15, [resize_param_9];
	ld.param.u32 	%r16, [resize_param_10];
	.loc 1 8 1
	mov.u32 	%r18, %ctaid.x;
	mov.u32 	%r19, %ntid.x;
	mov.u32 	%r20, %tid.x;
	mad.lo.s32 	%r21, %r19, %r18, %r20;
	.loc 1 9 1
	mov.u32 	%r22, %ntid.y;
	mov.u32 	%r23, %ctaid.y;
	mov.u32 	%r24, %tid.y;
	mad.lo.s32 	%r25, %r22, %r23, %r24;
	.loc 1 11 1
	setp.lt.s32	%p1, %r21, %r11;
	setp.lt.s32	%p2, %r25, %r17;
	and.pred  	%p3, %p1, %p2;
	.loc 1 11 1
	@!%p3 bra 	BB0_11;
	bra.uni 	BB0_1;

BB0_1:
	.loc 1 16 1
	setp.gt.s32	%p4, %r16, 0;
	@%p4 bra 	BB0_3;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	bra.uni 	BB0_10;

BB0_3:
	.loc 1 16 1
	mul.lo.s32 	%r35, %r16, %r25;
	mad.lo.s32 	%r36, %r14, %r13, %r35;
	mul.lo.s32 	%r37, %r12, %r36;
	mad.lo.s32 	%r57, %r15, %r21, %r37;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r58, 0;
	cvta.to.global.u64 	%rd6, %rd5;

BB0_4:
	mul.wide.s32 	%rd7, %r57, 4;
	add.s64 	%rd11, %rd6, %rd7;
	.loc 1 19 1
	setp.lt.s32	%p5, %r15, 1;
	@%p5 bra 	BB0_9;

	.loc 1 16 1
	mul.lo.s32 	%r59, %r15, %r21;
	mov.u32 	%r60, 0;

BB0_6:
	.loc 1 17 1
	mad.lo.s32 	%r47, %r25, %r16, %r58;
	setp.lt.s32	%p6, %r47, %r13;
	.loc 1 22 1
	setp.lt.s32	%p7, %r59, %r12;
	and.pred  	%p8, %p6, %p7;
	.loc 1 22 1
	@!%p8 bra 	BB0_8;
	bra.uni 	BB0_7;

BB0_7:
	.loc 1 23 1
	ld.global.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	.loc 1 24 1
	add.f32 	%f19, %f19, 0f3F800000;

BB0_8:
	add.s64 	%rd11, %rd11, 4;
	.loc 1 19 1
	add.s32 	%r59, %r59, 1;
	.loc 1 19 22
	add.s32 	%r60, %r60, 1;
	.loc 1 19 1
	setp.lt.s32	%p9, %r60, %r15;
	@%p9 bra 	BB0_6;

BB0_9:
	.loc 1 16 22
	add.s32 	%r58, %r58, 1;
	.loc 1 16 1
	setp.lt.s32	%p10, %r58, %r16;
	add.s32 	%r57, %r57, %r12;
	@%p10 bra 	BB0_4;

BB0_10:
	.loc 1 28 106
	mad.lo.s32 	%r56, %r25, %r11, %r21;
	cvta.to.global.u64 	%rd8, %rd4;
	mul.wide.s32 	%rd9, %r56, 4;
	add.s64 	%rd10, %rd8, %rd9;
	.loc 2 3608 3
	div.rn.f32 	%f18, %f20, %f19;
	.loc 1 28 106
	st.global.f32 	[%rd10], %f18;

BB0_11:
	.loc 1 30 2
	ret;
}


`
	resize_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<61>;
	.reg .f32 	%f<21>;
	.reg .s64 	%rd<12>;


	ld.param.u64 	%rd4, [resize_param_0];
	ld.param.u32 	%r11, [resize_param_1];
	ld.param.u32 	%r17, [resize_param_2];
	ld.param.u64 	%rd5, [resize_param_4];
	ld.param.u32 	%r12, [resize_param_5];
	ld.param.u32 	%r13, [resize_param_6];
	ld.param.u32 	%r14, [resize_param_8];
	ld.param.u32 	%r15, [resize_param_9];
	ld.param.u32 	%r16, [resize_param_10];
	.loc 1 8 1
	mov.u32 	%r18, %ctaid.x;
	mov.u32 	%r19, %ntid.x;
	mov.u32 	%r20, %tid.x;
	mad.lo.s32 	%r21, %r19, %r18, %r20;
	.loc 1 9 1
	mov.u32 	%r22, %ntid.y;
	mov.u32 	%r23, %ctaid.y;
	mov.u32 	%r24, %tid.y;
	mad.lo.s32 	%r25, %r22, %r23, %r24;
	.loc 1 11 1
	setp.lt.s32	%p1, %r21, %r11;
	setp.lt.s32	%p2, %r25, %r17;
	and.pred  	%p3, %p1, %p2;
	.loc 1 11 1
	@!%p3 bra 	BB2_11;
	bra.uni 	BB2_1;

BB2_1:
	.loc 1 16 1
	setp.gt.s32	%p4, %r16, 0;
	@%p4 bra 	BB2_3;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	bra.uni 	BB2_10;

BB2_3:
	.loc 1 16 1
	mul.lo.s32 	%r35, %r16, %r25;
	mad.lo.s32 	%r36, %r14, %r13, %r35;
	mul.lo.s32 	%r37, %r12, %r36;
	mad.lo.s32 	%r57, %r15, %r21, %r37;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r58, 0;
	cvta.to.global.u64 	%rd6, %rd5;

BB2_4:
	mul.wide.s32 	%rd7, %r57, 4;
	add.s64 	%rd11, %rd6, %rd7;
	.loc 1 19 1
	setp.lt.s32	%p5, %r15, 1;
	@%p5 bra 	BB2_9;

	.loc 1 16 1
	mul.lo.s32 	%r59, %r15, %r21;
	mov.u32 	%r60, 0;

BB2_6:
	.loc 1 17 1
	mad.lo.s32 	%r47, %r25, %r16, %r58;
	setp.lt.s32	%p6, %r47, %r13;
	.loc 1 22 1
	setp.lt.s32	%p7, %r59, %r12;
	and.pred  	%p8, %p6, %p7;
	.loc 1 22 1
	@!%p8 bra 	BB2_8;
	bra.uni 	BB2_7;

BB2_7:
	.loc 1 23 1
	ld.global.nc.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	.loc 1 24 1
	add.f32 	%f19, %f19, 0f3F800000;

BB2_8:
	add.s64 	%rd11, %rd11, 4;
	.loc 1 19 1
	add.s32 	%r59, %r59, 1;
	.loc 1 19 22
	add.s32 	%r60, %r60, 1;
	.loc 1 19 1
	setp.lt.s32	%p9, %r60, %r15;
	@%p9 bra 	BB2_6;

BB2_9:
	.loc 1 16 22
	add.s32 	%r58, %r58, 1;
	.loc 1 16 1
	setp.lt.s32	%p10, %r58, %r16;
	add.s32 	%r57, %r57, %r12;
	@%p10 bra 	BB2_4;

BB2_10:
	.loc 1 28 106
	mad.lo.s32 	%r56, %r25, %r11, %r21;
	cvta.to.global.u64 	%rd8, %rd4;
	mul.wide.s32 	%rd9, %r56, 4;
	add.s64 	%rd10, %rd8, %rd9;
	.loc 3 3608 3
	div.rn.f32 	%f18, %f20, %f19;
	.loc 1 28 106
	st.global.f32 	[%rd10], %f18;

BB2_11:
	.loc 1 30 2
	ret;
}


`
)
