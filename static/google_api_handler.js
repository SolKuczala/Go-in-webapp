function onSignIn(googleUser) {
    const id_token = googleUser.getAuthResponse().id_token
    const req = new XMLHttpRequest()
    req.open('POST', '/auth/google')
    req.setRequestHeader('Content-Type', 'application/json')
    req.onload = function() {
        console.log('Signed in as: ' + req.responseText)
    }
    req.send(JSON.stringify({"id_token": id_token}))
}