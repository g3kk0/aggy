"use strict";

function fetchValues() {
    fetch('/api')
	.then(resp => {
	    return resp.json();
        })
	.then(data => {
            document.getElementById("value").innerHTML = data.value.toFixed(2);
            document.getElementById("pnl").innerHTML = data.pnl.toFixed(2);
            document.getElementById("pnlPc").innerHTML = data.pnl_pc.toFixed(2);
            console.log(data);
        })
	.catch(err => {
            // Do something for an error here
        });
}
