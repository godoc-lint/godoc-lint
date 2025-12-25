package decl_func

// godoc
func Foo() {}

type Bar string

// godoc
func (Bar) BarFoo() {}

// godoc
func (*Bar) StarBarFoo() {}

type Baz[T any] struct{}

// godoc
func (Baz[T]) BazFoo() {}

// godoc
func (*Baz[T]) StarBazFoo() {}

type Yolo[T any, X any] struct{}

// godoc
func (Yolo[T, X]) YoloFoo() {}

// godoc
func (*Yolo[T, X]) StarYoloFoo() {}

type Zoo = Bar

// godoc
func (Zoo) ZooFoo() {}

// godoc
func (*Zoo) StarZooFoo() {}
