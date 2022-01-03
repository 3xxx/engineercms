function CreateLineType() {
    //返回控件的数据库对象
    var database =mxOcx.GetDatabase();
    //返回数据库中的标注样式表对象
    var mxlinetype = database.GetLinetypeTable();
    //增加新的标注样式
    var linetypestr = mxlinetype.Add("mylinetype");
    //设置样式
    linetypestr.Comments = "DOT";
    //返回数据库中的线型表对象
    database.CurrentlyLineTypeName  = "mylinetype";
}