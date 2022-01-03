<!-- 视频列表 视频主页-->
<!DOCTYPE html>

<head>
  <script src="https://cdn.jsdelivr.net/npm/jquery@3.4.1/dist/jquery.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
  <!-- <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script> -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
  <script src="/static/toast/toast.min.js"></script>
  <link rel="stylesheet" href="/static/toast/toast.min.css">
  <style>
    .jumbotron {
    background: url(/static/images/5ww.jpg);no-repeat;
    background-size:100% 100%;
  }

  .gotop {
    position: fixed;
    top: 50%;
    z-index: 100;
    width: 6rem;
    /*background: rgba(167, 5, 65, 0.38);*/
    color: #fff;
    padding: .5rem;
    border-radius: .5rem;
    font-size: .5rem;
    text-decoration: none;
  }
  
  .gotop:hover,
  .gotop:active {
    background: rgba(167, 5, 65, 0.56) ！important;
    color: #fff;
    text-decoration: none
  }
</style>
</head>

<body data-spy="scroll">
  <!--   <nav class="navbar fixed-top navbar-light bg-light">
  <a class="navbar-brand" href="#">Fixed top</a>
</nav> -->
  <!-- <nav class="navbar navbar-dark bg-dark sticky-top navbar-light bg-light">
  <a class="navbar-brand" href="#">Sticky top</a>
</nav> -->
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark sticky-top navbar-light bg-light">
    <a class="navbar-brand" href="/">首页</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarText">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item active">
          <a class="nav-link video">视频<span class="sr-only">(current)</span></a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/v1/flv/flvlist">直播</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">点播</a>
        </li>
      </ul>
      <span class="navbar-text">培训 工程 音乐
        <!-- <a class="nav-link" href="#"></a> -->
      </span>
    </div>
  </nav>
  <!--   <header class="navbar navbar-expand navbar-dark flex-column flex-md-row bd-navbar">
</header> -->
  <div class="jumbotron jumbotron-fluid">
    <div class="container-fluid">
      <h1 class="display-4 text-danger">FLV视频点播</h1>
      <!-- <img src="/static/images/1.jpg" class="w-100" alt="..."> -->
      <p class="lead text-warning">This is a simple hero unit, a simple jumbotron-style component for calling extra attention to featured content or information.</p>
      <hr class="my-4 text-success">
      <p>It uses utility classes for typography and spacing to space content out within the larger container.</p>
      <a class="btn btn-primary btn-lg" href="#" role="button">Learn more</a>
    </div>
  </div>
  <div class="container-fluid">
    <div class="row flex-xl-nowrap">
      <div id="list-example" class="list-group gotop d-none d-sm-block">
        <a class="list-group-item list-group-item-action" href="#list-item-1">直播</a>
        <a class="list-group-item list-group-item-action" href="#list-item-2">培训</a>
        <a class="list-group-item list-group-item-action" href="#list-item-3">工程</a>
        <a class="list-group-item list-group-item-action" href="#list-item-4">音乐</a>
      </div>
      <main class="col-md-12 col-xl-12 py-md-3 pl-md-5" role="main" data-target="#list-example">
        <div class="pos-f-t" id="list-item-1">
          <div class="collapse" id="navbarToggleExternalContent1">
            <div class="bg-dark p-4">
              <h5 class="text-white h4">Collapsed content</h5>
              <span class="text-muted">Toggleable via the navbar brand.</span>
            </div>
          </div>
          <nav class="navbar navbar-dark bg-dark">
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarToggleExternalContent1" aria-controls="navbarToggleExternalContent" aria-expanded="false" aria-label="Toggle navigation">
              <span class="navbar-toggler-icon"></span>直播
            </button>
          </nav>
        </div>
        <ul class="list-unstyled mt-3">
          <li class="media">
            <a href="/v1/flv?filepath=/static/download/VID20190517113037.mp4" target="-blank" title="直播" class="mr-3 w-25">
              <img src="/static/images/1.jpg" class="rounded w-100" alt="直播">
            </a>
            <div class="media-body">
              <h5 class="mt-0 mb-1">List-based media object</h5>
              Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante sollicitudin. Cras purus odio, vestibulum in vulputate at, tempus viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue felis in faucibus.
            </div>
          </li>
          <li class="media my-4">
            <a href="/v1/flv" target="-blank" title="直播" class="mr-3 w-25">
              <img src="/static/images/1.jpg" class="rounded w-100" alt="直播">
            </a>
            <div class="media-body">
              <h5 class="mt-0 mb-1">List-based media object</h5>
              Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante sollicitudin. Cras purus odio, vestibulum in vulputate at, tempus viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue felis in faucibus.
            </div>
          </li>
          <li class="media">
            <a href="/v1/flv" target="-blank" title="直播" class="mr-3 w-25">
              <img src="/static/images/1.jpg" class="rounded w-100" alt="直播">
            </a>
            <div class="media-body">
              <h5 class="mt-0 mb-1">List-based media object</h5>
              Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante sollicitudin. Cras purus odio, vestibulum in vulputate at, tempus viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue felis in faucibus.
            </div>
          </li>
        </ul>
        <div class="pos-f-t" id="list-item-2">
          <div class="collapse" id="navbarToggleExternalContent2">
            <div class="bg-dark p-4">
              <h5 class="text-white h4">Collapsed content</h5>
              <span class="text-muted">Toggleable via the navbar brand.</span>
            </div>
          </div>
          <nav class="navbar navbar-dark bg-dark">
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarToggleExternalContent2" aria-controls="navbarToggleExternalContent" aria-expanded="false" aria-label="Toggle navigation">
              <span class="navbar-toggler-icon"></span>培训
            </button>
          </nav>
        </div>
        <div class="row no-gutters bg-light shadow p-3 mb-5 bg-white rounded">
          <div class="col-md-6 mb-md-0 p-md-4">
            <!-- <img src="/static/images/1.jpg" class="w-100" alt="..."> -->
            <a href="/v1/flv" target="_blank" title="直播" class="mr-3 w-25">
              <img src="/static/images/1.jpg" class="rounded w-100" alt="直播">
            </a>
          </div>
          <div class="col-md-6 p-4 pl-md-0">
            <h5 class="mt-0">Columns with stretched link</h5>
            <p>Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante sollicitudin. Cras purus odio, vestibulum in vulputate at, tempus viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue felis in faucibus.</p>
            <a href="/v1/flv" class="stretched-link">Go somewhere</a>
          </div>
        </div>
        <div class="pos-f-t" id="list-item-3">
          <div class="collapse" id="navbarToggleExternalContent3">
            <div class="bg-dark p-4">
              <h5 class="text-white h4">Collapsed content</h5>
              <span class="text-muted">Toggleable via the navbar brand.</span>
            </div>
          </div>
          <nav class="navbar navbar-dark bg-dark">
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarToggleExternalContent23" aria-controls="navbarToggleExternalContent" aria-expanded="false" aria-label="Toggle navigation">
              <span class="navbar-toggler-icon"></span>工程
            </button>
          </nav>
        </div>
        <div class="row row-cols-1 row-cols-md-3">
          <div class="col mb-4">
            <div class="card h-100 shadow p-3 mb-5 bg-white rounded">
              <!-- <img src="/static/images/1.jpg" class="card-img-top" alt="..."> -->
              <a href="/v1/flv" target="_blank" title="直播">
                <img src="/static/images/1.jpg" class="card-img-top" alt="直播">
              </a>
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">This is a longer card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.</p>
                <a href="/v1/flv" class="stretched-link">Go somewhere</a>
              </div>
            </div>
          </div>
          <div class="col mb-4">
            <div class="card h-100 shadow p-3 mb-5 bg-white rounded">
              <img src="/static/images/1.jpg" class="card-img-top" alt="...">
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">This is a short card.</p>
              </div>
            </div>
          </div>
          <div class="col mb-4">
            <div class="card h-100 shadow p-3 mb-5 bg-white rounded">
              <img src="/static/images/1.jpg" class="card-img-top" alt="...">
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">This is a longer card with supporting text below as a natural lead-in to additional content.</p>
              </div>
            </div>
          </div>
          <div class="col mb-4">
            <div class="card h-100 shadow p-3 mb-5 bg-white rounded">
              <img src="/static/images/1.jpg" class="card-img-top" alt="...">
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">This is a longer card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.</p>
              </div>
            </div>
          </div>
        </div>
        <div class="pos-f-t" id="list-item-4">
          <div class="collapse" id="navbarToggleExternalContent4">
            <div class="bg-dark p-4">
              <h5 class="text-white h4">Collapsed content</h5>
              <span class="text-muted">Toggleable via the navbar brand.</span>
            </div>
          </div>
          <nav class="navbar navbar-dark bg-dark">
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarToggleExternalContent4" aria-controls="navbarToggleExternalContent" aria-expanded="false" aria-label="Toggle navigation">
              <span class="navbar-toggler-icon"></span>音乐
            </button>
          </nav>
        </div>
        <div id="carouselExampleCaptions" class="carousel slide" data-ride="carousel">
          <ol class="carousel-indicators">
            <li data-target="#carouselExampleCaptions" data-slide-to="0" class="active"></li>
            <li data-target="#carouselExampleCaptions" data-slide-to="1"></li>
            <li data-target="#carouselExampleCaptions" data-slide-to="2"></li>
          </ol>
          <div class="carousel-inner">
            <div class="carousel-item active">
              <!-- <img src="/static/images/1.jpg" class="d-block w-100" alt="..."> -->
              <a href="/v1/flv" target="_blank" title="直播">
                <img src="/static/images/1.jpg" class="card-img-top" alt="直播">
              </a>
              <div class="carousel-caption d-none d-md-block">
                <h5>First slide label</h5>
                <p>Nulla vitae elit libero, a pharetra augue mollis interdum.</p>
              </div>
            </div>
            <div class="carousel-item">
              <!-- <img src="/static/images/2.jpg" class="d-block w-100" alt="..."> -->
              <a href="/v1/flv" target="_blank" title="直播">
                <img src="/static/images/2.jpg" class="card-img-top" alt="直播">
              </a>
              <div class="carousel-caption d-none d-md-block">
                <h5>Second slide label</h5>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
              </div>
            </div>
            <div class="carousel-item">
              <img src="/static/images/3.jpg" class="d-block w-100" alt="...">
              <div class="carousel-caption d-none d-md-block">
                <h5>Third slide label</h5>
                <p>Praesent commodo cursus magna, vel scelerisque nisl consectetur.</p>
              </div>
            </div>
          </div>
          <a class="carousel-control-prev" href="#carouselExampleCaptions" role="button" data-slide="prev">
            <span class="carousel-control-prev-icon" aria-hidden="true"></span>
            <span class="sr-only">Previous</span>
          </a>
          <a class="carousel-control-next" href="#carouselExampleCaptions" role="button" data-slide="next">
            <span class="carousel-control-next-icon" aria-hidden="true"></span>
            <span class="sr-only">Next</span>
          </a>
        </div>
      </main>
    </div>
  </div>
  <!--   <div aria-live="polite" aria-atomic="true" class="d-flex justify-content-center align-items-center" style="min-height: 200px;">
    <div id="showtoasts" class="toast" role="alert" aria-live="assertive" aria-atomic="true">
      <div class="toast-header">
        <strong class="mr-auto">Bootstrap</strong>
        <small>11 mins ago</small>
        <button type="button" class="ml-2 mb-1 close" data-dismiss="toast" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="toast-body">
        Hello, world! This is a toast message.
      </div>
    </div>
  </div> -->
  <!-- <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script> -->
  <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script> -->
  <!-- <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script> -->
  
  <script>
  const TYPES = ['info', 'warning', 'success', 'error'],
    TITLES = {
      'info': 'Notice!',
      'success': 'Awesome!',
      'warning': 'Watch Out!',
      'error': 'Doh!'
    },
    CONTENT = {
      'info': 'Hello, world! This is a toast message.',
      'success': 'The action has been completed.',
      'warning': 'It\'s all about to go wrong',
      'error': 'It all went wrong.'
    },
    POSITION = ['top-right', 'top-left', 'top-center', 'bottom-right', 'bottom-left', 'bottom-center'];

  $.toastDefaults.position = 'top-center';
  $.toastDefaults.dismissible = true;
  $.toastDefaults.stackable = true;
  $.toastDefaults.pauseDelayOnHover = true;

  $('.snack').click(function() {
    var type = TYPES[Math.floor(Math.random() * TYPES.length)],
      content = CONTENT[type];
    $.snack(type, content);
  });

  $('.video').click(function() {
    var rng = Math.floor(Math.random() * 2) + 1,
      type = TYPES[Math.floor(Math.random() * TYPES.length)],
      title = TITLES[type],
      content = CONTENT[type];

    if (rng === 1) {
      $.toast({
        type: type,
        title: title,
        subtitle: '11 mins ago',
        content: content,
        delay: 5000
      });
    } else {
      $.toast({
        type: type,
        title: title,
        subtitle: '11 mins ago',
        content: content,
        delay: 5000,
        img: {
          src: 'https://via.placeholder.com/20',
          alt: 'Image'
        }
      });
    }
  });
  </script>
  <script type="text/javascript">
  // var offset = 60;
  // $('div a').click(function(event) {
  //   event.preventDefault();
  //   $($(this).attr('href'))[0].scrollIntoView();
  //   scrollBy(0, -offset);
  // });
  </script>
</body>

</html>