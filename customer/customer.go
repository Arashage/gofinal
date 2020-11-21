package customer

/*
Customer information
*/
type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

/*
func (c Customer) ID() int {
	return c.id
}

func (c *Customer) SetID(id int) {
	c.id = id
}

func (c Customer) Name() string {
	return c.name
}

func (c *Customer) SetName(name string) {
	c.name = name
}

func (c Customer) Email() string {
	return c.email
}

func (c *Customer) SetEmail(email string) {
	c.email = email
}

func (c Customer) Status() string {
	return c.status
}

func (c *Customer) SetStatus(status string) {
	c.status = status
}
*/
