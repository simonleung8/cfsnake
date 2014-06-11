var game = function() {
	
	const blockW = $("#gameCanvas").width() / 100;
	const blockH = $("#gameCanvas").height() / 100;



	function refresh(){
		$.get( "/update", function( data ) {
			console.log(data)
			jsonStr = $.parseJSON( data )			
			redraw(jsonStr);
		});		
	}

	function redraw(jsonStr) {
		var c = document.getElementById("gameCanvas");
		var ctx;
		jsonStr.forEach(function(player){
			player.Snake.forEach(function(d) {
				cord = d.split(",")		
				ctx = c.getContext("2d");
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
	refresh();
}