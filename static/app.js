// Mithril SPA

var params = {}

var fetchData = function() {
    m.request({
        method: "GET",
        url: "http://localhost:8081/phone",
    })
    .then(function(items) {
        this.App.countrySelect = [...new Set(items.map( row => ( row.country )))];
        this.App.data = items;
    })
    .catch(function (reason) {
        console.log(reason);
    })
};

var App = {
    oninit: fetchData,
    start : 1,
    pageSize : 10,
    max : 0,
    data: [],
    countrySelect: [],

    buttons: m("div",
        m('button.ui.button', {onclick: this.previousPage}, "< Prev"),
        m('button.ui.button', {onclick: this.nextPage}, "Next >"),
    ),

    thead: m('thead',
        m('tr', 
            m('td', 'Conuntry'), 
            m('td', 'State'), 
            m('td', 'Country Code'), 
            m('td', 'Phone Num.'), 
        )
    ),

    tr : function(row) {
        return m('tr',
            m('td', row.country),
            m('td', row.state),
            m('td', row.country_code),
            m('td', row.phone),
        );
    },

    view: function() {
        return m("div.ui.top.attached.segment", [
            m("h1", "Phone numbers"),
            m("div.ui.secondary.menu",
                m("select.ui.dropdown.item#country", { onchange: redraw },
                    m("option", { value: ""}, "Select country"),
                    this.countrySelect.map(
                        function (country) {
                            return m("option", { value: country }, country);
                        }
                    ),
                ),
                m("select.ui.dropdown.item#state", { onchange: redraw },
                    m("option", { value: ""}, "State"),
                    m("option", { value: "ok"}, "Valid phone numbers"),
                    m("option", { value: "nok"}, "Invalid phone numbers"),
                ),
            ),
            m('table.ui.celled.table', 
                this.thead,
                m("tbody", this.data.map(this.tr))
            ),
            this.buttons
        ]);
    },

    nextPage: function(){
        this.start += this.pageSize;
        if (this.start > this.max){
            this.start -= this.pageSize;
        }
        this.setStart();
    },

    prevPage: function() {
        this.start = Math.max(this.start - this.pageSize, 1);
        this.setStart();
    },
};


var redraw = function (e) {
    if (e.target.value !== "") {
        params[e.target.id] = e.target.value;
    } else {
        delete params[e.target.id];
    }
};

var root = document.body;
m.mount(root, App);

