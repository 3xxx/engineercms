<!DOCTYPE html>  
<html>  
<head>  
<meta charset="UTF-8">  
<title>Insert title here</title>  
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/> 
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script> 
  <script type="text/javascript"> 

  $(function(){
    //判断访问终端
    var browser={
      versions:function(){
        var u = navigator.userAgent, app = navigator.appVersion;
        return {
            trident: u.indexOf('Trident') > -1, //IE内核
            presto: u.indexOf('Presto') > -1, //opera内核
            webKit: u.indexOf('AppleWebKit') > -1, //苹果、谷歌内核
            gecko: u.indexOf('Gecko') > -1 && u.indexOf('KHTML') == -1,//火狐内核
            mobile: !!u.match(/AppleWebKit.*Mobile.*/), //是否为移动终端
            ios: !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/), //ios终端
            android: u.indexOf('Android') > -1 || u.indexOf('Adr') > -1, //android终端
            iPhone: u.indexOf('iPhone') > -1 , //是否为iPhone或者QQHD浏览器
            iPad: u.indexOf('iPad') > -1, //是否iPad
            webApp: u.indexOf('Safari') == -1, //是否web应该程序，没有头部与底部
            //webApp: ("standalone" in window.navigator)&&window.navigator.standalone, //是否web应该程序，没有头部与底部
            weixin: u.indexOf('MicroMessenger') > -1, //是否微信 （2015-01-22新增）
            qq: u.match(/\sQQ/i) == " qq" //是否QQ

        };
      }(),
      language:(navigator.browserLanguage || navigator.language).toLowerCase()
    }

    //判断是否IE内核
    if(browser.versions.trident){ alert("is IE"); }
    //判断是否webKit内核
    if(browser.versions.webKit){ alert("is webKit"); }
    //判断是否移动端
    if(browser.versions.mobile||browser.versions.android||browser.versions.ios){ alert("移动端"); }
    // 检测浏览器语言
    currentLang = navigator.language;   //判断除IE外其他浏览器使用语言
    if(!currentLang){//判断IE浏览器使用语言
        currentLang = navigator.browserLanguage;
    }
    alert(currentLang);
  })

    </script>  
</head>  
<body>  
    <p>欢迎使用设代系统移动端访问！</p>
<!-- </div>   -->
</body>  
</html>