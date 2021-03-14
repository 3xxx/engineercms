//绘制虚线直线
function DyDrawMyline() {
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
    var spDrawData = getSecondPt.InitUserDraw("TestDynDrawMyLine");
    // 设置动态绘制参数.
    spDrawData.SetPoint("BasePoint", point1);
    // 开始取第二个点。
    if (getSecondPt.go() != 1)
        return;
    var ret = spDrawData.Draw();

}

function DynWorldDrawMyLine( pCustomEntity, pWorldDraw, curPt) {
    //与用户交互在图面上提取一个点
    var vBasePt = pCustomEntity.GetPoint("BasePoint");
    //定义虚线数据据,"MyLineType"是线型名，"6,-8"是虚线的一个单位定义，6是实线长，-8是空格长
    mxOcx.AddLinetype("MyLineType", "6,-10");
    //设计当前线型为"MyLineType"
    pWorldDraw.LineType = "MyLineType";
    //绘制一个直线
    //参数一直线的开始点x坐标，参数二直线的开始点y坐标，参数三直线的结束点x坐标，参数四直线的结束点y坐标
    pWorldDraw.DrawLine (vBasePt.x, vBasePt.y, curPt.x, curPt.y);
}
