GOOF----LE-4-2.0D(      ] a 4     hë      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  oop¤	g  goops¤	g  dispatch¤		 ¤	
g  filenameS¤	f  oop/goops/dispatch.scm¤	g  importsS¤	 ¤	 ¤	g  util¤	 ¤	 ¤	g  compile¤	 ¤	 ¤	g  system¤	g  base¤	g  target¤	 ¤	 ¤	 ¤	g  exportsS¤	g  memoize-method!¤	 ¤	g  set-current-module¤	 ¤	  ¤	!g  current-module¤	"g  *dispatch-module*¤	#g  gensym¤	$f  a¤	%f  t¤	&g  append¤	'& ¤	(& ¤	)g  rest¤	*g  let¤	+g  map¤	,g  class-of¤	-f  p¤	.g  if¤	/g  and¤	0g  apply¤	1) ¤	2g  assq-ref¤	3g  eq?¤	4f  c¤	5g  
cache-miss¤	6g  cons*¤	7g  list¤	8g  emit-linear-dispatch¤	9g  make-vector¤	:f  g¤	;g  vector-length¤	<g  lambda¤	=g  cdr¤	>g  case-lambda¤	?g  car¤	@g  with-target¤	Ag  
%host-type¤	B ¤	CB ¤	DB ¤	Eg  envS¤	Fg  fromS¤	Gg  scheme¤	Hg  optsS¤	Ig  partial-eval?S¤	Jg  cse?S¤	KIJ ¤	Lg  args¤	ML ¤	Ng  max¤	Og  compute-dispatch-procedure¤	Pg  
timer-init¤	Qg  slot-ref¤	Rg  effective-methods¤	Sg  	slot-set!¤	Tg  	procedure¤	Ug  cache-dispatch¤	Vg  delayed-compile¤	Wg  n-specialized¤	X	, ¤	Y	, ¤	Zg  compute-cmethod¤	[g  memoize-effective-method!¤	\g  compute-applicable-methods¤	]g  %compute-applicable-methods¤	^g  no-applicable-method¤	_g  set-procedure-property!¤	`g  system-procedure¤C 5h!  Â   ]4	
5 4  >  "  G   4!i5 "R#   h8   ù   ]"   
$  C45"ÿÿà "ÿÿÔ      ñ       g  n
		2 g  stem		2 g  n			& g  syms			&  g  filenamef  oop/goops/dispatch.scm
	A			B		
	C	
		C			E			E			E		&	E	
	&	B		)	B		*	B		2	B	 		2	  g  nameg  gen-syms C$%()*+,     h   p   ]   C h       g  t
		 g  a		  g  filenamef  oop/goops/dispatch.scm
	Q			R	 			   C#-./01234567     hX  &  ]e4545" (  .$  45"  45
  	D"  ¶(  L4	
5	£$  45"  
 
	"ÿÿv45$  ! "ÿÿ{4	5 "ÿÿJ	£"ÿÿ4 $  45"   
	"ÿþÌ          g  gf-sym
	T g  nargs	T g  methods		T g  free		T g  rest?		T g  gen-syms		T g  args		T g  types		T g  methods		 $ g  free			 $ g  exp	
	 $ g  free		X g  types		X g  specs		X g  checks		X g  m-sym		e ª g  var	 ´ g  var	 ä  g  filenamef  oop/goops/dispatch.scm
	@
		F			F			F			F			G			G	 		G			F		 	H		&	N		,	P		-	P		=	P		>	Q		M	P		T	P		X	X		^	\	
	_	]		c	]	#	e	]		e	]		j	^		m	_	'	p	_		s	_		x	b	 	c	 	d	 	e	  	b	 ª	^	 «	g	 ²	g	( ´	g	 ´	g	 ¼	h	 Á	j	 Ä	k	 Æ	l	" É	l	) Î	l	" Ñ	l	 Ý	i	 Þ	n	 â	n	' ä	n	 ä	n	 é	o	% ì	o	 ò	p	 õ	q	 ÷	r	$ ú	r	+ ÿ	r	$	r		o		X		Z	(	Z		[	$	X	$	H	*	J	2	K	 4	L	$D	M	$J	J	T	H	 G	T	  g  nameg  emit-linear-dispatch C8R9#:;8        hp   b  ] 4L5$  KL £(    "ÿÿÔ4L >  G   "ÿÿ§L 4L56     Z      g  n
		k g  clauses		k g  free			k g  methods			[ g  clause		@	[ g  free		@	[  g  filenamef  oop/goops/dispatch.scm
 		 		 		 		 		 		 		  		. 		/ 		C 		J 		O 	&	[ 		^ 		e 		k 	 		k	  g  nameg  	emit-rest C<+=>?@ADE"FGHK     h       ]4L 5  L@          g  p
		  g  filenamef  oop/goops/dispatch.scm
 ¸		 ¹		 ¹		 »		 ¼		 ¹		 ¹		 ½	
 			
   C8 
 h     ] 
$  /45 45O 6L £(    "ÿÿ«4	L >  G   "ÿÿ~          g  n
	  g  clauses	  g  free		  g  exp		&	9 g  vals		&	9 g  methods		>  g  clause		i  g  free		i   g  filenamef  oop/goops/dispatch.scm
 		 		
 		 		 		 		 ¡		& 		9 ·		> ¢		> ¢		F £		I ¥		W ¥		X ª		l ¨		s ¬		x ¬	%  ¬	 	 	  g  nameg  emit-req C;L5MN     hH  ç  ]R" -(  ç4545"  Å(  y45O 	O 

	Q 	
Q 
	
45
$  "  
£$  	  "    6	£$  
££¤"ÿÿY
££¤"ÿÿ;"ÿÿ3	£$  4

£5"ÿþî4

£5"ÿþÓ	ÿ	ÿ"ÿþÃ     ß      g  gf
	C g  cache	C g  ls		3 g  nreq		3 g  nrest		3 g  req		 ó g  rest		 ó g  ls		& ë g  gf-sym		3 ¥ g  	emit-rest			K ¥ g  emit-req	
	K ¥ g  t		m  g  n	 ´ Í g  n	 Ò ë  g  filenamef  oop/goops/dispatch.scm
	u
		w			x			z			z			z	(		z			{			{			{	)		{			z		& 		, 		- 		1 		3 		3 		e ¯	 	l ¯		m ¯		~ °	 	 °	  ¯	  ±	  ±	  ²	   ³	 £ ³	" ¥ ®	 ¨ 	 « 	 ¯ 	 ² 	 ´ 	 ´ 	 ½ 	$ Â 	- Ã 	 Ä 	
 Ç 	 Í 	
 Ð 	 Ò 	 Ò 	 Û 	# à 	, á 	 â 	
 å 	 ë 	
 ë 	 ö	|	 ù	|	 ý	x	 	}		}	
	}	1	}	%	}		}					#		+%		'		3		3	w	 E	C	  g  nameg  compute-dispatch-procedure COR	PRPOQRSTU      hP   û   -  1  3 MNM
$  .4L 4L 554L >  "  G   @L  6    ó       g  args
			L g  dispatch	&	D  g  filenamef  oop/goops/dispatch.scm
 Ë		 Ì		 Ì		 Î		 Í		 Ï		 Ð		" Ð	)	$ Ð		& Ï		& Ï		) Ñ	
	/ Ñ		6 Ñ	
	D Ò	
	L Ö	 			L


   C  h      ]	H O C       g  gf
		 g  timer		  g  filenamef  oop/goops/dispatch.scm
 É
	 Ê	 		  g  nameg  delayed-compile CVR h8     ] 
$  "  $  C454L  5C        g  n
		6 g  f		6 g  ls			6 g  t			  g  filenamef  oop/goops/dispatch.scm
 Ù		 Ú		 Ú		 Ú		 Ú		 Û		  Ü		% Ü		' Ü		( Ü		- Ü	&	2 Ü	/	4 Ü		5 Ü	 		6	  g  nameg  	map-until C   h8   è   ] $  #$   &    "ÿÿÚCC (  CCà       g  x
		8 g  y		8  g  filenamef  oop/goops/dispatch.scm
 Ý		 Þ			 Þ		 Þ		 Þ		 ß		 ß	'	 Þ		 à	"	  à	*	( à		2 Þ		5 á	 		8	  g  nameg  equal? CQWY5R   h   å  ]"O O Q Q 4 5$  X44 55"  1(   64£5$  
	£@"ÿÿÏ4 5"ÿÿÀ 6       Ý      g  gf
	  g  args	  g  	map-until		  g  equal?		  g  types		;  g  cache		A	r  g  filenamef  oop/goops/dispatch.scm
 Ø
	 ã		# ã		% ã		) ã		* ä		- ä		3 ä	+	5 ä		; ä		; ä		A å		G æ	
	O ç		P è		U è	%	W è		[ è		_ æ	
	b é	$	e é		i é		l ê		r ê		r å		s å		y å	%	{ å	  å	  ë	 	 	  g  nameg  cache-dispatch CUR  h      ]4 5@        g  gf
		 g  args		  g  filenamef  oop/goops/dispatch.scm
 í
	 î			 î	 			  g  nameg  
cache-miss C5R     h0   ð   ]

$  "   $  C 4L  5C è       g  ls
		/ g  n		/ g  t			  g  filenamef  oop/goops/dispatch.scm
 ñ		 ò		 ò		 ò		 ò		 ó		" ô		# ô		( ô	 	+ ô	)	- ô		. ô	 		/	  g  nameg  first-n C+YZQRSTVW   h   T  ]3O "  ¡(  `4545 4 54 >  "  G  4 4	 5>  "  G  C4 
5$  4455"  "ÿÿ_"  V45 4 54 >  "  G  4 4	 5>  "  G  CQ 
"ÿþõ       L      g  gf
	 g  args	 g  
applicable		 g  n		 ¯ g  ls		 ¯ g  types			t g  cmethod		(	t g  cache		>	t g  len	 ³	 g  rest?	 ³	 g  types	 ³	 g  cmethod	 ¼	 g  cache	 Ó	 g  first-n		  g  filenamef  oop/goops/dispatch.scm
 ð
	 õ		 ö		 ÷		 ÷		  ý		( ý		4 þ		5 ÿ		; ÿ	%	= ÿ		> þ		> ý		A 		G 		N 		W		]		^		i		w ø		} ø		 ø	  ø	  ö	  ù	  ù	'  ù	 ¡ ù	 ¤ û	 § û	 ¯ û	 ³ ü	 ´ ý	 ¼ ý	 É þ	 Ê ÿ	 Ð ÿ	% Ò ÿ	 Ó þ	 Ó ý	 Ö 	 Ü 	 ã 	 ì	 ò	 ó	 þ		 2		  g  nameg  memoize-effective-method! C[R\][^     h8   ß   ]
4 &  "   5$  
 6 6     ×       g  gf
		3 g  args		3 g  
applicable			3  g  filenamef  oop/goops/dispatch.scm


									!		+		3	 			3	  g  nameg  memoize-method! CR_ii`6 º       g  m
		,  g  filenamef  oop/goops/dispatch.scm		
	-	 		6	 
Á	@
	u
 È
< É
P Ø
 í
Ï ð
!

!	)!
 	!
   C6 