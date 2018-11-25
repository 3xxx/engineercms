//绘制圆形的云标注.
function DyDoCloudCircleComment() {
    //清空当前显示内容
    mxOcx.NewFile();
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message = "点取第一点";
    if (getPt.go() != 1) {
        return;
    }
    var frstPt = getPt.value();
    if (frstPt == null)
        return;
    var getPt2 = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt2.message = "指定第二个点";
    var spDrawData = getPt2.InitUserDraw("DrawCircle");
    spDrawData.SetPoint("pt1", frstPt);
    if (getPt2.go() != 1) {
        return;
    }
    var secondPt = getPt2.value();
    if (secondPt == null)
        return;
    //调用文字编辑框
    var param = mxOcx.NewResbuf();
    param.AddString("");
    param.AddDouble(100);
    var ret = mxOcx.CallEx("Mx_ShowMTextDialog", param);
    if (ret.AtString(0) != "Ok") {
        return;
    }
    var txt = ret.AtString(1);
    var txtH = ret.AtDouble(2);
    var txtColorIndex = ret.AtLong(3);
    // 加入换行操作符\\P
    //txt = txt + "\\PTEST";
    // 创建一个文字样式,用于批注文字字体.
    mxOcx.AddTextStyle2("MyCommentFont", "黑体", 0.7);
    // 创建一个与用户交互取点的对象。
    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getSecondPt.message = "输入标注位置点";
    getSecondPt.basePoint = secondPt;
    getSecondPt.setUseBasePt(false);
    spDrawData = getSecondPt.InitUserDraw("DrawCloudComment");
    // 设置动态绘制参数.
    spDrawData.SetPoint("frstPt", frstPt);
    spDrawData.SetPoint("BasePoint", secondPt);
    spDrawData.SetString("Text", txt);
    spDrawData.SetDouble("TextHeight", txtH);
    spDrawData.SetLong("EdgeNumber", 2);
    spDrawData.SetDouble("ShapRadius", 1);
    // 开始取第二个点。
    if (getSecondPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
    // 设置绘制的批注文字样式。
    var ent = ret.AtObject(0);
    ent.colorIndex = 1;
}
