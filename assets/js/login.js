$('#form-login').on('submit', logar)


function logar(e){
    e.preventDefault();

    const body = {
        "login": $('#login').val(),
        "password": $('#password').val()
    }

    fetch("/login", {
        method: "POST",
        headers: {
            "Content-Type": "x-www-form-urlencoded",
        },
        body: JSON.stringify(body),

    }).then((R) => {
        if (R.status == 403){
            $('#login-erro').removeClass('invisible')
        }else if (R.status >= 400) {
            console.log("Erro ao fazer requisição à API:\n")
        }else {
            R.text().then((T) => {
                $('#login-erro').addClass('invisible')
                window.location.href = "/home"
            })
        }
    })
}