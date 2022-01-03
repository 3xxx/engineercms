//绘制文字
function DyDrawText()
{
    //清空当前显示内容
    mxOcx.NewFile();
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
    getPt.message = "点取插入点";
    // 设置动态绘制参数.
    var spDrawData = getPt.InitUserDraw("DrawText");
    spDrawData.SetString("Text", txt);
    spDrawData.SetDouble("TextHeight", 100);
    // 开始取第一个点。
    if (getPt.go() != 1)
        return;
    var ret = spDrawData.Draw();

}


function DynWorldDrawText( pCustomEntity, pWorldDraw, curPt) {
    // 得到绘制参数.
    var sText       = pCustomEntity.GetString("Text");
    var dTextHeight = pCustomEntity.GetDouble("TextHeight");
    // 绘制文字
    pWorldDraw.DrawText (curPt.x,curPt.y,sText,dTextHeight,0,1,1);
}
