
	{{template "post/menu.tpl" .}}
	<h2>post edit</h2>
  	<form method="POST" action="/post/{{.Data.Id}}">
		<div class="form-group">
			<label for="inputTitle" class="col-sm-2 control-label">標題</label>
			<div class="col-sm-10">
				<input type="text" class="form-control" id="title" placeholder="Title" name="title" value="{{.Data.Title}}">
			</div>
		</div>
		<div class="form-group">
			<label for="inputContent" class="col-sm-2 control-label">內容</label>
			<div class="col-sm-10">
				<textarea class="form-control" rows="5" id="content" name="content">{{.Data.Content}}</textarea>
			</div>
		</div>
		<div class="btn-group" role="group" aria-label="...">
			<button type="button" class="btn btn-default">確定</button>
		</div>
   	</form>

<script>
$(document).ready(function(){
    $("button").click(function() {
    	$.ajax({
    		url: $("form").attr("action"),
    		contentType: 'application/json; charset=utf-8',
    		dataType: 'json',
    		method: "PUT",
    		data: JSON.stringify($("form").serializeObject()),
    		success: function(e) {
    			console.log(e);
    		}
    	});
    });
});
</script>
