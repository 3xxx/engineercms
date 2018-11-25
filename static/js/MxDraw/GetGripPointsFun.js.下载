// 返回自定义实体夹点
function GetGripPointsFun(pCustomEntity) {
    var sGuid = pCustomEntity.Guid;
    //----------------------------------------------------------------------------------------------------------
    //交互绘制
    if (sGuid == "TestMxCustomEntity1") {
        if (!pCustomEntity.IsHave("First"))
            return;
        var stp = pCustomEntity.GetPoint("First");
        var ept = pCustomEntity.GetPoint("BasePoint");
        var dimpt = pCustomEntity.GetPoint("DimPoint");
        var ret = mxOcx.NewResbuf();
        ret.AddPoint(stp);
        ret.AddPoint(ept);
        ret.AddPoint(dimpt);
        mxOcx.SetEventRetEx(ret);
    }
    //---------------------------------------------------------------------------------------------------------
    //自定义实体绘制两条直线
    if (sGuid == "TestMxCustomEntity") {
        if (!pCustomEntity.IsHave("ept"))
            return;

        var stp = pCustomEntity.GetPoint("spt");
        if (stp == null)
            return;

        var ept = pCustomEntity.GetPoint("ept");
        if (ept == null)
            return;

        var ret = mxOcx.NewResbuf();

        ret.AddPoint(stp);
        ret.AddPoint(ept);

        mxOcx.SetEventRetEx(ret);
    }
    //--------------------------------------------------------------------------------------------------------
    //固定参数绘制
    if (sGuid == "TestMxCustomEntity2") {
        if (!pCustomEntity.IsHave("First"))
            return;
        var stp = pCustomEntity.GetPoint("First");
        var ept = pCustomEntity.GetPoint("BasePoint");
        var dimpt = pCustomEntity.GetPoint("DimPoint");
        var ret = mxOcx.NewResbuf();
        ret.AddPoint(stp);
        ret.AddPoint(ept);
        ret.AddPoint(dimpt);
        mxOcx.SetEventRetEx(ret);
    }

    //-------------------------------------------------------------------------------------------------------
    //自定义画三角形
    if (sGuid == "DrawEnt") {
        if (!pCustomEntity.IsHave("vFour"))
            return;
        var stp = pCustomEntity.GetPoint("vFirst");
        if (stp == null)
            return;
        var ept = pCustomEntity.GetPoint("vSecond");
        if (ept == null)
            return;
        var tpt = pCustomEntity.GetPoint("vThird");
        if (tpt == null)
            return;
        var fpt = pCustomEntity.GetPoint("vFour");
        if (fpt == null)
            return;
        var ret = mxOcx.NewResbuf();
        ret.AddPoint(stp);
        ret.AddPoint(ept);
        ret.AddPoint(tpt);
        ret.AddPoint(fpt);
        mxOcx.SetEventRetEx(ret);
    }
}
