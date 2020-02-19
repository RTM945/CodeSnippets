const peerMap = new Map()
const sendChannels = new Map()
const peer_sendChannel = new Map()
var sendChannelId = 0

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
            .then(_ => stompClient.send("/app/sdp", {}, JSON.stringify({ remote: remote, value: JSON.stringify(pc.localDescription) })))
    }
}

function oncandidate(msg) {
    console.log(msg.body)
    let dto = JSON.parse(msg.body)
    let remote = dto.remote
    let candidate = dto.value
    let pc = peerMap.get(remote)
    console.log('oncandidate get ' + pc)
    if (pc) {
        pc.addIceCandidate(new RTCIceCandidate(JSON.parse(candidate)))
    }
}

function createPeer(remote) {
    let pc = new RTCPeerConnection(servers)
    console.log(pc)
    peerMap.set(remote, pc)
    peer_sendChannel.set(remote, new Array())
    pc.onicecandidate = (event) => {
        if (event.candidate) {
            stompClient.send("/app/candidate", {}, JSON.stringify({ remote: remote, value: JSON.stringify(event.candidate) }))
        } else {
            console.log("Sent All Ice")
        }
    }
    pc.ondatachannel = async ({ channel }) => {
        const label = channel.label
        console.log(`answer data channel ${label} created!`)
        if (label == 'protocol') {
            channel.onmessage = ({ data }) => handler(data, channel, remote)
            const files = await listFiles()
            channel.send(JSON.stringify({ handler: 'listFiles', data: files }))
            pc.onconnectionstatechange = () => {
                switch (pc.connectionState) {
                    case "disconnected":
                    case "failed":
                    case "closed":
                        console.log(`answer data channel ${label} ${pc.connectionState}`)
                        peerMap.delete(remote)
                        let channelIds = peer_sendChannel.get(remote)
                        if (channelIds) {
                            channelIds.forEach(channelId => {
                                const channel = sendChannels.get(channelId)
                                if (channel) {
                                    channel.close()
                                }
                                sendChannels.delete(channelId)
                            })
                        }
                        peer_sendChannel.delete(remote)
                        break
                    default:
                        break
                }
            }
        }
    }
    return pc
}

async function handler(protocol, channel, remote) {
    console.log(protocol)
    protocol = JSON.parse(protocol)
    if (protocol.handler == 'listFiles') {
        const files = await listFiles(protocol.data)
        channel.send(JSON.stringify({ handler: 'listFiles', data: files }))
    }
    if (protocol.handler == 'download') {
        const fileInfo = await window.electron.invoke('fileInfo', protocol.data)
        console.log(fileInfo)
        if (!fileInfo) {
            console.log('fail access ' + protocol.data)
            return
        }
        channel.send(JSON.stringify({ handler: 'download', data: fileInfo }))
        download(protocol.data, fileInfo.size, remote)
    }
}

async function listFiles(dir) {
    const files = await window.electron.invoke('listFiles', dir)
    return files
}

async function download(filepath, size, remote) {
    let pc = peerMap.get(remote)
    if (!pc) {
        console.log('no peerconnection for ' + remote)
        return
    }
    let channel = pc.createDataChannel(filepath)
    const channelId = sendChannelId++
    sendChannels.set(channelId, channel)
    peer_sendChannel.get(remote).push(channelId)
    channel.binaryType = 'arraybuffer'
    channel.onclose = () => {
        console.log(`remote ${remote} channel ${channel.label} close`)
        sendChannels.delete(channelId)
    }
    const task = { filepath: filepath, size: size, received: 0, channelId: channelId, timeout: null }
    // tasks.push(task)
    window.electron.send('download', task)
}

window.electron.on('sendData', (task, data) => {
    const channel = sendChannels.get(task.channelId)
    if (channel && channel.readyState == 'open') {
        channel.send(data)
        console.log(`${task.received} / ${task.size}`)
        if (task.received == task.size) {
            console.log('send finished')
            channel.close()
        }
    }
})

window.electron.on('readFileErr', (task, err) => {
    console.log(err)
    const channel = sendChannels.get(task.channelId)
    if (channel && channel.readyState == 'open') {
        channel.send(JSON.stringify({ handler: 'downloaderr', data: err }))
        channel.close()
    }
})

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
