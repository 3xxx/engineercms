<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>EngineerCMS - A Fresh Mobile Systen</title>
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
  <!-- <link rel="shortcut icon" href="/static/voxes/img/favicon.png"> -->
    <!-- 收藏用logo图标 -->
  <link rel="bookmark" type="image/x-icon" href="/static/img/only.ico"/>
  <!-- 网站显示页logo图标 -->
  <link rel="shortcut icon" href="/static/img/only.ico">
</head>

<body class="login-register-wrap">
  <!-- login register -->
  <div class="login-register-wrap-home">
    <div class="container">
      <div class="content">
        <h1>E</h1>
        <h6>EngineerCMS</h6>
        <!-- <div id="warn">
          <span>!</span>
          <span id="warnText"></span>
        </div> -->
        <form action="">
          <input type="text" placeholder="Username" class="validate" required id="uname" onkeypress="getKey()">
          <input type="password" placeholder="Password" required id="pwd" onkeypress="getKey()">
          <span><a href="/regist">Forgot Password?</a>
          </span><a class="button-default" onclick="return login();">Sign in</a>
          <h6>Don't hava an account ? <a href="/regist">Sign up</a></h6>
        </form>
      </div>
    </div>
  </div><!-- end login register -->
  <!-- scripts -->
  <script src="/static/voxes/js/jquery.min.js"></script>
  <script src="/static/voxes/js/materialize.min.js"></script>
  <script src="/static/voxes/js/owl.carousel.min.js"></script>
  <script src="/static/voxes/js/contact-form.js"></script>
  <script src="/static/voxes/js/fakeLoader.min.js"></script>
  <script src="/static/voxes/js/main.js"></script>

  <script type="text/javascript">
//登陆功能
function login() {
  var uname = document.getElementById("uname");
  // if (uname.value.length == 0) {
  //   alert("请输入账号");
  //   return
  // }
  var pwd = document.getElementById("pwd");
  // if (pwd.value.length == 0) {
  //   alert("请输入密码");
  //   return
  // }

  $.ajax({
    type: 'post',
    url: '/loginpost',
    data: {
      "uname": $("#uname").val(),
      "pwd": $("#pwd").val()
    },
    success: function(result) {
      if (result.islogin == 0) {
        // $("#warn").html("登陆成功");
        // $('#modalNav').modal('hide');
        Materialize.toast('登陆成功', 3000, 'rounded')
        location.href = "/index";
        // window.location.reload();
      } else if (result.islogin == 1) {
        // $("#warn").html("用户名或密码错误！")
        Materialize.toast('用户名或密码错误！', 3000, 'rounded')
      } else if (result.islogin == 2) {
        // $("#warn").html("密码错误!")
        Materialize.toast('密码错误!', 3000, 'rounded')
      }
    }
  })
}
//登出功能
function logout() {
  $.ajax({
    type: 'get',
    url: '/logout',
    data: {},
    success: function(result) {
      if (result.islogin) {
        // $("#status").html("登出成功");
        alert("登出成功");
        window.location.reload();
      } else {
        // $("#status").html("登出失败");
        alert("登出失败")
      }
    }
  })
}

function getKey() {
  if (event.keyCode == 13) {
    login()
  }
}
</script>
</body>

</html>