function GetAllEntity()
{
    // 得到当前图纸空间
    var blkRec = mxOcx.GetDatabase().CurrentSpace();
    // 创建一个用于遍历当前图纸空间的遍历器
    var iter = blkRec.NewIterator();
    var iNum = 0;
    // 循环得到所有实体
    if (iter!= null)
    {
        for (; !iter.Done(); iter.Step(true, false))
        {
            iNum++;
            // 得到遍历器当前的实体
            var ent = iter.GetEntity();
            if (ent.ObjectName == "McDbLine")
            {
                // 当前实体是一个直线
                ;
            }
            else if (ent.ObjectName=="McDbBlockReference")
            {
                blkRef = ent;
                // 当前实体是一个块引用
                for (var j = 0; j < blkRef.AttributeCount; j++)
                {
                    // 得到块引用中所有的属性
                    var attrib = blkRef.AttributeItem(j);

                    var mxUtility = mxOcx.NewUtility();

                    mxUtility.Prompt("n Tag: " + attrib.Tag + "Text:" + attrib.TextString);
                }
            }
        }
    }
    var sT;
    sT = "发现" + iNum + "个实体";
    alert(sT);
    // 注意需要显示释放递代器.,不然会引起错误
    iter = null;
    CollectGarbage();
}