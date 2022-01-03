<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="initial-scale=1,user-scalable=no" />
  <title>SSO统一认证服务</title>
  <link rel="stylesheet" href="/static/css/sso/login.css" />
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
        <span class="monileHidden">认证登录</span>
        <span class="pcHidden">EngineerCMS知识管理登录平台</span>
      </div>
      <form action="" method="POST" id="login-form">
        <p>
          <span></span><input type="text" id="user" placeholder="请输入工号/域账号" value="" />
        </p>
        <p>
          <span></span><input type="password" id="psw" placeholder="请输入密码" value="" />
        </p>
        <div id="warn"><span>!</span><span id="warnText"></span></div>
        <div class="submit" onClick="login()">登录</div>
      </form>
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

  function login() {
    var warn = document.getElementById('warn');
    var user = document.getElementById('user');
    var psw = document.getElementById('psw');

    var text = !user.value && !psw.value ? '请输入账号和密码！' : !user.value ? '请输入账号！' : !psw.value ? '请输入密码！' : "";

    if (text) {
      document.getElementById("warnText").innerHTML = text;
      warn.style.display = 'block';
      return;
    }
    var options = {
      login_name: user.value,
      password: psw.value,
      service: "{{.service}}",
      renew: false,
      returnurl: ""
    };
    if ($("#service").length > 0) {
      options.service = $("#service").val();
    }
    if ($("#renew").length > 0) {
      options.renew = $("#renew").val();
    }
    if ($("#returnurl").length > 0) {
      options.returnurl = $("#returnurl").val();
    } else {
      warn.style.display = 'none';
      $.post("/v1/wx/ssologinpost", options, function(data) {
        $(".submit").html("登录");
        if (data.success == 0) {
          // console.log(data.service)
          // alert("onkeydown")
          location.href = data.service;
        } else {
          ShowMsg(data.Message);
        }
      });
    }
  }

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

</html>