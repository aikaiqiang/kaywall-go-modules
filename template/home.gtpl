<html>
<head>
    <title>Home</title>
</head>
<body>
Welcome !!!
<div>
    <p>上传文件</p>
    <form enctype="multipart/form-data" action="/upload" method="post">
        <input type="file" name="uploadfile" />
        <input type="hidden" name="token" value="{{.}}"/>
        <input type="submit" value="upload" />
    </form>
</div>
</body>
</html>