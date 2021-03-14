function DoFixCircleComment()
{
    //清空当前显示内容
    mxOcx.NewFile();
    var ent = mxOcx.DrawCustomEntity("TestMxCustomEntity2", "");
    var basept = mxOcx.NewPoint();
    basept.x = 100;
    basept.y = 100;
    ent.SetPoint("BasePoint", basept);
    ent.SetString("Text", "固定参数绘制圆形标注");
    ent.SetDouble("TextHeight", 100);
    //标注位置提示多边形的边数. 小于2不绘制，=2绘制圆，大于2绘制多边形
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
    ent.SetLong("isCircle", 1);
}
