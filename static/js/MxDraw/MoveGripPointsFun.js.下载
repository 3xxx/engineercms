// 移动自定义实体夹点
function MoveGripPointsFun(pCustomEntity, lGridIndex, dOffsetX, dOffsetY) {

    var sGuid = pCustomEntity.Guid;
    if (sGuid == "TestMxCustomEntity1") {
        if (!pCustomEntity.IsHave("First"))
            return;
        var stp = pCustomEntity.GetPoint("First");
        var ept = pCustomEntity.GetPoint("BasePoint");
        var dimpt = pCustomEntity.GetPoint("DimPoint");
        if (lGridIndex == 0) {
            stp.x = stp.x + dOffsetX;
            stp.y = stp.y + dOffsetY;
            pCustomEntity.SetPoint("First", stp);
        }
        else if (lGridIndex == 1)
        {
            ept.x = ept.x + dOffsetX;
            ept.y = ept.y + dOffsetY;
            pCustomEntity.SetPoint("BasePoint", ept);
        }
        else
        {
            dimpt.x = dimpt.x + dOffsetX;
            dimpt.y = dimpt.y + dOffsetY;
            pCustomEntity.SetPoint("DimPoint", dimpt);
        }
        mxOcx.SetEventRet(1);
    }
    if (sGuid == "TestMxCustomEntity") {
        if (!pCustomEntity.IsHave("ept"))
            return;

        var stp = pCustomEntity.GetPoint("spt");
        if (stp == null)
            return;

        var ept = pCustomEntity.GetPoint("ept");
        if (ept == null)
            return;


        if (lGridIndex == 0) {
            stp.x = stp.x + dOffsetX;
            stp.y = stp.y + dOffsetY;
            pCustomEntity.SetPoint("spt", stp);
        }
        else {
            ept.x = ept.x + dOffsetX;
            ept.y = ept.y + dOffsetY;
            pCustomEntity.SetPoint("ept", ept);
        }

        mxOcx.SetEventRet(1);
    }
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
        {
            stp.x = stp.x + dOffsetX;
            stp.y = stp.y + dOffsetY;
            pCustomEntity.SetPoint("vFirst", stp);
            ept.x = ept.x + dOffsetX;
            ept.y = ept.y + dOffsetY;
            pCustomEntity.SetPoint("vSecond", ept);
            tpt.x = tpt.x + dOffsetX;
            tpt.y = tpt.y + dOffsetY;
            pCustomEntity.SetPoint("vThird", tpt);
            fpt.x = fpt.x + dOffsetX;
            fpt.y = fpt.y + dOffsetY;
            pCustomEntity.SetPoint("vFour", fpt);
        }
        mxOcx.SetEventRet(1);
    }
}
