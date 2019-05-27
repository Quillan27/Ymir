function newWorld() {
	var needWorld = true;
	const xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if(this.readyState == 4 && this.status == 200) {
			console.log(this.responseText)
			document.getElementById("map").innerHTML = this.responseText;
			worldName();
		}
	};
	xhttp.open("POST", "/newWorld");
	xhttp.send();
	console.log("request sent");
}

function worldName() {
	const xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if(this.readyState == 4 && this.status == 200) {
			console.log(this.responseText)
			document.getElementById("worldName").innerHTML = this.responseText;
		}
	};
	xhttp.open("POST", "/worldName");
	xhttp.send();
	console.log("request sent");
}

function elevationView() {
	const xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if(this.readyState == 4 && this.status == 200) {
			console.log(this.responseText)
			document.getElementById("map").innerHTML = this.responseText;
		}
	};
	xhttp.open("POST", "/elevationView");
	xhttp.send();
	console.log("request sent");
}

function topographyView() {
	const xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if(this.readyState == 4 && this.status == 200) {
			console.log(this.responseText)
			document.getElementById("map").innerHTML = this.responseText;
		}
	};
	xhttp.open("POST", "/topographyView");
	xhttp.send();
	console.log("request sent");
}

function biomeView() {
	const xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if(this.readyState == 4 && this.status == 200) {
			console.log(this.responseText)
			document.getElementById("map").innerHTML = this.responseText;
		}
	};
	xhttp.open("POST", "/biomeView");
	xhttp.send();
	console.log("request sent");
}
	
function politicalView() {
	const xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if(this.readyState == 4 && this.status == 200) {
			console.log(this.responseText)
			document.getElementById("map").innerHTML = this.responseText;
		}
	};
	xhttp.open("POST", "/politicalView");
	xhttp.send();
	console.log("request sent");
}

function climateView() {
	const xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if(this.readyState == 4 && this.status == 200) {
			console.log(this.responseText)
			document.getElementById("map").innerHTML = this.responseText;
		}
	};
	xhttp.open("POST", "/climateView");
	xhttp.send();
	console.log("request sent");
}