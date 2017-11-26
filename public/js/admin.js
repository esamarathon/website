/**
 *  Adds a confirm alert when clicking an object with a .confirmable class
 *  Uses the data-message attribute as the message in the confirm-box
 */
(function () {
    var confirmables = document.getElementsByClassName("confirmable");
    for (el of confirmables) {
        if (typeof el !== 'object') return;

        el.addEventListener('click', function (e) {
            if (!confirm(this.dataset.message)) {
                e.preventDefault();
            }
        })
    }
})();

