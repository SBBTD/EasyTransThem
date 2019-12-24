package pages

var (
	PageIndex    string
	PageNotFound string
)

func init() {
	//TODO
	PageIndex = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>EasyTransThem</title>
		<!-- <link href="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet"> -->
	</head>
	<body>
		<form method="post" action="/api/upload" enctype="multipart/form-data">
			<input type="file" name="file"/>
			<button type="submit">submit</button>
		</form>
		
		<!-- <script src="https://cdn.staticfile.org/jquery/2.1.1/jquery.min.js"></script> -->
		<!-- <script src="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script> -->
	</body>
</html>

`
}

func init() {
	//TODO
	PageNotFound = `<!DOCTYPE html>
<h1>404 NOT FOUND.</h1>
`
}
