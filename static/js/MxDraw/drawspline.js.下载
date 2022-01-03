function DrawSpline() {
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //创建一个图层,名为"SplinelineLayer"
    mxOcx.AddLayer("SplinelineLayer");
    //设置当前图层为"SplinelineLayer"
    mxOcx.LayerName = "SplinelineLayer";

    //《---------------------------------------》
    //用绘制一个实线的样条线
    //定义一个路径的开始点
    mxOcx.PathMoveTo(0, 1700);
    //路径的一下个点
    mxOcx.PathLineTo(100, 1700);
    //路径的一下个点
    mxOcx.PathLineTo(100, 1800);
    //路径的一下个点
    mxOcx.PathLineTo(0, 1800);
    //把路径拟合成一个样线
    mxOcx.DrawPathToSpline();

    //《---------------------------------------》
    //用绘制一个虚线的样条线
    mxOcx.AddLinetype("MySplineType", "30,-10");
    mxOcx.LineType = ("MySplineType");
    mxOcx.DrawCADColor = 255;
    mxOcx.LineWidth = 5;
    //定义一个路径的开始点
    mxOcx.PathMoveTo(200, 1700);
    //路径的一下个点
    mxOcx.PathLineTo(300, 1700);
    //路径的一下个点
    mxOcx.PathLineTo(300, 1800);
    //路径的一下个点
    mxOcx.PathLineTo(400, 1800);
    mxOcx.DrawPathToSpline();
    mxOcx.ZoomAll();
    mxOcx.UpdateDisplay();

}

