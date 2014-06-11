var game = function() {
	
	const blockW = $("#gameCanvas").width() / 100;
	const blockH = $("#gameCanvas").height() / 100;
	var canvas = document.getElementById("gameCanvas");


	function refresh(){
		$(canvas).html();
		$.get( "/update", function( data ) {
			console.log(data)
			jsonStr = $.parseJSON( data )			
			redraw(jsonStr);
		});		
	}

	function redraw(jsonStr) {
		var ctx;
		jsonStr.forEach(function(player){
			player.Snake.forEach(function(d) {
				cord = d.split(",")		
				ctx = canvas.getContext("2d");
				ctx.fillStyle="#FF0000";
				ctx.fillRect(cord[0] * blockW,cord[1] * blockH,blockW, blockH);			
			})

		})
		
	}

	$(document).keydown(function(e) {
	    if (e.keyCode === 37) {
	    	$.get( "/left", function( data ) {});	
	    } else if (e.keyCode === 38) {
	    	$.get( "/up", function( data ) {});	
	    } else if (e.keyCode === 39) {
	    	$.get( "/right", function( data ) {});	
	    } else if (e.keyCode === 40) {
	    	$.get( "/down", function( data ) {});	
	    }
	});
	//setInterval(refresh,200)
	refresh()
}