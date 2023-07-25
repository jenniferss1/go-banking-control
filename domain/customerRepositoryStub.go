package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Jennifer", "Suzano", "087679", "16/06/2000", "1"},
		{"2", "Aurora", "Suzano", "087679", "19/08/2019", "1"},
		{"3", "Willow", "Suzano", "087679", "01/09/2022", "1"},
	}
	return CustomerRepositoryStub{customers}

}
