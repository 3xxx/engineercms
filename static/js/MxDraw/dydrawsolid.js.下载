/**
 * Created by admin on 2018/3/2.
 */
//绘制McDbSolid对象
function DyDrawSolid () {
    //清空当前显示内容
    mxOcx.NewFile();
    // 创建一个与用户交互取点的对象。
    var point1 = mxOcx.GetPoint(false,0,0,"点取第一点:");
    if(point1 == null)
    {
        return;
    }
    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getSecondPt.message = "点取第二点";
    getSecondPt.basePoint = point1;
    //开启皮筋效果
    getSecondPt.setUseBasePt(true);
    // 开始点第二点。
    if (getSecondPt.go() != 1)
        return;
    var getThirdPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getThirdPt.message = "点取第三点";
    getThirdPt.basePoint = getSecondPt.value();
    //开启皮筋效果
    getThirdPt.setUseBasePt(true);
    // 开始第三点。
    if (getThirdPt.go() != 1)
        return;
    var getFourthPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getFourthPt.message = "点取第四点";
    var spDrawData = getFourthPt.InitUserDraw("TestDynDrawMcDbSolid");
    spDrawData.SetPoint("vFirstPt", point1);
    spDrawData.SetPoint("vSecondPt", getSecondPt.value());
    spDrawData.SetPoint("vThirdPt", getThirdPt.value());
    getFourthPt.basePoint = getThirdPt.value();
    //不开启皮筋效果
    getFourthPt.setUseBasePt(false);
    //调用
    // 开始第四点。
    if (getFourthPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
}
function DynWorldDrawMcDbSolid( pCustomEntity, pWorldDraw, curPt) {
    //与用户交互在图面上提取一个点
    var vFirstPt = pCustomEntity.GetPoint("vFirstPt");
    var vSecondPt = pCustomEntity.GetPoint("vSecondPt");
    var vThirdPt = pCustomEntity.GetPoint("vThirdPt");
    pWorldDraw.DrawSolid(vFirstPt.x, vFirstPt.y,vSecondPt.x, vSecondPt.y, vThirdPt.x, vThirdPt.y,  curPt.x, curPt.y);
}
