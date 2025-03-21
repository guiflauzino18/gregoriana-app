$(document).ready(() => {

    $('#form-us-cadastro').on('submit', cadastroUsuarioRequest)
    $('.form-control').each(function() {
        if ($(this).val() == ""){
            $(this).addClass('is-invalid')
            $(this).removeClass('is-valid')
        }else {
            $(this).addClass('is-valid')
            $(this).removeClass('is-invalid')
        }
        $(this).on('keyup', confereCampo)
    })

})

function cadastroUsuarioRequest(e){
    e.preventDefault();

    alterasenha = $('#us-trocar-senha').val() == 'on' ? true : false
    
    const body = {
        "nome": $('#us-nome').val(),
        "sobrenome": $('#us-sobrenome').val(),
        "nascimento": $('#us-nascimento').val(),
        "login": $('#us-login').val(),
        "senha": $('#us-senha').val(),
        "email": $('#us-email').val(),
        "telefone": $('#us-telefone').val(),
        "endereco": $('#us-endereco').val(),
        "role": $('#us-role').val(),
        "alteraNextLogon": alterasenha,
    }


    fetch("/usuario", {
        method: "POST",
        headers: {
            "Content-Type": "x-www-form-urlencoded",
        },
        body: JSON.stringify(body),
    }).then((R) => {
        if (R.status == 409){
            alert("Já exite um usuário com este Login")
        }
        else if (R.status >= 400){
            alert("Erro ao cadastrar Usuário")
        }else {
            alert("Usuário Cadastrado com sucesso!" )
            window.location.href = "/usuarios"
        }
    })
}

function confereCampo(){
    if ($(this).val() == ""){
        $(this).addClass('is-invalid')
        $(this).removeClass('is-valid')
    }else {
        $(this).addClass('is-valid')
        $(this).removeClass('is-invalid')
    }
}