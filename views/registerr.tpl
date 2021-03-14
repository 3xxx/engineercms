<!DOCTYPE html>
{{template "header"}}
<title>注册 - MeritMS</title>
</head>

<body>
  <div id="content" class="col-md-8 col-md-offset-2">
    <div class="col-md-6 auth-page">
      <h3 class="title">
        <span class="glyphicon glyphicon-remove"></span>
        注册发生错误！
      </h3>
      <p class="well">可能用户名 或 邮箱已存在，请重新注册。</p>
      <form method="POST" action="/regist">
        <div class="form-group">
          <label class="control-label" for="LoginForm-UserName">用户名 或 邮箱</label>
          <input id="uname" name="uname" type="text" value="" class="form-control" placeholder="Enter account"></div>
        <div class="form-group">
          <label class="control-label" for="LoginForm-Password">密码</label>
          <input id="pwd" name="pwd" type="password" value="" class="form-control" placeholder="Password"></div>
        <button type="submit" class="btn btn-default" onclick="return checkInput();">
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
  </div>
  <script type="text/javascript">
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
  </script>

</body>

</html>