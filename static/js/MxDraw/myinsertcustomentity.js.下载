//绘制三角形
function MyInsertCustomEntity()
{ //清空当前显示内容
    mxOcx.NewFile();
    var spDrawData;
    var getPt1 = mxOcx.NewComObject("IMxDrawUiPrPoint");
    {
        getPt1.message = "点取第一点";
        spDrawData = getPt1.InitUserDraw("DynDrawEnt1");
        if (getPt1.go() != 1) {
            return;
        }
    }
    var getPt2 = mxOcx.NewComObject("IMxDrawUiPrPoint");
    {
        getPt2.message = "点取第二点";
        spDrawData = getPt2.InitUserDraw("DynDrawEnt2");
        spDrawData.SetPoint("vFirst", getPt1.value());
        if (getPt2.go() != 1) {
            return;
        }
    }
    var getPt3 = mxOcx.NewComObject("IMxDrawUiPrPoint");
    {
        getPt3.message = "点取第三点";
        spDrawData = getPt3.InitUserDraw("DynDrawEnt3");
        spDrawData.SetPoint("vFirst", getPt1.value());
        spDrawData.SetPoint("vSecond", getPt2.value());
        if (getPt3.go() != 1) {
            return;
        }
    }
    var getPt4 = mxOcx.NewComObject("IMxDrawUiPrPoint");
    {
        getPt4.message = "点取第四点";
        spDrawData = getPt4.InitUserDraw("DynDrawEnt4");
        spDrawData.SetPoint("vFirst", getPt1.value());
        spDrawData.SetPoint("vSecond", getPt2.value());
        spDrawData.SetPoint("vThird", getPt3.value());
        if (getPt4.go() != 1) {
            return;
        }
    }
    {
        spDrawData = mxOcx.DrawCustomEntity("DrawEnt", "");
        spDrawData.SetPoint("vFirst", getPt1.value());
        spDrawData.SetPoint("vSecond", getPt2.value());
        spDrawData.SetPoint("vThird", getPt3.value());
        spDrawData.SetPoint("vFour", getPt4.value());
    }
}
