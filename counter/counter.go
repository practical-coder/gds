package counter

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Inc() {
	c.count++
}

func (c *Counter) Dec() {
	if c.count > 0 {
		c.count--
	}
}

func (c *Counter) Reset() {
	c.count = 0
}

func (c Counter) Get() int {
	return c.count
}

func (c Counter) String() string {
	return fmt.Sprintf("counter: %d", c.count)
}
