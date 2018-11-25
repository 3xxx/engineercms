//绘制自定义实体
function DyInsertCustomEntity() {
    //清空当前显示内容
    mxOcx.NewFile();
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message = "点取第一点";
    if (getPt.go() != 1)
        return;
    var frstPt = getPt.value();
    if (frstPt == null)
        return;
    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getSecondPt.message = "点取第二点";
    getSecondPt.basePoint = frstPt;
    getSecondPt.setUseBasePt(true);
    if (getSecondPt.go() != 1)
        return;
    var secondPt = getSecondPt.value();
    if (secondPt == null)
        return;
    var ent = mxOcx.DrawCustomEntity("TestMxCustomEntity", "");
    ent.SetPoint("spt", frstPt);
    ent.SetPoint("ept", secondPt);
}