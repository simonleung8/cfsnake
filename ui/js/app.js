var game = function() {
	
	var canvas = document.getElementById("gameCanvas");
	var token;

	const blockW = canvas.width / 100;
	const blockH = canvas.height / 100;	

	function refresh(){		
		$.get( "/update", function( data ) {
			//console.log(data)
			jsonStr = $.parseJSON( data )			
			redraw(jsonStr);
		});		
	}

	function getToken(){		
		$.get( "/newPlayer", function( data ) {
			token = data;
		});		
	}

	function redraw(jsonStr) {
		var ctx;
		ctx = canvas.getContext("2d");
		ctx.fillStyle="#FF0000";
		ctx.clearRect(0, 0, canvas.width, canvas.height);
		jsonStr.forEach(function(player){
			player.Snake.forEach(function(d) {
				cord = d.split(",")										
				ctx.fillRect(cord[0] * blockW, cord[1] * blockH, blockW, blockH);			
			})

		})
		
	}

	$(document).keydown(function(e) {
	    if (e.keyCode === 37) {
	    	$.get( "/left/" + token, function( data ) {});	
	    } else if (e.keyCode === 38) {
	    	$.get( "/up/" + token, function( data ) {});	
	    } else if (e.keyCode === 39) {
	    	$.get( "/right/" + token, function( data ) {});	
	    } else if (e.keyCode === 40) {
	    	$.get( "/down/" + token, function( data ) {});	
	    }
	});

	getToken();
	setInterval(refresh,200);
	//refresh()
}