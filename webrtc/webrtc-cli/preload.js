const { dialog } = require('electron').remote;
const fs = require("fs");

window.dir = () => {
    return dialog.showOpenDialogSync({ properties: ['openDirectory'] })
}

window.checkDir = (dir) => {
    return fs.existsSync(dir) && fs.lstatSync(dir).isDirectory();
}

window.showMsg = (msg) => {
    dialog.showMessageBoxSync({message: msg});
}

window.showError = (msg) => {
    dialog.showErrorBox("webrtc-cli", msg)
}