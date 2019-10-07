package factory_mode

//client

func NewOperation(operate string) Operation {
	var oper Operation
	switch operate {
	case "+":
		oper = new(operationAdd)
		break
	case "-":
		oper = new(operationSub)
		break
	case "*":
		oper = new(operationMul)
		break
	case "/":
		oper = new(operationDiv)
		break
	}
	return oper
}

func Run(operate string) float64 {
	o := NewOperation(operate)
	o.SetNumberA(1)
	o.SetNumberB(2)
	return o.GetResult()
}
