// 自定义实体绘制函数
function ExplodeFun(pCustomEntity, pWorldDraw, txt) {
    var sGuid = pCustomEntity.Guid;
    if (sGuid == "TestMxCustomEntity") {
        if (!pCustomEntity.IsHave("ept"))
            return;
        var stp = pCustomEntity.GetPoint("spt");
        if (stp == null)
            return;
        var ept = pCustomEntity.GetPoint("ept");
        if (ept == null)
            return;
        var mxUtility = mxOcx.NewUtility();
        var vec = ept.SumVector(stp);
        vec.Mult(0.5);
        var midPt = mxOcx.NewPoint();
        midPt.x = stp.x;
        midPt.y = stp.y;
        midPt.Add(vec);
        var dAng = vec.Angle();
        //计算一个标注角度，使用文字对象始终头朝上
        dAng = mxUtility.GetDimAngle(dAng);
        var dDis = 0.0;
        dDis = stp.DistanceTo(ept);
        var sTxt = "L=" + formatNumber(dDis, '#.##');
        dAng = dAng * 180.0 / 3.14159265;
        vec.RotateByXyPlan(3.14159265 / 2.0);
        vec.Normalize();
        vec.Mult(10);
        stp.Add(vec);
        ept.Add(vec);
        pWorldDraw.DrawLine(stp.x, stp.y, ept.x, ept.y);
        vec.Mult(2);
        stp.Sum(vec);
        ept.Sum(vec);
        pWorldDraw.DrawLine(stp.x, stp.y, ept.x, ept.y);
        pWorldDraw.SetColorIndex(1);
        pWorldDraw.DrawText(midPt.x, midPt.y, sTxt, 5, dAng,
            1, 2);
        mxOcx.SetEventRet(1);

    }
    if (sGuid == "TestMxCustomEntity1") {
        if (!pCustomEntity.IsHave("First"))
            return;
        var stp = pCustomEntity.GetPoint("First");
        var ept = pCustomEntity.GetPoint("BasePoint");
        var dimpt = pCustomEntity.GetPoint("DimPoint");
        var txt = pCustomEntity.GetString("Text");
        var textH = pCustomEntity.GetDouble("TextHeight");
        var edgeNum = pCustomEntity.GetLong("EdgeNumber");
        var shapRadius = pCustomEntity.GetDouble("ShapRadius");
        var isCircle = pCustomEntity.GetLong("isCircle");
        var comment = mxOcx.NewEntity("IMxDrawComment");
        comment.Text = txt;
        comment.TextHeight = textH;
        comment.EdgeNumber = edgeNum;
        comment.ShapRadius = shapRadius;
        comment.basePoint = ept;
        comment.Position = dimpt;
        pWorldDraw.TextStyle = "MyCommentFont";
        // 动态绘制.
        pWorldDraw.DrawEntity(comment);
        // 绘制矩形框.
        if (isCircle) {
            var dR = stp.DistanceTo(ept) * 0.5;
            var vec = stp.SumVector(ept);
            vec.Mult(0.5);
            ept.Add(vec);
            pWorldDraw.DrawCircle(ept.x, ept.y, dR);
        }
        else {
            pWorldDraw.DrawLine(stp.x, stp.y, stp.x, ept.y);
            pWorldDraw.DrawLine(stp.x, ept.y, ept.x, ept.y);
            pWorldDraw.DrawLine(ept.x, ept.y, ept.x, stp.y);
            pWorldDraw.DrawLine(ept.x, stp.y, stp.x, stp.y);
        }
        mxOcx.SetEventRet(1);
    }
    if (sGuid == "TestMxCustomEntity2") {
        if (!pCustomEntity.IsHave("First"))
            return;
        var stp = pCustomEntity.GetPoint("First");
        var ept = pCustomEntity.GetPoint("BasePoint");
        var dimpt = pCustomEntity.GetPoint("DimPoint");
        var txt = pCustomEntity.GetString("Text");
        var textH = pCustomEntity.GetDouble("TextHeight");
        var edgeNum = pCustomEntity.GetLong("EdgeNumber");
        var shapRadius = pCustomEntity.GetDouble("ShapRadius");
        var isCircle = pCustomEntity.GetLong("isCircle");
        // 创建一个批注对象.
        var comment = mxOcx.NewEntity("IMxDrawComment");
        //标注文字
        comment.Text = txt;
        //标注文字高度
        comment.TextHeight = textH;
        //标注位置提示多边形的边数. 小于2不绘制，=2绘制圆，大于2绘制多边形
        comment.EdgeNumber = edgeNum;
        //标注位置提示多边形的半径
        comment.ShapRadius = shapRadius;
        //标注基点
        comment.basePoint = ept;
        //标注位置点
        comment.Position = dimpt;
        // 设置文字样式
        pWorldDraw.TextStyle = "MyCommentFont";
        // 动态绘制.
        pWorldDraw.DrawEntity(comment);

        if (isCircle) {
            var dR = stp.DistanceTo(ept) * 0.5;
            var vec = stp.SumVector(ept);
            vec.Mult(0.5);
            ept.Add(vec);
            //绘制一个圆
            //参数一为圆的中心X值 ，参数二为圆的中心Y值,
            //参数三为半径
            pWorldDraw.DrawCircle(ept.x, ept.y, dR);
        }
        // 绘制矩形框.
        else {
            //绘制一个直线
            //参数一直线的开始点x坐标，参数二直线的开始点y坐标，参数三直线的结束点x坐标，参数四直线的结束点y坐标
            pWorldDraw.DrawLine(stp.x, stp.y, stp.x, ept.y);
            pWorldDraw.DrawLine(stp.x, ept.y, ept.x, ept.y);
            pWorldDraw.DrawLine(ept.x, ept.y, ept.x, stp.y);
            pWorldDraw.DrawLine(ept.x, stp.y, stp.x, stp.y);
        }
        mxOcx.SetEventRet(1);
    }


    var dHeigth = mxOcx.ViewLongToDocCoord(5);




    if(sGuid == "DrawEnt"){
        {
            if (!pCustomEntity.IsHave("vFour"))
                return;

            if (!pCustomEntity.IsHave("vThird"))
                return;

            if (!pCustomEntity.IsHave("vSecond"))
                return;

            if (!pCustomEntity.IsHave("vFirst"))
                return;
        }
        {
            var vFourPt = pCustomEntity.GetPoint("vFour");
            var dX = vFourPt.x;
            var dY = vFourPt.y;

            var vThirdPt = pCustomEntity.GetPoint("vThird");
            pWorldDraw.DrawLine(vThirdPt.x, vThirdPt.y, dX, dY);
            pWorldDraw.DrawText(dX, dY, "(" + dX.toFixed(1) + "," + dY.toFixed(1) + ")", dHeigth, 0, 1, 1);
            var dX = vThirdPt.x;
            var dY = vThirdPt.y;

            var vSecondPt = pCustomEntity.GetPoint("vSecond");
            pWorldDraw.DrawLine(vSecondPt.x, vSecondPt.y, dX, dY);
            pWorldDraw.DrawText(dX, dY, "(" + dX.toFixed(1) + "," + dY.toFixed(1) + ")", dHeigth, 0, 1, 1);
            dX = vSecondPt.x;
            dY = vSecondPt.y;

            var vFirstPt = pCustomEntity.GetPoint("vFirst");
            pWorldDraw.DrawLine(vFirstPt.x, vFirstPt.y, dX, dY);
            pWorldDraw.DrawText(dX, dY, "(" + dX.toFixed(1) + "," + dY.toFixed(1) + ")", dHeigth, 0, 1, 1);
            dX = vFirstPt.x;
            dY = vFirstPt.y;

            pWorldDraw.DrawText(dX, dY, "(" + dX.toFixed(1) + "," + dY.toFixed(1) + ")", dHeigth, 0, 1, 1);
        }
        mxOcx.SetEventRet(1);
    }

}





