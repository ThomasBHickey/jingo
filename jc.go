package jingo

import (
	//"fmt"
	)

// from jc.h
const (
	CX = iota //0			 /* other                                   */
	CS        //1            /* space or tab                            */
	CA        //2            /* letter                                  */
	CN        //3            /* N (capital N)                           */
	CB        //4            /* B (capital B)                           */
	C9        //5            /* digit or sign (underscore)              */
	CD        //6            /* dot                                     */
	CC        //7            /* colon                                   */
	CQ        //8            /* quote                                   */
)
const (
	CESC1 = '.' /*  46 056 2e     1st escape char          */
	CESC2 = ':' /*  58 072 3a     2nd escape char*/
)
const (
	C0       = 0000 /*   0 000 00                              */
	C1       = 0001 /*   1 001 01                              */
	COFF     = 0004 /*   4 004 04     ctrl d                   */
	CTAB     = 0011 /*   9 011 09     tab                      */
	CLF      = 0012 /*  10 012 0a     line feed                */
	CCR      = 0015 /*  13 015 0d     carriage return          */
	CBW0000  = 0020 /*  16 020 10     bitwise fns              */
	CBW0001  = 0021 /*  17 021 11                              */
	CBW0010  = 0022 /*  18 022 12                              */
	CBW0011  = 0023 /*  19 023 13                              */
	CBW0100  = 0024 /*  20 024 14                              */
	CBW0101  = 0025 /*  21 025 15                              */
	CBW0110  = 0026 /*  22 026 16                              */
	CBW0111  = 0027 /*  23 027 17                              */
	CBW1000  = 0030 /*  24 030 18                              */
	CBW1001  = 0031 /*  25 031 19                              */
	CBW1010  = 0032 /*  26 032 1a                              */
	CBW1011  = 0033 /*  27 033 1b                              */
	CBW1100  = 0034 /*  28 034 1c                              */
	CBW1101  = 0035 /*  29 035 1d                              */
	CBW1110  = 0036 /*  30 036 1e                              */
	CBW1111  = 0037 /*  31 037 1f                              */
	CBANG    = '!'  /*  33 041 21                              */
	CQQ      = 0042 /*  34 042 22     double quote             */
	CPOUND   = '#'  /*  35 043 23                              */
	CDOLLAR  = '$'  /*  36 044 24                              */
	CDIV     = '%'  /*  37 045 25                              */
	CAMP     = '&'  /*  38 046 26                              */
	CQUOTE   = 0047 /*  39 047 27     single quote             */
	CLPAR    = '('  /*  40 050 28                              */
	CRPAR    = ')'  /*  41 051 29                              */
	CSTAR    = '*'  /*  42 052 2a                              */
	CPLUS    = '+'  /*  43 053 2b                              */
	CCOMMA   = ','  /*  44 054 2c                              */
	CMINUS   = '-'  /*  45 055 2d                              */
	CDOT     = '.'  /*  46 056 2e                              */
	CSLASH   = '/'  /*  47 057 2f                              */
	CNOUN    = '0'  /*  48 060 30                              */
	CHOOK    = '2'  /*  50 062 32                              */
	CFORK    = '3'  /*  51 063 33                              */
	CADVF    = '4'  /*  52 064 34     bonded conjunction       */
	CCOLON   = ':'  /*  58 072 3a                              */
	CSEMICO  = ';'  /*  59 073 3b                              */
	CRAZE    = ';'  /*  59 073 3b                              */
	CBOX     = '<'  /*  60 074 3c                              */
	CLT      = '<'  /*  60 074 3c                              */
	CEQ      = '='  /*  61 075 3d                              */
	COPE     = '>'  /*  62 076 3e                              */
	CGT      = '>'  /*  62 076 3e                              */
	CQUERY   = '?'  /*  63 077 3f                              */
	CAT      = '@'  /*  64 100 40                              */
	CLEFT    = '['  /*  91 133 5b                              */
	CBSLASH  = 0134 /*  92 134 5c \   backslash                */
	CRIGHT   = ']'  /*  93 135 5d                              */
	CEXP     = '^'  /*  94 136 5e                              */
	CSIGN    = '_'  /*  95 137 5f     minus sign               */
	CINF     = '_'  /*  95 137 5f     infinity                 */
	CGRAVE   = '`'  /*  96 140 60                              */
	CLBRACE  = '{'  /* 123 173 7b                              */
	CFROM    = '{'  /* 123 173 7b                              */
	CSTILE   = '|'  /* 124 174 7c                              */
	CRBRACE  = '}'  /* 125 175 7d                              */
	CAMEND   = '}'  /* 125 175 7d                              */
	CTILDE   = '~'  /* 126 176 7e                              */
	CASGN    = 0200 /* 128 200 80 =.                           */
	CGASGN   = 0201 /* 129 201 81 =:                           */
	CFLOOR   = 0202 /* 130 202 82 <.                           */
	CMIN     = 0202 /* 130 202 82 <.                           */
	CLE      = 0203 /* 131 203 83 <:                           */
	CCEIL    = 0204 /* 132 204 84 >.                           */
	CMAX     = 0204 /* 132 204 84 >.                           */
	CGE      = 0205 /* 133 205 85 >:                           */
	CUSDOT   = 0206 /* 134 206 86 _.                           */
	CPLUSDOT = 0210 /* 136 210 88 +.                           */
	CPLUSCO  = 0211 /* 137 211 89 +:                           */
	CSTARDOT = 0212 /* 138 212 8a *.                           */
	CSTARCO  = 0213 /* 139 213 8b *:                           */
	CNOT     = 0214 /* 140 214 8c -.                           */
	CLESS    = 0214 /* 140 214 8c -.                           */
	CHALVE   = 0215 /* 141 215 8d -:                           */
	CMATCH   = 0215 /* 141 215 8d -:                           */
	CDOMINO  = 0216 /* 142 216 8e %.                           */
	CSQRT    = 0217 /* 143 217 8f %:                           */
	CROOT    = 0217 /* 143 217 8f %:                           */
	CLOG     = 0220 /* 144 220 90 ^.                           */
	CPOWOP   = 0221 /* 145 221 91 ^:                           */
	CSPARSE  = 0222 /* 146 222 92 $.                           */
	CSELF    = 0223 /* 147 223 93 $:                           */
	CNUB     = 0224 /* 148 224 94 ~.                           */
	CNE      = 0225 /* 149 225 95 ~:                           */
	CREV     = 0226 /* 150 226 96 |.                           */
	CROT     = 0226 /* 150 226 96 |.                           */
	CCANT    = 0227 /* 151 227 97 |:                           */
	CEVEN    = 0230 /* 152 230 98 ..                           */
	CODD     = 0231 /* 153 231 99 .:                           */
	COBVERSE = 0232 /* 154 232 9a :.                           */
	CADVERSE = 0233 /* 155 233 9b ::                           */
	CCOMDOT  = 0234 /* 156 234 9c ,.                           */
	CLAMIN   = 0235 /* 157 235 9d ,:                           */
	CCUT     = 0236 /* 158 236 9e ;.                           */
	CWORDS   = 0237 /* 159 237 9f ;:                           */
	CBASE    = 0240 /* 160 240 a0 #.                           */
	CABASE   = 0241 /* 161 241 a1 #:                           */
	CFIT     = 0242 /* 162 242 a2 !.                           */
	CIBEAM   = 0243 /* 163 243 a3 !:                           */
	CSLDOT   = 0244 /* 164 244 a4 /.                           */
	CGRADE   = 0245 /* 165 245 a5 /:                           */
	CBSDOT   = 0246 /* 166 246 a6 \.                           */
	CDGRADE  = 0247 /* 167 247 a7 \:                           */
	CLEV     = 0250 /* 168 250 a8 [.                           */
	CCAP     = 0251 /* 169 251 a9 [:                           */
	CDEX     = 0252 /* 170 252 aa ].                           */
	CIDA     = 0253 /* 171 253 ab ]:                           */
	CHEAD    = 0254 /* 172 254 ac {.                           */
	CTAKE    = 0254 /* 172 254 ac {.                           */
	CTAIL    = 0255 /* 173 255 ad {:                           */
	CBEHEAD  = 0256 /* 174 256 ae }.                           */
	CDROP    = 0256 /* 174 256 ae }.                           */
	CCTAIL   = 0257 /* 175 257 af }:                           */
	CEXEC    = 0260 /* 176 260 b0 ".                           */
	CTHORN   = 0261 /* 177 261 b1 ":                           */
	CGRDOT   = 0262 /* 178 262 b2 `.                           */
	CGRCO    = 0263 /* 179 263 b3 `:                           */
	CATDOT   = 0264 /* 180 264 b4 @.                           */
	CATCO    = 0265 /* 181 265 b5 @:                           */
	CUNDER   = 0266 /* 182 266 b6 &.                           */
	CAMPCO   = 0267 /* 183 267 b7 &:                           */
	CQRYDOT  = 0270 /* 184 270 b8 ?.                           */
	CQRYCO   = 0271 /* 185 271 b9 ?:                           */

	CALP    = 0272 /* 186 272 ba a.                           */
	CATOMIC = 0273 /* 187 273 bb A.                           */
	CACE    = 0274 /* 188 274 bc a:                           */
	CBDOT   = 0275 /* 189 275 bd b.                           */
	CCDOT   = 0276 /* 190 276 be c.                           */
	CCYCLE  = 0300 /* 192 300 c0 C.                           */
	CDDOT   = 0301 /* 193 301 c1 d.                           */
	CDCAP   = 0302 /* 194 302 c2 D.                           */
	CDCAPCO = 0303 /* 195 303 c3 D:                           */
	CEPS    = 0304 /* 196 304 c4 e.                           */
	CEBAR   = 0305 /* 197 305 c5 E.                           */
	CFIX    = 0306 /* 198 306 c6 f.                           */
	CFCAPCO = 0307 /* 199 307 c7 F:                           */
	CHGEOM  = 0310 /* 200 310 c8 H.                           */
	CIOTA   = 0311 /* 201 311 c9 i.                           */
	CICO    = 0312 /* 202 312 ca i:                           */
	CICAP   = 0313 /* 203 313 cb I.                           */
	CICAPCO = 0314 /* 204 314 cc I:                           */
	CJDOT   = 0315 /* 205 315 cd j.                           */
	CLDOT   = 0316 /* 206 316 ce L.                           */
	CLCAPCO = 0317 /* 207 317 cf L:                           */
	CMDOT   = 0320 /* 208 320 d0 m.                           */
	CMCAP   = 0321 /* 209 321 d1 M.                           */
	CNDOT   = 0322 /* 210 322 d2 n.                           */
	CCIRCLE = 0323 /* 211 323 d3 o.                           */
	CPOLY   = 0324 /* 212 324 d4 p.                           */
	CPCO    = 0325 /* 213 325 d5 p:                           */
	CQCAPCO = 0326 /* 214 326 d6 Q:                           */
	CQCO    = 0327 /* 215 327 d7 q:                           */
	CRDOT   = 0330 /* 216 330 d8 r.                           */
	CSCO    = 0331 /* 217 331 d9 s:                           */
	CSCAPCO = 0332 /* 218 332 da S:                           */
	CTDOT   = 0333 /* 219 333 db t.                           */
	CTCO    = 0334 /* 220 334 dc t:                           */
	CTCAP   = 0335 /* 221 335 dd T.                           */
	CUDOT   = 0336 /* 222 336 de u.                           */
	CUCO    = 0337 /* 223 337 df u:                           */
	CVDOT   = 0340 /* 224 340 e0 v.                           */
	CXDOT   = 0341 /* 225 341 e1 x.                           */
	CXCO    = 0342 /* 226 342 e2 x:                           */
	CYDOT   = 0343 /* 227 343 e3 y.                           */

	CFCONS  = 0350 /* 232 350 e8 0: 1: 2: etc.                */
	CAMIP   = 0351 /* 233 351 e9 }   amend in place           */
	CCASEV  = 0352 /* 234 352 ea }   case in place            */
	CFETCH  = 0353 /* 235 353 eb {::                          */
	CMAP    = 0354 /* 236 354 ec {::                          */
	CEMEND  = 0355 /* 237 355 ed }::                          */
	CUNDCO  = 0356 /* 238 356 ee &.:                          */
	CPDERIV = 0357 /* 239 357 ef p..                          */
	CAPIP   = 0360 /* 240 360 f0 ,   append in place          */

	CFF = 0377 /* 255 377 ff                              */
)

// func init() {
// 	fmt.Println("in jc.go")
// }