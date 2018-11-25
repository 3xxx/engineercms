//绘制点
function DyDrawPoint() {
    //清空当前显示内容
    mxOcx.NewFile();

    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    var spDrawData = getPt.InitUserDraw("TestDynDrawPoint");
    // 开始取第二个点。
    if (getPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
}

