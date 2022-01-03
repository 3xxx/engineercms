function SelectByPolygon(){
    var pt1 = mxOcx.NewPoint();
    pt1.x = -200;
    pt1.y = -200;
    var pt2 = mxOcx.NewPoint();
    pt2.x = 200;
    pt2.y = 200;
    var pt3 = mxOcx.NewPoint();
    pt3.x = 200;
    pt3.y = -200;
    var pt4 = mxOcx.NewPoint();
    pt4.x = -200;
    pt4.y = 200;
    var points = mxOcx.NewComObject("IMxDrawPoints");
    points.Add2(pt1);
    points.Add2(pt2);
    points.Add2(pt3);
    points.Add2(pt4);
    //实例化一个构造选择集进行过滤,该类封装了选择集及其处理函数。
    var ss = mxOcx.NewSelectionSet();
    //构造一个过滤链表
    var spFilte = mxOcx.NewResbuf();
    //在多个点组合的闭合区域里，构造选择集
    ss.SelectByPolygon(7,points, spFilte);
    for (var i = 0; i < ss.Count; i++)
    {
        var ent = ss.Item(i);
        alert(ent.ObjectName);
    }
}