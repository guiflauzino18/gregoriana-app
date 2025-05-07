$(document).ready(() => {
    $('.btn-configuracao').each(function() {
        
        $(this).on('click', acessoConfiguracao)
    })
})

//Listenner dos botões da página de Configuração
function acessoConfiguracao(){
    switch ($(this).data('conf')) {
        case "usuario":
            window.location.href = "/usuarios"
            break;

        case "agenda":
            window.location.href = "/agenda"
            break

        case "profissional":
            window.location.href = "/profissional"
            break;

        case "empresa":
            window.location.href = "/empresa"
            break;
    
        default:
            break;
    }
}