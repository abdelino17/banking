package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Abdel", "Lome", "20701", "2020-03-12", "1"},
		{"2", "Aziz", "Shimla", "191009", "2021-03-12", "1"},
		{"3", "Fare", "Evry", "91000", "2022-03-12", "1"},
	}
	return CustomerRepositoryStub{customers}
}
