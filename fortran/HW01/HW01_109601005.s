	.file	"HW1_109601005.f90"
	.section	.rodata
.LC0:
	.string	"HW1_109601005.f90"
	.text
	.type	MAIN__, @function
MAIN__:
.LFB0:
	.cfi_startproc
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	subq	$496, %rsp
	movl	$1, -4(%rbp)
	movq	$.LC0, -488(%rbp)
	movl	$8, -480(%rbp)
	movl	$128, -496(%rbp)
	movl	$6, -492(%rbp)
	leaq	-496(%rbp), %rax
	movq	%rax, %rdi
	call	_gfortran_st_write
	leaq	-4(%rbp), %rcx
	leaq	-496(%rbp), %rax
	movl	$4, %edx
	movq	%rcx, %rsi
	movq	%rax, %rdi
	call	_gfortran_transfer_integer_write
	leaq	-496(%rbp), %rax
	movq	%rax, %rdi
	call	_gfortran_st_write_done
	movl	$0, %esi
	movl	$0, %edi
	call	_gfortran_stop_string
	.cfi_endproc
.LFE0:
	.size	MAIN__, .-MAIN__
	.globl	main
	.type	main, @function
main:
.LFB1:
	.cfi_startproc
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	subq	$16, %rsp
	movl	%edi, -4(%rbp)
	movq	%rsi, -16(%rbp)
	movq	-16(%rbp), %rdx
	movl	-4(%rbp), %eax
	movq	%rdx, %rsi
	movl	%eax, %edi
	call	_gfortran_set_args
	movl	$options.1.1884, %esi
	movl	$7, %edi
	call	_gfortran_set_options
	call	MAIN__
	movl	$0, %eax
	leave
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE1:
	.size	main, .-main
	.section	.rodata
	.align 16
	.type	options.1.1884, @object
	.size	options.1.1884, 28
options.1.1884:
	.long	68
	.long	1023
	.long	0
	.long	0
	.long	1
	.long	1
	.long	0
	.ident	"GCC: (GNU) 4.8.5 20150623 (Red Hat 4.8.5-44)"
	.section	.note.GNU-stack,"",@progbits
