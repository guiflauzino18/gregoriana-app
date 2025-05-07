$(document).ready(() => {
    //submit form cadastro de agenda
    $('#form-ag-cadastro').on('submit', cadastroAgenda)

    // clique de ação da agenda
    $('.ag-acao').each(function(){
        $(this).on('click', (e) => {
            id = e.target.dataset.agId
            acao = e.target.dataset.agAcao
            
            switch (acao) {
                case 'editar':
                    editarAgenda(id);
                    break;
                
                case 'bloquear':
                    showConfirma("Deseja Bloquear esta agenda?", () => {
                        bloquearAgenda(id);
                        hideConfirma();
                    });
                    break;

                case 'configurar':
                    idProfissional = e.target.dataset.prId
                    configuraAgenda(id, idProfissional);
                    break;

                case 'excluir':
                    showConfirma("Deseja excluir esta agenda?",  () => {
                        excluirAgenda(id);
                        hideConfirma();
                    })
                    break;
            
                default:
                    break;
            }
        })
    })

    //Ao abrir Modal de cadastro busca lista de profissionais para exibir no select
    $('#modal-agenda').on('show.bs.modal', () => {
        showLoading();
        fetch("/profissionais", {method: "GET"}).then((R) => {
            if (R.status >=400){
                showLoadingErro("Erro ao carregar dados de profissionais!")
                
            }else {
                R.text().then((T) => {
                    const profissionais = JSON.parse(T)

                    profissionais.forEach(profissional => {
                        $option = $('#clone-option').clone(true);
                        $option.val(profissional.id)
                        $option.text(profissional.nome)

                        hideLoading();

                        $('#profissional-id').append($option)
                    });
                })
                
            }
        })
    })

    $('#form-configura-agenda').on('submit', configuraAgendaRequest);

    $('.dia').on('change', (e) => {
        id = e.target.id;

        switch (id) {
            case 'segunda-feira':
                if ($(`#${id}`).is(':checked')){
                    
                    $('#segunda-intervalo').prop('disabled', false);
                    $('#segunda-duracao').prop('disabled', false);
                    $('#segunda-inicio').prop('disabled', false);
                    $('#segunda-fim').prop('disabled', false);
                }else {
                    $('#segunda-intervalo').prop('disabled', true);
                    $('#segunda-duracao').prop('disabled', true);
                    $('#segunda-inicio').prop('disabled', true);
                    $('#segunda-fim').prop('disabled', true);
                }
                break;

                case 'terca-feira':
                    if ($(`#${id}`).is(':checked')){
                        
                        $('#terca-intervalo').prop('disabled', false);
                        $('#terca-duracao').prop('disabled', false);
                        $('#terca-inicio').prop('disabled', false);
                        $('#terca-fim').prop('disabled', false);
                    }else {
                        $('#terca-intervalo').prop('disabled', true);
                        $('#terca-duracao').prop('disabled', true);
                        $('#terca-inicio').prop('disabled', true);
                        $('#terca-fim').prop('disabled', true);
                    }
                    break;
        
            default:
                break;
        }
    })

})

function cadastroAgenda(e){
    e.preventDefault();

    showLoading();

    if ($('#profissional-id').val() == 'invalido') {
        showLoadingErro("Selecione o profissional")
        return
    }

    const dados = {
        "nome": $('#ag-nome').val(),
        "idProfissional": parseInt($('#profissional-id').val())
    }

    fetch("/agenda", {
        method: "POST",
        headers: {
            "Content-Type": "x-www-form-urlencoded"
        },
        body: JSON.stringify(dados)

    }).then((R) => {

        if (R.status == 409){
            showLoadingErro("Já existe uma agenda com este nome!");

        }else if (R.status >= 400){
            showLoadingErro("Erro ao criar agenda!");

        }else {
            showLoadingSucesso("Agenda criada com Sucesso!");
            console.log("sucesso")
        }
    })
}

function excluirAgenda(id){
    showLoading();

    fetch("/agenda/"+id, {
        method: "DELETE"

    }).then((R) => {
        if (R.status >= 400){
            showLoadingErro("Erro ao excluir Agenda!")
        }else {
            showLoadingSucesso("Agenda excluída com sucesso!")
        }
    })
    
}

function editarAgenda(id){

}

function bloquearAgenda(id){
    console.log(`Agenda ${id} bloqueada!`);
}

function configuraAgenda(id, idProfissional){
    $('#id-agenda').val(id)
    $('#id-profissional').val(idProfissional)
    $('#modal-configura-agenda').modal('show')
}

function configuraAgendaRequest(e) {
    e.preventDefault();

    showLoading();
    
    var dias = [];

    if ($('#segunda-feira').is(':checked')){
        var segunda = {
            "nome":"Segunda-Feira",
			"intervaloSessaoInMinutes": parseInt($('#segunda-intervalo').val()),
			"duracaoSessaoInMinutes": parseInt($('#segunda-duracao').val()),
			"inicio": $('#segunda-inicio').val(),
			"fim": $('#segunda-fim').val()
        }
        dias.push(segunda)
    }
    
    if ($('#terca-feira').is(':checked')){
        var terca = {
            "nome":"Terca-Feira",
			"intervaloSessaoInMinutes": parseInt($('#terca-intervalo').val()),
			"duracaoSessaoInMinutes": parseInt($('#terca-duracao').val()),
			"inicio": $('#terca-inicio').val(),
			"fim": $('#terca-fim').val()
        }
        dias.push(terca)
    }

    dados = {
        "idAgenda": parseInt($('#id-agenda').val()),
        "idProfissional": parseInt($('#id-profissional').val()),
        "dias": dias
    }

    console.log(dados)

    fetch("/agenda/configure", {
        method: "PUT",
        headers:{
            "Content-Type": "x-www-form-urlencoded"
        },
        body: JSON.stringify(dados)
    }).then((R)=> {

        if (R.status >= 400){
            showLoadingErro("Erro ao salvar!")
        }else {
            showLoadingSucesso("Agenda configurada");
        }
    })
    

}