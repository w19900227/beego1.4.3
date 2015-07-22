
	{{template "post/menu.tpl" .}}
	<div class="panel panel-default">
		<!-- Default panel contents -->
		<div class="panel-heading">清單</div>

		<!-- Table -->
		<table class="table">
			<tr>
				<td width="5%"></td>
				<td width="5%">id</td>
				<td >標題</td>
			</tr>
			<form action="/post">
			{{range .Data}}
			<tr>
				<td width="5%"><input type="checkbox" name="id[]" value="{{.Id}}"></td>
				<td width="5%">{{.Id}}</td>
				<td >
					{{.Title}}
					<a href="/post/{{.Id}}">顯示</a>
					<a href="/post/edit/{{.Id}}">編輯</a>
				</td>
			</tr>
			{{end}}
			</form>
		</table>
	</div>	

<script>
$(document).ready(function(){
    $("#del").click(function() {
    	$.ajax({
    		url: $("form").attr("action"),
    		contentType: 'application/json; charset=utf-8',
    		dataType: 'json',
    		method: "DELETE",
    		data: JSON.stringify($("form").serializeObject()),
    		success: function(e) {
    			console.log(e);
    		}
    	});
    });
});
</script>
