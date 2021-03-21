
gosh
====

Gosh is a library of simple helpers that make using `go` as a scripting language
for quick system automation much more convenient.

It helps with things like these:

```go
import (
	sh "github.com/ljanyst/gosh"
)

func setupCommonLisp() {
	sh.Must(sh.Run("rm", "-rf", sh.Expand("~/Apps/common-lisp")))
	sh.Must(sh.RunS("sbcl --noinform --load setup-files/quicklisp.lisp --load setup-files/install-quicklisp.lisp"))
}
```
