function mostrarAlerta(action) {
  fetch(`http://192.168.0.110:8080/ping/${action}`)
}
