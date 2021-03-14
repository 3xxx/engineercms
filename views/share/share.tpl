<!-- 分享页面-->
<!-- <!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="initial-scale=1,user-scalable=no" />
    <title>文件提取</title>
    <link rel="stylesheet" href="/static/css/sso/login.css" />
    <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
    <script src="/static/js/sso/jquery-1.7.min.js"></script>
    <script src="/static/js/sso/placeholder.js"></script>
</head>

<body>
    <header class="pcHidden">
        <img src="/static/images/bg_mobile.png" />
    </header>
    <section>
        <div class="logo">
            <img src="/static/images/logo.png" alt="EngineerCMS logo" />
        </div>
        <div class="content">
            <div class="title">
                <span class="monileHidden">文件提取</span>
                <span class="pcHidden">EngineerCMS知识管理登录平台</span>
            </div>

            <form action="" method="POST" id="login-form">
                <p>
                    <i class="fa fa-key fa-2x"></i><input type="text" id="user" placeholder="请输入提取码" value="" />
                </p>
                <div id="warn"><span>!</span><span id="warnText"></span></div>
                <div class="submit" onClick="login()">提取文件</div>
            </form>

            <div class="title">
                <span class="monileHidden">请登录</span>
            </div>
            <div id="warn"><span>!</span><span id="warnText"></span></div>
            <div class="submit" onClick="login()"><i class="w14 fa fa-user-circle-o"></i>登录</div>

        </div>
        <div class="monileHidden footer">
            <img src="/static/images/bg_footer.png" alt="" />
        </div>
    </section>
    <div id="layer">
        <div class="layerBox">
        </div>
        <div class="shade"></div>
    </div>
    <script>
    $(function() {
        $(".layerBox").on("click", ".radio-name", function() {
            $(this).parent().find("input").click();
        })
        $("#loginform").attr("action", location.href);
        window.document.onkeydown = function(evt) {
            evt = window.event || evt;
            if (evt.keyCode == 13) {
                login();
            }
        }
    });

    function ShowMsg(title) {
        $(".submit").html("登录");
        $("#warnText").html(title);
        $("#warn").css("display", "block");
    }

    var browser = navigator.appName
    var b_version = navigator.appVersion
    var version = b_version.split(";");
    var trim_Version = version[1].replace(/[ ]/g, "");
    if (browser == "Microsoft Internet Explorer" && (trim_Version == "MSIE6.0" || trim_Version == "MSIE7.0" || trim_Version == "MSIE8.0")) {
        $(".pcHidden").hide();
        $("img").eq(0).hide();
        $(".logo").attr("style", "width:100%;text-align:center;margin-top:150px");
        $(".logo img").attr("style", "margin:auto;width:200px");
        $(".content").width("500px")
    }
    </script>
</body>

</html> -->
<!DOCTYPE html>
<html lang="en" class=" ">

<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <link rel="icon" href="https://tank.eyeblue.cn/favicon.ico">
  <title>EyeblueTank</title>
  <!-- <link href="/static/eyeblue/app.310dcbb8.css" rel="preload" as="style"> -->
  <!-- <link href="/static/eyeblue/chunk-vendors.fc2d671c.css" rel="preload" as="style"> -->
  <!-- <link href="/static/eyeblue/app.88391206.js" rel="preload" as="script"> -->
  <!-- <link href="/static/eyeblue/chunk-vendors.588116bf.js" rel="preload" as="script"> -->
  <!-- <link href="/static/eyeblue/chunk-vendors.fc2d671c.css" rel="stylesheet"> -->
  <link href="/static/eyeblue/app.310dcbb8.css" rel="stylesheet">
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <style type="text/css">
  #page-wrapper {
    position: fixed;
    left: 0;
    top: 45px;
    right: 0;
    bottom: 40px;
    overflow-y: auto;
    overflow-x: hidden;
    z-index: 10;
    padding: 10px;
    -webkit-transition: all .4s;
    transition: all .4s;
    background-color: #f3f3f4
  }

  @media (min-width: 768px) {
    #page-wrapper {
      left: 0
    }
  }

  @media (max-width: 767px) {
    #page-wrapper {
      left: 0;
      bottom: 0
    }
  }

  @media (min-width: 768px) {
    #page-wrapper.show-drawer {
      left: 0
    }
  }

  @media (max-width: 767px) {
    #page-wrapper.show-drawer {
      left: 0;
      bottom: 0
    }
  }
 
  </style>
</head>

<body>
  <div class="nb-app">
    <div id="body">
      <div>
        <div id="page-wrapper">
          <div class="share-detail">
            <div class="share-detail">
              <div>
                <div class="text-center" style="display: none;"><i class="fa fa-spinner fa-spin fa-3x fa-fw"></i></div>
                <div style="">
                  <div>
                    <div class="share-block" id="share-block">
                      <div class="upper">
                        <div class="left-box"><img src="/static/eyeblue/img/archive.77d78eb7.svg" class="share-icon"><span class="name">
                            {{.Share.Name}}
                            </span></div>
                        <div class="right-box"><button class="btn btn-primary btn-sm mr5" onclick="downloadzip()"><i class="fa fa-download"></i>
                            下载
                          </button>
                          <div class="el-dialog__wrapper" style="display: none;">
                            <div role="dialog" aria-modal="true" aria-label="分享详情" class="el-dialog" style="margin-top: 15vh;">
                              <div class="el-dialog__header"><span class="el-dialog__title">分享详情</span><button type="button" aria-label="Close" class="el-dialog__headerbtn"><i class="el-dialog__close el-icon el-icon-close"></i></button></div>
                              <div class="el-dialog__footer"><span class="dialog-footer"><button class="btn btn-primary btn-sm mr5">复制链接+提取码</button><button class="btn btn-default btn-sm mr5">关闭</button></span></div>
                            </div>
                          </div>
                        </div>
                      </div>
                      <div class="share-info"><span class="inline-block mr10">
                          分享者:{{.Share.Username}}
                        </span><span class="inline-block mr10">
                          创建时间:{{dateformat .Share.CreatedAt "2006-01-02 15:04:05"}}
                        </span><span class="inline-block mr10">
                          失效时间:{{dateformat .Share.ExpireTime "2006-01-02 15:04:05"}}
                        </span>
                      </div>
                    </div>


                    <div id="matter-panel">
                      <!-- <div class="widget-share-matter-panel">
                        <div>
                          <div class="media">
                            <div class="pull-left">
                              <div class="left-part"><span class="basic-span"><img src="/static/eyeblue/img/folder.f8d1b500.svg" class="matter-icon"></span></div>
                            </div>
                            <div class="pull-right hidden-sm hidden-xs">
                              <div class="right-part"><span class="matter-operation"><i title="下载" class="fa fa-download btn-action text-primary"></i></span><span class="matter-size">
                                  4.7 kB
                                </span><span class="matter-date">
                                  2020-03-02 17:16
                                </span></div>
                            </div>
                            <div class="pull-right hidden-lg hidden-md"><span class="more-btn"><i title="更多" class="fa fa-ellipsis-h btn-action"></i></span></div>
                            <div class="media-body">
                              <div class="middle-part"><span title="为什么自建的无法打开word" class="matter-name">
                                  为什么自建的无法打开word
                                </span></div>
                            </div>
                          </div>
                        </div>
                      </div> -->
                    </div>


                    <div class="mt20">
                      <div>
                        <div class="text-center" style="display: none;"><i class="fa fa-spinner fa-spin fa-3x fa-fw"></i></div>
                        <div style="display: none;">
                          <div class="italic text-center">
                            该目录下暂无任何内容
                          </div>
                        </div>
                        <div class="cursor" style="display: none;">
                          <div class="text-center">
                            <div><img src="" class="img-md"></div>
                            <div class="mt10">
                            </div>
                            <div>
                              点击刷新
                            </div>
                          </div>
                        </div>
                        <div class="text-center mt10" style="display: none;">
                          每页
                          <select>
                            <option value="5">5</option>
                            <option value="10">10</option>
                            <option value="15">15</option>
                            <option value="20">20</option>
                            <option value="30">30</option>
                            <option disabled="disabled" value="50">
                              50
                            </option>
                          </select>
                          条
                          共 1 条
                        </div>
                      </div>
                    </div>
                    <div class="col-md-4 col-md-offset-4 mt100" id="code">
                      <div class="input-group"><input type="text" placeholder="请输入提取码" class="form-control" id="password"><span class="input-group-btn">
                        <button type="button" class="btn btn-primary" onclick="browse()">
                            提取文件
                        </button></span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="top-navigation-bar">
          <nav class="navbar">
            <div class="navbar-title"><a href="https://tank.eyeblue.cn/" class="is-link-active"><img src="/static/eyeblue/logo.21458adb.png" class="logo"><span class="title">EyeblueTank</span></a></div><button class="drawer-trigger btn btn-primary btn-sm mr5"><i class="fa fa-bars"></i></button>
          </nav>
        </div>
        <div class="bottom-navigation text-center"><span class="item"><span></span></span><span class="item"><span></span></span><span class="item"><a href="javascript:void(0)">
              English
            </a></span><span class="brand">
            Powered by <a target="_blank" href="https://github.com/eyebluecn/tank"><img src="/static/eyeblue/logo.21458adb.png" class="w30">
              蓝眼云盘</a></span></div>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  $(document).ready(function() {
    browse()
  })

  function browse(){
    var code = $('#password').val()
    $.ajax({
      type: "get",
      url: "/v1/share/browse",
      data: { shareUuid: {{.Shareuuid }}, code: code },
      success: function(data, status) {
        // alert("查询到“" + data.data[0].Code + "-" + data.data[0].Title + "”成功！(status:" + status + ".)");
        // code: "NEED_SHARE_CODE"
        // data: null
        // msg: "提取码必填"
        if (data.code != "OK") {
          document.getElementById("share-block").style.display = "none"//不显示
          document.getElementById("matter-panel").style.display = "none"//不显示
          document.getElementById("code").style.display = "block"
          // document.getElementById("nav").style.display = "block"//显示
        }else{
          for(let i = 0;i<data.data.length;i++){
          $("#matter-panel").append('<div class="widget-share-matter-panel"><div><div class="media"><div class="pull-left"><div class="left-part"><span class="basic-span"><img src="/static/eyeblue/img/folder.f8d1b500.svg" class="matter-icon"></span></div></div><div class="pull-right hidden-sm hidden-xs"><div class="right-part"><span class="matter-operation"  onclick="download('+data.data[i].Id+')"><i title="下载" class="fa fa-download btn-action text-primary"></i></span><span class="matter-size">4.7 kB </span><span class="matter-date">'+moment(data.data[i].Created).format('YYYY-MM-DD HH:mm:ss')+'</span></div></div><div class="pull-right hidden-lg hidden-md"><span class="more-btn"><i title="更多" class="fa fa-ellipsis-h btn-action" onClick="show('+data.data[i].Id+')"></i></span></div><div class="media-body"><div class="middle-part"><span title="" class="matter-name">'+data.data[i].Code+'-'+data.data[i].Title+'</span></div></div></div></div><div class="hidden-lg hidden-md more-panel" style="display: none;" id="'+data.data[i].Id+'"><div class="cell-btn" style="border: none;"><span>'+moment(data.data[i].Created).format('YYYY-MM-DD HH:mm:ss')+'</span></div><div title="下载" class="cell-btn" onclick="download('+data.data[i].Id+')"><i class="fa fa-download"></i>下载</div></div></div>')
          }
          document.getElementById("share-block").style.display = "block"
          document.getElementById("matter-panel").style.display = "block"
          document.getElementById("code").style.display = "none"//不显示
        }
      },
      error: function(data, status) {
        alert(data.data);
      }
    });
  }

  function show(id){
    var _container = document.getElementsByClassName("hidden-lg hidden-md more-panel");
        // console.log(_container[0].style.display);
        for(let i = 0;i<_container.length;i++){
          _container[i].style.display = "none"
        }
    // document.getElementsByClassName("hidden-lg hidden-md more-panel").style.display = "none"
    // console.log(document.getElementById(id).style.display)
    // if (document.getElementById(id).style.display == "none"){
      document.getElementById(id).style.display = "block"
    // }else{
      // document.getElementById(id).style.display = "none"
    // }
  }

  function download(id){
    var code = $('#password').val();
    fetch('/v1/share/download',{
      // signal, // 在option中加入signal
      method: 'POST',
      // credentials:'include',
      headers: {
        'Content-Type': 'application/json;charset=UTF-8'
      },
      body: JSON.stringify({
        shareUuid: {{.Shareuuid }},
        code:code,
        id: id,
      })
    }).then(res => res.blob().then(blob => {
      // if (res.code === 0) {
        // console.log(blob)
      var a = document.createElement('a');
      var url = window.URL.createObjectURL(blob);
      var filename = res.headers.get('Content-Disposition').split(';')[1].split('=')[1];
      
      a.href = url;
      a.download = decodeURI(filename);
      a.click();
      window.URL.revokeObjectURL(url);
      // document.body.removeChild(a);
      // document.getElementById('status').innerHTML = '下载完成';
    }));
  }

  function downloadzip(id){
    var code = $('#password').val();
    fetch('/v1/share/downloadzip',{
      // signal, // 在option中加入signal
      method: 'POST',
      // credentials:'include',
      headers: {
        'Content-Type': 'application/json;charset=UTF-8'
      },
      body: JSON.stringify({
        shareUuid: {{.Shareuuid }},
        code:code,
        // id: id,
      })
    }).then(res => res.blob().then(blob => {
      // if (res.code === 0) {
      var a = document.createElement('a');
      var url = window.URL.createObjectURL(blob);
      var filename = res.headers.get('Content-Disposition').split(';')[1].split('=')[1];
      // console.log(filename);
      // console.log(decodeURI(filename));
      a.href = url;
      a.download = decodeURI(filename);
      a.click();
      window.URL.revokeObjectURL(url);
      // document.body.removeChild(a);
      // document.getElementById('status').innerHTML = '下载完成';
    }));
  }
//触发方法下载后
// const params = [new Date(),new Date(),new Date(),new Date()];
// var url = 'http://localhost:8080/api/file';
// fetch(url,{
//     method: 'POST',
//     body: window.JSON.stringify(params),
//     credentials: 'include',
//     headers: new Headers({
//         'Content-Type': 'application/json',
//     })
// }).then(res => res.blob().then(blob => {
//     var a = document.createElement('a');
//     var url = window.URL.createObjectURL(blob);
//     var filename = res.headers.get('Content-Disposition');
//     a.href = url;
//     a.download = filename;
//     a.click();
//     window.URL.revokeObjectURL(url);
//     document.getElementById('status').innerHTML = '下载完成';
// }))


  // var url = "http://localhost/sample/upload.php";
  // document.getElementById('btn').onclick = function() {
  //   document.getElementById('status').innerHTML = '下载中';
  //   fetch(url).then(res => res.blob().then(blob => {
  //       var a = document.createElement('a');
  //       var url = window.URL.createObjectURL(blob);
        
  //       var filename = res.headers.get('Content-Disposition');
  //       a.href = url;
  //       a.download = filename;
  //       a.click();
  //       window.URL.revokeObjectURL(url);
  //     document.getElementById('status').innerHTML = '下载完成';
  //   }));
  // };

//下载
// function downloadBlob(blob,fileName){
//   let blobUrl = window.URL.createObjectURL(blob);
//   let a = document.createElement('a');
//   a.href = blobUrl;
//   a.target = '_blank';
//   a.style.display = 'none'
//   document.body.appendChild(a)
//   a.download = fileName;
//   a.click();
//   window.URL.revokeObjectURL(blobUrl);
//   document.body.removeChild(a)
//   that.setState({
//     downloading:false
//   })
// }

// 下载excel文件，对象为一个二进制数据流
// fetch(myRequest).then(function(response) {
//     response.arrayBuffer().then(res=> {
//         var blob = new Blob([res], {type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"});
//         var objectUrl = URL.createObjectURL(blob);
//         var a = document.createElement("a");
//         document.body.appendChild(a);
//         a.style = "display: none";
//         a.href = objectUrl;
//         a.click();
//         document.body.removeChild(a);
//      })
// });
// 用 async/await 重新写了一版
// async function postDownload(url, params) {
//   const request = {
//     body: JSON.stringify(params),
//     method: 'POST',
//     headers: {
//       'Content-Type': 'application/json;charset=UTF-8'
//     }
//   }
//   const response = await fetch(url, request)
//   const filename = response.headers.get('content-disposition').split(';')[1].split('=')[1]
//   const blob = await response.blob()
//   const link = document.createElement('a')
//   link.download = decodeURIComponent(filename)
//   link.style.display = 'none'
//   link.href = URL.createObjectURL(blob)
//   document.body.appendChild(link)
//   link.click()
//   URL.revokeObjectURL(link.href)
//   document.body.removeChild(link)j
// }

// export function get(url,params){
// if (params) {
// let paramsArray = [];
// //拼接参数
// Object.keys(params).forEach(key => paramsArray.push(key + '=' + params[key]))
// if (url.search(/\?/) === -1) {
// url += '?' + paramsArray.join('&')
// } else {
// url += '&' + paramsArray.join('&')
// }
// }
// //fetch请求,get请求只能在url里拼接参数哦！！
// fetch(url,{
// method: 'GET',
// })
// .then((response) => response.json())
// .then((json) => {
// console.log("戴假发"+JSON.stringify(json));
// // this.setState({ discounts: json.data })
// })
// .catch((error) => {
// alert(error)
// })
// }
</script>
</body>

</html>
<!-- 1）get请求 -->
<script type="text/javascript">
$('.download').click(function() {
  var tt = new Date().getTime();
  var url = 'http://192.168.1.231:8080/91survey/ws/excel/download';
  /**
   * 使用form表单来发送请求
   * 1.method属性用来设置请求的类型——post还是get
   * 2.action属性用来设置请求路径。
   * 
   */
  var form = $("<form>"); //定义一个form表单
  form.attr("style", "display:none");
  form.attr("target", "");
  form.attr("method", "get"); //请求类型
  form.attr("action", url); //请求地址
  $("body").append(form); //将表单放置在web中
  /**
   * input标签主要用来传递请求所需的参数：
   *
   * 1.name属性是传递请求所需的参数名.
   * 2.value属性是传递请求所需的参数值.
   *
   * 3.当为get类型时，请求所需的参数用input标签来传递，直接写在URL后面是无效的。
   * 4.当为post类型时，queryString参数直接写在URL后面，formData参数则用input标签传递
   * 有多少数据则使用多少input标签
   *
   */
  var input1 = $("<input>");
  input1.attr("type", "hidden");
  input1.attr("name", "tt");
  input1.attr("value", tt);
  form.append(input1);
  var input2 = $("<input>");
  input2.attr("type", "hidden");
  input2.attr("name", "companyId");
  input2.attr("value", companyId);
  form.append(input2);
  form.submit(); //表单提交
})
// 2） post请求

$('.download').click(function() {
  var tt = newDate().getTime();
  var url = restUrl + '/excel/download?userId=' + userId;
  var form = $("<form>"); //定义一个form表单
  form.attr("style", "display:none");
  form.attr("target", "");
  form.attr("method", "post"); //请求类型
  form.attr("action", url); //请求地址
  $("body").append(form); //将表单放置在web中
  var input1 = $("<input>");
  input1.attr("type", "hidden");
  input1.attr("name", "tt");
  input1.attr("value", tt);
  form.append(input1);
  var input2 = $("<input>");
  input2.attr("type", "hidden");
  input2.attr("name", "companyId");
  input2.attr("value", companyId);
  form.append(input2);
  form.submit(); //表单提交
});
// 完成后， 在页面上面点击下载图标， 文件就会自动下载了。
</script>