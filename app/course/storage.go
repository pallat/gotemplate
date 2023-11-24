package course

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) NewCourses(c NewCourse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlStmt := `insert into courses(title) values(?)`
	_, err := s.db.ExecContext(ctx, sqlStmt, c.Title)
	if err != nil {
		return errors.WithMessagef(err, "new a course: %s", c.Title)
	}
	return nil
}

func (s *Storage) RecommendedCourses(vctx ViewerContext) (courses []Course, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "select title from courses")
	if err != nil {
		return nil, errors.WithMessage(err, "RecommendedCourses")
	}
	defer rows.Close()
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil, errors.WithMessage(err, "RecommendedCourses")
		}
		courses = append(courses, Course{Title: title})
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.WithMessage(err, "RecommendedCourses")
	}

	return
}
