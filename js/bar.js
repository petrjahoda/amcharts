const clicked = document.getElementById('clicked');
const userDowntimes = document.getElementById('usersDowntimes');
const downtimeUsers = document.getElementById('downtimeUsers');


let data = {
    Start: Date.now(),
    End: Date.now(),
    Data: "all",
};
fetch("/get_downtime_raw_data", {
    method: "POST",
    body: JSON.stringify(data)
}).then((response) => {
    response.text().then(function (data) {
        am4core.ready(function () {
            const chart = am4core.create("downtimes", am4charts.XYChart);
            chart.padding(40, 40, 40, 40);

            const categoryAxis = chart.yAxes.push(new am4charts.CategoryAxis());
            categoryAxis.renderer.grid.template.location = 0;
            categoryAxis.dataFields.category = "Name";
            categoryAxis.renderer.inversed = true;
            categoryAxis.renderer.grid.template.disabled = true;

            const valueAxis = chart.xAxes.push(new am4charts.ValueAxis());
            valueAxis.min = 0;
            valueAxis.renderer.baseGrid.disabled = true;
            valueAxis.renderer.labels.template.disabled = true;
            valueAxis.renderer.grid.template.disabled = true;

            const series = chart.series.push(new am4charts.ColumnSeries());
            series.dataFields.categoryY = "Name";
            series.dataFields.valueX = "Value";
            series.columns.template.strokeOpacity = 0;

            const labelBullet = series.bullets.push(new am4charts.LabelBullet());
            labelBullet.label.horizontalCenter = "right";
            labelBullet.label.dx = -10;
            labelBullet.label.text = "{values.valueX.workingValue.formatDuration('hh:mm')}";
            labelBullet.locationX = 0;

            series.fill = am4core.color("steelblue");

            categoryAxis.sortBySeries = series;

            const columnTemplate = series.columns.template;
            columnTemplate.togglable = true;
            const activeState = columnTemplate.states.create("active");
            activeState.properties.fill = am4core.color("#FF8686")


            let chartData = JSON.parse(data)
            chart.data = chartData.Downtimes;
            columnTemplate.events.on("hit", function (event) {
                clicked.textContent = event.target.dataItem.dataContext.Name
                let noActiveColumn = false
                series.columns.each(function (column) {
                    if (column !== event.target) {
                        column.setState("default");
                        column.isActive = false;
                    }
                    if (column.isActive) {
                        noActiveColumn = true
                    }
                })
                console.log(noActiveColumn)
                if (noActiveColumn) {
                    userDowntimes.innerHTML = ""
                    downtimeUsers.innerHTML = ""
                } else {
                    DownloadDowntimeUsers(event.target.dataItem.dataContext.Name);
                    userDowntimes.innerHTML = ""
                }
            });
        });

    });

}).catch((error) => {
    console.error('Error:', error);
});


function DownloadUsersDowntime(userName) {
    let data = {
        Start: Date.now(),
        End: Date.now(),
        Data: userName,
    };
    fetch("/get_users_downtimes", {
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            am4core.ready(function () {
                const chart = am4core.create("usersDowntimes", am4charts.XYChart);
                chart.padding(40, 40, 40, 40);

                const categoryAxis = chart.yAxes.push(new am4charts.CategoryAxis());
                categoryAxis.renderer.grid.template.location = 0;
                categoryAxis.dataFields.category = "Name";
                categoryAxis.renderer.inversed = true;
                categoryAxis.renderer.grid.template.disabled = true;

                const valueAxis = chart.xAxes.push(new am4charts.ValueAxis());
                valueAxis.min = 0;
                valueAxis.renderer.baseGrid.disabled = true;
                valueAxis.renderer.labels.template.disabled = true;
                valueAxis.renderer.grid.template.disabled = true;

                const series = chart.series.push(new am4charts.ColumnSeries());
                series.dataFields.categoryY = "Name";
                series.dataFields.valueX = "Value";
                series.columns.template.strokeOpacity = 0;

                const labelBullet = series.bullets.push(new am4charts.LabelBullet());
                labelBullet.label.horizontalCenter = "right";
                labelBullet.label.dx = -10;
                labelBullet.label.text = "{values.valueX.workingValue.formatDuration('hh:mm')}";
                labelBullet.locationX = 0;

                series.fill = am4core.color("steelblue");

                categoryAxis.sortBySeries = series;

                let chartData = JSON.parse(data)
                chart.data = chartData.Downtimes;
            });

        });

    }).catch((error) => {
        console.error('Error:', error);
    });
}

function DownloadDowntimeUsers(DowntimeName) {
    let data = {
        Start: Date.now(),
        End: Date.now(),
        DowntimeName: DowntimeName,
    };
    fetch("/get_downtime_users_data", {
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            am4core.ready(function () {
                const chart = am4core.create("downtimeUsers", am4charts.XYChart);
                chart.padding(40, 40, 40, 40);

                const categoryAxis = chart.yAxes.push(new am4charts.CategoryAxis());
                categoryAxis.renderer.grid.template.location = 0;
                categoryAxis.dataFields.category = "Name";
                categoryAxis.renderer.inversed = true;
                categoryAxis.renderer.grid.template.disabled = true;

                const valueAxis = chart.xAxes.push(new am4charts.ValueAxis());
                valueAxis.min = 0;
                valueAxis.renderer.baseGrid.disabled = true;
                valueAxis.renderer.labels.template.disabled = true;
                valueAxis.renderer.grid.template.disabled = true;

                const series = chart.series.push(new am4charts.ColumnSeries());
                series.dataFields.categoryY = "Name";
                series.dataFields.valueX = "Value";
                series.columns.template.strokeOpacity = 0;

                const labelBullet = series.bullets.push(new am4charts.LabelBullet());
                labelBullet.label.horizontalCenter = "right";
                labelBullet.label.dx = -10;
                labelBullet.label.text = "{values.valueX.workingValue.formatDuration('hh:mm')}";
                labelBullet.locationX = 0;

                series.fill = am4core.color("steelblue");

                categoryAxis.sortBySeries = series;

                const columnTemplate = series.columns.template;
                columnTemplate.togglable = true;
                const activeState = columnTemplate.states.create("active");
                activeState.properties.fill = am4core.color("#FF8686")


                let chartData = JSON.parse(data)
                chart.data = chartData.DowntimeUsers;
                columnTemplate.events.on("hit", function (event) {
                    clicked.textContent = event.target.dataItem.dataContext.Name
                    let noActiveColumn = false
                    series.columns.each(function (column) {
                        if (column !== event.target) {
                            column.setState("default");
                            column.isActive = false;
                        }
                        if (column.isActive) {
                            noActiveColumn = true
                        }
                    })
                    if (noActiveColumn) {
                        userDowntimes.innerHTML = ""
                    } else {
                        DownloadUsersDowntime(event.target.dataItem.dataContext.Name);
                    }
                });
            });

        });

    }).catch((error) => {
        console.error('Error:', error);
    });
}


