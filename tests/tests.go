package tests

import (
	"errors"

	"github.com/microservices-demo/user/users"
)

var (
	ErrFakeError = errors.New("Fake error")
	TestAddress  = users.Address{
		Street:  "street",
		Number:  "51b",
		Country: "Netherlands",
		City:    "Amsterdam",
		ID:      "000056",
	}
	TestUser users.User
)

func init() {
	u := users.New()
	u.Username = "user"
	u.Password = "737a21044f994ca25906702c27157ce3f7633f76"
	u.Salt = "6c1c6176e8b455ef37da13d953df971c249d0d8e"
	TestUser = u
}

type FakeDB struct {
	PassInit bool
}

func (f *FakeDB) Init() error {
	if f.PassInit {
		return nil
	}
	return ErrFakeError
}
func (f *FakeDB) GetUserByName(name string) (users.User, error) {
	if name == "user" {
		return TestUser, nil
	}
	return users.User{}, ErrFakeError
}
func (f *FakeDB) GetUser(id string) (users.User, error) {
	if id == "realuser" {
		return TestUser, nil
	}
	return users.User{}, ErrFakeError

}

func (f *FakeDB) GetUsers() ([]users.User, error) {
	return make([]users.User, 0), ErrFakeError
}

func (f *FakeDB) CreateUser(u *users.User) error {
	if u.Username == "passtest" {
		return nil
	}
	return ErrFakeError
}

func (f *FakeDB) GetUserAttributes(u *users.User) error {
	u.Addresses = append(u.Addresses, TestAddress)
	return nil
}

func (f *FakeDB) GetCard(id string) (users.Card, error) {
	if id == "realcard" {
		u := users.Card{}
		return u, nil
	}
	return users.Card{}, ErrFakeError
}

func (f *FakeDB) GetCards() ([]users.Card, error) {
	return make([]users.Card, 0), ErrFakeError
}

func (f *FakeDB) CreateCard(c *users.Card, id string) error {
	if c.LongNum == "passtest" {
		return nil
	}
	return ErrFakeError
}

func (f *FakeDB) GetAddress(id string) (users.Address, error) {
	if id == "realaddress" {
		u := users.Address{}
		return u, nil
	}
	return users.Address{}, ErrFakeError
}

func (f *FakeDB) GetAddresses() ([]users.Address, error) {
	return make([]users.Address, 0), ErrFakeError
}

func (f *FakeDB) CreateAddress(a *users.Address, id string) error {
	if a.Street == "passtest" {
		return nil
	}
	return ErrFakeError
}

func (f *FakeDB) Delete(ent string, id string) error {
	return nil
}

type TestServiceStruct struct{}

func (s *TestServiceStruct) Login(username, password string) (users.User, error) {
	return users.User{}, nil
}

func (s *TestServiceStruct) Register(username, password, email string) bool {
	return false
}

func (s *TestServiceStruct) GetUsers(id string) ([]users.User, error) {
	return make([]users.User, 0), nil
}

func (s *TestServiceStruct) PostUser(user users.User) bool {
	return false
}

func (s *TestServiceStruct) GetAddresses(id string) ([]users.Address, error) {
	return make([]users.Address, 0), nil
}

func (s *TestServiceStruct) PostAddress(add users.Address, userid string) bool {
	return false
}

func (s *TestServiceStruct) GetCards(id string) ([]users.Card, error) {
	return make([]users.Card, 0), nil
}

func (s *TestServiceStruct) PostCard(card users.Card, userid string) error {
	return nil
}

func (s *TestServiceStruct) Delete(ent, id string) error {
	return nil
}