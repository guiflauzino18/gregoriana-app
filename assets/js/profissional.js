$(document).ready(() => {
    // Submit do Form de Cadastro
    $('#form-pr-cadastro').on('submit', function(e){

        $('#loading').modal('show')
        
        if ($('#pr-id').val() == ''){
            cadastroProfissional(e)
            
        }else{
            console.log('editar')
        }
    })

    // Ao abrir modal de profissional já verifica campos validos e invalidos. Util na Edição
    $('#modal-profissional').on('show.bs.modal', function(e){
        $('.form-control').each(function() {
            if ($(this).val() == ""){
                $(this).addClass('is-invalid')
                $(this).removeClass('is-valid')
            }else {
                $(this).addClass('is-valid')
                $(this).removeClass('is-invalid')
            }
        })
    })
})

function cadastroProfissional(e){
    e.preventDefault()
    
    const dados = {
        "titulo": $('#pr-titulo').val(),
        "registro": $('#pr-registro').val(),
        "login": $('#pr-login').val()
    }

    fetch("/profissional",{
        method: "POST",
        headers: {
            "Content-Type": "x-www-form-urlencoded"
        },
        body: JSON.stringify(dados)

    }).then((R) => {
    
        $('#loading').modal('dispose') 

        if (R.status == 403){
            alert("Erro: Confira o campo login.")

        }else if (R.status >= 400){
            alert("Erro ao cadastrar usuário.")

        }else {
            location.reload();
        }

    })
}