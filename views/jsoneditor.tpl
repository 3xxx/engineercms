<!DOCTYPE HTML>
<html>

<head>
  <!-- when using the mode "code", it's important to specify charset utf-8 -->
  <meta http-equiv="Content-Type" content="text/html;charset=utf-8">
  <link href="/static/jsoneditor/dist/jsoneditor.min.css" rel="stylesheet" type="text/css">
  <script src="/static/jsoneditor/dist/jsoneditor.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <style type="text/css">
  /*body {
    font: 11pt arial;
  }

  #jsoneditor {
    width: 500px;
  }*/

  body {
    width: 600px;
    font: 11pt sans-serif;
  }

  #jsoneditor {
    width: 100%;
    height: 800px;
  }
  </style>
  <!-- <h1>JSON schema validation</h1> -->
  <!-- <p>
  This example demonstrates JSON schema validation. The JSON object in this example must contain properties <code>firstName</code> and <code>lastName</code>, can can optionally have a property <code>age</code> which must be a positive integer.
</p>
<p>
  See <a href="http://json-schema.org/" target="_blank">http://json-schema.org/</a> for more information.
</p>

<div id="jsoneditor"></div> -->
  <!-- <script>
  var schema = {
    "title": "Example Schema",
    "type": "object",
    "properties": {
      "firstName": {
        "type": "string"
      },
      "lastName": {
        "type": "string"
      },
      "gender": {
        "enum": ["male", "female"]
      },
      "age": {
        "description": "Age in years",
        "type": "integer",
        "minimum": 0
      }
    },
    "required": ["firstName", "lastName"]
  };

  var json = {
    firstName: 'John',
    lastName: 'Doe',
    gender: null,
    age: 28
  };

  var options = {
    schema: schema
  };

  // create the editor
  var container = document.getElementById('jsoneditor');
  var editor = new JSONEditor(container, options, json);
</script> -->
  <!-- <p>
  Switch editor mode using the mode box.
  Note that the mode can be changed programmatically as well using the method
  <code>editor.setMode(mode)</code>, try it in the console of your browser.
</p> -->
  <p>
    <button id="setJSON">Save JSON</button>
    <!-- <button id="getJSON">Get JSON</button> -->
  </p>
  <div class="col-lg-12">
    <div class="col-lg-6" id="jsoneditor"></div>
    <!-- <div class="col-lg-6" id="getjsoneditor"></div> -->
  </div>
  <script>
  var container = document.getElementById('jsoneditor');

  var options = {
    mode: 'code',
    modes: ['code', 'form', 'text', 'tree', 'view'], // allowed modes
    onError: function(err) {
      alert(err.toString());
    },
    onModeChange: function(newMode, oldMode) {
      console.log('Mode switched from', oldMode, 'to', newMode);
    }
  };

  var json = {
    // "array": [1, 2, 3],
    // "boolean": true,
    // "null": null,
    // "number": 123,
    // "object": { "a": "b", "c": "d" },
    // "string": "Hello World"
  };
  var editor = new JSONEditor(container, options, json);
  $(document).ready(function() {
    $.ajax({
      type: "get",
      url: "/v1/admin/getwxprojectconfig",
      data: { projectid: {{.ProjectId}} },
      success: function(data, status) {
        // alert("添加“"+data+"”成功！(status:"+status+".)");
        // json = data
        // editor = new JSONEditor(container, options, json);
        editor.set(data);
      }
    });
  })


  // create the editor
  // var container1 = document.getElementById('getjsoneditor');
  // var options = {
  //   mode: 'view'
  // };
  // var editor1 = new JSONEditor(container1, options);
  // set json
  document.getElementById('setJSON').onclick = function() {
    var json = editor.get();
    // alert(JSON.stringify(json, null, 2));
    // editor1.set(json);
    // var projectconfig=JSON.stringify(json)
    $.ajax({
      type: "POST",
      url: "/v1/admin/putwxprojectconfig?projectid="+{{.ProjectId}},
      data: JSON.stringify(json),//{ projectid: "1",projectconfig:projectconfig},
      success: function(data, status) {
        if (data.info=='SUCCESS'){
          alert("存储“"+data.json+"”成功！(status:"+status+".)");
        }else{
          alert(data.info);
        }
      }
    });
  };

  // get json
  // document.getElementById('getJSON').onclick = function() {
  //   var container = document.getElementById('getjsoneditor');
  //   var options = {
  //   };
  //   var editor = new JSONEditor(container, options, json);
  // };
  </script>
  </body>

</html>