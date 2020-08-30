function onSignIn(googleUser) {
    const auth = googleUser.getAuthResponse()
    const req = new XMLHttpRequest()
    req.open('POST', '/auth/google')
    req.setRequestHeader('Content-Type', 'application/json')
    req.onload = function() {
        console.log('Signed in as: ' + req.responseText)
    }
    req.send(JSON.stringify(auth))
}