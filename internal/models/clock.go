package models

import (
	"time"

	"github.com/rhuanpk/pointers-angle/pkg/utils/structure"
)

// Clock is the struct for ORM operations.
type Clock struct {
	ID     *uint      `gorm:"primaryKey" uri:"id" binding:"-"`
	Hour   *byte      `gorm:"not null" uri:"hour" binding:"required,numeric,gte=1,lte=12"`
	Minute *byte      `gorm:"not null" uri:"minute" binding:"omitempty,numeric,gte=0,lte=59"`
	Angle  *int16     `gorm:"not null" uri:"angle" binding:"-"`
	Date   *time.Time `gorm:"autoCreateTime" uri:"date" binding:"-"`
}

// NewClock return a *Clock with your fields memory allocated.
func NewClock() *Clock {
	clock := &Clock{}
	structure.Allocate(clock, "ID", "Date")
	return clock
}
