<!-- 这个是使用mxdraw控件预览dwg文件 -->
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<!-- saved from url=(0042)http://www.mxcad.net:2080/ie/database.html -->
<!-- <!DOCTYPE html PUBLIC "" ""> -->
<HTML lang="en" style="height: 100%;">
<HEAD>
  <META content="IE=10.000" http-equiv="X-UA-Compatible">
  <META charset="UTF-8">
  <META name="GENERATOR" content="MSHTML 10.00.9200.17457">
  <META name="ProgId" content="FrontPage.Editor.Document">
  <TITLE>MxDraw控件</TITLE>
  <!--引入打碎函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/ExplodeFun.js" type="text/javascript"></SCRIPT>
  <!--移动夹点-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MoveGripPointsFun.js" type="text/javascript"></SCRIPT>
  <!--返回夹点-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/GetGripPointsFun.js" type="text/javascript"></SCRIPT>
  <!--动态施放绘制事件回调函数指针-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/DoDynWordDrawFun.js" type="text/javascript"></SCRIPT>
  <!--引入参数绘制的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/Draw.js" type="text/javascript"></SCRIPT>
  <!--引入交互绘制的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxDyDraw.js" type="text/javascript"></SCRIPT>
  <!--引入光栅图处理的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxIamges.js" type="text/javascript"></SCRIPT>
  <!--引入界面控制的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxInterface.js" type="text/javascript"></SCRIPT>
  <!--引入控制事件的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxEvents.js" type="text/javascript"></SCRIPT>
  <!--引入打印控制的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxPrint.js" type="text/javascript"></SCRIPT>
  <!--引入选择集的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxSelect.js" type="text/javascript"></SCRIPT>
  <!--引入自定义命令的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxUserCustomCommand.js" type="text/javascript"></SCRIPT>
  <!--引入扩展数据的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxData.js" type="text/javascript"></SCRIPT>
  <!--引入图面搜索的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxMap.js" type="text/javascript"></SCRIPT>
  <!--引入图形数据库的相关函数-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/MxDataBase.js" type="text/javascript"></SCRIPT>
  <SCRIPT language="javascript" src="/static/js/mxdraw/mxcustom.js" type="text/javascript"></SCRIPT>
  <!--引入梦想控件-->
  <SCRIPT language="javascript" src="/static/js/mxdraw/mxocx.js" type="text/javascript"></SCRIPT>
  <SCRIPT src="/static/js/mxdraw/jquery.min.js"></SCRIPT>

  <STYLE>
        *{
            margin:0;
            padding:0;
        }
        #menu {
            overflow-y: auto;
            border-right: 1px solid #eee;
            min-width: 12%;
            /* min-width: 280px;*/
            float: left;
            overflow-y: scroll;
            height: 100%;
            background: #F7F9FD;
        }
        #menu ul {
            font-size: 14px;
            font-family: "微软雅黑";
            background-color: #f7f9fd;
            padding: 0 10px;
        }
        #menu ul > li {
            padding-right: 0;
        }
        #menu ul li {
            list-style: none;
            border-bottom: 1px solid #e9e9e9;
            padding: 15px 10px;
        }
        #menu>ul>li>a {
            width: 100%;
            height: 21px;
            line-height: 21px;
            display: block;
            color: #000;
            text-decoration: none;
        }
        #menu .t_close {
            background: url(result.png) 0 -96px no-repeat;
        }
        #menu i {
            width: 13px;
            height: 13px;
            display: block;
            margin-top: 4px;
            margin-right: 7px;
            float: left;
        }
        #menu .submenu {
            margin: 0;
            padding: 10px 0 0 6px;
            box-sizing: border-box;
            z-index: 9999;
            width: 100%;
        }
        #menu .submenu dl {
            list-style-type: square;
            border: none;
            padding: 0;
            margin: 0;
        }
        #menu .submenu dl dd {
            line-height: 35px;
            height: 35px;
            font-size: 14px;
            word-break: keep-all;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            margin-left: 1px;
        }
        #menu .submenu a {
            margin: 0;
            height: 35px;
            width: 100%;
            display: block;
            color: #000;
            text-decoration: none;
        }
        .one_head:hover{
            cursor: pointer;
        }
        #menu .submenu dl dd a:before {
            content: "・";
            font-size: 27px;
            float: left;
            margin-left: 2px;
            margin-right: 1px;
        }
        #menu .submenu dl dd:before {
            border-bottom: 1px dotted #b0b5c2;
            content: "";
            display: inline-block;
            float: left;
            width: 10px;
            margin: -15px 0 0;
            height: 33px;
            border-left: 1px dotted #b0b5c2;
        }
        #menu a:hover, .clickState {
            color: #06C!important;
            cursor: pointer;
        }
        ::-webkit-scrollbar {
            width: 8px;
        }
        ::-webkit-scrollbar-thumb {
            border-radius: 0;
            -webkit-box-shadow: inset 3px 3px 3px 10px #e4e4e4;
        }
        ::-webkit-scrollbar-track {
            box-shadow: inset 0 0 1px #d1cfcf;
            border-radius: 0;
        }
        #code_area {
            width: 30%;
            -webkit-user-select: none;
            float: left;
            height: 94%;
            z-index: 999;
            position: relative;
            border: 1px solid #F7F9FD;
        }
        #code_area #openbar {
            padding: 0;
            height: 36px;
            background: #F7F9FD;
            overflow: hidden;
            -webkit-user-select: none;
            -moz-user-select: none;
        }
        #openbar>p {
            line-height: 36px;
            margin-left: 10px;
        }
        .CodeMirror {
            line-height: 1;
            position: relative;
            overflow: hidden;
            background: white;
            color: black;
            height: 80%
        }
  </STYLE>

  <SCRIPT language="JavaScript">document.oncontextmenu = new Function('event.returnValue=false;'); //禁用右键
  </SCRIPT>
</HEAD>

<BODY>

<DIV style="height: 900px;"><!--  width: 150%; float: right; -->
  <SCRIPT type="text/javascript">
        LoadMxDrawX("http://127.0.0.1/static/img/05.dwg", "", "");
        // {{.DwgLink}}
        // LoadMxDrawX("http://127.0.0.1/static/img/05.dwg", "http://www.mxcad.net:2080/MxDrawX52.CAB#version=10,0,0,1", "http://www.mxdraw.com/MxDrawX52.msi");
        // var isShow = false;
        // var mxOcx = document.getElementById("MxDrawXCtrl");
        // mxOcx.ShowToolBar("常用工具",isShow);
            // mxOcx.ShowToolBar("绘图工具",isShow);
            // mxOcx.ShowToolBar("编辑工具",isShow);
            // mxOcx.ShowToolBar("特性",isShow);
            // mxOcx.ShowToolBar("ET工具",isShow);
            // isShow = !isShow;
        // mxOcx.ShowMenuBar(isShow);
  </SCRIPT>
</DIV>
<SCRIPT>
  
    changeMenu(11);
    function changeMenu(menuNum){
        var t_close=document.getElementsByClassName('t_close');
        t_close[menuNum].parentNode.style.color="#06C";
        t_close[menuNum].style.backgroundPositionX="0";
        t_close[menuNum].style.backgroundPositionY="-119px";
        var submenu=document.getElementsByClassName('submenu');
        submenu[menuNum].style.display='block';
        submenu[menuNum].children[0].children[0].children[0].classList.add("clickState");
    }
    function show(obj) {
        var _this = obj;
        _this.nextElementSibling.style.display = "block";
        _this.style.color="#06C";
        var one_head=document.getElementsByClassName('one_head');
        for(var j=0;j<one_head.length;j++){
            if(one_head[j]!=_this){
                one_head[j].style.color="#000";
                one_head[j].children[0].style.backgroundPositionX="0";
                one_head[j].children[0].style.backgroundPositionY="-96px";
                one_head[j].setAttribute("onclick", "show(this)");
            }
        }
        var submenus=document.getElementsByClassName('submenu');
        for(var i=0;i<submenus.length;i++){
            if(submenus[i]!=_this.nextElementSibling){
                submenus[i].style.display = "none";
            }
        }
        _this.children[0].style.backgroundPositionX="0";
        _this.children[0].style.backgroundPositionY="-119px";
        _this.setAttribute("onclick", "hide(this)");
    }

    function hide(obj) {
        var _this = obj;
        _this.style.color="#000";
        _this.nextElementSibling.style.display = "none";
        _this.children[0].style.backgroundPositionX="0";
        _this.children[0].style.backgroundPositionY="-96px";
        _this.setAttribute("onclick", "show(this)");
    }
    // var mxOcx = document.getElementById("MxDrawXCtrl");

    function DoCmd(obj,iCmd) {
        var _this=obj;
        _this.classList.add("clickState");
        var containers=document.getElementsByClassName('container');
        for(var i=0;i<containers.length;i++){
            if(containers[i]!=_this){
                containers[i].classList.remove("clickState");
            }
        }
        mxOcx.DoCommand(iCmd);
        var code=document.getElementById('code');
        // var test=document.getElementsByTagName('html')[0].outerHTML;
        var title=_this.innerText;
        var test='';
        test+=''+RetrenMethod(iCmd)+'\n';
        // test=test.replace(/\n/g,"<br/>").replace(/\s/g,"&nbsp;")
        code.innerText=test;
    }
    function ChangeSrc(src){
        var html="";
        $.ajax({
            type: "get",
            url: src,
            // dataType: "text",
            async: false
        }).done(function (data) {
            var data=data;
            console.log(data);
            // html=$("#one").html(data).html();//替换回车行成换行符
            //console.log(data.replace(/\r\n/ig,'</br>'));
            html=data;
        });
        return html;
    }
    //读文件
    function RetrenMethod(iCmd){
        var name="";
        //参数绘制-----------------------------------------------------------------------------
        if (iCmd == 1) {
            name=ChangeSrc('drawline.js');
        }
        else if (iCmd == 2) {
            name=ChangeSrc('drawpolyline.js');
        }
        else if (iCmd == 3) {
            name=ChangeSrc('drawpoint.js');
        }
        else if (iCmd == 4) {
            name=ChangeSrc('drawspline.js');
        }
        else if (iCmd == 5) {
            name=ChangeSrc('drawcircle.js');
        }
        else if (iCmd == 6) {
            name=ChangeSrc('drawarc.js');
        }
        else if (iCmd == 7) {
            name=ChangeSrc('drawellipse.js');
        }
        else if (iCmd == 8) {
            name=ChangeSrc('drawellipsearc.js');
        }
        else if (iCmd == 9) {
            name=ChangeSrc('drawtext.js');
        }
        else if (iCmd == 10) {
            name=ChangeSrc('drawmtext.js');
        }
        else if (iCmd == 11) {
            name=ChangeSrc('backgroundimage.js');
        }
        else if (iCmd == 12) {
            name=ChangeSrc('insertimage.js');
        }
        else if (iCmd == 13) {
            name=ChangeSrc('drawimage.js');
        }
        else if (iCmd == 14) {
            name=ChangeSrc('usersavejpg.js');
        }
        else if (iCmd == 15) {
            name=ChangeSrc('sethyperlink.js');
        }
        else if (iCmd == 16) {
            name=ChangeSrc('drawinsert.js');
        }
        else if (iCmd == 17) {
            name=ChangeSrc('drawsolid.js');
        }
        else if (iCmd == 18) {
            name=ChangeSrc('drawpathtohatch.js');
        }
        else if (iCmd == 19) {
            name=ChangeSrc('drawpathtohatch1.js');
        }
        else if (iCmd == 20) {
            name=ChangeSrc('drawpathtohatch2.js');
        }
        else if (iCmd == 21) {
            name=ChangeSrc('docommentfix.js');
        }
        else if (iCmd == 22) {
            name=ChangeSrc('dofixrectcomment.js');
        }
        else if (iCmd == 23) {
            name=ChangeSrc('dofixcirclecomment.js');
        }
        else if (iCmd == 24) {
            name=ChangeSrc('docloudcirclecommentfix.js');
        }
        else if (iCmd == 25) {
            name=ChangeSrc('drawdimrotated.js');
        }
        else if (iCmd == 26) {
            name=ChangeSrc('drawdimradial.js');
        }
        else if (iCmd == 27) {
            name=ChangeSrc('drawdimdiametric.js');
        }
        else if (iCmd == 28) {
            name=ChangeSrc('drawdimangular.js');
        }
        else if (iCmd == 29) {
            name=ChangeSrc('drawdimaligned.js');
        }
        else if (iCmd == 30) {
            name=ChangeSrc('drawqrcode.js');
        }
        else if (iCmd == 31) {
            name=ChangeSrc('drawflag.js');
        }
        //交互绘制----------------------------------------------------------------------------------
        else if (iCmd == 52) {
            name=ChangeSrc('dydrawline.js');
        }
        else if (iCmd == 53) {
            name=ChangeSrc('dydrawmyline.js');
        }
        else if (iCmd == 54) {
            name=ChangeSrc('dydrawwidthline.js');
        }
        else if (iCmd == 55) {
            name=ChangeSrc('dydrawpolyline.js');
        }
        else if (iCmd == 56) {
            name=ChangeSrc('dydrawpoint.js');
        }
        else if (iCmd == 57) {
            name=ChangeSrc('dydrawspline.js');
        }
        else if (iCmd == 58) {
            name=ChangeSrc('dydrawcircle.js');
        }
        else if (iCmd == 59) {
            name=ChangeSrc('dydrawarc.js');
        }
        else if (iCmd == 60) {
            ChangeSrc('.js');
            name='DyDrawEllipse()';
        }
        else if (iCmd == 61) {
            ChangeSrc('.js');
            name=$('#one').html();
            name='DyDrawEllipseArc()';
        }
        else if (iCmd == 62) {
            name=ChangeSrc('dydrawpathtopolyline.js');
        }
        else if (iCmd == 63) {
            name=ChangeSrc('dydrawtext.js');
        }
        else if (iCmd == 64) {
            name=ChangeSrc('dydyninsert.js');
        }
        else if (iCmd == 65) {
            name=ChangeSrc('dydrawsolid.js');
        }
        else if (iCmd == 66) {
            name=ChangeSrc('dydocomment.js');
        }
        else if (iCmd == 67) {
            name=ChangeSrc('dydocomment3.js');
        }
        else if (iCmd == 68) {
            name=ChangeSrc('dydocirclecomment.js');
        }
        else if (iCmd == 69) {
            name=ChangeSrc('dydocloudlinecomment.js');
        }
        else if (iCmd == 70) {
            name=ChangeSrc('dydocloudcirclecomment.js');
        }
        else if (iCmd == 72) {
            name=ChangeSrc('dyndrawmatrix.js');
        }

        //光栅图处理-------------------------------------------------------------------------------
        else  if (iCmd == 73) {
            name=ChangeSrc('backgroundimage.js');
        }
        else if (iCmd == 74) {
            name=ChangeSrc('insertimage.js');
        }
        else if (iCmd == 75) {
            name=ChangeSrc('modifyimage.js');
        }
        else if (iCmd == 76) {
            name=ChangeSrc('rotateimage.js');
        }
        else if (iCmd == 77) {
            name=ChangeSrc('savejpg.js');
        }
        else if (iCmd == 78) {
            name=ChangeSrc('savedxf.js');
        }
        else if (iCmd == 79) {
            name=ChangeSrc('savepdf.js');
        }
        else if (iCmd == 80) {
            name=ChangeSrc('savedwf.js');
        }
        else if (iCmd == 81) {
            name=ChangeSrc('saveEncryptionDWG.js');
        }
        else if (iCmd == 82) {
            name=ChangeSrc('openEncryptionDWG.js');
        }
        else if (iCmd == 83) {
            name=ChangeSrc('showwatermark.js');
        }
        else if (iCmd == 84) {
            name=ChangeSrc('drawgif.js');
        }
        else if (iCmd == 85) {
            name=ChangeSrc('drawimage.js');
        }
        else if (iCmd == 86) {
            name=ChangeSrc('usersavejpg.js');
        }
        else if (iCmd == 10000) {
            name=ChangeSrc('textoffsetposition.js');
        }
        else if (iCmd == 10001) {
            name=ChangeSrc('Transparent.js');
        }
        //自定义实体
        else if (iCmd == 71) {
            name=ChangeSrc('dyinsertcustomentity.js');
        }
        else if (iCmd == 87) {
            name=ChangeSrc('myinsertcustomentity.js');
        }
        //打印控制--------------------------------------------------------------------------
        else if (iCmd == 88) {
            name=ChangeSrc('print.js');
        }
        else if (iCmd == 89) {
            name=ChangeSrc('print2.js');
        }
        else if (iCmd == 90) {
            name=ChangeSrc('printhtml.js');
        }
        else if (iCmd == 91) {
            name=ChangeSrc('addpagecomment.js');
        }
        else if (iCmd == 92) {
            name=ChangeSrc('print1.js');
        }
        else if (iCmd == 93) {
            name=ChangeSrc('printall.js');
        }
        //界面控制-------------------------------------------------------------------------------
        else if (iCmd == 100) {
            name=ChangeSrc('hidetoolbar.js');
        }
        else if (iCmd == 101) {
            name=ChangeSrc('brownermode.js');
        }
        else if (iCmd == 102) {
            name=ChangeSrc('hidemenubar.js');
        }
        else if (iCmd == 103) {
            name=ChangeSrc('hiderulerwindow.js');

        }
        else if (iCmd == 104) {
            name=ChangeSrc('hidepropertywindow.js');
        }
        else if (iCmd == 105) {
            name=ChangeSrc('hidecommandwindow.js');
        }
        else if (iCmd == 106) {
            name=ChangeSrc('hidemodelbar.js');
        }
        else if (iCmd == 107) {
            name=ChangeSrc('hidestatusbar.js');
        }

        //控制事件------------------------------------------------------------------------------------
        //鼠标事件
        else if (iCmd == 200) {
            name=ChangeSrc('mouseevent.js');
        }
        //动态拖放时的绘制事件
        else if (iCmd == 203) {
            name=ChangeSrc('doinsert.js');
        }

        //选择集----------------------------------------------------------------------------------
        else if (iCmd == 300) {
            name=ChangeSrc('allselect.js');
        }
        else if (iCmd == 301) {
            name=ChangeSrc('selectwindow.js');
        }
        else if (iCmd == 302) {
            name=ChangeSrc('selectcross.js');
        }
        else if (iCmd == 303) {
            name=ChangeSrc('selectsetall.js');
        }
        else if (iCmd == 304) {
            name=ChangeSrc('selectuserselect.js');
        }
        else if (iCmd == 305) {
            name=ChangeSrc('selectimpliedselectselect.js');

        }
        else if (iCmd == 306) {
            name=ChangeSrc('selectassign.js');
        }
        else if (iCmd == 307) {
            name=ChangeSrc('currentselect.js');
        }
        else if (iCmd == 308) {
            name=ChangeSrc('selectatpoint.js');
        }
        else if (iCmd == 309) {
            name=ChangeSrc('selectbypolygon.js');
        }

        //图面搜索----------------------------------------------------------------------------------
        //查找文字
        else if (iCmd == 400) {
            name=ChangeSrc('findtext.js');
        }
        //查找所有文字
        else if (iCmd == 401) {
            name=ChangeSrc('findalltext.js');

        }
        //扩展数据----------------------------------------------------------------------------------
        //写扩展数据
        else if (iCmd == 500) {
            name=ChangeSrc('setxdata.js');
        }
        //读扩展数据
        else if (iCmd == 501) {
            name=ChangeSrc('getxdata.js');
        }
        //写扩展数据
        else if (iCmd == 502) {
            name=ChangeSrc('writexdata.js');
        }
        //读扩展数据
        else if (iCmd == 503) {
            name=ChangeSrc('readxdata.js');
        }
        //写扩展数据
        else if (iCmd == 504) {
            name=ChangeSrc('setxdatadouble.js');

        }
        //读扩展数据
        else if (iCmd == 505) {
            name=ChangeSrc('getxdatadouble.js');

        }
        //写扩展数据
        else if (iCmd == 506) {
            name=ChangeSrc('setxdatalong.js');

        }
        //读扩展数据
        else if (iCmd == 507) {
            name=ChangeSrc('getxdatalong.js');

        }
        //写扩展数据
        else if (iCmd == 508) {
            name=ChangeSrc('getallappname.js');

        }
        //读扩展数据
        else if (iCmd == 509) {
                name=ChangeSrc('deletexdata.js');

        }
        //---------------------------------------------------------------------------
        //自定义命令
        else if (iCmd == 600) {
                name=ChangeSrc('userprint.js');
            }
        else if (iCmd == 601) {
                name=ChangeSrc('useropendwg.js');
            }
        //---------------------------------------------------------------------------
        //图形数据库
            //图层
        else if (iCmd == 700) {
            name=ChangeSrc('createlayer.js');

        }
        else if (iCmd == 701) {
            name=ChangeSrc('hidelayer.js');

        }
        else if (iCmd == 702) {

            name=ChangeSrc('showlayer.js');
        }
        else if (iCmd == 703) {
            name=ChangeSrc('getalllayer.js');

        }
        else if (iCmd == 704) {
            name=ChangeSrc('openalllayer.js');

        }
        else if (iCmd == 705) {
            name=ChangeSrc('lockealllayer.js');

        }

        //文字样式
        else if (iCmd == 710) {
            name=ChangeSrc('createtext.js');

        }
        else if (iCmd == 711) {
            name=ChangeSrc('getalltext.js');

        }
        else if (iCmd == 712) {
            name=ChangeSrc('deltext.js');

        }
        //标注样式
        else if (iCmd == 720) {
                name=ChangeSrc('createdimy.js');
            }
        else if (iCmd == 721) {
                name=ChangeSrc('getalldim.js');
            }
        else if (iCmd == 722) {
            name=ChangeSrc('deldim.js');
        }
        else if (iCmd == 730) {
                name=ChangeSrc('getallentity.js');
            }
        else if (iCmd == 740) {
                name=ChangeSrc('createlinetype.js');
            }
        else if (iCmd == 750) {
            name=ChangeSrc('createblocktable.js');
        }
        else if (iCmd == 751) {
            name=ChangeSrc('getallblock.js');
        }
        else if (iCmd == 752) {
            name=ChangeSrc('modelspaceentity.js');
        }
        else if (iCmd == 753) {
            name=ChangeSrc('assignblock.js');
        }

        else if (iCmd == 760) {
                name=ChangeSrc('getalllayoutdictionary.js');
            }
        else if (iCmd == 770) {
            name=ChangeSrc('createlinetype.js');
        }
        else if (iCmd == 780) {
                name=ChangeSrc('mxorder.js');
            }
        else if (iCmd == 790) {
                name=ChangeSrc('readdwgtitle.js');
            }
        return name;
    }

    function DoCommandEventFunc(iCmd)
    {
        //参数绘制
        if (iCmd == 1) {
            DrawLine();
        }
        else if (iCmd == 2) {
            DrawPolyline();
        }
        else if (iCmd == 3) {
            DrawPoint();
        }
        else if (iCmd == 4) {
            DrawSpline();
        }
        else if (iCmd == 5) {
            DrawCircle();
        }
        else if (iCmd == 6) {
            DrawArc();
        }
        else if (iCmd == 7) {
            DrawEllipse();
        }
        else if (iCmd == 8) {
            DrawEllipseArc();
        }
        else if (iCmd == 9) {
            DrawText();
        }
        else if (iCmd == 10) {
            DrawMText();
        }
        else if (iCmd == 11) {
            BackGroundImage();
        }
        else if (iCmd == 12) {
            InsertImage();
        }
        else if (iCmd == 13) {
            DrawImage();
        }
        else if (iCmd == 14) {
            UserSaveJpg();
        }
        else if (iCmd == 15) {
            SetHyperlink();
        }

        else if (iCmd == 16) {
            DrawInsert();
        }
        else if (iCmd == 17) {
            DrawSolid();
        }

        else if (iCmd == 18) {
            DrawPathToHatch();
        }
        else if (iCmd == 19) {
            DrawPathToHatch1();
        }
        else if (iCmd == 20) {
            DrawPathToHatch2();
        }
        else if (iCmd == 21) {

            DoCommentFix();
        }
        else if (iCmd == 22) {
            AddTextStyle();
            DoFixRectComment();
        }
        else if (iCmd == 23) {
            AddTextStyle();
            DoFixCircleComment();
        }
        else if (iCmd == 24) {
            DoCloudCircleCommentFix();
        }
        else if (iCmd == 25) {
            DrawDimRotated();
        }
        else if (iCmd == 26) {
            DrawDimRadial();
        }
        else if (iCmd == 27) {
            DrawDimDiametric();
        }
        else if (iCmd == 28) {
            DrawDimAngular();
        }
        else if (iCmd == 29) {
            DrawDimAligned();
        }
        else if (iCmd == 30) {
            DrawQrCode();
        }
        else if (iCmd == 31) {
            DrawFlag();
        }

        //-------------------------------------------------------------------------------------------------------
        //交互绘制
        else if (iCmd == 52) {
            DyDrawLine();
        }
        else if (iCmd == 53) {
            DyDrawMyline();
        }
        else if (iCmd == 54) {
            DyDrawWidthline();
        }
        else if (iCmd == 55) {
            DyDrawPolyline();
        }
        else if (iCmd == 56) {
            DyDrawPoint();
        }
        else if (iCmd == 57) {
            DyDrawSpline();
        }
        else if (iCmd == 58) {
            DyDrawCircle();
        }
        else if (iCmd == 59) {
            DyDrawArc();
        }
        else if (iCmd == 60) {
            DyDrawEllipse();
        }
        else if (iCmd == 61) {
            DyDrawEllipseArc();
        }
        else if (iCmd == 62) {
            DyDrawPathToPolyline();
        }
        else if (iCmd == 63) {
            DyDrawText();
        }
        else if (iCmd == 64) {
            DyDynInsert();
        }
        else if (iCmd == 65) {
            DyDrawSolid();
        }
        else if (iCmd == 66) {
            AddTextStyle();
            // 绘制批注
            DyDoComment();
        }
        else if (iCmd == 67) {
            AddTextStyle();
            DyDoComment3();
        }
        else if (iCmd == 68) {
            AddTextStyle();
            DyDoCircleComment();
        }
        else if (iCmd == 69) {
            DyDoCloudLineComment();
        }
        else if (iCmd == 70) {
            DyDoCloudCircleComment();
        }

        else if (iCmd == 72) {
            DynDrawMatrix();
        }
        //--------------------------------------------------------------------------------------------------
        //光栅图处理
        else  if (iCmd == 73) {
            BackGroundImage();
        }
        else if (iCmd == 74) {
            InsertImage();
        }
        else if (iCmd == 75) {
            ModifyImage();
        }
        else if (iCmd == 76) {
            RotateImage();
        }
        else if (iCmd == 77) {
            SaveJpg();
        }
        else if (iCmd == 78) {
            savedxf();
        }
        else if (iCmd == 79) {
            savepdf();
        }
        else if (iCmd == 80) {
            savedwf();
        }
        else if (iCmd == 81) {
            saveEncryptionDWG();
        }
        else if (iCmd == 82) {
            openEncryptionDWG();
        }
        else if (iCmd == 83) {
            ShowWatermark();
        }
        else if (iCmd == 84) {
            DrawGif();
        }
        else if (iCmd == 85) {
            DrawImage();
        }
        else if (iCmd == 86) {
            UserSaveJpg();
        }
        //-------------------------------------------------------------------------------------------------------------
        //自定义实体
        else if (iCmd == 71) {
            DyInsertCustomEntity();
        }
        else if (iCmd == 87) {
            MyInsertCustomEntity();
        }

        //-------------------------------------------------------------------------------------------------------------
        //打印控制
        else if (iCmd == 92) {
            Print1();
        }
        else if (iCmd == 88) {
            Print();
        }
        else if (iCmd == 89) {
            Print2();
        }
        else if (iCmd == 90) {
            PrintHtml();
        }
        else if (iCmd == 91) {
            AddPageComment();
        }
        else if (iCmd == 93) {
            PrintAll();
        }

        //-------------------------------------------------------------------------------------------------------------
        //界面控制
        else if (iCmd == 100) {
            HideToolbar();
        }
        else if (iCmd == 101) {
            BrownerMode();
        }
        else if (iCmd == 102) {
            HideMenuBar();
        }
        else if (iCmd == 103) {
            HideRulerWindow();
        }
        else if (iCmd == 104) {
            HidePropertyWindow();
        }
        else if (iCmd == 105) {
            HideCommandWindow();
        }
        else if (iCmd == 106) {
            HideModelBar();
        }
        else if (iCmd == 107) {
            HideStatusBar();
        }

        //-------------------------------------------------------------------------------------------------------------
        //控制事件
        //鼠标事件
        else if (iCmd == 200) {
            MouseEvent();
        }
        //动态拖放时的绘制事件
        else if (iCmd == 203) {
            DoInsert();
        }
        //-------------------------------------------------------------------------------------------------------------
        //选择集
        //AllSelect得到当前空间的所有实体
        else if (iCmd == 300) {
            AllSelect();
        }
        else if (iCmd == 301) {
            SelectWindow();
        }
        else if (iCmd == 302) {
            SelectCross();
        }
        else if (iCmd == 303) {
            SelectSetAll();
        }
        else if (iCmd == 304) {
            SelectUserSelect();
        }
        else if (iCmd == 305) {
            SelectImpliedSelectSelect();
        }
        else if (iCmd == 306) {
            SelectAssign();
        }
        else if (iCmd == 307) {
            CurrentSelect();
        }
        else if (iCmd == 308) {
            SelectAtPoint();
        }
        else if (iCmd == 309) {
            SelectByPolygon();
        }
        //-------------------------------------------------------------------------------------------------------------
        //图面搜索
        //查找文字
        else if (iCmd == 400) {
            FindText();
        }
        //查找所有文字
        else if (iCmd == 401) {
            FindAllText();
        }
        //-------------------------------------------------------------------------------------------------------------
        //扩展数据
        //写扩展数据
        else if (iCmd == 500) {
            SetXData();
        }
        //读扩展数据
        else if (iCmd == 501) {
            GetXData();
        }
        //写扩展数据
        else if (iCmd == 502) {
            WriteXData();
        }
        //读扩展数据
        else if (iCmd == 503) {
            ReadXData();
        }
        //写扩展数据
        else if (iCmd == 504) {
            SetxDataDouble();
        }
        //读扩展数据
        else if (iCmd == 505) {
            GetxDataDouble();
        }
        //写扩展数据
        else if (iCmd == 506) {
            SetxDataLong();
        }
        //读扩展数据
        else if (iCmd == 507) {
            GetxDataLong();
        }
        //得到所有扩展数据名称
        else if (iCmd == 508) {
            GetAllAppName();
        }
        //删除展数据
        else if (iCmd == 509) {
            DeleteXData();
        }
        //-----------------------------------------------------------------------------------
        //设置自定义命令调用打印
        else if (iCmd == 600) {
            mxOcx.RegistUserCustomCommand("00",111);
            mxOcx.Focus();
        }
        else if (iCmd == 111) {
                UserPrint();
            }


        //设置自定义命令打开dwg
        else if (iCmd == 601) {
            mxOcx.RegistUserCustomCommand("22",222);
            mxOcx.Focus();
        }
        else if (iCmd == 222) {
                UserOpenDwg();
            }

        //-----------------------------------------------------------------------------------
        //光栅图处理
            //置为最顶层
        else if (iCmd == 10000) {
            TextOffsetPosition();
        }
        //设置图片的透明度
        else if (iCmd == 10001) {
            Transparent();
        }

        //-----------------------------------------------------------------------------------
        //图形数据库
            //图层操作
        else if (iCmd ==700) {
            CreateLayer();
        }
        else if (iCmd == 701) {
            HideLayer();
        }
        else if (iCmd ==702) {
            ShowLayer();
        }
        else if (iCmd == 703) {
            GetAllLayer();
        }
        else if (iCmd ==704) {
            OpenAllLayer();
        }
        else if (iCmd ==705) {
            LockeAllLayer();
        }
        //文字样式
        else if (iCmd ==710) {
                CreateText();
            }
        else if (iCmd ==711) {
            GetAllText();
        }
        else if (iCmd ==712) {
            DelText();
        }

        //标注样式
        else if (iCmd ==720) {
                CreateDim();
            }
        else if (iCmd ==721) {
            GetAllDim();
        }

        else if (iCmd ==723) {
            ModifyDim();
        }

        else if (iCmd ==730) {
                GetAllEntity();
            }
        else if (iCmd ==740) {
                CreateLineType();
            }
        else if (iCmd ==750) {
                CreateBlockTable();
            }

        else if (iCmd == 751) {
            GetAllBlock();
        }
        else if (iCmd == 752) {
            ModelSpaceEntity();
        }
        else if (iCmd == 753) {
            AssignBlock();
        }
        else if (iCmd ==760) {
            GetAllLayoutDictionary();
        }
        else if (iCmd ==770) {
                DataBaseInsert();
            }
        else if (iCmd ==780) {
                MxOrder();
            }
        else if (iCmd ==790) {
                ReadDwgTitle();
            }
    }

    mxtime = setInterval(InitMxDrawX, 100);
    //初始化
    function InitMxDrawX() {
        if (mxOcx) {
            if (!mxOcx.IsIniting())
            {
                clearInterval(mxtime);
            }
        }
    }
    function AddTextStyle()
    {
        mxOcx.AddTextStyle2("MyCommentFont", "黑体", 0.7);
    }
    //设置命令事件回调函数
    document.getElementById("MxDrawXCtrl").ImplementCommandEventFun = DoCommandEventFunc;
    //动态施放绘制事件回调函数指针
    document.getElementById("MxDrawXCtrl").ImpDynWorldDrawFun = DoDynWorldDrawFun;
    //打碎自定义实体事件回调函数指针
    document.getElementById("MxDrawXCtrl").ImpExplodeFun = ExplodeFun;
    document.getElementById("MxDrawXCtrl").ImpGetGripPointsFun = GetGripPointsFun;
    document.getElementById("MxDrawXCtrl").ImpMoveGripPointsAtFun = MoveGripPointsFun;
    //设置移动自定义实体夹点事件回调函数
    document.getElementById("MxDrawXCtrl").ImpTransformByFun = TransformByFun;
    //设置矩阵变换自定义实体事件回调函数
    document.getElementById("MxDrawXCtrl").ImpGetGeomExtentsFun = GetGeomExtentsFun;
    //设置得到自定义实体外包事件回调函数
    document.getElementById("MxDrawXCtrl").ImpInputPointToolTipFun = DoInputPointToolTipFun;
    //设置超连接事件回调函数
    document.getElementById("MxDrawXCtrl").ImpHyperlinkClickFun = DoHyperlinkClickFun;
    //设置鼠标事件回调函数
    document.getElementById("MxDrawXCtrl").ImplementMouseEventFun = MouseEvent;
    //得到自定义实体捕捉点事件回调函数指针
    document.getElementById("MxDrawXCtrl").ImpGetOsnapPointsFun = GetOsnapPointsFun;
    //控件的用户自定义事件回调函数指针，多用于js脚本使用
    /*document.getElementById("MxDrawXCtrl").ImplementCustomEvent = DoCustomEventEventFun;*/
</SCRIPT>
</BODY>
</HTML>