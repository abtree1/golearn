{{define "lender"}}

	<div>
	    <ul class="nav nav-pills">
		  <li role="presentation" {{if .Home}} class="active" {{end}} ><a href="/">Home</a></li>
		  <li role="presentation" {{if .Category}} class="active" {{end}} ><a href="/category">Category</a></li>
		  <li role="presentation" {{if .Files}} class="active" {{end}} ><a href="/files">Files</a></li>
		</ul>
	</div>

{{end}}