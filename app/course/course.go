package course

import "time"

type style string

type tag string

type level string

type Banner struct {
	URL string
}

type Instructor struct {
	Name       string
	PictureURL string
}

type Price struct {
	Price          uint
	PromotionPrice uint
}

type Rating struct {
	Rate     float64
	Feedback uint
}

type Course struct {
	Banner      Banner
	Style       style
	Title       string
	Description string
	Instructor  Instructor
	Tags        []tag
	Rating      Rating
	Price       string
	Duration    time.Duration
	Lectures    string
	Level       level
}

type storage interface {
	RecommendedCourses(ViewerContext) ([]Course, error)
	NewCourses(NewCourse) error
}

type Courses struct {
	storage storage
}

func New(s storage) *Courses {
	return &Courses{storage: s}
}
