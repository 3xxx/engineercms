// 动态绘制批注函数.
function DyDoCircleComment() {
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
    // 出来一个对象，让用户输入文字.
    var winWidth = 440;
    var winHeight = 140;
    var winLeft = (screen.width - winWidth) / 2;
    var winTop = (screen.height - winHeight) / 2 - 20;
    var str = 'dialogHeight:' + winHeight + 'px;dialogWidth:' + winWidth + 'px;dialogTop:' + winTop + 'px;dialogLeft:' + winLeft + 'px;resizable:yes;center:yes;status:no;'
    var rt = window.showModalDialog("Gettext.htm?tmp=" + Math.random(), "输入文字", str);
    var txt;
    if (typeof (rt) == "undefined") {
        return;
    } else {
        var arr = rt.split(",");
        txt = arr[0];
    }
    // 加入换行操作符\\P
    //txt = txt + "\\PTEST";
    // 创建一个文字样式,用于批注文字字体.
    mxOcx.AddTextStyle2("MyCommentFont", "黑体", 0.7);
    // 创建一个与用户交互取点的对象。
    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getSecondPt.message = "输入标注位置点";
    getSecondPt.basePoint = secondPt;
    getSecondPt.setUseBasePt(false);
    spDrawData = getSecondPt.InitUserDraw("DrawComment2");
    // 设置动态绘制参数.
    spDrawData.SetPoint("BasePoint", secondPt);
    spDrawData.SetString("Text", txt);
    spDrawData.SetDouble("TextHeight", 100);
    spDrawData.SetLong("EdgeNumber", 2);
    spDrawData.SetDouble("ShapRadius", 1);
    // 开始取第二个点。
    if (getSecondPt.go() != 1)
        return;
    var ent = mxOcx.DrawCustomEntity("TestMxCustomEntity1", "");
    ent.SetPoint("BasePoint", secondPt);
    ent.SetString("Text", txt);
    ent.SetDouble("TextHeight", 100);
    ent.SetLong("EdgeNumber", 2);
    ent.SetDouble("ShapRadius", 1);
    ent.SetPoint("DimPoint", getSecondPt.value());
    ent.SetPoint("First", frstPt);
    ent.SetLong("isCircle", 1);
}
