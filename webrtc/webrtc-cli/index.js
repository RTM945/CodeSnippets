const peerMap = new Map()
const servers = { iceServers: [{ "urls": ["stun:stun.l.google.com:19302"] }] }
var stompClient = null

function connect(signal) {
    let token = 1//makeid(6)
    let type = 'answer'
    let socket = new SockJS(signal)
    stompClient = Stomp.over(socket)
    stompClient.connect({
        "token": token,
        "type": type,
    }, _ => {
        $('#fileRoot').prop('readonly', true)
        $('#dirBtn').off()
        stompClient.subscribe('/user/queue/onsdp', onsdp)
        stompClient.subscribe('/user/queue/oncandidate', oncandidate)
        
        showMsg("connect success, your token: " + token)
        document.title = document.title + ' ' + token
    }, _ => {
        showMsg("connect failue")
        stompClient = null
    })
}

function onsdp(msg) {
    let dto = JSON.parse(msg.body)
    console.log(dto)
    let remote = dto.remote
    let desc = JSON.parse(dto.value)
    if (desc.type == 'offer') {
        let pc = createPeer(remote)
        pc.setRemoteDescription(new RTCSessionDescription(desc))
            .then(_ => pc.createAnswer())
            .then(answer => pc.setLocalDescription(answer))
            .then(_ => stompClient.send("/app/sdp", {}, JSON.stringify(pc.localDescription)))
    }
}

function oncandidate(msg) {
    console.log(msg.body)
    let dto = JSON.parse(msg.body)
    let remote = dto.remote
    let candidate = dto.value
    let pc = peerMap.get(remote)
    if (pc) {
        pc.addIceCandidate(new RTCIceCandidate(JSON.parse(candidate)))
    }
}

function createPeer(remote) {
    let pc = new RTCPeerConnection(servers)
    console.log(pc)
    pc.onicecandidate = (event => event.candidate ? stompClient.send("/app/candidate", {}, JSON.stringify(event.candidate)) : console.log("Sent All Ice"))
    pc.ondatachannel = async ({ channel }) => {
        console.log("answer data channel created!")
        channel.onmessage = ({ data }) => handler(data)
        const files = await window.electron.invoke('listFiles')
        channel.send(JSON.stringify({ handler: 'listFiles', data: files }))
    }
    peerMap.set(remote, pc)
    return pc
}

function handler(data) {

}

function showMsg(msg) {
    window.electron.send('showMsg', msg)
}

function makeid(length) {
    let result = ''
    let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    let charactersLength = characters.length
    for (var i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength))
    }
    return result
}

window.electron.on('app-close', _ => {
    if (stompClient != null) {
        stompClient.disconnect(_ => {
            showMsg("Bye~")
        })
    }
    window.electron.send('closed')
})

$('#connectBtn').on('click', async _ => {
    let signalServer = $('#signal').val()
    signalServer = 'http://localhost:8080/signalling' //for test
    if (signalServer == '') {
        showMsg("signal server address can't be null!")
        return
    }
    const check = await window.electron.invoke('checkDir', $('#fileRoot').val())
    if (!check) {
        showMsg("wrong file root!")
        return
    }
    if (stompClient != null) {
        window.showMsg("already connected!")
        return
    }
    connect(signalServer)

})

$('#dirBtn').on('click', async _ => {
    const path = await window.electron.invoke('dir')
    if (path) {
        $('#fileRoot').val(path)
    }
})
