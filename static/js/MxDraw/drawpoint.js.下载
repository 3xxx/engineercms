function DrawPoint() {
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    // 设置点样式，具体的值，可以看autocad帮助
    //pdmode是一个控制point的形式的系统变量，当pdmode=0时是可见的一个点，
    // 当pdmode=1时是一个不可见的点，但仍存在。pdmode=2时是一个十字，pdmode=3时是一个叉子
    //将值指定为 32、64 或 96，除了绘制通过点的图形外，还可以选择在点的周围绘制形。
    mxOcx.SetSysVarLong("PDMODE", 0);
    // 设置点大小，0.0表示它随屏幕
    mxOcx.SetSysVarDouble("PDSIZE", 15.0);
    mxOcx.DrawPoint(0,50);
    //-----------------------------------------------------------------------------------------------
    mxOcx.SetSysVarLong("PDMODE", 67);
    // 设置点大小，0.0表示它随屏幕
    mxOcx.SetSysVarDouble("PDSIZE", 15.0);
    mxOcx.DrawPoint(50,50);
    //-----------------------------------------------------------------------------------------------
    mxOcx.SetSysVarLong("PDMODE", 3);
    // 设置点大小，0.0表示它随屏幕
    mxOcx.SetSysVarDouble("PDSIZE", 15.0);
    mxOcx.DrawPoint(100,50);
    //-----------------------------------------------------------------------------------------------
    mxOcx.SetSysVarLong("PDMODE", 2);
    // 设置点大小，0.0表示它随屏幕
    mxOcx.SetSysVarDouble("PDSIZE", 15.0);
    mxOcx.DrawPoint(150,50);
    mxOcx.ZoomAll();
    mxOcx.UpdateDisplay();

}
