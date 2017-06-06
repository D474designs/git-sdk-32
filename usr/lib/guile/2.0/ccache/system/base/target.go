GOOF----LE-4-2.0�      ] i 4     h�      ] g  guile�	 �	g  define-module*�	 �	 �	g  system�	g  base�	g  target�		 �	
g  filenameS�	f  system/base/target.scm�	g  importsS�	g  rnrs�	g  bytevectors�	 �	 �	g  ice-9�	g  regex�	 �	 �	 �	g  exportsS�	g  target-type�	g  with-target�	g  
target-cpu�	g  target-vendor�	g  	target-os�	g  target-endianness�	g  target-word-size�	 �	g  set-current-module�	  �	! �	"g  foreign�	#" �	$g  sizeof�	%#$ �	&#$ �	'g  *�	(g  %native-word-size�	)g  
make-fluid�	*g  
%host-type�	+g  %target-type�	,g  native-endianness�	-g  %target-endianness�	.g  %target-word-size�	/g  string?�	0g  string-split�	1g  length�	2g  or-map�	3g  string-null?�	4g  error�	5f  invalid target�	6g  validate-target�	7g  triplet-cpu�	8g  cpu-endianness�	9g  triplet-pointer-size�	:g  string=?�	;g  string-match�	<f  
^i[0-9]86$�	=g  little�	>g  member�	?f  x86_64�	@f  ia64�	Af  	powerpcle�	Bf  powerpc64le�	Cf  mipsel�	Df  mips64el�	Ef  nios2�	Ff  sh3�	Gf  sh4�	Hf  alpha�	I?@ABCDEFGH 
�	Jf  sparc�	Kf  sparc64�	Lf  powerpc�	Mf  	powerpc64�	Nf  spu�	Of  mips�	Pf  mips64�	Qf  m68k�	Rf  s390x�	SJKLMNOPQR 	�	Tg  big�	Uf  ^arm.*el�	Vf  ^arm.*eb�	Wg  string-prefix?�	Xf  arm�	Yf  ^aarch64.*be�	Zf  aarch64�	[f  unknown CPU endianness�	\f  ^mips64.*-gnuabi64�	]f  ^mips64�	^f  ^x86_64-.*-gnux32�	_f  64$�	`f  64_?[lbe][lbe]$�	aJLOCEQFG �	bRH �	cf  ^arm.*�	df  unknown CPU word size�	eg  
triplet-os�	fg  	substring�	gg  string-index�	hg  triplet-vendor�C 5      hH  *  ]4	
5 4! >  "  G   4&'5(R4)i*i5+R4)i4,i5 5-R4)i(i5.R/012345   hX   �   ]4 5$  54 -545	�$  "  	45"  $   6C       �       g  target
		Q g  parts		> g  t		"	;  g  filenamef  system/base/target.scm�
	-
��		.	��		.	��		/	��		/	
��		0	��	"	0	��	"	0	��	0	1	��	C	.	
��	G	.	��	K	2	��	O	2	�� 		Q  g  nameg  validate-target� C6R67+-.89 hH   �   ]
4 >  "  G  4 5 454 5Y4>   ZCZF       �       g  target
		A g  thunk		A g  cpu			A  g  filenamef  system/base/target.scm�
	4
��		5	��		6	��		6	��	&	8	&��	-	9	%��	6	:	�� 		A	  g  nameg  with-target� CR:7*,;<=>ISTUVWXYZ4[  h�   -  ]4 455$  6 4 5$  C4 	5$  C4 
5$  C4 5$  C4 5$  C4 5$  C4 5$  C4 5$  C 6%      g  cpu
	 �  g  filenamef  system/base/target.scm�
	<
��		>	��		>	��		>	��		>	��		?	��		@	��		@	��	!	@	��	%	@	��	'	A	��	)	B	��	/	B	��	1	B	��	5	@	��	7	D	��	9	E	��	?	E	��	A	E	��	E	@	��	G	G	��	I	H	��	M	H	��	Q	H	��	U	@	��	W	I	��	Y	J	��	]	J	��	a	J	��	e	@	��	g	K	��	i	L	��	m	L	��	q	L	��	u	@	��	w	M	��	y	N	��	}	N	�� �	N	�� �	@	�� �	O	�� �	P	�� �	P	�� �	P	�� �	@	�� �	Q	�� �	S	�� �	S	�� 1	 �  g  nameg  cpu-endianness�g  documentationf  Return the endianness for CPU.� C8R7;<\]^_`>abc4d:*e(        h�   }  ]	4 5"  �45$  	C4 5$  	C45$  	C4 5$  	C45$  	C45$  	C4	
5$  	C4	5$  	C45$  	C64455$  44 5455$  C"��8"��4     u      g  triplet
	 � g  cpu		 �  g  filenamef  system/base/target.scm�
	U
��		W	��			W	��		\	��		\	��		\	��		X	��	 	c	��	$	c	��	(	c	��	,	X	��	0	d	��	4	d	��	8	d	��	<	X	��	@	f	��	D	f	��	H	f	��	L	X	��	P	h	��	T	h	��	X	h	��	\	X	��	`	i	��	d	i	��	h	i	��	l	X	��	p	j	��	v	j	��	x	j	��	|	X	�� �	k	�� �	k	�� �	k	�� �	X	�� �	l	�� �	l	�� �	l	�� �	X	�� �	m	�� �	m	�� �	X	�� �	X	�� �	X	�� �	X	�� �	X	�� �	Y	�� �	Y	�� �	Y	/�� �	Y	�� �	X	�� 4	 �  g  nameg  triplet-pointer-size�g  documentationf  1Return the size of pointers in bytes for TRIPLET.� C9Rfg        h   �   ] 
4 -56      x       g  t
		  g  filenamef  system/base/target.scm�
	o
��		p	��		p	�� 		  g  nameg  triplet-cpu� C7Rgf     h(   �   ]	4 -5� 4 -56       �       g  t
		! g  start		!  g  filenamef  system/base/target.scm�
	r
��		s	��		s	��		s	��		t	��	!	t	�� 		!  g  nameg  triplet-vendor� ChRgf        h    �   ]	4 -4 -5�5� 6�       g  t
		  g  start		   g  filenamef  system/base/target.scm�
	v
��		w	��	
	w	+��		w	'��		w	��		w	��		w	��	 	x	�� 			   g  nameg  
triplet-os� CeR+      h   �   ] [C  �       g  filenamef  system/base/target.scm�
	{
�� 		
  g  nameg  target-type�g  documentationf  <Return the GNU configuration triplet of the target platform.� CR7   h   �   ] 45 6     �       g  filenamef  system/base/target.scm�
	
��	 �	��	 �	�� 		
  g  nameg  
target-cpu�g  documentationf  +Return the CPU name of the target platform.� CRh   h   �   ] 45 6     �       g  filenamef  system/base/target.scm�
 �
��	 �	��	 �	�� 		
  g  nameg  target-vendor�g  documentationf  .Return the vendor name of the target platform.� CRe    h   �   ] 45 6     �       g  filenamef  system/base/target.scm�
 �
��	 �	��	 �	�� 		
  g  nameg  	target-os�g  documentationf  8Return the operating system name of the target platform.� CR-        h   �   ] [C  �       g  filenamef  system/base/target.scm�
 �
�� 		
  g  nameg  target-endianness�g  documentationf  4Return the endianness object of the target platform.� CR.      h   �   ] [C  �       g  filenamef  system/base/target.scm�
 �
�� 		
  g  nameg  target-word-size�g  documentationf  7Return the word size, in bytes, of the target platform.� CRC"      g  m
		,  g  filenamef  system/base/target.scm�		
��	-	'	��	/	'	��	1	'	��	3	'	��	6	$
��	7	)	��	D	)
��	E	*	��	J	*	'��	R	*	��	U	*
��	V	+	��	c	+
���	-
��	4
��	<
��	�	U
��
S	o
��@	r
��4	v
��	{
���	
��� �
��� �
��t �
��F �
�� 	H
   C6 