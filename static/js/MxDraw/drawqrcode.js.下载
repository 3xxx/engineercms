function DrawQrCode() {
    //清空当前显示内容
    mxOcx.NewFile();
    var winWidth = 440;
    var winHeight = 140;
    var winLeft = (screen.width - winWidth) / 2;
    var winTop = (screen.height - winHeight) / 2 - 20;

    var str = 'dialogHeight:' + winHeight + 'px;dialogWidth:' + winWidth + 'px;dialogTop:' + winTop + 'px;dialogLeft:' + winLeft + 'px;resizable:yes;center:yes;status:no;'

    var rt = window.showModalDialog("Gettext.htm?tmp=" + Math.random(), "输入二维码文本", str);

    var txt;

    if (typeof (rt) == "undefined") {
        return;
    } else {
        var arr = rt.split(",");
        txt = arr[0];
    }
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message = "点取插入点";
    if (getPt.go() != 1) {
        return;
    }
    var pos = getPt.value();
    if (pos == null)
        return;

    var qrCode = mxOcx.NewEntity("IMxDrawQrCode");
    qrCode.Text = txt;
    qrCode.Center = pos;
    mxOcx.DrawEntity(qrCode);
}
