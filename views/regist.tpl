<!DOCTYPE html>
{{template "header"}}
<title>注册 - EngineerCMS</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1  maximum-scale=1 user-scalable=no">
<meta name="mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-capable" content="yes">
<meta name="apple-touch-fullscreen" content="yes">
<meta name="HandheldFriendly" content="True">
<link rel="stylesheet" href="/static/voxes/css/materialize.css">
<link rel="stylesheet" href="/static/voxes/font-awesome/css/font-awesome.min.css">
<link rel="stylesheet" href="/static/voxes/css/normalize.css">
<link rel="stylesheet" href="/static/voxes/css/owl.carousel.css">
<link rel="stylesheet" href="/static/voxes/css/owl.theme.css">
<link rel="stylesheet" href="/static/voxes/css/owl.transitions.css">
<link rel="stylesheet" href="/static/voxes/css/fakeLoader.css">
<link rel="stylesheet" href="/static/voxes/css/style.css">
<link rel="shortcut icon" href="/static/voxes/img/favicon.png">
</head>

<body>
  <!-- login register -->
  <div class="login-register-wrap-home">
    <div class="container">
      <div class="content">
        <h1>E</h1>
        <h6>EngineerCMS</h6>
        <form method="POST" action="/regist">
          <input type="text" id="uname" name="uname" placeholder="Full Name" class="validate" required>
          <input type="text" id="nickname" name="nickname" placeholder="Nick Name" class="validate" required>
          <input type="email" id="email" name="email" placeholder="Email" class="validate" required>
          <input type="password" id="pwd" name="pwd" placeholder="password" class="validate" required>
          <button class="button-default" type="submit" οnclick="return checkInput()">Sing up</button>
          <h6>Already hava an account ? <a href="/login">Sign in</a></h6>
        </form>
      </div>
    </div>
  </div><!-- end login register -->
  <!--   <div id="content" class="col-md-8 col-md-offset-2">
    <div class="col-md-6 auth-page">
      <h3 class="title">
        <span class="glyphicon glyphicon-user"></span>
        注册
      </h3>
      <form method="POST" action="/regist">
        <div class="form-group">
          <label class="control-label" id="regis" for="LoginForm-UserName">*用户名</label>
          <input id="uname" name="uname" type="text" value="" class="form-control" placeholder="Enter account"></div>
        <div class="form-group">
          <label class="control-label" for="LoginForm-UserName">邮箱</label>
          <input id="email" name="email" type="text" value="" class="form-control" placeholder="Enter Email"></div>
        <div class="form-group">
          <label class="control-label" for="LoginForm-UserName">昵称</label>
          <input id="nickname" name="nickname" type="text" value="" class="form-control" placeholder="Enter NickName"></div>
        <div class="form-group">
          <label class="control-label" for="LoginForm-Password">*密码</label>
          <input id="pwd" name="pwd" type="password" value="" class="form-control" placeholder="Password"></div>
        <button id="regist" name="regist" type="submit" class="btn btn-default" onclick="return checkInput();">
          注册&nbsp;&nbsp;
          <span class="glyphicon glyphicon-circle-arrow-right"></span>
        </button>
        <a href="./forgot" class="pull-right">
          <span class="glyphicon glyphicon-question-sign"></span>
          忘记密码
        </a>
        <button class="btn btn-default" onclick="return backToHome();">返回&nbsp;&nbsp; <span class="glyphicon glyphicon-circle-arrow-left"></span></button>
      </form>
    </div>
    <div class="col-md-6 auth-page">
      <div class="auth-page">
        <h3 class="title">
          <span class="glyphicon glyphicon-question-sign"></span>
          帮助
        </h3>
        <p class="well">如果您已经注册的话，请登陆。</p>
        <p>
          <a href="/login" class="btn btn-default">
            立即登陆&nbsp;&nbsp; <i class="icon-chevron-sign-right"></i>
          </a>
        </p>
      </div>
    </div>
  </div> -->
  <script type="text/javascript">
  // $(document).ready(function() {
  $("#uname").blur(function() {
    $.ajax({
      type: "post",
      url: "/regist/checkuname",
      data: { uname: $("#uname").val() },
      success: function(data, status) {
        if ($("#uname").val().length > 0 && data == "false") {
          $("#regis").html("用户名已存在！");
          $("#regis").css("color", "red");
        } else if ($("#uname").val().length > 0 && data == "true") {
          $("#regis").html("用户名可用！");
          $("#regis").css("color", "green");
        } else if ($("#uname").val().length == 0) {
          $("#regis").html("用户名");
          $("#regis").css("color", "black");
        }
      }
    });
  });
  // });

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
    if ($("#uname").val().length > 0 && $("#regis").html() == "用户名已存在！") {
      alert("用户名已存在！");
      return false;
    }
    // return true
  }

  function backToHome() {
    window.location.href = "/";
    return false;
  }
  </script>
  <!-- scripts -->
  <script src="/static/voxes/js/jquery.min.js"></script>
  <script src="/static/voxes/js/materialize.min.js"></script>
  <script src="/static/voxes/js/owl.carousel.min.js"></script>
  <script src="/static/voxes/js/contact-form.js"></script>
  <script src="/static/voxes/js/fakeLoader.min.js"></script>
  <script src="/static/voxes/js/main.js"></script>
</body>

</html>