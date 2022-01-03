function DrawEllipseArc() {
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //创建一个图层,名为"EllipseArcLayer"
    mxOcx.AddLayer("EllipseArcLayer");
    //设置当前图层为"EllipseArcLayer"
    mxOcx.LayerName = "EllipseArcLayer";
    //绘制椭圆弧
    //参数一为椭圆的中心点X坐标，参数二为椭圆的中心点Y坐标
    //参数三为椭圆的主轴向量X值，参数四为椭圆的主轴向量Y值
    //参数五为椭圆的副轴长度与主轴长度的比值
    //参数六为椭圆弧的开始角度，参数七为椭圆弧的结束角度
    mxOcx.DrawEllipseArc(0, 1600, 80, 30, 0.4, 30, 200);
    //----------------------------------------------------------------------------------------------------------
    mxOcx.AddLinetype("MyEllipseArcType", "6,-10");
    //设计当前线型为"MyEllipseArcType"
    mxOcx.LineType = ("MyEllipseArcType");
    mxOcx.DrawCADColor = 16711680;
    mxOcx.DrawEllipseArc(200, 1600, -80, 30, 0.5, 60, 300);
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}