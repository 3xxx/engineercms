//动态绘制圆弧
function DyDrawArc() {
    //清空当前显示内容
    mxOcx.NewFile();
    // 新建一个COM组件对象
    var getPt1 = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt1.message = "指定圆弧的起点";
    if (getPt1.go() != 1) {
        return;
    }
    // 新建一个COM组件对象
    var getPt2 = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt2.message = "指定圆弧的第二个点";
    //基点
    getPt2.basePoint = getPt1.value();
    //设置是否启用拖放基点  true：使用，false：不使用
    getPt2.setUseBasePt(true);
    if (getPt2.go() != 1) {
        return;
    }
    // 新建一个COM组件对象
    var getPt3 = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt3.message = "指定圆弧的端点";
    //初始动态绘制；  动态绘制Id,在动态绘制函数里，这个值用于判断是那个动态绘制
    var spDrawData = getPt3.InitUserDraw("DrawArc");
    //设置point类型的属性
    spDrawData.SetPoint("pt1", getPt1.value());
    spDrawData.SetPoint("pt2", getPt2.value());
    if (getPt3.go() != 1) {
        return;
    }
    //把实体绘制到图上。一般用于动态拖放时，在拖放结束时，绘制到图上；返回新绘制的实体的id
    spDrawData.Draw();
}

