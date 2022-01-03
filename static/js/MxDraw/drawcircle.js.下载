function DrawCircle() {
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //创建一个图层,名为"CircleLayer"
    mxOcx.AddLayer("CircleLayer");
    //设置当前图层为"CircleLayer"
    mxOcx.LayerName = "CircleLayer";
    //《---------------------------------------》
    //绘制一个实线黑色的圆
    mxOcx.DrawCircle(0, 800, 70);
    //《---------------------------------------》
    //绘制一个带颜色的圆
    mxOcx.AddLinetype("CircleLineType", "");
    //设计当前线型为"CircleLineType"
    mxOcx.LineType = ("CircleLineType");
    mxOcx.DrawColor = 255;
    mxOcx.DrawCircle(200, 800, 70);
    //《---------------------------------------》
    //绘制一个虚线线的圆
    mxOcx.AddLinetype("CircleLineType", "30,-5,7,-7");
    //设计当前线型为"CircleLineType"
    mxOcx.LineType = ("CircleLineType");
    mxOcx.DrawColor = 255;
    mxOcx.DrawCircle(400, 800, 70);
    //《---------------------------------------》
    //绘制一个带有宽度的圆
    mxOcx.AddLinetype("CircleLineType", "30,-5,7,-7");
    //设计当前线型为"CircleLineType"
    mxOcx.LineType = ("CircleLineType");
    mxOcx.LineWidth = 40;
    mxOcx.DrawColor = 255;
    mxOcx.DrawCircle(600, 800, 70);
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}