//绘制矩形框
function DyDrawPathToPolyline() {
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
    var spDrawData = getSecondPt.InitUserDraw("TestDynDrawMatrix1");
    // 设置动态绘制参数.
    spDrawData.SetPoint("BasePoint", point1);
    // 开始取第二个点。
    if (getSecondPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
}
function DynWorldDrawMatrix1( pCustomEntity, pWorldDraw, curPt) {
    //与用户交互在图面上提取一个点
    var vBasePt = pCustomEntity.GetPoint("BasePoint");
    //绘制一个直线
    //参数一直线的开始点x坐标，参数二直线的开始点y坐标，参数三直线的结束点x坐标，参数四直线的结束点y坐标
    pWorldDraw.DrawLine (vBasePt.x, vBasePt.y, vBasePt.x, curPt.y);
    pWorldDraw.DrawLine (vBasePt.x, curPt.y, curPt.x, curPt.y);
    pWorldDraw.DrawLine (curPt.x, curPt.y, curPt.x, vBasePt.y);
    pWorldDraw.DrawLine (curPt.x, vBasePt.y, vBasePt.x, vBasePt.y);
}

