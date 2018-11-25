//绘制直线
function DyDrawLine() {
    //清空当前显示内容
    mxOcx.NewFile();
    // 创建一个与用户交互取点的对象。
    var point1 = mxOcx.GetPoint(false,0,0,"\n 点取开始点:");
    if(point1 == null)
    {
        return;
    }
    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getSecondPt.message = "第二点";
    getSecondPt.basePoint = point1;
    //不开启皮筋效果
    getSecondPt.setUseBasePt(false);
    //调用
    var spDrawData = getSecondPt.InitUserDraw("TestDynDrawLine");
    // 设置动态绘制参数.
    spDrawData.SetPoint("BasePoint", point1);
    // 开始取第二个点。
    if (getSecondPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
}

function DynWorldDrawLine( pCustomEntity, pWorldDraw, curPt) {
    //与用户交互在图面上提取一个点
    var vBasePt = pCustomEntity.GetPoint("BasePoint");
    var length=vBasePt.DistanceTo(curPt);
    var length1=length.toFixed(2);
    //计算中心点位置
    var centerPt = mxOcx.NewPoint();
    {
        centerPt.x = curPt.x + (vBasePt.x - curPt.x) / 2;
        centerPt.y = curPt.y + (vBasePt.y - curPt.y) / 2;
    }
    //计算this点减去pPoint，得到的向量
    var vec = vBasePt.SumVector(curPt);
    //返回向量顺时针转到与X轴的角度.[0,2PI)
    var dAng = vec.Angle();
    //变换视区长度
    var dHeigth = mxOcx.ViewLongToDocCoord(15);
    //绘制一个单行文字
    //参数一为文字的位置的X坐标 ；参数二为文字的位置的Y坐标 ；参数三为文字内容字符串
    //参数四为文字高度；参数五为文字的旋转角度
    //参数六为文字的水平对齐方式,0=kTextLeft,1=kTextCenter,2=kTextRight
    //参数七文字的竖直对齐方式,1=kTextBottom,2=kTextVertMid,3=kTextTop
    pWorldDraw.DrawText(centerPt.x, centerPt.y, "长度：" + length1, dHeigth, dAng * 180.0 / 3.14159265, 1, 1);
    //绘制一个直线
    //参数一直线的开始点x坐标，参数二直线的开始点y坐标，参数三直线的结束点x坐标，参数四直线的结束点y坐标
    pWorldDraw.DrawLine (vBasePt.x, vBasePt.y, curPt.x, curPt.y);
}
