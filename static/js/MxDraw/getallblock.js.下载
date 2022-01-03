function GetAllBlock() {
    var database = mxOcx.GetDatabase();
    var blkTab = database.GetBlockTable();
    var iter = blkTab.NewIterator();
    for (; !iter.Done(); iter.Step())
    {
        var blkRec = iter.GetRecord();
        alert(blkRec.Name);

    }
}