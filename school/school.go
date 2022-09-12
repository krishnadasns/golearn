package school

type School struct {
	Name    string
	Address string
}

func New(name string, address string) School {
	return (School{Name: name, Address: address})
}

func (s *School) GetSchoolName() string {
	return s.Name
}

func (s *School) GetSchoolAddress() string {
	return s.Address
}
