function SetXData(){
    var ent = mxOcx.GetEntity("选择要写扩展数据的实体:");
    if(ent == null)
    {
        return;
    }
    var exData = mxOcx.Call("Mx_NewResbuf","");
    exData.AddStringEx("TestApp",1001);
    exData.AddStringEx("甲",1000);
    ent.SetXData(exData);
}
