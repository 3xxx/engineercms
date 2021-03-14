
function LoadMxDrawX(dwgfile,cabpath,msipath) {

    var s, classid, Sys = {}, ua = navigator.userAgent.toLowerCase();
    (s = ua.match(/msie ([\d.]+)/)) ? Sys.ie = s[1] : (s = ua.match(/trident\/([\d.]+)/)) ? Sys.ie9 = s[1] : (s = ua.match(/firefox\/([\d.]+)/)) ? Sys.firefox = s[1] : (s = ua.match(/chrome\/([\d.]+)/)) ? Sys.chrome = s[1] : (s = ua.match(/opera.([\d.]+)/)) ? Sys.opera = s[1] : (s = ua.match(/version\/([\d.]+).*safari/)) ? Sys.safari = s[1] : 0,
    classid = "74A777F8-7A8F-4e7c-AF47-7074828086E2",

    Sys.ie || Sys.ie9 ? (document.write("<!-- 用来产生编辑状态的ActiveX控件的JS脚本-->   "),
    document.write("<!-- 因为微软的ActiveX新机制，需要一个外部引入的js-->   "),
    document.write('<object id="MxDrawXCtrl" classid="clsid:' + classid + '" '),
	//document.write('width="85%" height="85%" align="left">   '),
    document.write('codebase=" ' + cabpath + 'width="85%" height="85%" align="left">   '),
    document.write('<param name="_Version" value="65536">  '),
    document.write('<param name="_ExtentX" value="24262">  '),
    document.write('<param name="_ExtentY" value="16219">  '),
    document.write('<param name="_StockProps" value="0">'),
    document.write('<param name="DwgFilePath" value="' + dwgfile + '" > '),
    document.write('<param name="IsRuningAtIE" value="1">'),
    document.write('<param name="EnablePrintCmd" value="1">  '),
    document.write('<param name="ShowCommandWindow" value="1">   '),
    document.write('<param name="ShowToolBars" value="1">  '),
    document.write('<param name="ShowModelBar" value="1">'),
    document.write('<param name="Iniset" value="">  '),
    document.write('<param name="ToolBarFiles" value="">'),
    document.write('<param name="ShowMenuBar" value="1">'),
    document.write('<param name="EnableUndo" value="1">'),
    document.write('<param name="ShowPropertyWindow" value="1">'),
    document.write('<SPAN STYLE="color:red">不能装载文档控件。请在检查浏览器的选项中检查浏览器的安全设置。请点击<a href=' + msipath + '>安装控件</a></SPAN>'),
    document.write('</object>')) :

    Sys.chrome ? (document.write('<object id="MxDrawXCtrl" clsid="{' + classid + '}" '),
    document.write('type="application/mxdraw-activex" width="85%" height="85%" align="left" '),
    document.write('<param name="_Version" value="65536">  '),
    document.write('<param name="_ExtentX" value="24262">  '),
    document.write('<param name="_ExtentY" value="16219">  '),
    document.write('<param name="_StockProps" value="0">'),
    document.write('<param name="DwgFilePath" value="' + dwgfile + '" > '),
    document.write('<param name="IsRuningAtIE" value="1">'),
    document.write('<param name="EnablePrintCmd" value="1">  '),
    document.write('<param name="ShowCommandWindow" value="1">   '),
    document.write('<param name="ShowToolBars" value="1">  '),
    document.write('<param name="ShowModelBar" value="1">'),
    document.write('<param name="Iniset" value="">  '),
    document.write('<param name="ToolBarFiles" value="">'),
    document.write('<param name="ShowMenuBar" value="1">'),
    document.write('<param name="EnableUndo" value="1">'),
    document.write('<param name="ShowPropertyWindow" value="1">'),
    document.write('<param name="Event_ImplementCommandEvent" value="DoCommandEventFunc">'),
    document.write('<SPAN STYLE="color:red">不能装载文档控件。请在检查浏览器的选项中检查浏览器的安全设置。请点击<a href=' + msipath + '>安装控件</a></SPAN>'),
    document.write('</object>')) :

    Sys.firefox ? (document.write("<!-- 需要安装ieTab插件才能使用-->   "),
    document.write("<!-- 右键弹出菜单，点击使用ieTab浏览> -->  "),
   
    document.write('<div class="no_title">  不能装载文档控件。1.请在安装ieTab插件，2.然后请点击<a href=' + msipath + '>安装控件</a>, 3.然后右键弹出菜单，点击使用ieTab浏览 </div> ')

    ) :
    Sys.opera ? alert("sorry,ntko 暂时不支持opera!") :
    Sys.safari && alert("sorry,ntko 暂时不支持safari!");
}


function isFireFox()
{
    var s, classid, Sys = {}, ua = navigator.userAgent.toLowerCase();
    (s = ua.match(/msie ([\d.]+)/)) ? Sys.ie = s[1] : (s = ua.match(/trident\/([\d.]+)/)) ? Sys.ie9 = s[1] : (s = ua.match(/firefox\/([\d.]+)/)) ? Sys.firefox = s[1] : (s = ua.match(/chrome\/([\d.]+)/)) ? Sys.chrome = s[1] : (s = ua.match(/opera.([\d.]+)/)) ? Sys.opera = s[1] : (s = ua.match(/version\/([\d.]+).*safari/)) ? Sys.safari = s[1] : 0;

    return  Sys.firefox;

}