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

// Javascript dates are pretty silly so this function is necessary
function getMonthName(number) {
    var monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
    return monthNames[number]
}
function getWeekday(number) {
    var weekdays = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
    return weekdays[number]
}

function addLeadingZero(number) {
    return (number < 10) ? "0" + number : number
}

function addScheduleDate(el, date) {
    // Format: Saturday, July 22. 2017
    var timestamp = `${getWeekday(date.getDay())}, ${getMonthName(date.getMonth())} ${date.getDate()}, ${date.getFullYear()}`
    el.parentNode.parentNode.insertAdjacentHTML('beforebegin', `<tr><td colspan="99" class="new-date"><strong>${timestamp}</strong></td></tr>`);
}

(function () {
    var timestamps = document.getElementsByTagName('time')
    if (!timestamps.length) return

    var schedule = document.querySelector('.schedule')
    // Define prevDate in case we're on the schedule
    var prevDate

    // Loop through all the time-elements
    for (let el of timestamps) {
        if (typeof el !== 'object') return
        // Get the timestamp
        var utcTimestamp = el.getAttribute('datetime')
        // Create a local timestamp
        var date = new Date(utcTimestamp)

        // Define timestamp
        var timestamp

        // Formats the timestamp correctly
        if (el.className.indexOf('time-only') !== -1) {
            // Format: 20:08
            timestamp = `${addLeadingZero(date.getHours())}:${addLeadingZero(date.getMinutes())}`
        } else {
            // Format: November 24. 2017, 20:08
            timestamp = `${getMonthName(date.getMonth())} ${date.getDate()}. ${date.getFullYear()}, ${addLeadingZero(date.getHours())}:${addLeadingZero(date.getMinutes())}`
        }
        // Update the DOM
        el.innerHTML = timestamp

        // Add new schedule date row if we're on a new date
        // and on the schedule page
        if (schedule !== null && date.getDate() !== prevDate) {
            addScheduleDate(el, date)
            prevDate = date.getDate()
        }
    }
}())