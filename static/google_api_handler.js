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
    
    auth.access_token = auth.id_token
    req.send(JSON.stringify(auth))
}

function Edit(){
    const req = new XMLHttpRequest()
    req.open('GET', '/enter-profile-info')
    //reditect to enter profile info
}