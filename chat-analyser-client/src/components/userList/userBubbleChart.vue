<template>
    <heatmap-chart :data="data" :options="options" ref="heatmap"></heatmap-chart>
</template>

<script>
import 'tui-chart/dist/tui-chart.css'
import {heatmapChart} from '@toast-ui/vue-chart'
export default {
    name : 'userBubbleChart',
    components : {
        'heatmap-chart' : heatmapChart
    },
    data() {
        return {
            data : {
                categories: {
                    x: ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11',
                        '12', '13', '14', '15', '16', '17', '18', '19', '20', '21', '22', '23'],
                    y: ['월', '화', '수', '목', '금', '토', '일']
                },
                series: []
            },
            options : {
                width : 600,
                chart: {           
                    title: '사용자 요일별/시간대별 분포'
                },
                yAxis: {
                    title: '요일'
                },
                xAxis: {
                    title: '시간대'
                },
                series: {
                    showLabel: true
                },
                tooltip: {
                    suffix: '회'
                }
            }
        }
    },
    created() {
        this.$http.get(`${this.api_url}/user/dist`)
            .then((result) => {
                let resultDatas = result.data.aggregations.dayOfWeek
                let isFirst = false;
                for(let dayOfWeek of resultDatas.buckets) {
                    let data = [];
                    for(let hourOfDay of dayOfWeek.hourOfDay.buckets) {
                        data.push(hourOfDay.doc_count)
                    }
                    this.data.series.push(data)
                }
            })
    }
}
</script>

<style>
.tui-heatmap-chart {
    display: inline-block;
    /* width: 100% !important; */
}
/* svg {
    width: 100%;
}
svg {
    width: 100%;
} */
.tui-chart-series-custom-event-area {
    width: 100% !important;
}
</style>
