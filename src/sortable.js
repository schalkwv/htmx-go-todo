import Sortable from 'sortablejs';

htmx.onLoad(function (elt) {
    var sortables = elt.querySelectorAll("#items");
    for (var i = 0; i < sortables.length; i++) {
        var sortable = sortables[i];
        new Sortable(sortable, {
            draggable: ".draggable",
            handle: ".handle",
        });
    }
});