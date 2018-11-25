
document.write("<script type='text/javascript'  src='dydrawline.js'></script>");
document.write("<script type='text/javascript'  src='dydrawmyline.js'></script>");
document.write("<script type='text/javascript'  src='dydrawwidthline.js'></script>");
document.write("<script type='text/javascript'  src='dydrawcircle.js'></script>");
document.write("<script type='text/javascript'  src='dydrawpoint.js'></script>");
document.write("<script type='text/javascript'  src='dydrawarc.js'></script>");
document.write("<script type='text/javascript'  src='dydrawpathtopolyline.js'></script>");
document.write("<script type='text/javascript'  src='dydrawpolyline.js'></script>");
document.write("<script type='text/javascript'  src='dydrawspline.js'></script>");
document.write("<script type='text/javascript'  src='dydrawsolid.js'></script>");
document.write("<script type='text/javascript'  src='dydyninsert.js'></script>");
document.write("<script type='text/javascript'  src='dyinsertcustomentity.js'></script>");
document.write("<script type='text/javascript'  src='dyndrawmatrix.js'></script>");
document.write("<script type='text/javascript'  src='dydrawtext.js'></script>");
document.write("<script type='text/javascript'  src='myinsertcustomentity.js'></script>");
document.write("<script type='text/javascript'  src='dydocloudcirclecomment.js'></script>");
document.write("<script type='text/javascript'  src='dydocloudlinecomment.js'></script>");
document.write("<script type='text/javascript'  src='dydocomment.js'></script>");
document.write("<script type='text/javascript'  src='dydocomment3.js'></script>");
document.write("<script type='text/javascript'  src='dydocirclecomment.js'></script>");
// 交互绘制批注时的动态绘制1
function DynWorldDrawComment( pCustomEntity, pWorldDraw, curPt) {
    // 得到绘制参数.

    var sText       = pCustomEntity.GetString("Text");
    var dTextHeight = pCustomEntity.GetDouble("TextHeight");
    var iEdgeNumber  = pCustomEntity.GetLong("EdgeNumber");
    var dShapRadius = pCustomEntity.GetDouble("ShapRadius");
    // 创建一个批注对象.
    var comment = mxOcx.NewEntity("IMxDrawComment");
    comment.Text = sText;
    comment.TextHeight = dTextHeight;
    comment.EdgeNumber = iEdgeNumber;
    comment.ShapRadius = dShapRadius;
    comment.basePoint = curPt;
    comment.Position = curPt;
    pWorldDraw.TextStyle = "MyCommentFont";
    // 动态绘制.
    pWorldDraw.DrawEntity(comment);
}

// 交互绘制批注时的动态绘制1
function DynWorldDrawCloudComment(pCustomEntity, pWorldDraw, curPt) {
    // 得到绘制参数.
    var sText = pCustomEntity.GetString("Text");
    var dTextHeight = pCustomEntity.GetDouble("TextHeight");
    var iEdgeNumber = pCustomEntity.GetLong("EdgeNumber");
    var dShapRadius = pCustomEntity.GetDouble("ShapRadius");
    var basePoint = pCustomEntity.GetPoint("BasePoint");
    var frstPt = pCustomEntity.GetPoint("frstPt");
    // 创建一个批注对象.
    var comment = mxOcx.NewEntity("IMxDrawComment");
    comment.Text = sText;
    comment.TextHeight = dTextHeight;
    comment.EdgeNumber = iEdgeNumber;
    comment.ShapRadius = dShapRadius;
    comment.basePoint = basePoint;
    comment.Position = curPt;
    comment.BoundingBoxType = 5;
    comment.ArrowType = 2;
    comment.ShapRadius = 40;
    var points = mxOcx.NewComObject("IMxDrawPoints");
    //向点数组中增加一个点坐标
    points.Add2(frstPt);
    points.Add2(basePoint);
    //标注框坐标数据
    comment.Bounding = points;
    var tmpp = mxOcx.NewResbuf();
    tmpp.AddLong(1);
    comment.SetProp("isAlwaysShowLineWeight", tmpp);
    //线重
    comment.Lineweight = 25;
    //实体的CAD颜色索引值属性
    comment.colorIndex = 1;
    // 设置文字样式。
    pWorldDraw.TextStyle = "MyCommentFont";
    // 动态绘制.
    pWorldDraw.DrawEntity(comment);
}
// 交互绘制批注时的动态绘制1
function  DynWorldDrawComment2( pCustomEntity,pWorldDraw,  curPt) {
    // 得到绘制参数.

    var sText = pCustomEntity.GetString("Text");
    var dTextHeight = pCustomEntity.GetDouble("TextHeight");
    var iEdgeNumber = pCustomEntity.GetLong("EdgeNumber");
    var dShapRadius = pCustomEntity.GetDouble("ShapRadius");
    var  basePoint = pCustomEntity.GetPoint("BasePoint");
    // 创建一个批注对象.
    var comment = mxOcx.NewEntity("IMxDrawComment");
    //标注文字
    comment.Text        = sText;
    //标注文字高度
    comment.TextHeight  = dTextHeight;
    //标注位置提示多边形的边数. 小于2不绘制，=2绘制圆，大于2绘制多边形
    comment.EdgeNumber  = iEdgeNumber;
    //标注位置提示多边形的半径
    comment.ShapRadius  = dShapRadius;
    //标注基点
    comment.basePoint   = basePoint;
    //标注位置点
    comment.Position = curPt;
    // 设置文字样式。
    pWorldDraw.TextStyle = "MyCommentFont";
    // 动态绘制.
    pWorldDraw.DrawEntity(comment);
}

// 动态绘制批注函数.

// 动态绘制批注函数.







//绘制椭圆
function  DyDrawEllipse()
{
    //清空当前显示内容
    mxOcx.NewFile();
    // 创建一个与用户交互取点的对象。
    var point1 = mxOcx.GetPoint(false,0,0,"点取主轴开始点:");
    if(point1 == null)
    {
        return;
    }
    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getSecondPt.message = "点取主轴结束点";
    getSecondPt.basePoint = point1;
    //开启皮筋效果
    getSecondPt.setUseBasePt(true);
    // 开始点取主轴结束点。
    if (getSecondPt.go() != 1)
        return;
    var getThirdPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getThirdPt.message = "点取主轴结束点";
    getThirdPt.setUseBasePt(true);
    //计算中心点位置
    var centerPt = mxOcx.NewPoint();
    {
        centerPt.x = point1.x + (getSecondPt.value().x - point1.x) / 2;
        centerPt.y = point1.y + (getSecondPt.value().y - point1.y) / 2;
    }
    //调用
    var spDrawData = getThirdPt.InitUserDraw("TestDynDrawEllipse");
    spDrawData.SetPoint("vCenterPt", centerPt);
    spDrawData.SetPoint("vSecondPt", getSecondPt.value());
    getThirdPt.basePoint = centerPt;
    if (getThirdPt.go() != 1)
        return;
    var ret = spDrawData.Draw();
}
function DynWorldDrawEllipse( pCustomEntity, pWorldDraw, curPt) {
    //与用户交互在图面上提取一个点
    var vCenterPt = pCustomEntity.GetPoint("vCenterPt");
    var vSecondPt = pCustomEntity.GetPoint("vSecondPt");

    var dMainLen = vCenterPt.DistanceTo(vSecondPt);
    var dVice = vCenterPt.DistanceTo(curPt);

    var vVec = vSecondPt.SumVector(vCenterPt);
    var vTemp = mxOcx.NewEntity("IMxDrawVector3d");

    vTemp.x = vVec.x;
    vTemp.y = vVec.y;
    vTemp.Normalize();
    vTemp = vTemp.rotate(PI / 2);
    vTemp.Mult(dVice);

    var stu = pWorldDraw.DrawEllipse(vCenterPt.x, vCenterPt.y, vTemp.x, vTemp.y, dVice / dMainLen);
}
//绘制椭圆弧
function DyDrawEllipseArc() {
    //清空当前显示内容
    mxOcx.NewFile();
    // 与用户交互到在图上提取一个点
    var point1 = mxOcx.GetPoint(false,0,0,"\n 点取椭圆的中心点:");
    if(point1 == null)
    {
        return;
    }
    // 与用户交互到在图上提取一个点
    var point2 = mxOcx.GetPoint(true,point1.x,point1.y,"\n 点取椭圆的主轴向量:");
    if(point2 == null)
    {
        return;
    }
    //绘制椭圆弧
    //参数一为椭圆的中心点X坐标，参数二为椭圆的中心点Y坐标
    //参数三为椭圆的主轴向量X值，参数四为椭圆的主轴向量Y值
    //参数五为椭圆的副轴长度与主轴长度的比值
    //参数六为椭圆弧的开始角度，参数七为椭圆弧的结束角度
    mxOcx.DrawEllipseArc(point1.x,point1.y,point2.x,point2.y,0.3,60 * 180.0 / 3.14159265,45 * 180.0 / 3.14159265)
}