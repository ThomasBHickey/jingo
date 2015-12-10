// http://www.jsoftware.com/jwiki/Essays/Incunabulum
// One summer weekend in 1989, Arthur Whitney visited Ken Iverson at Kiln Farm
// and produced—on one page and in one afternoon—an interpreter fragment on the
// AT&T 3B1 computer. I studied this interpreter for about a week for its
// organization and programming style; and on Sunday, August 27, 1989, at about
// four o'clock in the afternoon, wrote the first line of code that became the
// implementation described in this document.

// Arthur's one-page interpreter fragment is as follows:
// Contributed by RogerHui. From An Implementation of J, Appendix A: Incunabulum, 1992-01-27.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// typedef char C;typedef long I;
// typedef struct a{I t,r,d[3],p[2];}*A;
type Vector interface{}
type Array struct {
	Type   AType
	Length int
	Shape  []int
	Data   Vector
}
type AType int

const (
	IntArray AType = iota
	BoxArray
	VerbArray
	ZeroArray
)

// #define P printf
// #define R return
// #define V1(f) A f(w)A w;
// #define V2(f) A f(a,w)A a,w;
type vMonad func(x Array) Array
type vDyad func(x, y Array) Array

// #define DO(n,x) {I i=0,_n=(n);for(;i<_n;++i){x;}}
// I *ma(n){R(I*)malloc(n*4);}mv(d,s,n)I *d,*s;{DO(n,d[i]=s[i]);}
// tr(r,d)I *d;{I z=1;DO(r,z=z*d[i]);R z;}
func size(shape []int) int {
	sz := 1
	for _, sp := range shape {
		sz *= sp
	}
	return sz
}

// A ga(t,r,d)I *d;{A z=(A)ma(5+tr(r,d));z->t=t,z->r=r,mv(z->d,d,r);
//  R z;}
func getArray(typ AType, shape []int) (na Array) {
	na.Type = typ
	na.Length = size(shape)
	na.Data = make([]int, na.Length)
	return
}
func getIntArray(typ AType, i int) Array {
	na := getArray(typ, []int{1})
	na.Data = i
	return na
}

// V1(iota){I n=*w->p;A z=ga(0,1,&n);DO(n,z->p[i]=i);R z;}
func iot(w Array) (z Array) { // iota conflicts with Go
	n := w.Data.([]int)[0]
	d := make([]int, n)
	z = getArray(0, []int{n})
	for i := 0; i < n; i++ {
		d[i] = i
	}
	z.Data = d
	return
}
func plus(a, w Array) (z Array) {
	fmt.Println("Plus not implemented yet")
	return z
}

// V2(plus){I r=w->r,*d=w->d,n=tr(r,d);A z=ga(0,r,d);
//  DO(n,z->p[i]=a->p[i]+w->p[i]);R z;}
// V2(from){I r=w->r-1,*d=w->d+1,n=tr(r,d);
//  A z=ga(w->t,r,d);mv(z->p,w->p+(n**a->p),n);R z;}
// V1(box){A z=ga(1,0,0);*z->p=(I)w;R z;}
// V2(cat){I an=tr(a->r,a->d),wn=tr(w->r,w->d),n=an+wn;
//  A z=ga(w->t,1,&n);mv(z->p,a->p,an);mv(z->p+an,w->p,wn);R z;}
// V2(find){}
// V2(rsh){I r=a->r?*a->d:1,n=tr(r,a->p),wn=tr(w->r,w->d);
//  A z=ga(w->t,r,a->p);mv(z->p,w->p,wn=n>wn?wn:n);
//  if(n-=wn)mv(z->p+wn,z->p,n);R z;}
// V1(sha){A z=ga(0,1,&w->r);mv(z->p,w->d,w->r);R z;}
// V1(id){R w;}V1(size){A z=ga(0,0,0);*z->p=w->r?*w->d:1;R z;}
// pi(i){P("%d ",i);}nl(){P("\n");}
func prInt(i int) {
	fmt.Print(i)
}
func newLine() {
	fmt.Println()
}

// pr(w)A w;{I r=w->r,*d=w->d,n=tr(r,d);DO(r,pi(d[i]));nl();
//  if(w->t)DO(n,P("< ");pr(w->p[i]))else DO(n,pi(w->p[i]));nl();}
func pr(w Array) {
	fmt.Println("Just called 'pr'")
	for _, d := range w.Shape {
		prInt(d)
	}
	newLine()
	switch w.Type {
	case IntArray:
		for i := 0; i < w.Length; i++ {
			prInt(w.Data.([]int)[i])
		}
		newLine()
	case BoxArray:
		for i := 0; i < w.Length; i++ {
			pr(w.Data.([]Array)[i])
		}
	}
}

// C vt[]="+{~<#,";
var vt = "+{~<#,"

// A(*vd[])()={0,plus,from,find,0,rsh,cat},
//  (*vm[])()={0,id,size,iota,box,sha,0};
var vDyads = []vDyad{}
var vMonads = []vMonad{nil, iot, iot}
// I st[26]; qp(a){R  a>='a'&&a<='z';}qv(a){R a<'a';}
var stack [26]int

func isAlpha(a byte) bool { return a >= 'a' && a <= 'z' }
func isOp(a byte) bool    { return a < 'a' }

// A ex(e)I *e;{I a=*e;
//  if(qp(a)){if(e[1]=='=')R st[a-'a']=ex(e+2);a= st[ a-'a'];}
//  R qv(a)?(*vm[a])(ex(e+1)):e[1]?(*vd[e[1]])(a,ex(e+2)):(A)a;}
func execute(e Array) (z Array) {
	fmt.Println("just called execute", e)
	return
}

// noun(c){A z;if(c<'0'||c>'9')R 0;z=ga(0,0,0);*z->p=c-'0';R z;}
func mkNoun(c byte) (z Array, ok bool) {
	if c < '0' || c > '9' {
		return z, false
	}
	z = getArray(0, []int{1})
	z.Data = make([]int, 1)
	z.Data.([]int)[0] = int(c - '0')
	return z, true
}

// verb(c){I i=0;for(;vt[i];)if(vt[i++]==c)R i;R 0;}
func verbPos(ct byte) (pos int, ok bool) {
	pos = strings.IndexByte(vt, ct)
	if pos < 0 {
		return 0, false
	}
	return pos, true
}

// I *wd(s)C *s;{I a,n=strlen(s),*e=ma(n+1);C c;
//  DO(n,e[i]=(a=noun(c=s[i]))?a:(a=verb(c))?a:c);e[n]=0;R e;}
func words(s string) (z Array) {
	fmt.Println("just called words")
	n := len(s)
	e := make([]Array, n+1)
	for i := 0; i < n; i++ {
		c := s[i]
		if a, ok := mkNoun(c); ok {
			fmt.Println("wordsA", a)
			e[i]=a
		}else if a, ok := verbPos(c); ok {
			fmt.Println("wordsB", getIntArray(VerbArray, a))
			e[i] = getIntArray(VerbArray, a)
		}
	}
	e[n] = getIntArray(ZeroArray, 0)
	z.Data = e
	fmt.Println("wordsC", z)
	return
}

// main(){C s[99];while(gets(s))pr(ex(wd(s)));}

func getString(reader *bufio.Reader) string {
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
func main() {
	//fmt.Println("In Main")
	reader := bufio.NewReader(os.Stdin)
	for {
		w := words(getString(reader))
		fmt.Println("words:", w)
		execute(w)
		break
		//if w=="quit" || w=="exit"{break}
		pr(execute(words(getString(reader))))
	}
}
