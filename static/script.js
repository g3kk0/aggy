"use strict";

function fetchValues() {
    fetch('/api')
	.then(resp => {
	    return resp.json();
        })
	.then(data => {
	    document.getElementById("loading").style.display = "none";
            document.getElementById("title").style.display = "flex";
            document.getElementById("subtitle-1").style.display = "flex";
            document.getElementById("subtitle-2").style.display = "flex";
            if (data.pnl > 0) {
                document.getElementById("pnlPc").style.color = "#00C853";
                document.getElementById("pnl").style.color = "#00C853";
            } else {
                document.getElementById("pnlPc").style.color = "#D50000";
                document.getElementById("pnl").style.color = "#D50000";
            }
            document.getElementById("pnlPc").innerHTML = data.pnl_pc.toFixed(2) + " %";
            document.getElementById("value").innerHTML = "£" + data.value.toFixed(2);
            document.getElementById("pnl").innerHTML = "£" + data.pnl.toFixed(2);
            console.log(data);
        })
	.catch(err => {
	    console.log(err);
        });
}
