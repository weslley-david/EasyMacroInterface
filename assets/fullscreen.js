function activateFullScreen() {
    var fullscreenElement = document.getElementById('fullscreen-element');
    if (fullscreenElement.requestFullscreen) {
        fullscreenElement.requestFullscreen();
    } else if (fullscreenElement.mozRequestFullScreen) {
        fullscreenElement.mozRequestFullScreen();
    } else if (fullscreenElement.webkitRequestFullscreen) {
        fullscreenElement.webkitRequestFullscreen();
    } else if (fullscreenElement.msRequestFullscreen) {
        fullscreenElement.msRequestFullscreen();
    }
}
