package pages

var PageIndex string

func init() {
	//TODO
	PageIndex = `<!DOCTYPE html>
<html>
	
<head>
	<meta charset="utf-8">
	<title>EasyTransThem</title>
	<style>.hidden {display: none;}</style>
	<script src="https://cdn.staticfile.org/jquery/3.4.1/jquery.min.js"></script>
	<script src="https://cdn.staticfile.org/jquery.form/4.2.2/jquery.form.min.js"></script>
	<script>
		function upload() {
			var progress = $("#progressBar");
			var file = $("#uploadFile");
			var size = file[0].files[0].size
			if (size > 512 * 1024 * 1024)
				if (!confirm("This file is big(>512MB).Would you like to continue?"))
					return;
			$("#uploadForm").ajaxSubmit({
				url: "./api/upload",
				type: "post",
				beforeSend: function () {
					console.log("beforeSend");
					progress.removeClass("hidden");
				},
				uploadProgress: function (event, position, total, percentComplete) {
					console.log("uploadProgress:" + percentComplete + "%");
					progress[0].value = percentComplete;
				},
				success: function (res) {
					console.log(res)
					alert("Upload Success!");
				},
				error: function (XMLHttpRequest, textStatus, errorThrown) {
					console.log(errorThrown)
					alert("Upload FAILED!");
				}
			});
		}
	</script>
</head>

<body>
	<form id="uploadForm" enctype="multipart/form-data">
		<input id="uploadFile" type="file" name="file"/>
		<input type="button" onclick="upload();" value="提交"></input>
	</form>
	<progress class="hidden" id="progressBar" max="100" value="0"></progress>
</body>

</html>
`
}
