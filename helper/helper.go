package helper

// var Name string = "Yohanes"
// var Address string = "Kendal"

type Student struct {
	Name    string
	Address string
	Age     uint
	Gender  string
	Married bool
}

func SayHello(name string) string {
	return "Hello " + name
}

func SayGoodBye(name string) string {
	return "Goodbye " + name
}

func (student *Student) Welcome() string {
	return "Hello, my name is " + student.Name + ". I'm from " + student.Address
}

func CreateNewStudent(name string, address string) (string, string) {
	student := Student{
		Name:    name,
		Address: address,
		Age:     26,
		Gender:  "Male",
		Married: false,
	}

	student2 := &student
	student2.Name = "Alta"

	return student.Welcome() + " - ", student2.Welcome()
}

func UpdateStudent(name string, address string) string {
	student := &Student{Name: name, Address: address}
	return student.Welcome()
}
