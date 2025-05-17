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

    //Faz request à API para salvar a configuração da agenda
    $('#form-configura-agenda').on('submit', configuraAgendaRequest);

    //Ao marcar o dia libera os campos dos horários.
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

            case 'quarta-feira':
                if ($(`#${id}`).is(':checked')){
                    
                    $('#quarta-intervalo').prop('disabled', false);
                    $('#quarta-duracao').prop('disabled', false);
                    $('#quarta-inicio').prop('disabled', false);
                    $('#quarta-fim').prop('disabled', false);
                }else {
                    $('#quarta-intervalo').prop('disabled', true);
                    $('#quarta-duracao').prop('disabled', true);
                    $('#quarta-inicio').prop('disabled', true);
                    $('#quarta-fim').prop('disabled', true);
                }
                break;

            case 'quinta-feira':
                if ($(`#${id}`).is(':checked')){
                    
                    $('#quinta-intervalo').prop('disabled', false);
                    $('#quinta-duracao').prop('disabled', false);
                    $('#quinta-inicio').prop('disabled', false);
                    $('#quinta-fim').prop('disabled', false);
                }else {
                    $('#quinta-intervalo').prop('disabled', true);
                    $('#quinta-duracao').prop('disabled', true);
                    $('#quinta-inicio').prop('disabled', true);
                    $('#quinta-fim').prop('disabled', true);
                }
                break;

            case 'sexta-feira':
                if ($(`#${id}`).is(':checked')){
                    
                    $('#sexta-intervalo').prop('disabled', false);
                    $('#sexta-duracao').prop('disabled', false);
                    $('#sexta-inicio').prop('disabled', false);
                    $('#sexta-fim').prop('disabled', false);
                }else {
                    $('#sexta-intervalo').prop('disabled', true);
                    $('#sexta-duracao').prop('disabled', true);
                    $('#sexta-inicio').prop('disabled', true);
                    $('#sexta-fim').prop('disabled', true);
                }
                break;

            case 'sabado':
                if ($(`#${id}`).is(':checked')){
                    
                    $('#sabado-intervalo').prop('disabled', false);
                    $('#sabado-duracao').prop('disabled', false);
                    $('#sabado-inicio').prop('disabled', false);
                    $('#sabado-fim').prop('disabled', false);
                }else {
                    $('#sabado-intervalo').prop('disabled', true);
                    $('#sabado-duracao').prop('disabled', true);
                    $('#sabado-inicio').prop('disabled', true);
                    $('#sabado-fim').prop('disabled', true);
                }
                break;

            case 'domingo':
                if ($(`#${id}`).is(':checked')){
                    
                    $('#domingo-intervalo').prop('disabled', false);
                    $('#domingo-duracao').prop('disabled', false);
                    $('#domingo-inicio').prop('disabled', false);
                    $('#domingo-fim').prop('disabled', false);
                }else {
                    $('#domingo-intervalo').prop('disabled', true);
                    $('#domingo-duracao').prop('disabled', true);
                    $('#domingo-inicio').prop('disabled', true);
                    $('#domingo-fim').prop('disabled', true);
                }
                break;
        
            default:
                break;
        }
    })

    //adiciona clique do botao 
    $('.abreModalConfiguraHorarios').on('click', configuraHorarios)

    // Faz request para cadastro do status hora
    $('#form-cadastro-status-hora').on('submit', cadastroStatusHora)

    //Busca status das horas
    $('#modal-lista-status-hora').on("show.bs.modal", listaStatusHora)
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

//Abre Modal de configuração de agenda, buscando fias e horários se houver.
function configuraAgenda(id, idProfissional){
    $('#id-agenda').val(id)
    $('#id-profissional').val(idProfissional)

    fetch("/agenda/"+id).then ((R) => {
        if (R.status >= 400){
            showLoadingErro("Erro ao carregar dados da agenda.");
            return;
        }

        R.text().then((T) => {
            var agenda = JSON.parse(T)
            console.log(agenda)

            $('#agenda-nome').text(agenda.nome)
            $('#profissional-nome').text(agenda.NomeProfissional)

            for (dia of agenda.dias){
                switch (dia.nome){
                    case 'Segunda-Feira':
                        $('#segunda-id').val(dia.id);
                        $('#segunda-intervalo').val(dia.intervaloSessaoInMinutes);
                        $('#segunda-duracao').val(dia.duracaoSessaoInMinutes);
                        $('#segunda-inicio').val(dia.inicio);
                        $('#segunda-fim').val(dia.fim);
                        $('#segunda-feira').prop('checked', true);
                        $('#segunda-intervalo').prop('disabled', false);
                        $('#segunda-duracao').prop('disabled', false);
                        $('#segunda-inicio').prop('disabled', false);
                        $('#segunda-fim').prop('disabled', false);

                        $('#segunda-id-hora').attr('data-segunda-id', dia.id)
                        break;

                    case 'Terça-Feira':
                        $('#terca-id').val(dia.id);
                        $('#terca-intervalo').val(dia.intervaloSessaoInMinutes);
                        $('#terca-duracao').val(dia.duracaoSessaoInMinutes);
                        $('#terca-inicio').val(dia.inicio);
                        $('#terca-fim').val(dia.fim);
                        $('#terca-feira').prop('checked', true);
                        $('#terca-intervalo').prop('disabled', false);
                        $('#terca-duracao').prop('disabled', false);
                        $('#terca-inicio').prop('disabled', false);
                        $('#terca-fim').prop('disabled', false);
                        break;

                    case 'Quarta-Feira':
                        $('#quarta-id').val(dia.id);
                        $('#quarta-intervalo').val(dia.intervaloSessaoInMinutes);
                        $('#quarta-duracao').val(dia.duracaoSessaoInMinutes);
                        $('#quarta-inicio').val(dia.inicio);
                        $('#quarta-fim').val(dia.fim);
                        $('#quarta-feira').prop('checked', true);
                        $('#quarta-intervalo').prop('disabled', false);
                        $('#quarta-duracao').prop('disabled', false);
                        $('#quarta-inicio').prop('disabled', false);
                        $('#quarta-fim').prop('disabled', false);
                        break;

                    case 'Quinta-Feira':
                        $('#quinta-id').val(dia.id);
                        $('#quinta-intervalo').val(dia.intervaloSessaoInMinutes);
                        $('#quinta-duracao').val(dia.duracaoSessaoInMinutes);
                        $('#quinta-inicio').val(dia.inicio);
                        $('#quinta-fim').val(dia.fim);
                        $('#quinta-feira').prop('checked', true);
                        $('#quinta-intervalo').prop('disabled', false);
                        $('#quinta-duracao').prop('disabled', false);
                        $('#quinta-inicio').prop('disabled', false);
                        $('#quinta-fim').prop('disabled', false);
                        break;

                    case 'Sexta-Feira':
                        $('#sexta-id').val(dia.id);
                        $('#sexta-intervalo').val(dia.intervaloSessaoInMinutes);
                        $('#sexta-duracao').val(dia.duracaoSessaoInMinutes);
                        $('#sexta-inicio').val(dia.inicio);
                        $('#sexta-fim').val(dia.fim);
                        $('#sexta-feira').prop('checked', true);
                        $('#sexta-intervalo').prop('disabled', false);
                        $('#sexta-duracao').prop('disabled', false);
                        $('#sexta-inicio').prop('disabled', false);
                        $('#sexta-fim').prop('disabled', false);
                        break;

                    case 'Sábado':
                        $('#sabado-id').val(dia.id);
                        $('#sabado-intervalo').val(dia.intervaloSessaoInMinutes);
                        $('#sabado-duracao').val(dia.duracaoSessaoInMinutes);
                        $('#sabado-inicio').val(dia.inicio);
                        $('#sabado-fim').val(dia.fim);
                        $('#sabado-feira').prop('checked', true);
                        $('#sabado-intervalo').prop('disabled', false);
                        $('#sabado-duracao').prop('disabled', false);
                        $('#sabado-inicio').prop('disabled', false);
                        $('#sabado-fim').prop('disabled', false);
                        break;

                    case 'Domingo':
                        $('#domingo-id').val(dia.id);
                        $('#domingo-intervalo').val(dia.intervaloSessaoInMinutes);
                        $('#domingo-duracao').val(dia.duracaoSessaoInMinutes);
                        $('#domingo-inicio').val(dia.inicio);
                        $('#domingo-fim').val(dia.fim);
                        $('#domingo-feira').prop('checked', true);
                        $('#domingo-intervalo').prop('disabled', false);
                        $('#domingo-duracao').prop('disabled', false);
                        $('#domingo-inicio').prop('disabled', false);
                        $('#domingo-fim').prop('disabled', false);
                        break;
                }
            }

            $('#modal-configura-agenda').modal('show')
        })
    })

}

//Faz o request à API para salvar a agenda.
function configuraAgendaRequest(e) {
    e.preventDefault();

    showLoading();
    
    var dias = [];

    if ($('#segunda-feira').is(':checked')){
        var segunda = {
            "id": parseInt($('#segunda-id').val()),
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
            "nome":"Terça-Feira",
			"intervaloSessaoInMinutes": parseInt($('#terca-intervalo').val()),
			"duracaoSessaoInMinutes": parseInt($('#terca-duracao').val()),
			"inicio": $('#terca-inicio').val(),
			"fim": $('#terca-fim').val()
        }
        dias.push(terca)
    }

    if ($('#quarta-feira').is(':checked')){
        var quarta = {
            "nome":"Quarta-Feira",
			"intervaloSessaoInMinutes": parseInt($('#quarta-intervalo').val()),
			"duracaoSessaoInMinutes": parseInt($('#quarta-duracao').val()),
			"inicio": $('#quarta-inicio').val(),
			"fim": $('#quarta-fim').val()
        }
        dias.push(quarta)
    }

    if ($('#quinta-feira').is(':checked')){
        var quinta = {
            "nome":"Quinta-Feira",
			"intervaloSessaoInMinutes": parseInt($('#quinta-intervalo').val()),
			"duracaoSessaoInMinutes": parseInt($('#quinta-duracao').val()),
			"inicio": $('#quinta-inicio').val(),
			"fim": $('#quinta-fim').val()
        }
        dias.push(quinta)
    }

    if ($('#sexta-feira').is(':checked')){
        var sexta = {
            "nome":"Sexta-Feira",
			"intervaloSessaoInMinutes": parseInt($('#sexta-intervalo').val()),
			"duracaoSessaoInMinutes": parseInt($('#sexta-duracao').val()),
			"inicio": $('#sexta-inicio').val(),
			"fim": $('#sexta-fim').val()
        }
        dias.push(sexta)
    }

    if ($('#sabado').is(':checked')){
        var sabado = {
            "nome":"Sábado",
			"intervaloSessaoInMinutes": parseInt($('#sabado-intervalo').val()),
			"duracaoSessaoInMinutes": parseInt($('#sabado-duracao').val()),
			"inicio": $('#sabado-inicio').val(),
			"fim": $('#sabado-fim').val()
        }
        dias.push(sabado)
    }

    if ($('#domingo').is(':checked')){
        var domingo = {
            "nome":"Domingo",
			"intervaloSessaoInMinutes": parseInt($('#domingo-intervalo').val()),
			"duracaoSessaoInMinutes": parseInt($('#domingo-duracao').val()),
			"inicio": $('#domingo-inicio').val(),
			"fim": $('#domingo-fim').val()
        }
        dias.push(domingo)
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

function configuraHorarios(e){
    showLoading
    diaId = e.currentTarget.dataset.segundaId;

    fetch("/horas/"+id).then((R) => {
        if (R.status >= 400){
            showLoadingErro("Erro ao buscar horas do dia")
        }else {
            R.text().then((T) => {
                var horas = JSON.parse(T)
                console.log(horas)

                //Busca status das horas cadastrados no sistema
                fetch("agenda/horas/status").then((S) => {
                    if (S.status == 200){
                        S.text().then((U) => {
                            var statusList = JSON.parse(U)

                        $('.clonado').remove()

                        horas.forEach((hora) => {
                            $row = $('#clone-row-horas').clone(true)
                            $parent = $('#clone-row-horas').parent()

                            // Cria options do select com status cadastrados
                            $select = $row.find(".form-select")

                            $select.attr('data-hora-id', hora.id)

                            statusList.forEach((status) => {

                                $select.html(`${$select.html()} <option value="${status.id}">${status.nome}</option>`)
                            })

                            $select.on('change', alteraStatusHoraRequest)

                            //Seleciona status da hora no select
                            $select.val(hora.statusHora.id)

                            $row.removeAttr('id')
                            $row.removeClass('d-none')
                            $row.addClass("clonado")
                            $row.attr("data-id-hora", hora.id)
                            $row.find(".idHora").html(hora.id)
                            $row.find(".InicioHora").html(hora.inicio)
                            $row.find(".FimHora").html(hora.fim)
                            

                            $parent.append($row)
                        })


                        $('#modal-configura-agenda').modal('hide')
                        $('#modal-configura-horarios').modal('show')

                        })
                    }
                })
            })
        }
    })

}

function fechaModalHorariosEAbreDias(){
    $('#modal-configura-horarios').modal('hide')
    $('#modal-configura-agenda').modal('show')
}

function cadastroStatusHora(e){
    e.preventDefault();

    showLoading();

    const body = {
        "nome": $('#status-hora-nome').val(),
    }

    fetch("/status/hora", {
        method: "POST",
        headers: {
            "Content-Type": "x-www-form-urlencoded"
        },
        body: JSON.stringify(body)
        
    }).then((R) => {
        if (R.status == 409){
            showLoadingErro("Status já existe");
        }else if (R.status >= 400){
            showLoadingErro("Erro no cadastro do status");
        }else {
            R.text().then((T) => {
                showLoadingSucesso("Status cadastrado com sucesso")
                $('#modal-cadastro-status-hora').modal('dispose');
            })
        }
    })


}

function alteraStatusHoraRequest(e){

    
    idStatus = e.target.value
    idHora = e.target.dataset.horaId

    var body = {
        "idHora": idHora,
        "idStatusHora": idStatus,
    }

    fetch("/agenda/hora/status", {
        method: "PUT",
        headers: {
            "Content-Type": "x-www-form-urlencoded"
        },
        body: JSON.stringify(body)

    }).then((R) => {
        if (R.status != 200){
            showAlert("Erro ao alterar status da hora", "danger")
        }else {
            showAlert("Status alterado", "success")
        }
    })

    console.log(body)

}

function listaStatusHora(){

    showLoading();
    fetch("/agenda/horas/status").then((R) => {
        if (R.status >= 400){
            showLoadingErro("Erro ao buscar status")
        }else {

            R.text().then((T) => {

                var statusList = JSON.parse(T)

                $('.clonado').remove();

                statusList.forEach((status) => {
                    
                    $tr = $('#clone-row-status-hora').clone(true);
                    console.log($('#clone-row-status-hora'))
                    $tbody = $('#clone-row-status-hora').parent();

                    $tr.find('.id-status-hora').text(status.id);
                    $tr.find('.status-hora').text(status.nome);
                    $tr.find('ul').attr('data-status-id', status.id)

                    $tr.find('.dropdown-item').on('click', (e) => {
                        acao = e.target.dataset.acao
                        id = e.target.parentNode.parentNode.dataset.statusId

                        switch(acao){
                            case 'editar':
                                console.log('Editar')
                                break;

                            case 'excluir':
                                excluirStatusHora(id)
                                break;
                        }
                    })

                    $tr.removeClass('d-none');
                    $tr.addClass('clonado')

                    $tbody.append($tr)

                    hideLoading();
                })

            })


        }
    })

}

function excluirStatusHora(id){
    
    showConfirma("Deseja excluir este status?", () => {
        fetch("/agenda/hora/status/"+id, {
            method: "DELETE",

        }).then((R) => {
            showLoading();
            if (R.status >= 400){
                showLoadingErro("Erro ao excluir status")
            }else {
                R.text((T) => {

                    showLoadingSucesso(T.erro)
                })
            }
        })
    })
}