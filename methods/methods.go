package methods

import (
	"log"
	"time"

	"github.com/google/uuid"
)

//structs methods examples

type Company struct {
	Name    string
	Address string
	Sfera   string
}

func (c *Company) GetInfo() {
	log.Println(c.Address, c.Name, c.Sfera)
}

//base class
type Human struct {
	id        uuid.UUID //incapsulate
	Name      string
	Sex       string
	Race      string
	Country   string
	National  string
	DateBirth time.Time
}

//parentclass methods
func (h *Human) GetDocumentID() uuid.UUID {
	return h.id
}

func (h *Human) SetDocumentID(newID uuid.UUID) {
	h.id = newID
}
func BornHuman(name, sex, race, country, national string) *Human {
	return &Human{
		Name:      name,
		Sex:       sex,
		Race:      race,
		Country:   country,
		National:  national,
		DateBirth: time.Now(),
	}
}

//embedded
type Programmer struct {
	ColorEyes string
	Company
	Human
}

type Builder struct {
	Power int
	*Company
	*Human
}

type Saler struct {
	*Company
	*Human
	Mouth string
}

func NewCompany(name, address, sfera string) *Company {
	return &Company{
		Name:    name,
		Address: address,
		Sfera:   sfera,
	}
}

//1: company; Human; - go to work -> to company
func NewBuilder(h *Human, c *Company) *Builder {
	return &Builder{
		Company: c,
		Human:   h,
	}
}

//polymorph
func (b *Builder) Work() {
	b.Company.GetInfo() // get company info
	log.Println("строитель: на работе занимается стройкой", b.id, b.Human.Name, b.Human.DateBirth)
}

//polymorph
func (s *Saler) Work() {
	log.Println("продавец : на работе занимается продажей ")
}

//polymorph
func (p *Programmer) Work() {
	log.Println("прогер : на работе занимается хуйней ")
}
