## Footgun

Shoot yourself in the foot with very fast, very unsafe Malloc/Free replacements for cgo.

### Why?

Cgo overhead is a little higher than many are comfortable with (at the time of this writing, a simple call tends to run between 4-6x an equivalent JNI call).  Where they really get you, though, is the data marshalling.  Each individual call to malloc or free is another cgo call with a 30-50ns overhead.  Each assignment to a c pointer is a 50ns cgopointercheck.  There's also a live bug at this time where cgopointercheck will accidentally clone slices it tries to check.  

This library does its part to correct these mistakes by making larger, scarier mistakes.  It uses a completely forbidden and poorly-thought-out assembler trampoline to skip the cgo call process.  For the record, if these methods take too long to run, your program will deadlock or explode, so you may consider using C.malloc for big ticket items.  This does more or less eliminate the cgo overhead for Malloc & Free calls (and methods that use them, such as CBytes and CString).

### Should I Use This For X?

Nobody should use this for anything, under any circumstances.