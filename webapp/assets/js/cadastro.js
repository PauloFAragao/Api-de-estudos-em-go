$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) 
{
    evento.preventDefault();
    //console.log("Dentro da função usuário!")
    
    if ($('#senha').val() != $('#confirmar-senha').val()){
        //alert("As senhas não coincidem!");
        Swal.fire('Ops...', 'As senhas não coincidem!', 'error');
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data:{
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {// status resposta da pagina 201, 200, 204
        //alert("usuário cadastrado com sucesso!");
        Swal.fire('Sucesso', 'Usuário cadastrado com sucesso', 'success')
        .then(function(){
            $.ajax({
                url: "/login",
                method: "POST",
                data:{
                    email: $('#email').val(),
                    senha: $('#senha').val()
                }
            }).done(function(){
                window.location = "/home";
            }).fail(function(erro){// status resposta da pagina 400, 404, 401, 403, 500
                Swal.fire('Ops...', 'Erro ao autenticar usuário!', 'error');
            });
        })
    }).fail(function(erro){// status resposta da pagina 400, 404, 401, 403, 500
        //console.log(erro);
        //alert("Erro ao cadastrar usuário!");
        Swal.fire('Ops...', 'Erro ao cadastrar usuário!', 'error');
    });

}