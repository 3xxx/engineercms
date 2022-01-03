//动态施放绘制事件回调函数指针
//动态拖放时的绘制事件
//参数一为当前鼠标位置X坐标；参数为二当前鼠标位置Y坐标；
// 参数三为IMxDrawWorldDraw对象，用于动态绘制；参数四为IMxDrawCustomEntity对象，动态绘制数据
function DoDynWorldDrawFun(dX,dY,pWorldDraw,pData) {
    //自定义实体的GUID标识符
    var sGuid = pData.Guid;
    //设置自定义事件的返回值
    mxOcx.SetEventRet(0);

    var curPt = mxOcx.NewPoint();
    curPt.x = dX;
    curPt.y = dY;




    if (sGuid == "TestMxCustomEntity2") {
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
    if (sGuid == "DrawComment") {

        DynWorldDrawComment(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    else if (sGuid == "DrawComment2") {

        DynWorldDrawComment2(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    else if (sGuid == "DrawCloudComment") {
        DynWorldDrawCloudComment(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    else if (sGuid == "DrawCircle") {
        var pt1 = pData.GetPoint("pt1");
        var dR = curPt.DistanceTo(pt1) * 0.5;
        var vec = curPt.SumVector(pt1);
        vec.Mult(0.5);
        pt1.Add(vec);
        pWorldDraw.DrawCircle(pt1.x, pt1.y, dR);
    }
    //三点画圆弧
    else if (sGuid == "DrawArc") {
        //与用户交互到在图上提取一个点
        var pt1 = pData.GetPoint("pt1");
        if (pt1 == null)
            return;
        var pt2 = pData.GetPoint("pt2");
        if (pt2 == null)
            return;
        //由圆弧上的三点绘制一个圆弧
        //参数一为开始点X坐标 ；参数二为开始点Y坐标 ；参数三为结束点X坐标
        //参数四为结束点Y坐标 ；参数五为圆弧上的一点X坐标 ；参数六为圆弧上的一点Y坐标
        pWorldDraw.DrawArc2(pt1.x, pt1.y, pt2.x, pt2.y, dX, dY);
    }
    else if (sGuid == "TestDynDraw") {
        //动态插入图块
        // 与用户交互到在图上提取一个点
        var firstPt = pData.GetPoint("pt1");
        if (firstPt == null)
            return;
        //与用户交互，从命令中取字符串
        var sBlkName = pData.GetString("BlkName");
        var secondPt = mxOcx.NewPoint();
        secondPt.x = dX;
        secondPt.y = dY;
        //计算this点减去pPoint，得到的向量
        var vec = firstPt.SumVector(secondPt);
        //返回向量顺时针转到与X轴的角度.[0,2PI)
        var dAng = vec.Angle();
        //绘制块引用对象
        //参数一为插入点的X坐标；参数二为插入点的Y坐标；
        //参数三为图块名；参数四为图块缩放比例；参数五为图块的旋转角度
        pWorldDraw.DrawBlockReference(firstPt.x, firstPt.y, sBlkName, 1.0, dAng * 180.0 / 3.14159265 + 90.0);
        //设置自定义事件的返回值
        mxOcx.SetEventRet(1);
    }
    //绘制块引用对象
    else if (sGuid == "TestDynDraw1") {
        var sBlkName = pData.GetString("BlkName");
        //绘制块引用对象
        pWorldDraw.DrawBlockReference(dX, dY, sBlkName, 1.0, 0.0);
    }
    //画矩形
    else if (sGuid == "TestDynDrawMatrix1") {
        DynWorldDrawMatrix1(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    //画带面积周长的圆
    else if (sGuid == "TestDynDrawCircle") {
        var pt1 = pData.GetPoint("pt1");
        var dR = curPt.DistanceTo(pt1);
        //计算面积（取绝对值）
        var arec = dR*dR*Math.PI;
        //计算周长
        var girth = 2*dR * Math.PI;
        //变换视区长度
        var dHeigth = mxOcx.ViewLongToDocCoord(15);
        //绘制一个单行文字
        //参数一为文字的位置的X坐标 ；参数二为文字的位置的Y坐标 ；参数三为文字内容字符串
        //参数四为文字高度；参数五为文字的旋转角度
        //参数六为文字的水平对齐方式,0=kTextLeft,1=kTextCenter,2=kTextRight
        //参数七文字的竖直对齐方式,1=kTextBottom,2=kTextVertMid,3=kTextTop
        pWorldDraw.DrawText(pt1.x, pt1.y, "面积：" + arec.toFixed(2), dHeigth, 0, 1, 1);
        pWorldDraw.DrawText(pt1.x, pt1.y, "周长：" + girth.toFixed(2), dHeigth, 0, 1, 3);
        pWorldDraw.DrawCircle(pt1.x, pt1.y, dR);
    }
    //画带面积周长的矩形
    else if (sGuid == "TestDynDrawMatrix") {
        DynWorldDrawMatrix(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    //画直线
    else if (sGuid == "TestDynDrawLine") {
        DynWorldDrawLine(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    //画虚线
    else if (sGuid == "TestDynDrawMyLine") {
        DynWorldDrawMyLine(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    //绘制带颜色线宽的直线
    else if (sGuid == "TestDynDrawWidthline") {
        DynWorldDrawWidthline(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    //绘制点
    else if (sGuid == "TestDynDrawPoint") {
        var p = pWorldDraw.DrawPoint(dX - 10, dY - 10);
        mxOcx.SetEventRet(1);
    }
    //绘制椭圆
    else if (sGuid == "TestDynDrawEllipse") {
        DynWorldDrawEllipse(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    //绘制McDbSolid对象
    else if (sGuid == "TestDynDrawMcDbSolid") {
        DynWorldDrawMcDbSolid(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }
    //绘制文字
    if (sGuid == "DrawText") {

        DynWorldDrawText(pData, pWorldDraw, curPt);
        mxOcx.SetEventRet(1);
    }








    var dHeigth = mxOcx.ViewLongToDocCoord(5);
    //自定义实体的GUID标识符
    switch (pData.Guid) {
        case "DynDrawEnt4":
        {
            var vThirdPt = pData.GetPoint("vThird");
            pWorldDraw.DrawLine(vThirdPt.x, vThirdPt.y, dX, dY);
            pWorldDraw.DrawText(dX, dY, "(" + dX.toFixed(1) + "," + dY.toFixed(1) + ")", dHeigth, 0, 1, 1);
            dX = vThirdPt.x;
            dY = vThirdPt.y;
        }
        case "DynDrawEnt3":
        {
            var vSecondPt = pData.GetPoint("vSecond");
            pWorldDraw.DrawLine(vSecondPt.x, vSecondPt.y, dX, dY);
            pWorldDraw.DrawText(dX, dY, "(" + dX.toFixed(1) + "," + dY.toFixed(1) + ")", dHeigth, 0, 1, 1);
            dX = vSecondPt.x;
            dY = vSecondPt.y;
        }
        case "DynDrawEnt2":
        {
            var vFirstPt = pData.GetPoint("vFirst");
            pWorldDraw.DrawLine(vFirstPt.x, vFirstPt.y, dX, dY);
            pWorldDraw.DrawText(dX, dY, "(" + dX.toFixed(1) + "," + dY.toFixed(1) + ")", dHeigth, 0, 1, 1);
            dX = vFirstPt.x;
            dY = vFirstPt.y;
        }
        case "DynDrawEnt1":
        {
            pWorldDraw.DrawText(dX, dY, "(" + dX.toFixed(1) + "," + dY.toFixed(1) + ")", dHeigth, 0, 1, 1);
        }
    }
    //设置自定义事件的返回值
    mxOcx.SetEventRet(0);


}
