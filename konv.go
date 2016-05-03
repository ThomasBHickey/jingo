// Copyright 2016 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file

package jingo

func cvt(jt *J, t AType, w A) (y A) {
	oq := jt.rank
	*jt.rank = 0
	b := ccvt(jt, t, w, &y)
	jt.Log.Println("cvt b:", b)
	jt.rank = oq
	return
}

func xcvt(jt *J, param XMode, a A)(y A) {
	jt.Log.Println("xcvt is not implemented")
	return a
}

// static B jtccvt(J jt,I t,A w,A*y){A d;I n,r,*s,wt,*wv,*yv;
func ccvt(jt *J, t AType, w A, y *A)bool {
//  RZ(w);
	if w.Type==NoAType {return false}
//  r=AR(w); s=AS(w);
	//s := w.Shape
	//r:= len(s)
	if (t&SPARSE)!=0 || (w.Type&SPARSE)!=0{
		jt.Log.Println("SPARSE arrays not supported")
		return false
	}

//  n=AN(w); wt=AT(w); wv=AV(w);
//  if(t==wt){RZ(*y=ca(w)); R 1;}
	if (t==w.Type){}
//  // else if(n&&t&JCHAR){ASSERT(HOMO(t,wt),EVDOMAIN); RZ(*y=uco1(w)); R 1;}
//  GA(*y,t,n,r,s); yv=AV(*y); 
//  if(t&CMPX)fillv(t,n,(C*)yv); 
//  if(!n)R 1;
//  switch(CVCASE(t,wt)){
//   case CVCASE(LIT, C2T ): R C1fromC2(w,yv);
//   case CVCASE(C2T, LIT ): R C2fromC1(w,yv);
//   case CVCASE(BIT ,B01 ): R cvt2bit(w,yv);
//   case CVCASE(INT ,B01 ): {I*x=    yv;B*v=(B*)wv; DO(n,*x++   =*v++;);} R 1;
//   case CVCASE(XNUM,B01 ): R XfromB(w,yv);
//   case CVCASE(RAT ,B01 ): GA(d,XNUM,n,r,s); R XfromB(w,AV(d))&&QfromX(d,yv);
//   case CVCASE(FL  ,B01 ): {D*x=(D*)yv;B*v=(B*)wv; DO(n,*x++   =*v++;);} R 1;
//   case CVCASE(CMPX,B01 ): {Z*x=(Z*)yv;B*v=(B*)wv; DO(n,x++->re=*v++;);} R 1;
//   case CVCASE(BIT ,INT ): R cvt2bit(w,yv);
//   case CVCASE(B01 ,INT ): R BfromI(w,yv);
//   case CVCASE(XNUM,INT ): R XfromI(w,yv);
//   case CVCASE(RAT ,INT ): GA(d,XNUM,n,r,s); R XfromI(w,AV(d))&&QfromX(d,yv);
//   case CVCASE(FL  ,INT ): {D*x=(D*)yv;I*v=    wv; DO(n,*x++   =(D)*v++;);} R 1;
//   case CVCASE(CMPX,INT ): {Z*x=(Z*)yv;I*v=    wv; DO(n,x++->re=(D)*v++;);} R 1;
//   case CVCASE(BIT ,FL  ): R cvt2bit(w,yv);
//   case CVCASE(B01 ,FL  ): R BfromD(w,yv);
//   case CVCASE(INT ,FL  ): R IfromD(w,yv);
//   case CVCASE(XNUM,FL  ): R XfromD(w,yv);
//   case CVCASE(RAT ,FL  ): R QfromD(w,yv);
//   case CVCASE(CMPX,FL  ): {Z*x=(Z*)yv;D t,*v=(D*)wv; DO(n, t=*v++; x++->re=t||_isnan(t)?t:0.0;);} R 1;  /* -0 to 0*/
//   case CVCASE(BIT ,CMPX): GA(d,FL,n,r,s); RZ(DfromZ(w,AV(d))); R cvt2bit(d,yv);
//   case CVCASE(B01 ,CMPX): GA(d,FL,n,r,s); RZ(DfromZ(w,AV(d))); R BfromD(d,yv);
//   case CVCASE(INT ,CMPX): GA(d,FL,n,r,s); RZ(DfromZ(w,AV(d))); R IfromD(d,yv);
//   case CVCASE(XNUM,CMPX): GA(d,FL,n,r,s); RZ(DfromZ(w,AV(d))); R XfromD(d,yv);
//   case CVCASE(RAT ,CMPX): GA(d,FL,n,r,s); RZ(DfromZ(w,AV(d))); R QfromD(d,yv);
//   case CVCASE(FL  ,CMPX): R DfromZ(w,yv);
//   case CVCASE(B01 ,XNUM): R BfromX(w,yv);
//   case CVCASE(INT ,XNUM): R IfromX(w,yv);
//   case CVCASE(RAT ,XNUM): R QfromX(w,yv);
//   case CVCASE(FL  ,XNUM): R DfromX(w,yv);
//   case CVCASE(CMPX,XNUM): GA(d,FL,  n,r,s); RZ(DfromX(w,AV(d))); R ccvt(t,d,y);
//   case CVCASE(B01 ,RAT ): GA(d,XNUM,n,r,s); RZ(XfromQ(w,AV(d))); R BfromX(d,yv);
//   case CVCASE(INT ,RAT ): GA(d,XNUM,n,r,s); RZ(XfromQ(w,AV(d))); R IfromX(d,yv);
//   case CVCASE(XNUM,RAT ): R XfromQ(w,yv);
//   case CVCASE(FL  ,RAT ): R DfromQ(w,yv);
//   case CVCASE(CMPX,RAT ): GA(d,FL,  n,r,s); RZ(DfromQ(w,AV(d))); R ccvt(t,d,y);
//   default:                ASSERT(0,EVDOMAIN);
// }}
	return true
}