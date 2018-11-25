function DoFixRectComment() {
    //清空当前显示内容
    mxOcx.NewFile();
    var ent = mxOcx.DrawCustomEntity("TestMxCustomEntity2", "");
    var basept = mxOcx.NewPoint();
    basept.x = 100;
    basept.y = 100;
    ent.SetPoint("BasePoint", basept);
    ent.SetString("Text", "固定矩形标注");
    ent.SetDouble("TextHeight", 100);
    ent.SetLong("EdgeNumber", 2);
    ent.SetDouble("ShapRadius", 1);
    var dimpt = mxOcx.NewPoint();
    dimpt.x = 200;
    dimpt.y = 200;
    ent.SetPoint("DimPoint", dimpt);
    var frstPt = mxOcx.NewPoint();
    frstPt.x = 0;
    frstPt.y = 0;
    ent.SetPoint("First", frstPt);
    ent.TextStyle = "MyCommentFont";
    /*var comment = mxOcx.NewEntity("IMxDrawComment");
     //标注文字
     comment.Text = "固定矩形标注";
     //标注文字高度
     comment.TextHeight = 100;
     //标注位置提示多边形的边数. 小于2不绘制，=2绘制圆，大于2绘制多边形
     comment.EdgeNumber = 2;
     //标注位置提示多边形的半径
     comment.ShapRadius = 1;
     //标注基点
     var basept = mxOcx.NewPoint();
     basept.x = 100;
     basept.y = 100;
     comment.basePoint = basept;
     //标注位置点
     var dimpt = mxOcx.NewPoint();
     dimpt.x = 200;
     dimpt.y = 200;
     comment.Position = dimpt;
     /!*  // 设置文字样式
     pWorldDraw.TextStyle = "MyCommentFont";
     // 动态绘制.
     pWorldDraw.DrawEntity(comment);
     *!/
     // 绘制矩形框.
     //绘制一个直线
     //参数一直线的开始点x坐标，参数二直线的开始点y坐标，参数三直线的结束点x坐标，参数四直线的结束点y坐标

     //定义一个路径的开始点
     comment.TextHeight = 50;
     var frstPt = mxOcx.NewPoint();
     frstPt.x = -200;
     frstPt.y = -200;
     var basepos = mxOcx.NewPoint();
     basepos.x = 0;
     basepos.y = 10;
     var points = mxOcx.NewComObject("IMxDrawPoints");
     points.Add2(frstPt);
     points.Add2(basepos);
     //标注框坐标数据
     comment.Bounding = points;

     var comobj = mxOcx.ObjectIdToObject(mxOcx.DrawEntity(comment));
     comobj.TextStyle = "MyCommentFont";

     mxOcx.SetEventRet(1);*/
}
