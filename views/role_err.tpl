<!-- 权限不够界面 -->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
</head>

<body>
  <div id="content" class="col-md-8 col-md-offset-2">
    <div class="col-md-6 auth-page">
      <h3 class="title">
        <span class="glyphicon glyphicon-remove"></span>
        权限不够！
      </h3>
      <p class="well">请重新登陆。</p>
      <form method="POST" action="/post">
        <span style="color: #ff0000;"><input type="hidden" name="url" value="{{.Url}}" /></span>
        <div class="form-group">
          <label class="control-label" for="LoginForm-UserName">用户名 或 邮箱</label>
          <input id="uname" name="uname" type="text" value="" class="form-control" placeholder="Enter account" list="cars"></div>
        <div id='datalistDiv'>
          <datalist id="cars" name="cars">
          </datalist>
        </div>
        <div class="form-group">
          <label class="control-label" for="LoginForm-Password">密码</label>
          <input id="pwd" name="pwd" type="password" value="" class="form-control" placeholder="Password"></div>
        <div class="checkbox">
          <label>
            <input type="checkbox">自动登陆</label>
        </div>
        <button type="submit" class="btn btn-default" onclick="return checkInput();">
          登录&nbsp;&nbsp; <span class="glyphicon glyphicon-circle-arrow-right"></span>
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
        <p class="well">如果您还没有注册帐户的话，请先注册。</p>
        <p>
          <a href="./regist" class="btn btn-default">
            立即注册&nbsp;&nbsp; <i class="icon-chevron-sign-right"></i>
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
  $('#uname').attr("autocomplete", "off");
  $(document).ready(function() {
    $("#uname").keyup(function(event) {
      var uname1 = document.getElementById("uname");
      // alert(event.keyCode);
      if (event.keyCode != 38 && event.keyCode != 40 && uname1.value.length == 2) {
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: "/regist/getuname",
          data: { uname: $("#uname").val() },
          dataType: 'json', //dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
          success: function(data, status) {
            $(".option").remove();
            $.each(data, function(i, d) {
              $("#cars").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
            });
          }
        });

      }
    });
  });
  </script>
</body>

</html>