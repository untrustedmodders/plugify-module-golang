package cross_call_master

// Generated from cross_call_master

type Example = int32

const (
	Example_First Example = 1
	Example_Second Example = 2
	Example_Third Example = 3
	Example_Forth Example = 4
)


// noCopy prevents copying via go vet
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// ownership indicates whether the instance owns the underlying handle
type ownership bool

const (
	Owned    ownership = true
	Borrowed ownership = false
)

