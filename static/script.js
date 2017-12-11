function requestDownload(url) {
    xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/download", true);
    //xhr.setRequestHeader("Content-type", "application/text");
    xhr.send(url);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            console.log("Done")
        }
    };
}

function sendDownload() {
    let url = document.getElementById("downloadInputBox").value;
    console.log(url);
    requestDownload(url)
}
