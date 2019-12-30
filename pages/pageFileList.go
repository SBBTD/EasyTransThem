package pages

//PageFileList = PageFileList1 + FileList + PageFileList2
var (
	PageFileList1 string
	PageFileList2 string
)

func init() {
	//TODO
	PageFileList1 = `<!DOCTYPE html>
<html>

	<head>
		<meta charset="utf-8">
		<title>FileList</title>
		<script src="https://cdn.staticfile.org/jquery/3.4.1/jquery.min.js"></script>
		<script>
			function del(f) {
			    if(confirm("Are you sure to DELETE?"))
					$.get("/api/delete?f=" + f, function(r) {
				    	alert(r);
				    	window.location.reload();
					})
			}
		</script>
	</head>

	<body>
		<table>
`
}

func init() {
	//TODO
	PageFileList2 = `
		</table>
	</body>

</html>
`
}
