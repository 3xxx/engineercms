//绘制多线段
function DyDrawPolyline() {
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //与用户交互到在图上提取一个点
    var point1 = mxOcx.GetPoint(false,0,0,"\n 点取开始点:");
    if(point1 == null)
    {
        return;
    }
    //把路径的开始位置移动指定的点
    //参数一为点的X坐标 ，参数二为点的Y坐标
    mxOcx.PathMoveTo(point1.x,point1.y);
    //与用户交互到在图上提取一个点
    var point2 = mxOcx.GetPoint(true,point1.x,point1.y,"\n 点取下一个:");
    if(point2 == null)
    {
        return;
    }
    //建一个数组
    var tmpobj = new Array();
    //把路径下一个点移到指定位置
    //参数一为点的X坐标 ，参数二为点的Y坐标
    mxOcx.PathLineTo(point2.x,point2.y);
    var iCount = 0;
    //绘制一个直线
    //参数一直线的开始点x坐标，参数二直线的开始点y坐标，参数三直线的结束点x坐标，参数四直线的结束点y坐标
    var id = mxOcx.DrawLine(point1.x,point1.y,point2.x,point2.y);
    tmpobj[iCount] = id;
    iCount = iCount + 1;
    point1 = point2;
    while(true)
    {
        var point2 = mxOcx.GetPoint(true,point1.x,point1.y,"\n 点取下一个:");
        if(point2 == null)
        {
            break;
        }
        mxOcx.PathLineTo(point2.x,point2.y);
        var id = mxOcx.DrawLine(point1.x,point1.y,point2.x,point2.y);
        tmpobj[iCount] = id;
        iCount = iCount + 1;
        point1 = point2;
    }
    var i = 0;
    for(;i < iCount;i++)
    {
        //由id删除一个对象
        mxOcx.Erase(tmpobj[i] );
    }
    //把路径变成一个Polyline
    mxOcx.DrawPathToPolyline();

}
