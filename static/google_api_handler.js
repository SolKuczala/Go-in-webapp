function onSignIn(googleUser) {
    var id_token = googleUser.getAuthResponse().id_token;

    var req2Server = new XMLHttpRequest();
    req2Server.open('POST', '/auth/google');
    req2Server.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    req2Server.onload = function() {
        console.log('Signed in as: ' + req2Server.responseText);
    };

    req2Server.send('idtoken=' + id_token);
}