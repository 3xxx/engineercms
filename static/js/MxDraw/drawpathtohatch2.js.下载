function DrawPathToHatch2() {
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //创建一个图层,名为"HatchLayer"
    mxOcx.AddLayer("HatchLayer2");
    //设置当前图层为"HatchLayer"
    mxOcx.LayerName = "HatchLayer2";

    //绘制一个有图案的填充
    //angle, x-origin,y-origin, delta-x,delta-y,dash-1,dash-2, …
    //45 = angle 是图案线角度.
    //0 = x-origin 是第一个填充线经过的点位置X坐标
    //0 = y-origin 是第一个填充线经过的点位置Y坐标
    //0 = delta-x   是下一个填充线相对前一个线的X方向偏移
    //0.125 = delta-y   是下一个填充线相对前一个线的Y方向偏移
    mxOcx.AddPatternDefinition("MyHatchPattern1", "((45, 0,0, 0,0.125))");
    //把路径变成一个填充对象
    mxOcx.PatternDefinition = "MyHatchPattern1";
    //把路径的开始位置移动指定的点
    //参数一为点的X坐标 ，参数二为点的Y坐标，参数三为该点处开始宽度,对Polyline线的绘制有效
    //参数四为该点处结束宽度,对Polyline线的绘制有效
    //参数五为该点处的凸度,对Polyline线的绘制有效
    mxOcx.PathMoveToEx(0, 30, 0.3, 0.3, 0.3) ;
    //路径的一下个点
    mxOcx.PathLineTo( 100, 30);
    //路径的一下个点
    mxOcx.PathLineTo(100, 60);
    //路径的一下个点
    mxOcx.PathLineTo(0, 30);
    //把路径变成一个填充,80,是填充图案的缩放比例.
    mxOcx.DrawPathToHatch(100);
    //----------------------------------------------------------------
    //绘制图案填充2
    mxOcx.AddPatternDefinition("MyHatchPattern2", "((0,0,0,0,8)(90,0,0,8,8,8,-8))");
    //把路径变成一个填充对象
    mxOcx.PatternDefinition = "MyHatchPattern2";
    //把路径的开始位置移动指定的点
    //参数一为点的X坐标 ，参数二为点的Y坐标，参数三为该点处开始宽度,对Polyline线的绘制有效
    //参数四为该点处结束宽度,对Polyline线的绘制有效
    //参数五为该点处的凸度,对Polyline线的绘制有效
    mxOcx.PathMoveToEx(100, 30, 0, 0, 0.3) ;
    //路径的一下个点
    mxOcx.PathLineTo( 200, 30);
    //路径的一下个点
    mxOcx.PathLineTo(200, 60);
    //路径的一下个点
    mxOcx.PathLineTo(100, 30);
    //把路径变成一个填充,80,是填充图案的缩放比例.
    mxOcx.DrawPathToHatch(1);
    mxOcx.ZoomAll();

    mxOcx.UpdateDisplay();
}