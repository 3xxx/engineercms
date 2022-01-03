//保存为加密文件
function saveEncryptionDWG() {
    //保存为加密文件，文件路径与文件名有用户设置
    var ret = mxOcx.Call("MxET_SaveEncryptionDWG", "c:\\TmpEncryption.dwg");
    if(ret.AtString(0) == "Ok")
    {
        alert("加密文件已经成功保存到C:\\TmpEncryption.dwg");
    }
    else
    {
        alert("保存失败");
    }
    ret = null;
    CollectGarbage();
}
