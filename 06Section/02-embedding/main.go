package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (addr Address) FullAddress() string {
	if addr.Street == "" || addr.City == "" {
		return "No address provided"
	}
	return fmt.Sprintf("%s, %s, %s %s", addr.Street, addr.City, addr.State, addr.ZipCode)
}

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)

	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.Street)

	fmt.Printf("Email (promoted): %s\n", c.Email)
	fmt.Printf("Business Type: %s\n", c.BusinessType)
}

type CompanyWithOwnEmail struct {
	Name string
	Address
	ContactInfo
	Email string // This email field will shadow the one in ContactInfo
}

func (c CompanyWithOwnEmail) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)

	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.Street)

	fmt.Printf("Email (promoted): %s\n", c.Email)
}
func main() {

	fmt.Println(" ----- Struct Embedded ----- ")
	comp := Company{
		Name: "Innovate Solutions Inc.",
		Address: Address{
			Street:  "789 Innovation Drive",
			City:    "Techville",
			State:   "TS",
			ZipCode: "12345",
		},
		ContactInfo: ContactInfo{
			Email: "contact@innovate.com",
			Phone: "555-0100",
		},
		BusinessType: "Technology",
	}

	comp.GetProfile()

	fmt.Printf("\nDirect access to comp.City: %s\n", comp.City) // Promoted from Address
	fmt.Printf("Direct access to comp.Phone: %s\n", comp.Phone) // Promoted from ContactInfo

	fmt.Printf("Embedded Address struct: %+v\n", comp.Address)
	fmt.Printf("Embedded ContactInfo struct: %+v\n", comp.ContactInfo)

	compWithEmail := CompanyWithOwnEmail{
		Name: "MNIT Jaipur",
		Address: Address{
			Street:  "Malaviya Nagar JLN Marg",
			City:    "Jaipur",
			State:   "Raj",
			ZipCode: "01010",
		},
		ContactInfo: ContactInfo{
			Email: "staff@mnit.ac.in",
			Phone: "555-0100",
		},
		Email: "prof@mnit.ac.in",
	}

	compWithEmail.GetProfile()

}
