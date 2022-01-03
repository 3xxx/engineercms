function DyDoComment() {
    //清空当前显示内容
    mxOcx.NewFile();
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
    mxOcx.AddTextStyle2("MyCommentFont","黑体",0.7);
    // 创建一个与用户交互取点的对象。
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message = "输入标注插入基点";
    // 设置动态绘制参数.
    var spDrawData = getPt.InitUserDraw("DrawComment");
    spDrawData.SetString("Text", txt);
    spDrawData.SetDouble("TextHeight", 100);
    spDrawData.SetLong("EdgeNumber", 0);
    spDrawData.SetDouble("ShapRadius", 300);
    // 开始取第一个点。
    if (getPt.go() != 1)
        return;
    // 创建一个与用户交互取点的对象。
    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getSecondPt.message = "输入标注位置点";
    getSecondPt.basePoint = getPt.value();
    getSecondPt.setUseBasePt(false);
    spDrawData = getSecondPt.InitUserDraw("DrawComment2");
    // 设置动态绘制参数.
    spDrawData.SetPoint("BasePoint", getPt.value());
    spDrawData.SetString("Text", txt);
    spDrawData.SetDouble("TextHeight", 100);
    spDrawData.SetLong("EdgeNumber", 0);
    spDrawData.SetDouble("ShapRadius", 300);
    // 开始取第二个点。
    if (getSecondPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
    // 设置绘制的批注文字样式。
    for (var i = 0; i < ret.Count; i++)
    {
        var ent = ret.AtObject(i);
        ent.TextStyle = "MyCommentFont";
        ent.colorIndex = 4;
    }
}
