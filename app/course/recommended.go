package course

import "gitdev.devops.krungthai.com/techcoach/template/goapi.git/app"

type ViewerContext struct {
	Lang     string
	Currency string
}

func (c *Courses) Recommended(ctx app.Context) {
	var viewer ViewerContext
	if err := ctx.Bind(&viewer); err != nil {
		ctx.BadRequest(err)
		return
	}

	courses, err := c.storage.RecommendedCourses(viewer)
	if err != nil {
		ctx.StoreError(err)
		return
	}
	ctx.OK(courses)
}
