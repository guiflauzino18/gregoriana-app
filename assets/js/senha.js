$(document).ready(() => {
    $('#form-troca-senha').on('submit', trocaSenha)
})

function trocaSenha(e){
    e.preventDefault();
    if ($('#new-senha').val() != $('#test-new-senha').val()){
        $('#new-senha').addClass('is-invalid');
        $('#test-new-senha').addClass('is-invalid');
        return
    }else {
   
        $('#new-senha').removeClass('is-invalid')
        $('#test-new-senha').removeClass('is-invalid')
    }

    showLoading();

    const data = {
        "id": $('#id').val(),
        "senha": $('#new-senha').val(),
        "alteraNextLogon": false,
    }

    fetch("/senha",{
        method: "POST",
        headers: {
            "Content-Type": "x-www-form-urlencoded"
        },
        body: JSON.stringify(data)

    }).then((R) => {
        if (R.status >= 400){
            showLoadingErro("Erro ao alterar senha!")

        }else {
            
            window.location.href = "/"
        }
    })

}