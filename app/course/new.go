package course

import "gitdev.devops.krungthai.com/techcoach/template/goapi.git/app"

type NewCourse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c *Courses) NewCourse(ctx app.Context) {
	var course NewCourse
	if err := ctx.Bind(&course); err != nil {
		ctx.BadRequest(err)
		return
	}

	if err := c.storage.NewCourses(course); err != nil {
		ctx.StoreError(err)
		return
	}
	ctx.OK(nil)
}
