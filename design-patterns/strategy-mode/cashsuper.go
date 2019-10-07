package strategy_mode

type cashSuper interface {
	AcceptCash(money float64) float64
}

//正常收费
type cashNormal struct {
}

func (c cashNormal) AcceptCash(money float64) float64 {
	return money
}

//打折收费
type cashRebate struct {
	moneyRebate float64
}

func (c cashRebate) SetMoneyRebate(moneyRebate float64) {
	c.moneyRebate = moneyRebate
}

func (c cashRebate) AcceptCash(money float64) float64 {
	return money * c.moneyRebate
}

type cashReturn struct {
	moneyCondition float64
	moneyReturn    float64
}

func (c cashReturn) cashReturn(moneyCondition, moneyReturn float64) {

	c.moneyCondition = moneyCondition
	c.moneyReturn = moneyReturn

}

func (c cashReturn) AcceptCash(money float64) float64 {
	var result float64
	if money >= c.moneyCondition {
		result = money - c.moneyReturn
	}

	return result
}