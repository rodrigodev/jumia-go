// Mithril SPA
var filters = {
    country: "",
    state: ""
}

var fetchData = function() {
    m.request({
        method: "GET",
        url: "http://localhost:8081/phone",
    })
    .then(function(items) {
        this.App.countrySelect = [...new Set(items.map( row => ( row.country )))];
        Data.phones.list = items;
        redraw(new Event("init"));
    })
    .catch(function (reason) {
        console.log(reason);
    })
};

var Data = {
    phones: {
        list: [],
    }
}

var App = {
    oninit: fetchData,
    start : 1,
    pageSize : 10,
    max : 0,
    data: [],
    countrySelect: [],

    thead: m('thead',
        m('tr',
            m('td', 'Country'),
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
                    m("option", { value: "OK"}, "Valid phone numbers"),
                    m("option", { value: "NOK"}, "Invalid phone numbers"),
                ),
            ),
            m('table.ui.celled.table',
                this.thead,
                m("tbody", this.data.map(this.tr))
            )
        ]);
    },
};


var redraw = function (e) {
    if (e.type === "change") {
        filters[e.target.id] = e.target.value;
    }

    App.data = Data.phones.list.filter(function (item) {
        return (item.country === filters.country || filters.country === "") &&
            (item.state === filters.state || filters.state === "") ;
    });
};

var root = document.body;
m.mount(root, App);

