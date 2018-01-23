package bottles

import "github.com/gokit/rpkit/examples/bottles/cellar"

//@rp
type BottleCellar interface {
	AddBottle(string) error
	Cellar() cellar.Cellar
}
