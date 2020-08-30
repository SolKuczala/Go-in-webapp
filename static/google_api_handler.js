function onSignIn(googleUser) {
    const auth = googleUser.getAuthResponse()
    const req = new XMLHttpRequest()
    req.open('POST', '/auth/google')
    req.setRequestHeader('Content-Type', 'application/json')
    
    req.onerror = function(params) {
        console.log('Signed in FAILED :(')
    }
    
    req.onload = function() {
        console.log('Signed in')
        window.location = "/enter-profile-info";
    }
    
    req.send(JSON.stringify(auth))
}