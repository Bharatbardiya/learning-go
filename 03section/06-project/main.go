package main


import (
	"fmt"
)

type Contact struct {
	ID int
	name string
	phone string
	email string
}

var ContactList []Contact;
var ContactIndexByName map[string]int;
var nextID int;

func init(){
	ContactList = make([]Contact,0);
	ContactIndexByName = make(map[string]int);
	nextID = 0;
}

func AddContact (name, email, phone string) {
	_, ok := ContactIndexByName[name];

	if ok {
		fmt.Printf("%s is already present\n");
		return;
	}

	newContact := Contact{
		ID: nextID,
		name : name,
		phone : phone,
		email : email,
	}
	ContactList = append(ContactList, newContact);
	ContactIndexByName[name] = nextID;
	nextID++;
	fmt.Printf("%s added successfully\n", name);
}

func FindContact(name string) *Contact{
	contactID, ok := ContactIndexByName[name];
	if !ok {
		return nil;
	}
	return &ContactList[contactID];
}

func ListContact(){

	for _, Contact := range ContactList{
		fmt.Printf("ID: %d, name: %s, email: %s, phone: %s\n",
		 Contact.ID, Contact.name, Contact.email, Contact.phone);
	}
}


func main(){

	fmt.Println(" --------- Contact project --------- ");

	AddContact("Bharat", "bharat@google.com", "123-23434");
	AddContact("Jhon", "Jhon@amazon.com", "123-67823");
	AddContact("Krishna", "kriishn@amazon.com", "123-78982");

	var bharatPtr *Contact = FindContact("Bharat");
	if bharatPtr!=nil {
		fmt.Printf("Bharat's Detail : %d, %s[%s, %s]\n",
		 bharatPtr.ID, bharatPtr.name, bharatPtr.email, bharatPtr.phone);
	}

}