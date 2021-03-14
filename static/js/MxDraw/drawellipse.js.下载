function  DrawEllipse()
{
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //创建一个图层,名为"LineLayer"
    mxOcx.AddLayer("LineLayer");
    //设置当前图层为"LineLayer"
    mxOcx.LayerName = "LineLayer";
    // 直接绘制一个椭圆
    mxOcx.DrawEllipse(0, 1400, 100, 0, 0.7);
    //------------------------------------------------------------------------------------------------------------
    //绘制一个倾斜的椭圆
    mxOcx.AddLinetype("MyEllipseType", "6,-10");
    //设计当前线型为"MyEllipseType"
    mxOcx.LineType = ("MyEllipseType");
    mxOcx.DrawCADColor = 16711680;
    mxOcx.DrawEllipse(200, 1400, 80, 30, 0.5);
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();

}
