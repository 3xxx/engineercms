function CreateLineType2() {
    //增加一个带有形的线型
    mxOcx.AddTextStyle1("MyLineTypeTextStyle", "txt.shx", "hztxt.shx", 1);
    var linetypeId= mxOcx.AddLinetypeEx ("MyLineType2", "(12.7,(\"T=MxDraw\",\"S=2.54\",\"L=-5.08\",\"R=0.0\"," +
        "\"X=-2.54\",\"Y=-1.27\"),-10.08)","MyLineTypeTextStyle");
    //返回控件的数据库对象
    var database = mxOcx.GetDatabase();
    //实体id返回实体对象
    var linetype = database.ObjectIdToObject(linetypeId);
    //设计当前线型为"MyLineType2"
    database.CurrentlyLineTypeName  = "MyLineType2";
    alert("添加成功");
}