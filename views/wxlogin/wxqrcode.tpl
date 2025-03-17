<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
</head>

<body>
  <div id="login_container"></div>
</body>
<!-- <script src="/static/vue.js/js/jquery.min.js"></script> -->
<script src="https://downs.yaoulive.com/liveJs/axios.min.js"></script>
<script src="https://res.wx.qq.com/connect/zh_CN/htmledition/js/wxLogin.js"></script>
<script>
var state = parseInt(new Date().getTime() / 1000);
var obj = new WxLogin({
  self_redirect: false,
  id: "login_container",
  appid: "*******",
  scope: "snsapi_login",
  redirect_uri: encodeURIComponent("https://zsj.itdos.net/v1/wx/redirecturi"),
  state: state,
  style: "white",
  href: "data:text/css;base64,LmxvZ2luUGFuZWwubm9ybWFsUGFuZWwgLnRpdGxlIHsNCiAgZGlzcGxheTogbm9uZTsNCn0NCi5xcmNvZGUubGlnaHRCb3JkZXIgew0KICB3aWR0aDogMTc0cHg7DQogIGhlaWdodDogMTc0cHg7DQogIG1hcmdpbi10b3A6IDA7DQogIGJveC1zaXppbmc6IGJvcmRlci1ib3g7DQp9DQouaW1wb3dlckJveCAuaW5mbyB7DQogIGRpc3BsYXk6IG5vbmU7DQp9DQoud2ViX3FyY29kZV90eXBlX2lmcmFtZSB7DQogIHdpZHRoOiAxNzRweDsNCn0NCg=="
});

// 这一步可以直接调用接口返回给后端code，由后端来进行这一步
console.log(window.location.search.substring(6).split('&')[0])
// if (window.location.search.substring(6).split('&')[0]) {
//   axios({
//     'url': `https://api.weixin.qq.com/sns/oauth2/access_token?appid=*****&secret=******d&code=${window.location.search.substring(6).split('&')[0]}&grant_type=authorization_code`,
//     'method': 'get'
//   })
// }
</script>

</html>