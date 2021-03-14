function DoCloudCircleCommentFix() {
    //清空当前显示内容
    mxOcx.NewFile();
    var comment = mxOcx.NewEntity("IMxDrawComment");
    comment.Text = "固定参数绘圆形云线标注";
    comment.TextHeight = 50;
    var frstPt = mxOcx.NewPoint();
    frstPt.x = -200;
    frstPt.y = -200;
    var basepos = mxOcx.NewPoint();
    basepos.x = 0;
    basepos.y = 10;
    var pos = mxOcx.NewPoint();
    pos.x = 200;
    pos.y = 300;
    //基点
    comment.BasePoint = basepos;
    //位置
    comment.Position = pos;
    //标注位置提示多边形的半径
    comment.ShapRadius = 50;
    //标注框类型
    comment.BoundingBoxType = 5;
    //标注箭头类型
    comment.ArrowType = 2;
    var points = mxOcx.NewComObject("IMxDrawPoints");
    points.Add2(frstPt);
    points.Add2(basepos);
    //标注框坐标数据
    comment.Bounding = points;
    var comobj = mxOcx.ObjectIdToObject(mxOcx.DrawEntity(comment));
    comobj.TextStyle = "MyCommentFont";
    mxOcx.ZoomAll();
}
