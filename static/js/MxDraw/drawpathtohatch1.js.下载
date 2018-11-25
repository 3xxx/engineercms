function DrawPathToHatch1() {
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //创建一个图层,名为"HatchLayer"
    mxOcx.AddLayer("HatchLayer1");
    //设置当前图层为"HatchLayer"
    mxOcx.LayerName = "HatchLayer1";

    mxOcx.AddPatternDefinition("SOLID", "((0, 0,0, 0,0.125))");
    //把路径变成一个填充对象
    mxOcx.PatternDefinition = "SOLID";

    //把路径的开始位置移动指定的点
    //参数一为点的X坐标 ，参数二为点的Y坐标，参数三为该点处开始宽度,对Polyline线的绘制有效
    //参数四为该点处结束宽度,对Polyline线的绘制有效
    //参数五为该点处的凸度,对Polyline线的绘制有效
    mxOcx.PathMoveToEx(0, 30, 0.3, 0.3, 0.3) ;
    //路径的一下个点
    mxOcx.PathLineTo( 100, 30);
    //路径的一下个点
    mxOcx.PathLineTo(100, 90);
    //路径的一下个点
    mxOcx.PathLineTo(0, 30);
    //设置颜色
    mxOcx.AddLinetype( "MLineType1", "");
    mxOcx.LineType = "MLineType1";
    mxOcx.DrawColor = 255;
    // 把路径闭合
    mxOcx.DrawPathToHatch(1);
//-----------------------------------------------------------------------------------------------
    mxOcx.PathMoveToEx(50, 120, 0.3, 0.3, 0.3) ;
    //路径的一下个点
    mxOcx.PathLineTo( 150, 120);
    //路径的一下个点
    mxOcx.PathLineTo(150, 180);
    //路径的一下个点
    mxOcx.PathLineTo(50, 120);
    //设置颜色
    mxOcx.AddLinetype( "MLineType2", "");
    mxOcx.LineType = "MLineType2";
    mxOcx.DrawColor = 65280;
    // 把路径闭合
    mxOcx.DrawPathToHatch(1);
//-----------------------------------------------------------------------------------------------
    mxOcx.PathMoveToEx(110, 50, 0.3, 0.3, 0.3) ;
    //路径的一下个点
    mxOcx.PathLineTo( 200, 30);
    //路径的一下个点
    mxOcx.PathLineTo(200, 90);
    //路径的一下个点
    mxOcx.PathLineTo(110,50);
    //设置颜色
    mxOcx.AddLinetype( "MLineType3", "");
    mxOcx.LineType = "MLineType3";
    mxOcx.DrawColor = 65535;
    // 把路径闭合
    mxOcx.DrawPathToHatch(1);
//-----------------------------------------------------------------------------------------------
    mxOcx.PathMoveToEx(250, 30, 0.3, 0.3, 0.3) ;
    //路径的一下个点
    mxOcx.PathLineTo( 300, 30);
    //路径的一下个点
    mxOcx.PathLineTo(300, 90);
    //路径的一下个点
    mxOcx.PathLineTo(250, 30);
    //设置颜色
    mxOcx.AddLinetype( "MLineType4", "");
    mxOcx.LineType = "MLineType4";
    mxOcx.DrawColor = 16711680;
    // 把路径闭合
    mxOcx.DrawPathToHatch(1);

    mxOcx.ZoomAll();

    mxOcx.UpdateDisplay();

}
