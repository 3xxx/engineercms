function CreateDim(){
    //返回控件的数据库对象
    var database =mxOcx.GetDatabase();
    //返回数据库中的标注样式表对象
    var mxdim = database.GetDimStyleTable();
    //增加新的标注样式
    var dimstr = mxdim.Add("mydim");
    //设置样式
    dimstr.SetDimVarDouble(40, 20.0);
    //返回数据库中的标注样式表对象
    database.CurrentlyDimstyleName = "mydim";

}
