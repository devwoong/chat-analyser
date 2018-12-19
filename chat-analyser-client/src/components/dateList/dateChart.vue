<template>
    <div>
        <div ref="dateChartBase" style="width:100%; height:50%">
            <area-chart v-if="dateChart.data.categories.length != 0" :data="dateChart.data" :options="dateChart.options" ref="dateChart"></area-chart>
        </div>
        <div style="width:100%; height:50%" ref="dataBarChartBase">
            <bar-chart v-if="dateBarChart.data.categories.length != 0" :data="dateBarChart.data" :options="dateBarChart.options" ref="dateBarChart"></bar-chart>
        </div>
    </div>
</template>
    
<script>
import 'tui-chart/dist/tui-chart.css'
import {areaChart, barChart} from '@toast-ui/vue-chart'
export default {
    name : 'dateChart',
    components : {
        'area-chart' : areaChart,
        'bar-chart' : barChart
    },
    created() {
        this.$http.get(`${this.api_url}/user`)
            .then(function(result){
                this.userList = result.data.aggregations.author.buckets;
                for(var user of this.userList) {
                    this.dateChart.data.series.push({ name : user.key, data:[] })
                    this.dateBarChart.data.series.push({ name : user.key, data:[] })
                }
                this.search();
            }.bind(this))

    },
    beforeDestroy() {
        window.removeEventListener('resize', this.onResize);
    },
    mounted() {
        this.$nextTick(function() {
            window.addEventListener('resize', this.onResize)
        });
    },
    data() {
        return {
            interval : "month",
            userList : [],
            dateChart : {
                data : {
                    categories: [],
                    series: []
                },
                options : {
                    chart: {
                        width: 1000,
                        // height: 540,
                        title: '날짜별 채팅 수'
                    },
                    yAxis: {
                        title: 'hit',
                        pointOnColumn: true
                    },
                    xAxis: {
                        title: '날짜',
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
            dateBarChart : {
                data : {
                    categories : [],
                    series : []
                },
                options : {
                    chart: { 
                        width : 1000,
                        title: '날짜별 채팅 분포'
                    },
                    series: {
                        showLabel: true,
                        zoomable: true,
                        stackType: 'normal'
                    },
                    tooltip: {
                        suffix: '회'
                    }
                }
            }
        }
    },
    methods : {
        onResize(e) {
            this.dateChart.options.width = this.$refs.dateChartBase.clientWidth - 10;
            this.dateBarChart.options.width = this.$refs.dateChartBase.clientWidth - 10;
            this.$refs.dateChart.$data.chartInstance.resize({width:this.dateChart.options.width})
            this.$refs.dateBarChart.$data.chartInstance.resize({width:this.dateBarChart.options.width})
        },
        search() {
            this.$http.get(`${this.api_url}/date/dateChart`, {
                params:{
                    interval:this.interval,
                    toDate : this.toDate,
                    fromDate : this.fromDate
                    },
                    headers:{}, timeout:10000})
                .then(function(result){
                    let resultChats = result.data.aggregations.daily.buckets;
                    for(var date of resultChats) {
                        this.dateChart.data.categories.push(date.key_as_string);
                        this.dateBarChart.data.categories.push(date.key_as_string);
                        for(var series of this.dateChart.data.series) {
                            var isFind = false;
                            for(var author of date.author.buckets) {
                                if( author.key == series.name ){
                                    series.data.push(author.doc_count);
                                    isFind = true;
                                    break;
                                }
                            }
                            if(isFind == false) {
                                series.data.push(0);
                            }
                        }
                        for(var seriesBar of this.dateBarChart.data.series) {
                            var isFind = false;
                            for(var author of date.author.buckets) {
                                if( author.key == seriesBar.name ){
                                    seriesBar.data.push(author.doc_count);
                                    isFind = true;
                                    break;
                                }
                            }
                            if(isFind == false) {
                                seriesBar.data.push(0);
                            }
                        }
                    } 
                }.bind(this))
        }
    },
    props : ['fromDate', 'toDate', 'searchOrder'],
    watch : {
        searchOrder() {
            this.dateChart.data = {categories : [], series:[]}
            this.dateBarChart.data = {categories : [], series:[]}
            // this.dateChart.data.series.splice(0,this.dateChart.data.series.length)
            // this.dateBarChart.data.series.splice(0,this.dateBarChart.data.series.length)
            for(var user of this.userList) {
                this.dateChart.data.series.push({ name : user.key, data:[] })
                this.dateBarChart.data.series.push({ name : user.key, data:[] })
            }
            this.search();
        }
    }
}
</script>

<style>

</style>
