function DrawLine() {
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
    // 直接绘制一个实线
    mxOcx.DrawLine(0, 0, 100, 0);
    // 绘制一个实斜线
    mxOcx.DrawLine(200, 0, 300, 100);
    //----------------------------------------------------------------------------------------------------------
    //绘制一个虚线
    //定义虚线数据据,"MyLineType"是线型名，"6,-8"是虚线的一个单位定义，6是实线长，-8是空格长。
    mxOcx.AddLinetype("MyLineType", "6,-10");
    //设计当前线型为"MyLineType"
    mxOcx.LineType = "MyLineType";
    // 绘制一个虚线
    mxOcx.DrawLine(0, 30, 100, 30);
    // 绘制一个斜虚线
    mxOcx.DrawLine(200, 30, 300, 130);
    //---------------------------------------------------------------------------------------------------------
    // 修改绘线的颜色为 16711680(蓝色),  16711680转成16进制是0xFF 00 00,其中，FF是蓝色，00是绿色，第个二00是红色。
    mxOcx.DrawCADColor = 16711680;
    // 绘制一个蓝色的虚线
    mxOcx.DrawLine(0, 60, 100, 60);
    // 绘制一个蓝色的斜虚线
    mxOcx.DrawLine(200, 60, 300, 160);
    //---------------------------------------------------------------------------------------------------------
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 4;
    //绘制一个带宽度的直线。
    mxOcx.DrawLine(0, 90, 100, 90);
    // 绘制一个带宽度的斜线
    mxOcx.DrawLine(200, 90, 300, 190);
    //---------------------------------------------------------------------------------------------------------
    //绘制一个点划线虚线
    mxOcx.AddLinetype("MyLineType2", "10,-2,3,-2");
    //把线型改点划线
    mxOcx.LineType = "MyLineType2";
    // 修改绘线的颜色为 255(红色),  255转成16进制是0x00 00 FF,其中，00是蓝色，第个二00是绿色，FF是红色。
    mxOcx.DrawCADColor = 255;
    //绘制一个带宽度的红色点划线。
    mxOcx.DrawLine(0, 120, 100, 120);
    // 绘制一个带宽度红色点划斜线
    mxOcx.DrawLine(200, 120, 300, 220);
    //---------------------------------------------------------------------------------------------------------
    //增加一个带有形的线型
    mxOcx.AddTextStyle1("MyLineTypeTextStyle", "txt.shx", "hztxt.shx", 1);
    mxOcx.AddLinetypeEx("MyLineType3", "(12.7,(\"T=MxDraw\",\"S=2.54\",\"L=-5.08\",\"R=0.0\",\"X=-2.54\",\"Y=-1.27\"),-10.08)", "MyLineTypeTextStyle");
    mxOcx.LineType = "MyLineType3";
    mxOcx.LineWidth = 0;
    //绘制一个带宽度的红色点划线。
    mxOcx.DrawLine(350, 120, 600, 120);
    //---------------------------------------------------------------------------------------------------------
    //增加一个带有形的线型
    //把颜色改回黑白色
    mxOcx.DrawCADColorIndex = 0;
    mxOcx.LineType = "FENCELINE1";
    mxOcx.LineWidth = 0;
    //绘制一个带宽度的红色点划线。
    mxOcx.DrawLine(350, 220, 600, 220);
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}