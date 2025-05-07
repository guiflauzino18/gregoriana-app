$(document).ready(() => {
    // Submit do Form de Cadastro
    $('#form-pr-cadastro').on('submit', function(e){

        $('#loading').modal('show')
        
        if ($('#pr-id').val() == ''){
            cadastroProfissional(e)
            
        }else{
            editarProfissionalRequest(e)
        }
    })

    // Ao abrir modal de profissional já verifica campos validos e invalidos. Util na Edição.
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

    //limpa campos ao fechar modal de Cadstro ou edição
    $('#modal-profissional').on('hide.bs.modal', function(){
        recarregaPagina();
    })

    $('.prof-acao').each(function(e){
        $(this).on('click', (e) => {
            acao = e.target.dataset.profAcao
            id = e.target.dataset.profId
            switch (acao) {
                case 'editar':
                    editarProfissional(id)
                    break;
            
                case 'excluir':
                    showConfirma("Deseja excluir este profissional?", ()=> {
                        excluirProfissional(id)
                        hideConfirma();
                    })
                    break;

                default:
                    break;
            }
        })
    })
})

function cadastroProfissional(e){
    e.preventDefault()
    showLoading();
    
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
    

        if (R.status == 403){
            showLoadingErro("Erro no cadastro do profissional! Confira o login");

        }else if (R.status >= 400){
            showLoadingErro("Erro ao cadastrar usuário.");

        }else {
            showLoadingSucesso("Profissional cadastrado com sucesso!");
        }

    })
}

function editarProfissional(id){
    //Busca Profissional pelo ID
    fetch("/profissional/"+id, {method: "GET"}).then((R) => {

        if (R.status >= 400){
            alert("Erro ao carregar dados do usuário")

        }else {
            R.text().then((T) => {
                profissional = JSON.parse(T)
                console.log(profissional)
                $('#pr-id').val(profissional.id)
                $('#pr-titulo').val(profissional.titulo)
                $('#pr-registro').val(profissional.registro)
                $('#pr-login').val(profissional.login)

                $('#modal-profissional').modal('show')
            })
        }
    })
}

function editarProfissionalRequest(e){
    e.preventDefault();

    const dados = {
        "id": $('#pr-id').val(),
        "titulo": $('#pr-titulo').val(),
        "registro": $('#pr-registro').val(),
        "login": $('#pr-login').val(),
    }

    fetch("/profissional", {
        method: "PUT",
        headers: {
            "Content-Type": "x-www-form-urlencoded"
        },
        body: JSON.stringify(dados)

    }).then((R)=> {
        if (R.status >= 400){
            showLoadingErro("Erro ao salvar!")

        }else {
            showLoadingSucesso("Profissional salvo com sucesso!");
        }
    })
}

function excluirProfissional(id){

    $('#loading').modal('show')

    fetch("/profissional/"+id, {method: "DELETE"}).then((R) => {

        if (R.status >= 400){
            showLoadingErro("Não foi possível excluir este profissional!")
            $('#loading').modal('hide');

        }else {
            showLoadingSucesso("Profissional exluído com sucesso!")
        }
    })

}