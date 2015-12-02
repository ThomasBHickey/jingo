package jingo

import (
//"fmt"
)

// from jc.h
type CBType int

const (
	CX CBType = iota //0			 /* other                                   */
	CS               //1            /* space or tab                            */
	CA               //2            /* letter                                  */
	CN               //3            /* N (capital N)                           */
	CB               //4            /* B (capital B)                           */
	C9               //5            /* digit or sign (underscore)              */
	CD               //6            /* dot                                     */
	CC               //7            /* colon                                   */
	CQ               //8            /* quote                                   */
)

type ESCType int

const (
	CESC1 ESCType = '.' /*  46 056 2e     1st escape char          */
	CESC2         = ':' /*  58 072 3a     2nd escape char*/
)

type IDType int

const (
	C0       IDType = 0000 /*   0 000 00                              */
	C1       IDType = 0001 /*   1 001 01                              */
	COFF     IDType = 0004 /*   4 004 04     ctrl d                   */
	CTAB     IDType = 0011 /*   9 011 09     tab                      */
	CLF      IDType = 0012 /*  10 012 0a     line feed                */
	CCR      IDType = 0015 /*  13 015 0d     carriage return          */
	CBW0000  IDType = 0020 /*  16 020 10     bitwise fns              */
	CBW0001  IDType = 0021 /*  17 021 11                              */
	CBW0010  IDType = 0022 /*  18 022 12                              */
	CBW0011  IDType = 0023 /*  19 023 13                              */
	CBW0100  IDType = 0024 /*  20 024 14                              */
	CBW0101  IDType = 0025 /*  21 025 15                              */
	CBW0110  IDType = 0026 /*  22 026 16                              */
	CBW0111  IDType = 0027 /*  23 027 17                              */
	CBW1000  IDType = 0030 /*  24 030 18                              */
	CBW1001  IDType = 0031 /*  25 031 19                              */
	CBW1010  IDType = 0032 /*  26 032 1a                              */
	CBW1011  IDType = 0033 /*  27 033 1b                              */
	CBW1100  IDType = 0034 /*  28 034 1c                              */
	CBW1101  IDType = 0035 /*  29 035 1d                              */
	CBW1110  IDType = 0036 /*  30 036 1e                              */
	CBW1111  IDType = 0037 /*  31 037 1f                              */
	CBANG    IDType = '!'  /*  33 041 21                              */
	CQQ      IDType = 0042 /*  34 042 22     double quote             */
	CPOUND   IDType = '#'  /*  35 043 23                              */
	CDOLLAR  IDType = '$'  /*  36 044 24                              */
	CDIV     IDType = '%'  /*  37 045 25                              */
	CAMP     IDType = '&'  /*  38 046 26                              */
	CQUOTE   IDType = 0047 /*  39 047 27     single quote             */
	CLPAR    IDType = '('  /*  40 050 28                              */
	CRPAR    IDType = ')'  /*  41 051 29                              */
	CSTAR    IDType = '*'  /*  42 052 2a                              */
	CPLUS    IDType = '+'  /*  43 053 2b                              */
	CCOMMA   IDType = ','  /*  44 054 2c                              */
	CMINUS   IDType = '-'  /*  45 055 2d                              */
	CDOT     IDType = '.'  /*  46 056 2e                              */
	CSLASH   IDType = '/'  /*  47 057 2f                              */
	CNOUN    IDType = '0'  /*  48 060 30                              */
	CHOOK    IDType = '2'  /*  50 062 32                              */
	CFORK    IDType = '3'  /*  51 063 33                              */
	CADVF    IDType = '4'  /*  52 064 34     bonded conjunction       */
	CCOLON   IDType = ':'  /*  58 072 3a                              */
	CSEMICO  IDType = ';'  /*  59 073 3b                              */
	CRAZE    IDType = ';'  /*  59 073 3b                              */
	CBOX     IDType = '<'  /*  60 074 3c                              */
	CLT      IDType = '<'  /*  60 074 3c                              */
	CEQ      IDType = '='  /*  61 075 3d                              */
	COPE     IDType = '>'  /*  62 076 3e                              */
	CGT      IDType = '>'  /*  62 076 3e                              */
	CQUERY   IDType = '?'  /*  63 077 3f                              */
	CAT      IDType = '@'  /*  64 100 40                              */
	CLEFT    IDType = '['  /*  91 133 5b                              */
	CBSLASH  IDType = 0134 /*  92 134 5c \   backslash                */
	CRIGHT   IDType = ']'  /*  93 135 5d                              */
	CEXP     IDType = '^'  /*  94 136 5e                              */
	CSIGN    IDType = '_'  /*  95 137 5f     minus sign               */
	CINF     IDType = '_'  /*  95 137 5f     infinity                 */
	CGRAVE   IDType = '`'  /*  96 140 60                              */
	CLBRACE  IDType = '{'  /* 123 173 7b                              */
	CFROM    IDType = '{'  /* 123 173 7b                              */
	CSTILE   IDType = '|'  /* 124 174 7c                              */
	CRBRACE  IDType = '}'  /* 125 175 7d                              */
	CAMEND   IDType = '}'  /* 125 175 7d                              */
	CTILDE   IDType = '~'  /* 126 176 7e                              */
	CASGN    IDType = 0200 /* 128 200 80 =.                           */
	CGASGN   IDType = 0201 /* 129 201 81 =:                           */
	CFLOOR   IDType = 0202 /* 130 202 82 <.                           */
	CMIN     IDType = 0202 /* 130 202 82 <.                           */
	CLE      IDType = 0203 /* 131 203 83 <:                           */
	CCEIL    IDType = 0204 /* 132 204 84 >.                           */
	CMAX     IDType = 0204 /* 132 204 84 >.                           */
	CGE      IDType = 0205 /* 133 205 85 >:                           */
	CUSDOT   IDType = 0206 /* 134 206 86 _.                           */
	CPLUSDOT IDType = 0210 /* 136 210 88 +.                           */
	CPLUSCO  IDType = 0211 /* 137 211 89 +:                           */
	CSTARDOT IDType = 0212 /* 138 212 8a *.                           */
	CSTARCO  IDType = 0213 /* 139 213 8b *:                           */
	CNOT     IDType = 0214 /* 140 214 8c -.                           */
	CLESS    IDType = 0214 /* 140 214 8c -.                           */
	CHALVE   IDType = 0215 /* 141 215 8d -:                           */
	CMATCH   IDType = 0215 /* 141 215 8d -:                           */
	CDOMINO  IDType = 0216 /* 142 216 8e %.                           */
	CSQRT    IDType = 0217 /* 143 217 8f %:                           */
	CROOT    IDType = 0217 /* 143 217 8f %:                           */
	CLOG     IDType = 0220 /* 144 220 90 ^.                           */
	CPOWOP   IDType = 0221 /* 145 221 91 ^:                           */
	CSPARSE  IDType = 0222 /* 146 222 92 $.                           */
	CSELF    IDType = 0223 /* 147 223 93 $:                           */
	CNUB     IDType = 0224 /* 148 224 94 ~.                           */
	CNE      IDType = 0225 /* 149 225 95 ~:                           */
	CREV     IDType = 0226 /* 150 226 96 |.                           */
	CROT     IDType = 0226 /* 150 226 96 |.                           */
	CCANT    IDType = 0227 /* 151 227 97 |:                           */
	CEVEN    IDType = 0230 /* 152 230 98 ..                           */
	CODD     IDType = 0231 /* 153 231 99 .:                           */
	COBVERSE IDType = 0232 /* 154 232 9a :.                           */
	CADVERSE IDType = 0233 /* 155 233 9b ::                           */
	CCOMDOT  IDType = 0234 /* 156 234 9c ,.                           */
	CLAMIN   IDType = 0235 /* 157 235 9d ,:                           */
	CCUT     IDType = 0236 /* 158 236 9e ;.                           */
	CWORDS   IDType = 0237 /* 159 237 9f ;:                           */
	CBASE    IDType = 0240 /* 160 240 a0 #.                           */
	CABASE   IDType = 0241 /* 161 241 a1 #:                           */
	CFIT     IDType = 0242 /* 162 242 a2 !.                           */
	CIBEAM   IDType = 0243 /* 163 243 a3 !:                           */
	CSLDOT   IDType = 0244 /* 164 244 a4 /.                           */
	CGRADE   IDType = 0245 /* 165 245 a5 /:                           */
	CBSDOT   IDType = 0246 /* 166 246 a6 \.                           */
	CDGRADE  IDType = 0247 /* 167 247 a7 \:                           */
	CLEV     IDType = 0250 /* 168 250 a8 [.                           */
	CCAP     IDType = 0251 /* 169 251 a9 [:                           */
	CDEX     IDType = 0252 /* 170 252 aa ].                           */
	CIDA     IDType = 0253 /* 171 253 ab ]:                           */
	CHEAD    IDType = 0254 /* 172 254 ac {.                           */
	CTAKE    IDType = 0254 /* 172 254 ac {.                           */
	CTAIL    IDType = 0255 /* 173 255 ad {:                           */
	CBEHEAD  IDType = 0256 /* 174 256 ae }.                           */
	CDROP    IDType = 0256 /* 174 256 ae }.                           */
	CCTAIL   IDType = 0257 /* 175 257 af }:                           */
	CEXEC    IDType = 0260 /* 176 260 b0 ".                           */
	CTHORN   IDType = 0261 /* 177 261 b1 ":                           */
	CGRDOT   IDType = 0262 /* 178 262 b2 `.                           */
	CGRCO    IDType = 0263 /* 179 263 b3 `:                           */
	CATDOT   IDType = 0264 /* 180 264 b4 @.                           */
	CATCO    IDType = 0265 /* 181 265 b5 @:                           */
	CUNDER   IDType = 0266 /* 182 266 b6 &.                           */
	CAMPCO   IDType = 0267 /* 183 267 b7 &:                           */
	CQRYDOT  IDType = 0270 /* 184 270 b8 ?.                           */
	CQRYCO   IDType = 0271 /* 185 271 b9 ?:                           */

	CALP    IDType = 0272 /* 186 272 ba a.                           */
	CATOMIC IDType = 0273 /* 187 273 bb A.                           */
	CACE    IDType = 0274 /* 188 274 bc a:                           */
	CBDOT   IDType = 0275 /* 189 275 bd b.                           */
	CCDOT   IDType = 0276 /* 190 276 be c.                           */
	CCYCLE  IDType = 0300 /* 192 300 c0 C.                           */
	CDDOT   IDType = 0301 /* 193 301 c1 d.                           */
	CDCAP   IDType = 0302 /* 194 302 c2 D.                           */
	CDCAPCO IDType = 0303 /* 195 303 c3 D:                           */
	CEPS    IDType = 0304 /* 196 304 c4 e.                           */
	CEBAR   IDType = 0305 /* 197 305 c5 E.                           */
	CFIX    IDType = 0306 /* 198 306 c6 f.                           */
	CFCAPCO IDType = 0307 /* 199 307 c7 F:                           */
	CHGEOM  IDType = 0310 /* 200 310 c8 H.                           */
	CIOTA   IDType = 0311 /* 201 311 c9 i.                           */
	CICO    IDType = 0312 /* 202 312 ca i:                           */
	CICAP   IDType = 0313 /* 203 313 cb I.                           */
	CICAPCO IDType = 0314 /* 204 314 cc I:                           */
	CJDOT   IDType = 0315 /* 205 315 cd j.                           */
	CLDOT   IDType = 0316 /* 206 316 ce L.                           */
	CLCAPCO IDType = 0317 /* 207 317 cf L:                           */
	CMDOT   IDType = 0320 /* 208 320 d0 m.                           */
	CMCAP   IDType = 0321 /* 209 321 d1 M.                           */
	CNDOT   IDType = 0322 /* 210 322 d2 n.                           */
	CCIRCLE IDType = 0323 /* 211 323 d3 o.                           */
	CPOLY   IDType = 0324 /* 212 324 d4 p.                           */
	CPCO    IDType = 0325 /* 213 325 d5 p:                           */
	CQCAPCO IDType = 0326 /* 214 326 d6 Q:                           */
	CQCO    IDType = 0327 /* 215 327 d7 q:                           */
	CRDOT   IDType = 0330 /* 216 330 d8 r.                           */
	CSCO    IDType = 0331 /* 217 331 d9 s:                           */
	CSCAPCO IDType = 0332 /* 218 332 da S:                           */
	CTDOT   IDType = 0333 /* 219 333 db t.                           */
	CTCO    IDType = 0334 /* 220 334 dc t:                           */
	CTCAP   IDType = 0335 /* 221 335 dd T.                           */
	CUDOT   IDType = 0336 /* 222 336 de u.                           */
	CUCO    IDType = 0337 /* 223 337 df u:                           */
	CVDOT   IDType = 0340 /* 224 340 e0 v.                           */
	CXDOT   IDType = 0341 /* 225 341 e1 x.                           */
	CXCO    IDType = 0342 /* 226 342 e2 x:                           */
	CYDOT   IDType = 0343 /* 227 343 e3 y.                           */

	CFCONS  IDType = 0350 /* 232 350 e8 0: 1: 2: etc.                */
	CAMIP   IDType = 0351 /* 233 351 e9 }   amend in place           */
	CCASEV  IDType = 0352 /* 234 352 ea }   case in place            */
	CFETCH  IDType = 0353 /* 235 353 eb {::                          */
	CMAP    IDType = 0354 /* 236 354 ec {::                          */
	CEMEND  IDType = 0355 /* 237 355 ed }::                          */
	CUNDCO  IDType = 0356 /* 238 356 ee &.:                          */
	CPDERIV IDType = 0357 /* 239 357 ef p..                          */
	CAPIP   IDType = 0360 /* 240 360 f0 ,   append in place          */

	CFF IDType = 0377 /* 255 377 ff                              */
)

// func init() {
// 	fmt.Println("in jc.go")
// }
