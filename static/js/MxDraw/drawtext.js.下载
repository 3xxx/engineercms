function DrawText()
{
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawColor = 0;
    //创建一个图层,名为"TextLayer"
    mxOcx.AddLayer("TextLayer");
    //设置当前图层为"TextLayer"
    mxOcx.LayerName = "TextLayer";
    //绘制一个单行文字
    //参数一为文字的位置的X坐标 ；参数二为文字的位置的Y坐标 ；参数三为文字内容字符串
    //参数四为文字高度；参数五为文字的旋转角度
    //参数六为文字的水平对齐方式,0=kTextLeft,1=kTextCenter,2=kTextRight
    //参数七文字的竖直对齐方式,1=kTextBottom,2=kTextVertMid,3=kTextTop
    //-------------------------------------------------------------------------------------------------
    //写一个文字,0,1是左对齐.
    mxOcx.DrawColor = 65280;
    mxOcx.DrawText(0, 1900, "梦想绘图控件3.0 www.mxdraw.com", 100, 0, 0, 1);
    //---------------------------------------------------------------------------------------------------
    //写一个文字,2,1是右下对齐.
    mxOcx.DrawColor = 4556677;
    mxOcx.DrawText(3000, 2100, "梦想绘图控件3.0 www.mxdraw.com", 100, -20, 2, 1);
    //--------------------------------------------------------------------------------------------------
    //按指定样式绘制文字
    mxOcx.AddTextStyle1("MyTextStyle", "italicc.shx", "gbcbig.shx", 0.7);
    mxOcx.TextStyle = "MyTextStyle";
    mxOcx.DrawColor = 255;
    mxOcx.DrawText(0, 3000, "梦想绘图控件3.0 www.mxdraw.com", 100, 0, 0, 1);
    //--------------------------------------------------------------------------------------------------
    //写一个TureType字体
    mxOcx.AddTextStyle2("MyTrueTypeStyle", "黑体", 0.7);
    mxOcx.TextStyle = "MyTrueTypeStyle";
    mxOcx.DrawText(0, 3150, "梦想绘图控件3.0 www.mxdraw.com", 100, 0, 0, 1);
    mxOcx.ZoomAll();
    mxOcx.UpdateDisplay();
}
