<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>

  <!-- <link rel="stylesheet" href="/static/css/index.css"> -->
  {{assets_css "/static/css/index.css"}}
  <!-- <script src="/static/js/base.js"></script> -->
  <!-- {{assets_js "/static/js/base.js"}} -->

</head>
<body>

  <p>
    {{.now}}
  </p>
  <p>
    {{date .now "Y-m-d H:i:s"}}
  </p>
  <p>
    {{.title}}
    <br>
    {{substr .title 2 8}}
  </p>
  <p>
    {{.html}}
  </p>
  <p>
    {{.html | html2str}}
    <br>
    {{html2str .html}}
  </p>
  <p>
    {{str2html .html}}
    <br>
    {{.html |str2html}}
    <br>
    {{htmlquote .html}}
  </p>

  <h2>获取端口号</h2>
    {{config "String" "httpport" ""}}


  <h2>获取map对象数据</h2>
    {{map_get .userinfo "username"}} <br>
    {{map_get .userinfo "age"}} <br>

    {{map_get .userinfo "hobby" 1}} <br>
    {{map_get .userinfo "hobby" 2}} <br>
    {{map_get .userinfo "hobby" 3}} <br>


    {{.title}} <br>
    {{.title | hello}} <br>
<hr>
    {{.unix}} <br>
    {{.unix | unixToDate}} <br>

    {{.str}} <br>

    {{.date}} <br>
    {{.date | dateToUnix}} <br>

</body>
</html>