
$(document).on('click','.btn_submit', function(){
	console.log("here")
	var form = $(this).closest('form')
	var data = new FormData(form[ 0 ] );
	if (data.get("id") != "-1"){
		$(this).closest('form').toggle()
	}
	$.post('/comment/create',{ author: data.get("author"), comment: data.get("comment"), parentId: data.get("id")},
		function(returnedData){
			$('#comment' + data.get("id")).append($('<div id="comment' + returnedData + '">')
				.append($('<div class="commentCard">')
					.append($('<h3>').append('Author: ' + data.get("author")))
					.append($('<p>').append(data.get("comment"))))
				.append($('<button class="btnAnswer" id=' + returnedData + '>Reply</button>'))
				.append($('<button class="btnDelete" id=' + returnedData + '>Delete</button>'))
				.append($('<form id="addComment"  method=\'POST\' style="display: none" onsubmit="event.preventDefault();">')
					.append($('<div>')
						.append($('<label>Author:</label>'))
						.append($('<input type=\'text\' name=\'author\' class="author">')))
					.append($('<div>')
						.append($('<label>Comment:</label>'))
						.append($('<textarea name=\'comment\' class="comment"></textarea>')))
					.append($('<div>')
						.append($('<input type=\'hidden\' name=\'id\' value=' + returnedData+ '>'))
						.append($('<button class="btn_submit">Publish </button>')))
				))
			$('.comment').val('')
			$('.author').val('')

		});
});
$(document).ready(function() {
	$("#comment-1").on('click', '.btnDelete', function () {
		$.post('/delete', { id: this.id},
			function(returnedData){
			}).fail(function(){
			console.log("error");
		});
		$(this).closest('div').remove();
	});
});
$(document).on('click','.btnAnswer', function () {
	$('.comment').val('')
	$('.author').val('')
	$(this).closest("div").find('form').toggle()
});
function nested(id, arr) {
	var html = ''
	for (var i=0; i<arr.length; i++) {
		if (id == arr[i]["ParentID"]) {
			html = html + '<div id="comment'+ arr[i]["ID"] + '">' +
				'<div class="commentCard">' +
				'<h3> Author:'+ arr[i]["Author"] + '</h3>' +
				'<p>' + arr[i]["Comments"] +'</p>' +
				'</div>' +
				'<button class="btnAnswer" id="' + arr[i]["ID"] +'">Reply</button>' +
				'<button class="btnDelete" id="' + arr[i]["ID"] +'">Delete</button>' +
				'<form id="addComment"  method=\'POST\' style="display: none" onsubmit="event.preventDefault();">' +
				'<div>' +
				'<label>Author:</label>' +
				'<input type=\'text\' name=\'author\' class="author">' +
				'</div>' +
				'<div>' +
				'<label>Comment:</label>' +
				'<textarea name=\'comment\' class="comment"></textarea>' +
				'</div>' +
				'<div >' +
				'<input type=\'hidden\' name=\'id\' value="' + arr[i]["ID"] +'">' +
				'<button class="btn_submit">Publish</button>' +
				'</div>' +
				'</form>'
				html += nested(arr[i]["ID"], arr)
				html += '</div>'
		}
	}
	return html

}