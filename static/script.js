

function Check() {

    console.log("hello")
    var input = document.querySelectorAll(".input");

    for (var i = 0; i < input.length; i++){

    if (input[i].value=="") {
        alert("Please Fill input Values!");
        return false;
    }
    }    
}