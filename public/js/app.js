// Toggles menu active class
function toggleMenu (e) {
    var header = document.getElementsByTagName("header")[0]
    var toggleClass = 'active';
    var body = document.getElementsByTagName('body')[0];
    body.className = body.className.indexOf(toggleClass) !== -1 ?
    body.className.replace(toggleClass, '') :
    (body.className + ' ' + toggleClass).trim();
    
    header.className = header.className.indexOf(toggleClass) !== -1 ?
    header.className.replace(toggleClass, '') :
    (header.className + ' ' + toggleClass).trim();
}

// Adds the evenlisteners to the menu-toggle elements
document.getElementsByClassName('menu-toggle')[0]
    .addEventListener('click', toggleMenu);
document.getElementsByClassName('close')[0]
    .addEventListener('click', toggleMenu);

    // Enable livemode if element is present and window.Twitch is defined
if (!!document.getElementById('twitch-embed') && typeof window.Twitch !== 'undefined') {
    new Twitch.Embed("twitch-embed", {
        width: 854,
        height: 480,
        channel: "esamarathon"
    });
}
