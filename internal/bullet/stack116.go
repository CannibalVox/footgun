// +build go1.16 go1.17

package bullet

type stack struct {
	lo uintptr
	hi uintptr
}

type gobuf struct {
	sp   uintptr
	pc   uintptr
	g    uintptr
	ctxt uintptr
	ret  uintptr
	lr   uintptr
	bp   uintptr
}

type m struct {
	g0 *g
}

type g struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr

	_panic uintptr
	_defer uintptr
	m      *m
	sched  gobuf
}
