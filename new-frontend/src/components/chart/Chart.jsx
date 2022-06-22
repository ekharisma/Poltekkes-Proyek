import axios from 'axios';
import React, { useState, useEffect, useContext } from 'react'
import { Line } from 'react-chartjs-2'
import 'chartjs-adapter-date-fns';
import 'chartjs-plugin-zoom';
import zoomPlugin from 'chartjs-plugin-zoom';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    TimeScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
ChartJS.register(
    CategoryScale,
    LinearScale,
    TimeScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    zoomPlugin,
    Legend
);

const Chart = () => {
    const [label, setLabels] = useState([])
    const [dataset, setDataset] = useState([])
    const bulan = 6
    const testChart = () => {
        let label = [];
        let dataset = [];

        axios.post("http://35.232.8.249:8081/api/temperature", {
            "month": bulan,
            "year": 2022
        })
            .then(res => {
                console.log(res);
                for (const dataObj of res.data.payload) {
                    // console.log(dataObj.temperature);
                    dataset.push(parseFloat(dataObj.temperature));
                    label.push(dataObj.timestamp);
                }
                setLabels(label)
                setDataset(dataset)
            },
            )
            .catch(err => {
                console.log("error nih bang");
                console.log(err);
            })
    }
    useEffect(() => {
        testChart();
    }, []);
    const options = {
        scales: {
            x: {
                type: 'time',
                time: {
                    unit: "minute",
                }
            },
        },
        responsive: true,
        plugins: {
            legend: {
                position: 'top',
            },
            title: {
                display: true,
                text: 'Chart.js Line Chart',
            },
            zoom: {
                pan: {
                    enabled: true,
                    mode: 'xy'
                },
                zoom: {
                    wheel: {
                        enabled: true,
                    },
                    mode: 'xy',

                },
            }
        }
    };
    const labels = label;
    const data = {
        labels,
        datasets: [
            {
                label: 'Data Suhu',
                data: dataset,
                borderColor: 'rgb(255, 99, 132)',
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
            },
        ],
    };
    return (
        <div className="App">
            <Line options={options} data={data} />
        </div>
    )
}

export default Chart;