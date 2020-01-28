$('#dirBtn').on('click', () => {
    let path = window.dir()
    if (path != undefined) {
        $('#fileRoot').val(path)
    }
});

let stompClient = null;
function connect(signal) {
    let token = makeid(6);
    let type = 'answer';
    let socket = new SockJS(signal);
    stompClient = Stomp.over(socket);
    stompClient.connect({
        "token": token,
        "type": type,
    }, function (frame) {
        stompClient.subscribe('/user/queue/onsdp', function (msg) {
            stompClient.send("/app/sdp", {}, "answer's sdp");
        });
        window.showMsg("connect success, you token: " + token)
        document.title = document.title + ' ' + token
    }, function (frame) {
        window.showMsg("connect failue");
        stompClient = null
    });
}

$('#connectBtn').on('click', () => {
    let signalServer = $('#signal').val()
    signalServer = "http://localhost:8080/signalling" //for test
    if (signalServer == '') {
        window.showMsg("signal server address can't be null!")
        return
    }
    if (!window.checkDir($('#fileRoot').val())) {
        window.showMsg("wrong file root!")
        return
    }
    if(stompClient != null) {
        window.showMsg("already connected!")
        return
    }
    connect(signalServer)
});

function makeid(length) {
    let result = '';
    let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let charactersLength = characters.length;
    for (var i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}