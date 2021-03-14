//打开加密文件
function openEncryptionDWG() {
    //打开加密文件，文件路径与文件名有用户设置
    var ret =mxOcx.Call("MxET_OpenEncryptionDWG", "c:\\TmpEncryption.dwg");
    if(ret.AtString(0) == "Ok")
    {
        alert("加密文件已经成功打开");
    }
    else
    {
        alert("打开失败");
    }
    ret = null;
    CollectGarbage();

}