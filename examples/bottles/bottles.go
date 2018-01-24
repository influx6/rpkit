package bottles

import (
	"context"

	"github.com/gokit/rpkit/examples/bottles/cellar"
)

//@rp
type BottleCellar interface {
	AddBottle(string) error
	Cellar() *cellar.Cellar
}

type IBox interface {
	Store(int) error
}

//@rp
type IBoxCellar interface {
	StoreBottle(context.Context, IBox) error
	GetBox(context.Context, string) (IBox, error)
}
