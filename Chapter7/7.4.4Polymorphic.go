package main

import "fmt"

type Income interface {
	calculate() float64
	source() string
}

//固定账单项目
type FixedBilling struct {
	projectName  string
	biddedAmount float64 //招标总额
}

//定时和材料项目（定时生产项目）
type TimeAndMaterial struct {
	projectName string
	workHours   float64 //工作时长
	hourlyRate  float64 //每小时工资率
}

//固定收入项目
func (fb FixedBilling) calculate() float64 {
	return fb.biddedAmount
}
func (fb FixedBilling) source() string {
	return fb.projectName
}

//定时生产项目
func (tm TimeAndMaterial) calculate() float64 {
	return tm.workHours * tm.hourlyRate
}
func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

type Advertisement struct {
	adName       string
	costPerclick float64
	clickCount   int
}

func (a Advertisement) calculate() float64 {
	return a.costPerclick * float64(a.clickCount)
}
func (a Advertisement) source() string {
	return a.adName
}

//计算和打印总收入的 calculateNetIncome 函数
func calculateNetIncome(ic []Income) {
	netincome := 0.0
	for _, income := range ic {
		fmt.Printf("%T\n", income)
		fmt.Printf("收入来源： %s = $%.2f \n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("公司净收入合计 = $%.2f ", netincome)
}
func main() {
	project1 := FixedBilling{projectName: "项目 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "项目 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "项目 3", workHours: 100, hourlyRate: 40}
	project4 := TimeAndMaterial{projectName: "项目 4", workHours: 250, hourlyRate: 20}
	project5 := Advertisement{adName: "广告 5", costPerclick: 0.1, clickCount: 10000}
	incomeStreams := []Income{project1, project2, project3, project4, project5}
	calculateNetIncome(incomeStreams)
}
