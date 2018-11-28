<template>
<div style="float:left; padding-left:20px"  v-cloak>
    <line-chart v-if="dayOfWeekChart.data.series.length != 0" :data="dayOfWeekChart.data" :options="dayOfWeekChart.options"></line-chart>
    <line-chart v-if="hourOfDayChart.data.series.length != 0" :data="hourOfDayChart.data" :options="hourOfDayChart.options"></line-chart>
</div>
</template>

<script>
import 'tui-chart/dist/tui-chart.css'
import {lineChart} from '@toast-ui/vue-chart'
export default {
    name : 'userDateDistChart',
    components : {
        'line-chart' : lineChart
    },
    data() {
        return {
            weeks : ['월', '화', '수', '목', '금', '토', '일'],
            dayOfWeekChart : {
                data : {
                    categories: [],
                    series: []
                },
                options : {
                    chart: {
                        // width: 1160,
                        // height: 540,
                        title: '요일별 채팅 트랜드'
                    },
                    yAxis: {
                        title: 'hit',
                        pointOnColumn: true
                    },
                    xAxis: {
                        title: '요일',
                        // type: 'datetime',
                        // dateFormat: 'MMM'
                    },
                    series: {
                        showDot: false,
                        zoomable: true
                    },
                    tooltip: {
                        suffix: '회'
                    }
                }
            },
            hourOfDayChart : {
                data : {
                    categories: [],
                    series: []
                },
                options : {
                    chart: {
                        // width: 1160,
                        // height: 540,
                        title: '시간대별 채팅 트랜드'
                    },
                    yAxis: {
                        title: 'hit',
                        pointOnColumn: true
                    },
                    xAxis: {
                        title: '시간',
                        // type: 'datetime',
                        // dateFormat: 'MMM'
                    },
                    series: {
                        showDot: false,
                        zoomable: true
                    },
                    tooltip: {
                        suffix: '회'
                    }
                }
            }
        }
    },
    created() {
        this.$http.get(`${this.api_url}/user/day/${this.author}`)
            .then((result) => {
                var resultDatas = result.data.aggregations.dayOfWeek
                var data = []
                var categories = []
                for(var i = 0; i<resultDatas.buckets.length; i++) {
                    categories.push(this.weeks[resultDatas.buckets[i].key-1])
                    data.push(resultDatas.buckets[i].doc_count)
                }
                this.dayOfWeekChart.data.categories = categories;
                this.dayOfWeekChart.data.series.push({name:this.author, data:data})
                this.$http.get(`${this.api_url}/user/day`)
                    .then((result) => {
                        var resultDatas = result.data.aggregations.dayOfWeek
                        var data = []
                        var categories = []
                        for(var i = 0; i<resultDatas.buckets.length; i++) {
                            categories.push(this.weeks[resultDatas.buckets[i].key-1])
                            data.push(resultDatas.buckets[i].doc_count)
                        }
                        this.dayOfWeekChart.data.series.push({name:'전체', data:data})
                    })
            })


        this.$http.get(`${this.api_url}/user/hour/${this.author}`)
            .then((result) => {
                var resultDatas = result.data.aggregations.hourOfDay
                var data = []
                var categories = []
                for(var i = 0; i<resultDatas.buckets.length; i++) {
                    categories.push(resultDatas.buckets[i].key)
                    data.push(resultDatas.buckets[i].doc_count)
                }
                this.hourOfDayChart.data.categories = categories;
                this.hourOfDayChart.data.series.push({name:this.author, data:data})
                this.$http.get(`${this.api_url}/user/hour`)
                    .then((result) => {
                        var resultDatas = result.data.aggregations.hourOfDay
                        var data = []
                        var categories = []
                        for(var i = 0; i<resultDatas.buckets.length; i++) {
                            categories.push(resultDatas.buckets[i].key)
                            data.push(resultDatas.buckets[i].doc_count)
                        }
                        this.hourOfDayChart.data.series.push({name:'전체', data:data})
                    })
            })
    },
    methods : {

    },
    props : ['author']
}
</script>
