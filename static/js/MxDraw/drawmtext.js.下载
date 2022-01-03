function DrawMText()
{
    //清空当前显示内容
    mxOcx.NewFile();
    //把颜色改回黑白色
    mxOcx.DrawColor = 0;
    //绘制多行文字
    //参数一为多行文字位置的X坐标  ；参数二为多行文字位置的Y坐标  ；参数三为多行文字的内容
    //参数四为多行文字的高度 ；参数五为多行文字的显示宽度
    //参数六为多行文字的旋转角度
    //参数七为多行文字的对齐方式
    mxOcx.DrawColor = 255;
    mxOcx.DrawMText( 100, 100,"cad绘图\\P梦想软件",15,500,0.0,7);
    mxOcx.DrawColor = 16711680;
    mxOcx.DrawMText(0,0,"\\A2;平方方法 :m{\\H0.3x;\\S 2^;}",15,100,0.0,7);
    //按指定样式绘制文字
    mxOcx.AddTextStyle1("MyTextStyle", "italicc.shx", "gbcbig.shx", 0.7);
    mxOcx.TextStyle = "MyTextStyle";
    mxOcx.DrawColor = 65280;
    mxOcx.DrawMText(0, 50,"\\A2;平方方法 :m{\\H0.3x;\\S 2^;}",15,100,0.0,7);
    //--------------------------------------------------------------------------------------------------
    //写一个TureType字体
    mxOcx.AddTextStyle2("MyTrueTypeStyle", "黑体", 0.7);
    mxOcx.TextStyle = "MyTrueTypeStyle";
    mxOcx.DrawColor = 4556677;
    mxOcx.DrawMText(0,100,"\\A2;平方方法 :m{\\H0.3x;\\S 2^;}",15,100,0.0,7);
    mxOcx.ZoomAll();
    mxOcx.UpdateDisplay();

}