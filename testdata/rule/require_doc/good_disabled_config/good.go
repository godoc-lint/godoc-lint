package good

// (NG: no godoc)

const FooNG = 0

type TFooNG int

func FFooNG() {}

func (*TFooNG) FooFooNG() {}

func (*TFooNG) fooFooNG() {}

const fooNG = 0

type tFooNG int

func fFooNG() {}

func (*tFooNG) fooFooNG() {}

func (*tFooNG) FooFooNG() {}
