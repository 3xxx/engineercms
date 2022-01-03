function DrawPathToHatch () {
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //创建一个图层,名为"HatchLayer"
    mxOcx.AddLayer("HatchLayer");
    //设置当前图层为"HatchLayer"
    mxOcx.LayerName = "HatchLayer";
    mxOcx.AddPatternDefinition("SOLID", "((0, 0,0, 0,0.125))");
    //把路径变成一个填充对象
    mxOcx.PatternDefinition = "SOLID";
    mxOcx.PathMoveTo (0, 30);
    //路径的一下个点
    mxOcx.PathLineTo (100, 50);
    //路径的一下个点
    mxOcx.PathLineTo (100, 30);
    //路径的一下个点
    mxOcx.PathLineTo(0, 30) ;
    mxOcx.AddLinetype( "MLineType1", "");
    mxOcx.LineType = "MLineType1";
    mxOcx.DrawColor = 255;
    mxOcx.DrawPathToHatch(1);
    //-----------------------------------------------------------------------------------------------
    mxOcx.PathMoveTo (30, 70);
    //路径的一下个点
    mxOcx.PathLineTo (200, 70);
    //路径的一下个点
    mxOcx.PathLineTo (200, 120);
    //路径的一下个点
    mxOcx.PathLineTo(30, 70) ;
    mxOcx.AddLinetype( "MLineType2", "");
    mxOcx.LineType = "MLineType2";
    mxOcx.DrawColor = 65280;
    mxOcx.DrawPathToHatch(1);
    //-----------------------------------------------------------------------------------------------
    mxOcx.PathMoveTo (110, 30);
    //路径的一下个点
    mxOcx.PathLineTo (150, 60);
    //路径的一下个点
    mxOcx.PathLineTo (150, 30);
    //路径的一下个点
    mxOcx.PathLineTo(110, 30) ;
    mxOcx.AddLinetype( "MLineType3", "");
    mxOcx.LineType = "MLineType3";
    mxOcx.DrawColor = 65535;
    mxOcx.DrawPathToHatch(1);
    //-----------------------------------------------------------------------------------------------
    mxOcx.PathMoveTo (200, 30);
    //路径的一下个点
    mxOcx.PathLineTo (300, 30);
    //路径的一下个点
    mxOcx.PathLineTo (300, 60);
    //路径的一下个点
    mxOcx.PathLineTo(200, 30) ;
    mxOcx.AddLinetype( "MLineType4", "");
    mxOcx.LineType = "MLineType4";
    //设置颜色
    mxOcx.DrawColor = 16711680;
    // 把路径闭合
    mxOcx.DrawPathToHatch(1);

    mxOcx.ZoomAll();

    mxOcx.UpdateDisplay();
}
