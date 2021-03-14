function DynDrawMatrix() {
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
    var spDrawData = getSecondPt.InitUserDraw("TestDynDrawMatrix");
    // 设置动态绘制参数.
    spDrawData.SetPoint("BasePoint", point1);
    // 开始取第二个点。
    if (getSecondPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
}
function DynWorldDrawMatrix( pCustomEntity, pWorldDraw, curPt) {
    //与用户交互在图面上提取一个点
    var vBasePt = pCustomEntity.GetPoint("BasePoint");
    //计算中心点位置
    var centerPt = mxOcx.NewPoint();
    {
        centerPt.x = curPt.x + (vBasePt.x - curPt.x) / 2;
        centerPt.y = curPt.y + (vBasePt.y - curPt.y) / 2;
    }
    //计算面积（取绝对值）
    var arec = Math.abs(vBasePt.x - curPt.x) * Math.abs(vBasePt.y - curPt.y);
    //计算周长
    var girth = 2 * (Math.abs(vBasePt.x - curPt.x) + Math.abs(vBasePt.y - curPt.y));
    //变换视区长度
    var dHeigth = mxOcx.ViewLongToDocCoord(15);
    //绘制一个单行文字
    //参数一为文字的位置的X坐标 ；参数二为文字的位置的Y坐标 ；参数三为文字内容字符串
    //参数四为文字高度；参数五为文字的旋转角度
    //参数六为文字的水平对齐方式,0=kTextLeft,1=kTextCenter,2=kTextRight
    //参数七文字的竖直对齐方式,1=kTextBottom,2=kTextVertMid,3=kTextTop
    pWorldDraw.DrawText(centerPt.x, centerPt.y, "面积：" + arec.toFixed(2), dHeigth, 0, 1, 1);
    pWorldDraw.DrawText(centerPt.x, centerPt.y, "周长：" + girth.toFixed(2), dHeigth, 0, 1, 3);
    //绘制一个直线
    //参数一直线的开始点x坐标，参数二直线的开始点y坐标，参数三直线的结束点x坐标，参数四直线的结束点y坐标
    pWorldDraw.DrawLine (vBasePt.x, vBasePt.y, vBasePt.x, curPt.y);
    pWorldDraw.DrawLine (vBasePt.x, curPt.y, curPt.x, curPt.y);
    pWorldDraw.DrawLine (curPt.x, curPt.y, curPt.x, vBasePt.y);
    pWorldDraw.DrawLine (curPt.x, vBasePt.y, vBasePt.x, vBasePt.y);
}
