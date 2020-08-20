package visitor

type UserInterface interface {
	Accept(Visitor) string
}

type User struct {
	idNumber string
	realName string
}

func (u *User) Accept(visitor Visitor) string {
	return visitor.VisitUser(u)
}

type Company struct {
	regNumber string
	name      string
}

func (c *Company) Accept(visitor Visitor) string {
	return visitor.VisitCompany(c)
}

type Visitor interface {
	VisitUser(*User) string
	VisitCompany(*Company) string
}

type NameVisitor struct{}

func (v *NameVisitor) VisitUser(user *User) string {
	return user.realName
}
func (v *NameVisitor) VisitCompany(company *Company) string {
	return company.name
}

type IdNumberVisitor struct{}

func (v *IdNumberVisitor) VisitUser(user *User) string {
	return user.idNumber
}

func (v *IdNumberVisitor) VisitCompany(company *Company) string {
	return company.regNumber
}
