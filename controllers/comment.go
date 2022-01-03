package controllers

type CommentController struct {
	MindocBaseController
}

func (c *CommentController) Lists() {

}

func (c *CommentController) Create() {

	c.JsonResult(0, "ok")
}

func (c *CommentController) Index() {
	c.Prepare()
	c.TplName = "comment/index.tpl"
}
