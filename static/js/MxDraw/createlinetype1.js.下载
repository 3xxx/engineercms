function CreateLineType1() {
    //定义虚线数据据,"MyLineType"是线型名，"6,-8"是虚线的一个单位定义，6是实线长，-8是空格长。
    var linetypeId= mxOcx.AddLinetype("MyLineType1", "10,-2,3,-2");
    //返回控件的数据库对象
    var database = mxOcx.GetDatabase();
    //实体id返回实体对象
    var linetype = database.ObjectIdToObject(linetypeId);
    //设计当前线型为"MyLineType1"
    database.CurrentlyLineTypeName  = "MyLineType1";
    alert("添加成功");
}
