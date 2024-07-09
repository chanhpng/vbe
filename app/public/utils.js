const { app } = require('electron');
const path = require('path');

const osShortName = function() {
    switch (process.platform) {
        case "win32":
            return "win"
        case "darwin":
            return "mac"
        case "linux":
            return "linux"
        default:
            return null
    }
}();

module.exports = {
    resourcesPath: function () {
        if (!app.isPackaged) {
            return path.join(__dirname, "..", "resources", osShortName);
        }
        return process.resourcesPath;
    },
    defaultServerBinary: function () {
        if (!app.isPackaged) {
            return {
                "mac": path.join(__dirname, "..", "..", "dist", "vbe_darwin_amd64", "vbe"),
                "win": path.join(__dirname, "..", "..", "dist", "vbe_windows_amd64_v1", "vbe.exe"),
                "linux": path.join(__dirname, "..", "..", "dist", "vbe_linux_amd64", "vbe"),
            }[osShortName]
        }

        return {
            "mac": path.join(process.resourcesPath, "server", "vbe"),
            "win": path.join(process.resourcesPath, "server", "vbe.exe"),
            "linux": path.join(process.resourcesPath, "server", "vbe"),
        }[osShortName]
    },
    selectByOS: function (x) {
        return x[osShortName]
    },
}