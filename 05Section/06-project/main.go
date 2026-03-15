package main

import "fmt"

type Payable interface {
	CalculatePay() float64
	fmt.Stringer
}

type SalaryEmployee struct {
	Name         string
	AnnualSalary float64
}

func (emp SalaryEmployee) CalculatePay() float64 {
	return emp.AnnualSalary / 12
}

func (emp SalaryEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual salary: %.2f)", emp.Name, emp.AnnualSalary)
}

type HourlyEmployee struct {
	Name       string
	HourlyRate float64
	HourlyWork float64
}

func (emp HourlyEmployee) CalculatePay() float64 {
	return emp.HourlyRate * emp.HourlyWork
}

func (emp HourlyEmployee) String() string {
	return fmt.Sprintf("Freelancer: %s (Hourly work: %.2f, Hourly rate: %.2f)",
		emp.Name, emp.HourlyRate, emp.HourlyWork)
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64 // monthly
	CommissionRate float64 // 0.05 for 5%
	SalesAmount    float64
}

func (emp CommissionEmployee) CalculatePay() float64 {
	return emp.BaseSalary + emp.SalesAmount*emp.CommissionRate
}

func (emp CommissionEmployee) String() string {
	return fmt.Sprintf("CommissionEmployee: %s (base salary: %.2f, Sales: %.2f, commission rate: %.2f)",
		emp.Name, emp.BaseSalary, emp.SalesAmount, 100*emp.CommissionRate)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf("- Processing: %s\n", employee)
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("------ Processing Payroll ------")
	totalAmount := 0.0
	for _, emp := range employees {
		PrintEmployeeSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("Monthly Payable: %.2f\n", pay)
		totalAmount += pay
	}
	fmt.Printf("Total amount to pay %f\n", totalAmount)
	fmt.Println("------ ********** ------")
}
func main() {

	Bharat := SalaryEmployee{
		Name:         "Bharat Bardiya",
		AnnualSalary: 12000,
	}

	Jhon := HourlyEmployee{
		Name:       "Jhon carner",
		HourlyRate: 12,
		HourlyWork: 100,
	}

	Alysa := CommissionEmployee{
		Name:           "Alysa Liu",
		BaseSalary:     1200,
		CommissionRate: 0.15,
		SalesAmount:    2000,
	}

	Employees := []Payable{Bharat, Jhon, Alysa}
	ProcessPayroll(Employees)
}
