//新建一个图层
function CreateLayer() {
    //增加新的图层
    var layerId = mxOcx.AddLayer("自定义图层");
    //返回控件的数据库对象
    var database = mxOcx.GetDatabase();
    //实体id返回实体对象
    var layer = database.ObjectIdToObject(layerId);
    //返回图层的颜色
    var newColor = layer.Color;
    //设置图层颜色
    newColor.SetRGB(255,0,0);
    //将设置的颜色赋值给图层
    layer.Color = newColor;
    //当前绘图函数的使用的图层名
    mxOcx.LayerName = "自定义图层";
    //当前控件绘图函数使用的CAD颜色索引值
    mxOcx.DrawCADColorIndex = 256;
}