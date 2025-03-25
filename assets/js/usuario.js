$(document).ready(() => {

    $('#form-us-cadastro').on('submit', function(e){
        if ($("#us-id").val() == ''){
            cadastroUsuarioRequest(e)
        }else {
            editarUsuarioRequest(e)
        }
    })

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

    $('.acao').each(function(){
        $(this).on('click', (e) => {
            acao = e.target.dataset.acao
            id = e.target.dataset.userId
            switch (acao) {
                case 'editar':
                    editarUsuario(id)
                    break;
                
                case 'resetar':
                    $('#id-reset-senha').val(id)
                    $('#modal-reset-senha').modal('show')
                    break;

                case 'excluir':
                    if (confirm("Deletar Usuário?")){
                        $('#loading').modal('show');

                        fetch("/usuario/"+id,{
                            method: "DELETE",
                            headers: {
                                "Content-Type": "x-www-form-urlencoded"
                            }
                        }).then ((R) => {
                            if (R.status >= 400){
                                alert("Erro ao tentar excluir usuário!");
                            }else {
                                $('#loading').modal('hide');
                                location.reload();
                            }
                        })
                    }

                    break;
            
                default:
                    break;
            }
        })
    })

    $('#modal-usuarios').on('show.bs.modal', function(e){
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

    $('#modal-usuarios').on('hide.bs.modal', () => {
        location.reload() //Recarrega página ao fechar o modal para limpar campos e atualizar tela
    })

    $('#form-reset-senha').on('submit', resetSenha)

})

function cadastroUsuarioRequest(e){
    e.preventDefault();

    if ($('#us-senha').val() == $('#us-senha-2').val()){


        alterasenha = $('#us-trocar-senha').is(":checked") ? true : false

        if ($('#us-role').val() == 'invalido') {
            $('#us-role').css({
                "border": "1px solid red"
            })
            return
        }
        
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
    }else {
        $('#us-senha').addClass('is-invalid')
        $('#us-senha').removeClass('is-valid')
        $('#us-senha-2').addClass('is-invalid')
        $('#us-senha-2').removeClass('is-valid')
    }
}

function editarUsuarioRequest(e){
    e.preventDefault()

    const body = {
        "id": $('#us-id').val(),
        "nome": $('#us-nome').val(),
        "sobrenome": $('#us-sobrenome').val(),
        "nascimento": $('#us-nascimento').val(),
        "email": $('#us-email').val(),
        "telefone": $('#us-telefone').val(),
        "endereco": $('#us-endereco').val(),
        "role": $('#us-role').val(),
        "status": $('#us-status').is(":checked") ? 1 : 0,
        "alteraNextLogon": false,
    }

    fetch("/usuario",{
        method: "PUT",
        headers: {
            "Content-Type": "x-www-form-urlencoded",
        },
        body: JSON.stringify(body),

    }).then((R) => {
        if (R.status >= 400){
            alert("Erro ao salvar usuário")
        }else {
            alert("Usuário salvo com Sucesso")
            location.reload()
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

function editarUsuario(id){
                    
    fetch("/usuario/"+id).then((R) => {
        if (R.status >= 400){
            alert("Erro ao Buscar Dados do Usuário!")
            return
        }
        else {
            R.text().then((T) => {
                u = JSON.parse(T)
                //Prepara Modal para 
                $('#us-senha').parent().addClass('d-none')
                $('#us-senha').removeAttr("required")
                $('#us-senha-2').parent().addClass('d-none')
                $('#us-senha-2').removeAttr("required")
                $('#us-status').prop("disabled", false)
                $('#us-trocar-senha').parent().parent().parent().addClass('d-none')

                //Preenche os campos
                $('#us-id').val(u.id)
                $('#us-nome').val(u.nome)
                $('#us-sobrenome').val(u.sobrenome)
                $('#us-nascimento').val(u.nascimento.split("T")[0])
                $('#us-telefone').val(u.telefone)
                $('#us-email').val(u.email)
                $('#us-endereco').val(u.endereco)
                statusU = u.status == 1 ? true : false
                $('#us-status').prop("checked", statusU)
                $('#us-role').val(u.role)
                $('#us-login').val(u.login)

                $('#modal-usuarios').modal('show')
            })
        }
    })
}

function resetSenha(e){
    e.preventDefault()

    if ($('#nova-senha-reset-senha').val() != $('#test-nova-senha-reset-senha').val()){
        $('#nova-senha-reset-senha').addClass('is-invalid')
        $('#test-nova-senha-reset-senha').addClass('is-invalid')
        return
    }

    $('#nova-senha-reset-senha').removeClass('is-invalid')
    $('#test-nova-senha-reset-senha').removeClass('is-invalid')

    const dados = {
        "id": $('#id-reset-senha').val(),
        "senha": $('#nova-senha-reset-senha').val(),
        "alteraNextLogon": $('#troca-senha-reset-senha').is(":checked") ? true : false
    }

    
    fetch("/senha", {
        method: "POST",
        headers:{
            "Content-Type": "x-www-form-urlencoded"
        },
        body: JSON.stringify(dados)
    }).then((R) => {
        if (R.status >= 400){
            alert("Erro ao resetar senha!")
        }else {
            alert("Senha alterada com sucesso")
            location.reload()
        }
    })
}

