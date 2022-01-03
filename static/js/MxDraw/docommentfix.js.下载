function DoCommentFix() {
    //清空当前显示内容
    mxOcx.NewFile();
    var comment = mxOcx.NewEntity("IMxDrawComment");
    comment.Text = "固定参数绘批注";
    comment.TextHeight = 10;
    var basepos = mxOcx.NewPoint();
    basepos.x = 0;
    basepos.y = 10;
    var pos = mxOcx.NewPoint();
    pos.x = 20;
    pos.y = 30;
    //基点
    comment.BasePoint = basepos;
    //位置
    comment.Position = pos;
    var comobj = mxOcx.ObjectIdToObject( mxOcx.DrawEntity(comment) );
    comobj.TextStyle = "MyCommentFont";
    mxOcx.ZoomAll();
}