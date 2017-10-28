document.getElementsByClassName('menu-toggle')[0]
    .addEventListener('click', toggleMenu);
document.getElementsByClassName('close')[0]
    .addEventListener('click', toggleMenu);

// Toggles menu active class
function toggleMenu (e) {
    var header = document.getElementsByTagName("header")[0]
    var toggleClass = 'active';
    header.className = header.className.indexOf(toggleClass) !== -1 ?
        header.className.replace(toggleClass, '') :
        (header.className + ' ' + toggleClass).trim();
}