
	{{template "post/menu.tpl" .}}
<h2>post edit</h2>
<div class="form-group">
	<label for="inputTitle" class="col-sm-2 control-label">標題</label>
	<div class="col-sm-10">
		{{.Data.Title}}
	</div>
</div>
<div class="form-group">
	<label for="inputContent" class="col-sm-2 control-label">內容</label>
	<div class="col-sm-10">
		{{.Data.Content}}
	</div>
</div>

