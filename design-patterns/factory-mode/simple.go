package factory_mode

type Operation interface {
	GetNumberA() float64
	GetNumberB() float64
	SetNumberA(float64)
	SetNumberB(float64)
	GetResult() float64
}

//base
type operation struct {
	numberA, numberB, Result float64
}

func (o *operation) GetNumberA() float64 {
	return o.numberA
}

func (o *operation) SetNumberA(newNumberA float64) {

	o.numberA = newNumberA
	return
}

func (o *operation) GetNumberB() float64 {
	return o.numberB
}

func (o *operation) SetNumberB(newNumberB float64) {

	o.numberB = newNumberB
	return
}

func (o *operation) GetResult() float64 {
	return o.Result
}

//add
type operationAdd struct {
	operation
}

func (oa *operationAdd) GetResult() float64 {
	return oa.numberA + oa.numberB
}

//sub
type operationSub struct {
	operation
}

func (os *operationSub) GetResult() float64 {
	return os.numberA - os.numberA
}

//mul
type operationMul struct {
	operation
}

func (om *operationMul) GetResult() float64 {
	return om.numberA * om.numberB
}

//Div
type operationDiv struct {
	operation
}

func (od *operationDiv) GetResult() float64 {
	if od.numberB == 0 {
		panic("除数不能为零")
	}
	return od.numberA / od.numberB
}
