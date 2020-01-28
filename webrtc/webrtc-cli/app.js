const { app, BrowserWindow, Menu } = require('electron')
const path = require('path')
const url = require('url')

let win

app.once('ready', () => {
    win = new BrowserWindow({
        width: 400,
        height: 250,
        backgroundColor: "#D6D8DC",
        // resizable: false,
        // minimizable: false,
        // maximizable: false,
        show: false,
        webPreferences: {
            preload: path.join(__dirname, "preload.js")
        }
    })
    Menu.setApplicationMenu(null)

    win.loadURL(url.format({
        pathname: path.join(__dirname, 'index.html'),
        protocol: 'file:',
        slashes: true
    }))
    win.webContents.openDevTools()
    win.webContents.on('devtools-opened', () => {
        win.focus()
        setImmediate(() => {
            win.focus()
        })
    })
    win.once('ready-to-show', () => {
        win.show()
    })
})
