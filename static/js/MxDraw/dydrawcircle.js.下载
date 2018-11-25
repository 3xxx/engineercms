//绘制圆
function DyDrawCircle() {
    //清空当前显示内容
    mxOcx.NewFile();
    // 创建一个与用户交互取点的对象。
    var point1 = mxOcx.GetPoint(false,0,0,"\n 点取圆心:");
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
    var spDrawData = getSecondPt.InitUserDraw("TestDynDrawCircle");
    // 设置动态绘制参数.
    spDrawData.SetPoint("pt1", point1);
    // 开始取第二个点。
    if (getSecondPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
}

