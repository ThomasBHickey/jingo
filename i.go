// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
	"fmt"
)

type EventType int

const (
	EVATTN    EventType = iota + 1
	EVBREAK             //     2
	EVDOMAIN            //     3
	EVILNAME            //    4
	EVILNUM             //  5
	EVINDEX             //    6
	EVFACE              //    7
	EVINPRUPT           //    8
	EVLENGTH            //    9
	EVLIMIT             //    10
	EVNONCE             //    11
	EVASSERT            //    12
	EVOPENQ             //    13
	EVRANK              //    14
	EVSPELL             //    16
	EVSTACK             //    17
	EVSTOP              //    18
	EVSYNTAX            //    19
	EVSYSTEM            //    20
	EVVALUE             //    21
	EVWSFULL            //    22
	EVCTRL              //    23
	EVFACCESS           //    24
	EVFNAME             //    25
	EVFNUM              //    26
	EVTIME              //    27
	EVSECURE            //    28
	EVSPARSE            //    29
	EVLOCALE            //    30
	EVRO                //    31
	EVALLOC             //    32
	EVNAN               //    33
	NEVM      = EVNAN   /* number of event codes       */
)

/* The following codes are never displayed to the user         */
const (
	EWOV    EventType = iota + 50 /* integer overflow            */
	EWIMAG                        //  51      /* imaginery  result           */
	EWIRR                         //   52      /* irrational result           */
	EWRAT                         //   53      /* rational   result           */
	EWDIV0                        //   54      /* division by zero            */
	EWTHROW                       //   55      /* throw. executed             */
)

var Event2String = map[EventType]string{}

func init() {
	Event2String[EVALLOC] = "allocation error"
	Event2String[EVASSERT] = "assertion failure"
	Event2String[EVATTN] = "attention interrupt"
	Event2String[EVBREAK] = "break"
	Event2String[EVCTRL] = "control error"
	Event2String[EVDOMAIN] = "domain error"
	Event2String[EVFACCESS] = "file access error"
	Event2String[EVFNAME] = "file name error"
	Event2String[EVFNUM] = "file number error"
	Event2String[EVILNAME] = "ill-formed name"
	Event2String[EVILNUM] = "ill-formed number"
	Event2String[EVINDEX] = "index error"
	Event2String[EVINPRUPT] = "input interrupt"
	Event2String[EVFACE] = "interface error"
	Event2String[EVLENGTH] = "length error"
	Event2String[EVLIMIT] = "limit error"
	Event2String[EVLOCALE] = "locale error"
	Event2String[EVNAN] = "NaN error"
	Event2String[EVNONCE] = "nonce error"
	Event2String[EVSPARSE] = "non-unique sparse elements"
	Event2String[EVOPENQ] = "open quote"
	Event2String[EVWSFULL] = "out of memory"
	Event2String[EVRANK] = "rank error"
	Event2String[EVRO] = "read-only data"
	Event2String[EVSECURE] = "security violation"
	Event2String[EVSPELL] = "spelling error"
	Event2String[EVSTACK] = "stack error"
	Event2String[EVSTOP] = "stop"
	Event2String[EVSYNTAX] = "syntax error"
	Event2String[EVSYSTEM] = "system error"
	Event2String[EVTIME] = "time limit"
	Event2String[EVVALUE] = "value error"
}

func jsignal2(event EventType, pos wp) {
	fmt.Println("jsignal2", Event2String[event], pos)
}
