const ctx = document.getElementById('temp_chart');

async function getLogTemp() {
    try {
        const response = await fetch("http://localhost:6969/temperature-data", { credentials: 'omit' });
        const data = await response.json();
        console.log(data);

        renderChart(data);
    } catch (error) {
        console.error('Failed to fetch temperature data:', error);
    }
}

function renderChart(dataFromAPI) {
    const labels = [];
    const temperaturaData = [];

    dataFromAPI.forEach(objeto => {
        const temperatura = objeto.temperature;
        const datahora = objeto.date_measured;

        labels.push(datahora);
        temperaturaData.push(temperatura);
    });

    chart.data.labels = labels;
    chart.data.datasets = [{
        fill: 'origin',
        label: 'Temperature',
        data: temperaturaData,
        backgroundColor: 'rgba(255, 99, 132, 0.2)',
        borderColor: 'rgba(255, 99, 132, 1)',
        borderWidth: 1
    }];

    chart.update();
}

let chart = new Chart(ctx, {
    type: 'line',
    data: [],
    options: {
        scales: {
            y: {
                beginAtZero: true
            }
        }
    }
});

setInterval(getLogTemp, 10000);
window.addEventListener("load", getLogTemp)