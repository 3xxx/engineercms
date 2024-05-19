<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="initial-scale=1,user-scalable=no" />
  <title>知识管理系统</title>
  <!-- 收藏用logo图标 -->
  <link rel="bookmark" type="image/x-icon" href="/static/img/only1.ico" />
  <!-- 网站显示页logo图标 -->
  <link rel="shortcut icon" href="/static/img/only1.ico">
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
      <img src="/static/images/ecmslogo.png" alt="EngineerCMS logo" />
    </div>
    <div class="content">
      <div class="title">
        <span class="monileHidden">认证登录</span>
        <span class="pcHidden">知识管理系统</span>
      </div>
      <input type="hidden" name="url" value="{{.Url}}" />
      <input id="referrer" type="text" name="referrer" class="form-control" style="display:none;">
      <form action="" method="POST" id="login-form">
        <p>
          <span></span><input type="text" id="uname" placeholder="请输入用户名/警号" value="" />
        </p>
        <p>
          <span></span><input type="password" id="pwd" placeholder="请输入密码" value="" onkeypress="getKey()" />
        </p>
        <div id="warn"><span>!</span><span id="warnText"></span></div>
        <div class="submit" onclick="return login();">登录</div>
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
  $('#referrer').val(document.referrer);

  function checkInput() {
    var uname = document.getElementById("uname");
    if (uname.value.length == 0) {
      alert("请输入账号");
      return false;
    }
    var pwd = document.getElementById("pwd");
    if (pwd.value.length == 0) {
      alert("请输入密码");
      return false;
    }
    return true
  }

  function backToHome() {
    window.location.href = "/";
    return false;
  }

  //监听输入框中回车键
  function getKey() {
    if (event.keyCode == 13) {
      login()
    }
  }

  //登陆功能
  function loginback() {
    var uname = document.getElementById("uname");
    if (uname.value.length == 0) {
      alert("请输入账号");
      return
    }
    var pwd = document.getElementById("pwd");
    if (pwd.value.length == 0) {
      alert("请输入密码");
      return
    }

    $.ajax({
      type: 'post',
      url: '/loginpost',
      data: {
        "uname": $("#uname").val(),
        "pwd": $("#pwd").val()
      },
      success: function(result) {
        if (result.islogin == 0) {
          $("#status").html("登陆成功");
          // $('#modalNav').modal('hide');
          window.location.reload();
        } else if (result.islogin == 1) {
          $("#status").html("用户名或密码错误！")
        } else if (result.islogin == 2) {
          $("#status").html("密码错误")
        }
      }
    })
  }


  // $(function() {
  //   $(".layerBox").on("click", ".radio-name", function() {
  //     $(this).parent().find("input").click();
  //   })
  //   $("#loginform").attr("action", location.href);
  //   window.document.onkeydown = function(evt) {
  //     evt = window.event || evt;
  //     if (evt.keyCode == 13) {
  //       login();
  //     }
  //   }
  // });

  function login() {
    var warn = document.getElementById('warn');
    var user = document.getElementById('uname');
    var psw = document.getElementById('pwd');
    var text = !user.value && !psw.value ? '请输入账号和密码！' : !user.value ? '请输入账号！' : !psw.value ? '请输入密码！' : "";
    if (text) {
      document.getElementById("warnText").innerHTML = text;
      warn.style.display = 'block';
      return;
    }
    var options = {
      uname: user.value,
      pwd: psw.value,
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
      $.post("/loginpost", options, function(data) {
        $(".submit").html("登录");
        if (data.islogin == 0) {
          // console.log(data.service)
          alert("登录成功，点确定跳转到首页……")
          // location.href = "/index";
          location.href = "{{.Url}}";
          // window.location.reload();
        } else if (data.islogin == 1) {
          ShowMsg("用户名或密码错误！");
        } else if (data.islogin == 2) {
          ShowMsg("密码错误！");
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