package clock

import (
	"fmt"
	"math"
)

//Clock - define the clock type
type Clock struct {
	Hours   int
	Minutes int
}

//New - Creates a new clock
func New(hour, minute int) Clock {
	return Clock{}.Add((hour * 60) + minute)
}

//String - Get the clock hours string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

//Add - Add minutes to the clock
func (c Clock) Add(minutes int) Clock {
	time := ((c.Hours * 60) + c.Minutes) + (24 * 60) + (minutes % (24 * 60))
	hours := int(math.Floor(float64(time) / 60))
	c.Minutes = time - (hours * 60)
	c.Hours = hours % 24
	return c
}

//Subtract - Sutract minutes of the clock
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
