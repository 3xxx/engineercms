function DrawPolyline() {
    //清空当前显示内容
    mxOcx.NewFile();
    //Polyline 又称多义线，就是可以绘制多条首尾相连的线,也可以是圆弧相连，同是可以指定线的开始和结束宽。
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //创建一个图层,名为"PolylineLayer"
    mxOcx.AddLayer("PolylineLayer"); ;
    //设置当前图层为"PolylineLayer"
    mxOcx.LayerName = "PolylineLayer";
    //《---------------------------------------》
    //用Plyline显示绘制矩形框
    //定义一个路径的开始点
    mxOcx.PathMoveTo(0, 300);
    //路径的一下个点
    mxOcx.PathLineTo(100, 300);
    //路径的一下个点
    mxOcx.PathLineTo(100, 400);
    //路径的一下个点
    mxOcx.PathLineTo(0, 400);
    //把路径设置成闭合
    mxOcx.PathMakeClosed();
    //把路径闭合，生成一个矩形框
    mxOcx.DrawPathToPolyline();
    //《---------------------------------------》
    //绘制一个有宽度，带线型的矩形框
    //定义一个路径的开始点
    mxOcx.PathMoveTo(200, 300);
    //路径的一下个点
    mxOcx.PathLineTo(300, 300);
    //路径的一下个点
    mxOcx.PathLineTo(300, 400);
    //路径的一下个点
    mxOcx.PathLineTo(200, 400);
    //把路径设置成闭合
    mxOcx.PathMakeClosed();
    mxOcx.AddLinetype("MyLineType3", "30,-10");
    mxOcx.LineType = "MyLineType3";
    mxOcx.DrawCADColorIndex = 255;
    mxOcx.LineWidth = 5;
    mxOcx.DrawPathToPolyline();
    //《---------------------------------------》
    //绘制一个有宽度的矩形框
    //定义一个路径的开始点
    mxOcx.PathMoveTo(400, 300);
    //路径的一下个点
    mxOcx.PathLineTo(500, 300);
    //路径的一下个点
    mxOcx.PathLineTo(500, 400);
    //路径的一下个点
    mxOcx.PathLineTo(400, 400);
    //把路径设置成闭合
    mxOcx.PathMakeClosed();
    mxOcx.LineType = ("");
    mxOcx.DrawCADColor = 65280;
    mxOcx.LineWidth = 10;
    mxOcx.DrawPathToPolyline();
    //《---------------------------------------》
    //绘制一个有宽度的矩形框
    //定义一个路径的开始点
    mxOcx.PathMoveTo(400, 300);
    //路径的一下个点
    mxOcx.PathLineTo(500, 300);
    //路径的一下个点
    mxOcx.PathLineTo(500, 400);
    //路径的一下个点
    mxOcx.PathLineTo(400, 400);
    //把路径设置成闭合
    mxOcx.PathMakeClosed();
    mxOcx.LineType = "";

    mxOcx.DrawCADColor = 65280;

    mxOcx.LineWidth = 10;

    mxOcx.DrawPathToPolyline();

    //《---------------------------------------》
    //绘制一个有宽度的，不闭合的矩形框
    //定义一个路径的开始点
    mxOcx.PathMoveTo(600, 300);

    //路径的一下个点
    mxOcx.PathLineTo(700, 300);

    //路径的一下个点
    mxOcx.PathLineTo(700, 400);

    //路径的一下个点
    mxOcx.PathLineTo(600, 400);

    //把路径设置成闭合
    mxOcx.LineType = "";

    mxOcx.DrawCADColor = 65280;

    mxOcx.LineWidth = 10;

    mxOcx.DrawPathToPolyline();


    //《---------------------------------------》
    //绘制一个有宽度的，三角形
    //定义一个路径的开始点
    mxOcx.PathMoveTo(800, 300);

    //路径的一下个点
    mxOcx.PathLineTo(900, 300);

    //路径的一下个点
    mxOcx.PathLineTo(900, 400);

    mxOcx.PathMakeClosed();

    //把路径设置成闭合
    mxOcx.LineType = "";

    mxOcx.DrawCADColor = 65280;

    mxOcx.LineWidth = 10;

    mxOcx.DrawPathToPolyline();


    //《---------------------------------------》
    //绘制一个箭头
    mxOcx.PathMoveToEx(1000, 300, 10, 10, 0);

    mxOcx.PathLineToEx(1000, 500, 30, 0, 0);

    mxOcx.PathLineTo(1000, 580);

    mxOcx.DrawPathToPolyline();


    //《---------------------------------------》
    //绘制一个虚线箭头
    mxOcx.PathMoveToEx(1200, 300, 10, 10, 0);

    mxOcx.PathLineToEx(1200, 500, 30, 0, 0);

    mxOcx.PathLineTo(1200, 580);

    mxOcx.DrawCADColor = 65535;

    mxOcx.LineType = ("MyLineType3");

    mxOcx.DrawPathToPolyline();

    //《---------------------------------------》
    //绘制一个圆弧箭头,-0.34是圆弧的凸度.
    //axMxDrawX1.PathMoveToEx 1400, 300, 10, 10, -0.34);
    mxOcx.PathMoveToEx(1400, 300, 10, 10, -0.34);
    mxOcx.PathLineToEx(1400, 500, 30, 0, 0.34);
    mxOcx.PathLineTo(1400, 600);
    mxOcx.DrawCADColor = 16776960;
    mxOcx.LineType = "MyLineType3";
    mxOcx.DrawPathToPolyline();
    //《---------------------------------------》
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}
