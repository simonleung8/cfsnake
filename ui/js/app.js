var game = function() {
	
	const blockW = $("#gameBoard").width() / 100;
	const blockH = $("#gameBoard").height() / 100;

	function refresh(){
		$.get( "/update", function( data ) {
		  console.log(JSON.stringify(data))
		});

		var tmp = ["1,1","1,2","1,3"]
		var c = document.getElementById("gameBoard");
		var ctx = c.getContext("2d");
		ctx.rect(20,20,150,100);
		ctx.stroke();
	}

	refresh();
}